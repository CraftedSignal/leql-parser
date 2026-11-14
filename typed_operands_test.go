package leql

import (
	"context"
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestTypedOperandsRegression(t *testing.T) {
	for _, tt := range []struct {
		query  string
		rhs    Operand
		op     string
		errors int
	}{
		{`where(n=42)`, Operand{Kind: OperandNumber, Value: "42"}, "=", 0},
		{`where(n="42")`, Operand{Kind: OperandString, Value: "42"}, "=", 0},
		{`where(a==b)`, Operand{Kind: OperandField, Value: "b"}, "==", 0},
		{`where(a==/x/i)`, Operand{Kind: OperandRegex, Value: "x", RegexFlags: "i"}, "==", 1},
		{`where(a!==IP("10.0.0.0/8"))`, Operand{Kind: OperandString, Value: "10.0.0.0/8", Wrapper: OperandIP}, "!==", 1},
		{`where(a IN [42,"42",/x/i])`, Operand{Kind: OperandList, Items: []Operand{{Kind: OperandNumber, Value: "42"}, {Kind: OperandString, Value: "42"}, {Kind: OperandRegex, Value: "x", RegexFlags: "i"}}}, "in", 0},
	} {
		t.Run(tt.query, func(t *testing.T) {
			a, s := ExtractConditions(tt.query), AnalyzeQuery(tt.query)
			for name, r := range map[string]struct {
				c []Condition
				e []string
			}{"extract": {a.Conditions, a.Errors}, "analyze": {s.Conditions, s.Errors}} {
				if len(r.e) != tt.errors || len(r.c) != 1 {
					t.Fatalf("%s conditions=%+v errors=%v", name, r.c, r.e)
				}
				c := r.c[0]
				if c.Operator != tt.op || !reflect.DeepEqual(c.RHS, &tt.rhs) {
					t.Errorf("%s got %+v rhs=%+v want %+v", name, c, c.RHS, tt.rhs)
				}
				if tt.errors > 0 && (!strings.Contains(r.e[0], "unsupported key-comparison RHS") || c.CaseInsensitive) {
					t.Errorf("invalid key: %+v %v", c, r.e)
				}
			}
		})
	}
}

func requireTypedConditions(t *testing.T, query string, operands []Operand, shape string) []Condition {
	t.Helper()
	extracted, semantic := ExtractConditions(query), AnalyzeQuery(query)
	for name, result := range map[string]struct {
		conditions []Condition
		errors     []string
	}{"extract": {extracted.Conditions, extracted.Errors}, "analyze": {semantic.Conditions, semantic.Errors}} {
		if len(result.errors) > 0 || len(result.conditions) != len(operands) {
			t.Fatalf("%s conditions=%+v errors=%v", name, result.conditions, result.errors)
		}
		for i, c := range result.conditions {
			if !reflect.DeepEqual(c.RHS, &operands[i]) {
				t.Errorf("%s condition %d RHS=%+v want %+v", name, i, c.RHS, operands[i])
			}
			if operands[i].Kind == OperandList {
				if len(c.Alternatives) != len(operands[i].Items) {
					t.Fatalf("unaligned alternatives: %+v", c)
				}
				for j, item := range operands[i].Items {
					if c.Alternatives[j] != item.Value {
						t.Errorf("item %d projection=%q want %q", j, c.Alternatives[j], item.Value)
					}
				}
			} else if c.Value != operands[i].Value {
				t.Errorf("legacy projection=%q want %q", c.Value, operands[i].Value)
			}
		}
	}
	if !reflect.DeepEqual(extracted.Conditions, semantic.Conditions) {
		t.Errorf("semantic clone differs: %+v %+v", extracted.Conditions, semantic.Conditions)
	}
	if shape != "" && literalParseShape(extracted.Expression) != shape {
		t.Errorf("shape=%s want %s", literalParseShape(extracted.Expression), shape)
	}
	var checkLeaves func(*ConditionExpression, *SemanticExpression)
	checkLeaves = func(p *ConditionExpression, s *SemanticExpression) {
		if p == nil {
			if s != nil {
				t.Fatal("unexpected semantic expression")
			}
			return
		}
		if s == nil || p.Operator != s.Operator || len(p.Children) != len(s.Children) {
			t.Fatal("semantic topology differs")
		}
		if p.Operator == BooleanCondition && s.Condition != &semantic.Conditions[p.ConditionIndex] {
			t.Fatal("semantic leaf not bound to cloned condition")
		}
		for i := range p.Children {
			checkLeaves(p.Children[i], s.Children[i])
		}
	}
	checkLeaves(extracted.Expression, semantic.Expression)
	return extracted.Conditions
}

func TestTypedOperandLexicalKinds(t *testing.T) {
	for _, tt := range []struct {
		source string
		kind   OperandKind
		value  string
	}{
		{"42", OperandNumber, "42"}, {`"42"`, OperandString, "42"}, {`'42'`, OperandString, "42"},
		{"-1e-03", OperandNumber, "-1e-03"}, {"9007199254740993", OperandNumber, "9007199254740993"}, {"1.0e+03", OperandNumber, "1.0e+03"},
		{"null", OperandString, "null"}, {"true", OperandString, "true"}, {"false", OperandString, "false"}, {`"null"`, OperandString, "null"},
		{`"""multi word"""`, OperandString, "multi word"}, {`'''multi word'''`, OperandString, "multi word"},
		{`${name}`, OperandVariable, `${name}`}, {"30m", OperandTimeUnit, "30m"}, {"10.0.0.0/8", OperandCIDR, "10.0.0.0/8"},
		{`""`, OperandString, ""}, {`"a\nb"`, OperandString, `a\nb`},
	} {
		t.Run(tt.source, func(t *testing.T) {
			requireTypedConditions(t, "where(n="+tt.source+")", []Operand{{Kind: tt.kind, Value: tt.value}}, "0")
		})
	}
	for _, query := range []string{`where(run)`, `where(42)`, `where(${name})`, `where(/x/i)`} {
		t.Run(query, func(t *testing.T) {
			a, s := ExtractConditions(query), AnalyzeQuery(query)
			if len(a.Errors) > 0 || len(s.Errors) > 0 || len(a.Conditions) != 1 || len(s.Conditions) != 1 || a.Conditions[0].RHS != nil || s.Conditions[0].RHS != nil {
				t.Fatalf("bare search acquired RHS: %+v %+v", a, s)
			}
		})
	}
}

func TestTypedValidKeyOperands(t *testing.T) {
	for _, op := range []string{"==", "!=="} {
		for _, tt := range []struct{ source, value string }{
			{"b", "b"}, {`"b"`, "b"}, {`'b'`, "b"}, {"42", "42"}, {`"42"`, "42"}, {"1e3", "1e3"}, {"1.0", "1.0"}, {"null", "null"}, {"obj.field", "obj.field"}, {`"key with space"`, "key with space"}, {`"loose"`, "loose"},
		} {
			t.Run(op+tt.source, func(t *testing.T) {
				cs := requireTypedConditions(t, "where(a"+op+tt.source+")", []Operand{{Kind: OperandField, Value: tt.value}}, "0")
				if cs[0].Operator != op || cs[0].SourceField != "" {
					t.Fatalf("key role/operator lost: %+v", cs[0])
				}
			})
		}
	}
	if result := ExtractConditions(`where(a==loose)`); len(result.Errors) == 0 {
		t.Fatal("broadened existing bare LOOSE syntax")
	}
	requireTypedConditions(t, `where(a=b)`, []Operand{{Kind: OperandString, Value: "b"}}, "0")
}

func TestTypedKeyPredicateTopology(t *testing.T) {
	for _, tt := range []struct {
		query, shape, op string
		count            int
		negated          bool
	}{
		{`where(a,b==c)`, "or(0,1)", "==", 2, false},
		{`where(ALL(a,b)!=="c")`, "and(0,1)", "!==", 2, false},
		{`where(NOT a==c)`, "0", "==", 1, true},
		{`where(a NOT !==c)`, "0", "!==", 1, true},
		{`where(NOT (a==c OR b!==c))`, "and(0,1)", "", 2, true},
	} {
		t.Run(tt.query, func(t *testing.T) {
			rhs := make([]Operand, tt.count)
			for i := range rhs {
				rhs[i] = Operand{Kind: OperandField, Value: "c"}
			}
			cs := requireTypedConditions(t, tt.query, rhs, tt.shape)
			for _, c := range cs {
				if c.Negated != tt.negated || tt.op != "" && c.Operator != tt.op {
					t.Fatalf("key predicate changed: %+v", c)
				}
			}
		})
	}
	cs := requireTypedConditions(t, `from(a==b) where(c!==d, loose)`, []Operand{{Kind: OperandField, Value: "b"}, {Kind: OperandField, Value: "d"}}, "and(0,1)")
	if cs[0].Loose || !cs[1].Loose || cs[0].Operator != "==" || cs[1].Operator != "!==" {
		t.Fatalf("scoped key metadata lost: %+v", cs)
	}
}

func TestTypedInvalidKeyOperands(t *testing.T) {
	for _, tt := range []struct {
		source string
		rhs    Operand
	}{
		{`/x/i`, Operand{Kind: OperandRegex, Value: "x", RegexFlags: "i"}},
		{`/admin/msxU`, Operand{Kind: OperandRegex, Value: "admin", RegexFlags: "msxU"}},
		{`NOCASE(b)`, Operand{Kind: OperandString, Value: "b", Wrapper: OperandNoCase}},
		{`NOCASE(42)`, Operand{Kind: OperandNumber, Value: "42", Wrapper: OperandNoCase}},
		{`NOCASE(/x/i)`, Operand{Kind: OperandRegex, Value: "x", RegexFlags: "i", Wrapper: OperandNoCase}},
		{`IP("10.0.0.0/8")`, Operand{Kind: OperandString, Value: "10.0.0.0/8", Wrapper: OperandIP}},
		{`IP(10.0.0.0/8)`, Operand{Kind: OperandCIDR, Value: "10.0.0.0/8", Wrapper: OperandIP}},
		{`IP(/x/i)`, Operand{Kind: OperandRegex, Value: "x", RegexFlags: "i", Wrapper: OperandIP}},
		{`"""b"""`, Operand{Kind: OperandString, Value: "b"}},
		{`'''b'''`, Operand{Kind: OperandString, Value: "b"}},
		{`${name}`, Operand{Kind: OperandVariable, Value: `${name}`}},
		{`30m`, Operand{Kind: OperandTimeUnit, Value: "30m"}},
		{`10.0.0.0/8`, Operand{Kind: OperandCIDR, Value: "10.0.0.0/8"}},
	} {
		for _, op := range []string{"==", "!=="} {
			for _, form := range []struct {
				prefix, suffix string
				count          int
				negated        bool
			}{
				{"a", "", 1, false}, {"a,b", "", 2, false}, {"ALL(a,b)", "", 2, false}, {"NOT a", "", 1, true}, {"a NOT ", "", 1, true}, {"NOT (a", ")", 1, true},
			} {
				// Wrappers have no postfix-NOT grammar production; IP also lacks
				// an ALL-fields production. Keep those syntax-recovery cases below.
				if tt.rhs.Wrapper != "" && form.prefix == "a NOT " || tt.rhs.Wrapper == OperandIP && form.prefix == "ALL(a,b)" {
					continue
				}
				query := "where(" + form.prefix + op + tt.source + form.suffix + ")"
				t.Run(query, func(t *testing.T) {
					a, s := ExtractConditions(query), AnalyzeQuery(query)
					for name, result := range map[string]struct {
						c []Condition
						e []string
					}{"extract": {a.Conditions, a.Errors}, "analyze": {s.Conditions, s.Errors}} {
						if len(result.e) != 1 || !strings.HasPrefix(result.e[0], "unsupported key-comparison RHS for "+op+": expected field name") || len(result.c) != form.count {
							t.Fatalf("%s conditions=%+v errors=%v", name, result.c, result.e)
						}
						for _, c := range result.c {
							if c.Operator != op || c.CaseInsensitive || c.Negated != form.negated || c.SourceField != "" || !reflect.DeepEqual(c.RHS, &tt.rhs) || c.Value != tt.rhs.Value {
								t.Fatalf("%s key metadata changed: %+v rhs=%+v", name, c, c.RHS)
							}
							flags := ""
							if tt.rhs.Kind == OperandRegex && tt.rhs.Wrapper == "" {
								flags = tt.rhs.RegexFlags
							}
							if c.RegexFlags != flags {
								t.Errorf("legacy flags=%q want %q", c.RegexFlags, flags)
							}
						}
					}
					if a.Expression == nil || s.Expression == nil {
						t.Fatal("invalid key lost best-effort tree")
					}
				})
			}
		}
	}
	a := ExtractConditions(`where(a==/x/i AND ok=1)`)
	s := AnalyzeQuery(`where(a==/x/i AND ok=1)`)
	if len(a.Errors) != 1 || len(a.Conditions) != 2 || literalParseShape(a.Expression) != "and(0,1)" || len(s.Errors) != 1 || len(s.Conditions) != 2 || s.Expression == nil {
		t.Fatalf("invalid key destroyed valid sibling/tree: %+v %+v", a, s)
	}
	// Preserve the two generated prefix-NOT key variants migrated from #19.
	for _, op := range []string{"==", "!=="} {
		result := extractConditionRuleForTest(t, "NOT user,actor"+op+"/admin/im")
		if len(result.Errors) != 1 || len(result.Conditions) != 2 {
			t.Fatalf("prefix key errors/expansion=%+v", result)
		}
		for _, c := range result.Conditions {
			if c.Operator != op || !c.Negated || c.CaseInsensitive || c.RHS.Kind != OperandRegex || c.RHS.RegexFlags != "im" {
				t.Fatalf("prefix key metadata=%+v", c)
			}
		}
	}
}

func TestTypedOperandVisitorFamilies(t *testing.T) {
	// Separate typed expectations supplement the older tests of legacy projections.
	for _, tt := range []struct {
		predicate string
		rhs       Operand
		count     int
	}{
		{`a="x"`, Operand{Kind: OperandString, Value: "x"}, 1},
		{`a=~/x/mi`, Operand{Kind: OperandRegex, Value: "x", RegexFlags: "mi"}, 1},
		{`NOT a=~/x/mi`, Operand{Kind: OperandRegex, Value: "x", RegexFlags: "mi"}, 1},
		{`a NOT !~/x/mi`, Operand{Kind: OperandRegex, Value: "x", RegexFlags: "mi"}, 1},
		{`a CONTAINS "x"`, Operand{Kind: OperandString, Value: "x"}, 1},
		{`NOT a CONTAINS "x"`, Operand{Kind: OperandString, Value: "x"}, 1},
		{`a NOT CONTAINS "x"`, Operand{Kind: OperandString, Value: "x"}, 1},
		{`a IN [1,"x"]`, Operand{Kind: OperandList, Items: []Operand{{Kind: OperandNumber, Value: "1"}, {Kind: OperandString, Value: "x"}}}, 1},
		{`NOT a IN [1,"x"]`, Operand{Kind: OperandList, Items: []Operand{{Kind: OperandNumber, Value: "1"}, {Kind: OperandString, Value: "x"}}}, 1},
		{`a NOT IN [1,"x"]`, Operand{Kind: OperandList, Items: []Operand{{Kind: OperandNumber, Value: "1"}, {Kind: OperandString, Value: "x"}}}, 1},
		{`a CONTAINS-ANY ["x"]`, Operand{Kind: OperandList, Items: []Operand{{Kind: OperandString, Value: "x"}}}, 1},
		{`NOT a CONTAINS-ANY ["x"]`, Operand{Kind: OperandList, Items: []Operand{{Kind: OperandString, Value: "x"}}}, 1},
		{`a NOT CONTAINS-ANY ["x"]`, Operand{Kind: OperandList, Items: []Operand{{Kind: OperandString, Value: "x"}}}, 1},
		{`a=NOCASE(42)`, Operand{Kind: OperandNumber, Value: "42", Wrapper: OperandNoCase}, 1},
		{`a=IP("10.0.0.0/8")`, Operand{Kind: OperandString, Value: "10.0.0.0/8", Wrapper: OperandIP}, 1},
		{`ALL(a,b)=1`, Operand{Kind: OperandNumber, Value: "1"}, 2},
		{`ALL(a,b) IN [1]`, Operand{Kind: OperandList, Items: []Operand{{Kind: OperandNumber, Value: "1"}}}, 2},
		{`ALL(a,b)=NOCASE("x")`, Operand{Kind: OperandString, Value: "x", Wrapper: OperandNoCase}, 2},
	} {
		t.Run(tt.predicate, func(t *testing.T) {
			rhs := make([]Operand, tt.count)
			for i := range rhs {
				rhs[i] = tt.rhs
			}
			requireTypedConditions(t, "where("+tt.predicate+")", rhs, "")
		})
	}
	// Exercise generated prefix-NOT visitor contexts directly, independently of
	// whole-query expression-level NOT selection.
	for _, tt := range []struct {
		predicate string
		rhs       Operand
	}{
		{`NOT a,b=1`, Operand{Kind: OperandNumber, Value: "1"}},
		{`NOT a,b=~/x/i`, Operand{Kind: OperandRegex, Value: "x", RegexFlags: "i"}},
		{`NOT a,b CONTAINS "x"`, Operand{Kind: OperandString, Value: "x"}},
		{`NOT a,b IN [1]`, Operand{Kind: OperandList, Items: []Operand{{Kind: OperandNumber, Value: "1"}}}},
		{`NOT a,b CONTAINS-ANY ["x"]`, Operand{Kind: OperandList, Items: []Operand{{Kind: OperandString, Value: "x"}}}},
	} {
		t.Run(tt.predicate, func(t *testing.T) {
			r := extractConditionRuleForTest(t, tt.predicate)
			if len(r.Errors) > 0 || len(r.Conditions) != 2 {
				t.Fatalf("prefix result=%+v", r)
			}
			for _, c := range r.Conditions {
				if !c.Negated || !reflect.DeepEqual(c.RHS, &tt.rhs) {
					t.Fatalf("prefix RHS=%+v", c.RHS)
				}
			}
		})
	}
}

func TestTypedMixedLists(t *testing.T) {
	mixed := Operand{Kind: OperandList, Items: []Operand{{Kind: OperandNumber, Value: "42"}, {Kind: OperandString, Value: "42"}, {Kind: OperandString, Value: "plain"}, {Kind: OperandRegex, Value: "x", RegexFlags: "i"}, {Kind: OperandRegex, Value: "x", RegexFlags: "m"}, {Kind: OperandString, Value: "10.0.0.0/8", Wrapper: OperandIP}, {Kind: OperandCIDR, Value: "10.0.0.0/8", Wrapper: OperandIP}}}
	for _, op := range []string{"IN", "IIN"} {
		for _, prefix := range []string{"a ", "NOT a ", "a NOT ", "ALL(a,b) ", "NOT (ALL(a,b) "} {
			query := "where(" + prefix + op + ` [42,"42",plain,/x/i,/x/m,IP("10.0.0.0/8"),IP(10.0.0.0/8)]`
			if strings.Contains(prefix, "(") && strings.HasPrefix(prefix, "NOT (") {
				query += ")"
			}
			query += ")"
			count := 1
			if strings.Contains(prefix, "ALL") {
				count = 2
			}
			rhs := make([]Operand, count)
			for i := range rhs {
				rhs[i] = mixed
			}
			t.Run(query, func(t *testing.T) {
				cs := requireTypedConditions(t, query, rhs, "")
				for _, c := range cs {
					if c.RegexFlags != "" || c.Operator != strings.ToLower(op) {
						t.Fatalf("list projected heterogeneous flags/operator: %+v", c)
					}
				}
			})
		}
	}
	// This positive case replaces the precise IN rejection from #19.
	requireTypedConditions(t, `where(user IN [/admin/i])`, []Operand{{Kind: OperandList, Items: []Operand{{Kind: OperandRegex, Value: "admin", RegexFlags: "i"}}}}, "0")
	cs := requireTypedConditions(t, `where(a IN [IP(10.0.0.0/8),IP("192.168.0.0/16")])`, []Operand{{Kind: OperandList, Items: []Operand{{Kind: OperandCIDR, Value: "10.0.0.0/8", Wrapper: OperandIP}, {Kind: OperandString, Value: "192.168.0.0/16", Wrapper: OperandIP}}}}, "0")
	if cs[0].Operator != "cidr" {
		t.Fatalf("lost all-IP normalization: %+v", cs)
	}
	for _, query := range []string{`where(a CONTAINS-ANY [/x/i])`, `where(a IN [IP(/x/i)])`, `where(a=NOCASE(/x/i))`, `where(a=/x/i AND b>/x/i)`} {
		for name, errs := range map[string][]string{"extract": ExtractConditions(query).Errors, "analyze": AnalyzeQuery(query).Errors} {
			if len(errs) != 1 || !strings.Contains(errs[0], "regex literal") {
				t.Errorf("%s %s errors=%v", name, query, errs)
			}
		}
	}
}

func TestTypedOperandSyntaxRecoveryBoundaries(t *testing.T) {
	for _, query := range []string{`where(a NOT ==NOCASE(b))`, `where(a NOT !==IP("10.0.0.0/8"))`, `where(ALL(a,b)==IP(10.0.0.0/8))`, `where(a==[1,2])`} {
		a, s := ExtractConditions(query), AnalyzeQuery(query)
		if len(a.Errors) == 0 || len(s.Errors) == 0 {
			t.Fatalf("unsupported grammar syntax accepted: %s", query)
		}
	}
}

func TestTypedOperandDedupIdentity(t *testing.T) {
	for _, tt := range []struct {
		query string
		want  int
	}{
		{`where(a=42 OR a="42" OR a='42' OR a=42)`, 2},
		{`where(a==b OR a=="b" OR a=='b')`, 1},
		{`where(a==42 OR a=="42" OR a=42 OR a="42")`, 3},
		{`where(a IN [/x/i] OR a IN [/x/m] OR a IN [/x/i])`, 2},
		{`where(a IN [42,"42"] OR a IN ["42",42] OR a IN [42,"42"])`, 2},
		{`where(a=42 OR a=NOCASE(42))`, 2},
		{`where(a IN [IP("10.0.0.0/8")] OR a IN ["10.0.0.0/8"])`, 2},
		{`from(a==b) where(a==b, loose)`, 2},
	} {
		t.Run(tt.query, func(t *testing.T) {
			r := ExtractConditions(tt.query)
			if len(r.Errors) > 0 {
				t.Fatal(r.Errors)
			}
			if got := DeduplicateConditions(r.Conditions); len(got) != tt.want {
				t.Fatalf("dedup got %d want %d: %+v", len(got), tt.want, got)
			}
		})
	}
	base := Condition{Field: "a", Operator: "=", Value: "42"}
	typed := base
	typed.RHS = &Operand{Kind: OperandNumber, Value: "42"}
	stringValue := typed
	stringValue.RHS = &Operand{Kind: OperandString, Value: "42"}
	field := typed
	field.RHS = &Operand{Kind: OperandField, Value: "42"}
	wrapped := typed
	wrapped.RHS = &Operand{Kind: OperandNumber, Value: "42", Wrapper: OperandNoCase}
	if got := DeduplicateConditions([]Condition{base, typed, stringValue, field, wrapped, typed}); len(got) != 5 {
		t.Fatalf("nil/type/wrapper identity collapsed: %+v", got)
	}
	// Ordered typed items have a structured identity, even when legacy strings
	// contain delimiters previously used to build dedup keys.
	left := base
	left.RHS = &Operand{Kind: OperandList, Items: []Operand{{Kind: OperandString, Value: "a\x00b"}, {Kind: OperandString, Value: "c"}}}
	right := base
	right.RHS = &Operand{Kind: OperandList, Items: []Operand{{Kind: OperandString, Value: "a"}, {Kind: OperandString, Value: "b\x00c"}}}
	if got := DeduplicateConditions([]Condition{left, right, left}); len(got) != 2 {
		t.Fatalf("typed list identity ambiguous: %+v", got)
	}
}

func TestTypedOperandCloneIsolation(t *testing.T) {
	for _, query := range []string{`where(a,b=42)`, `where(a,b IN [42,/x/i,IP("10.0.0.0/8")])`} {
		t.Run(query, func(t *testing.T) {
			extracted := ExtractConditions(query)
			if len(extracted.Errors) > 0 || len(extracted.Conditions) != 2 {
				t.Fatalf("extraction=%+v", extracted)
			}
			semantic := semanticFromExtraction(extracted, nil)
			want := cloneOperand(semantic.Conditions[0].RHS, nil)
			extracted.Conditions[0].RHS.Value = "mutated"
			if extracted.Conditions[0].RHS.Kind == OperandList {
				extracted.Conditions[0].RHS.Items[1].Value = "mutated"
				extracted.Conditions[0].RHS.Items[1].RegexFlags = "U"
				extracted.Conditions[0].RHS.Items[2].Wrapper = OperandNoCase
				extracted.Conditions[0].RHS.Items = append(extracted.Conditions[0].RHS.Items, Operand{Kind: OperandNumber, Value: "7"})
			}
			if !reflect.DeepEqual(semantic.Conditions[0].RHS, want) || !reflect.DeepEqual(semantic.Conditions[1].RHS, want) || !reflect.DeepEqual(extracted.Conditions[1].RHS, want) {
				t.Fatalf("mutable operand aliased across ownership: extraction=%+v semantics=%+v", extracted.Conditions, semantic.Conditions)
			}
			if semantic.Expression == nil || len(semantic.Expression.Children) != 2 || semantic.Expression.Children[0].Condition != &semantic.Conditions[0] || semantic.Expression.Children[1].Condition != &semantic.Conditions[1] {
				t.Fatal("semantic expression not bound to cloned conditions")
			}
			semantic.Conditions[0].RHS.Kind = OperandVariable
			if semantic.Conditions[1].RHS.Kind != want.Kind {
				t.Fatal("semantic siblings share RHS")
			}
		})
	}
	// Deep copying also preserves ownership for manual operands with nested items.
	manual := &Operand{Kind: OperandList, Items: []Operand{{Kind: OperandList, Items: []Operand{{Kind: OperandString, Value: "inner"}}}}}
	clone := cloneOperand(manual, nil)
	manual.Items[0].Items[0].Value = "changed"
	if clone.Items[0].Items[0].Value != "inner" {
		t.Fatal("nested manual items were not cloned")
	}
	legacy := semanticFromExtraction(&ParseResult{Conditions: []Condition{{Field: "a", Value: "42"}}}, nil)
	if legacy.Conditions[0].RHS != nil {
		t.Fatal("legacy unknown RHS was inferred")
	}
}

func TestTypedOperandJSONAndContextParity(t *testing.T) {
	legacy, err := json.Marshal(Condition{Field: "a", Value: "42"})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(legacy), `"rhs"`) {
		t.Fatalf("nil RHS appeared in JSON: %s", legacy)
	}
	for _, query := range []string{`where(a=9007199254740993)`, `where(a="")`, `where(a==b)`, `where(a IN [42,"42",/x/mi])`, `where(a==NOCASE(/x/i))`} {
		a, s := ExtractConditions(query), AnalyzeQuery(query)
		ac, sc := ExtractConditionsContext(context.Background(), query), AnalyzeQueryContext(context.Background(), query)
		if !reflect.DeepEqual(a, ac) || !reflect.DeepEqual(s, sc) {
			t.Fatalf("context API parity differs for %s", query)
		}
		data, err := json.Marshal(a.Conditions[0])
		if err != nil {
			t.Fatal(err)
		}
		var decoded Condition
		if err := json.Unmarshal(data, &decoded); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(decoded, a.Conditions[0]) {
			t.Fatalf("operand JSON roundtrip differs: %s", data)
		}
		if strings.Contains(query, "9007199254740993") && !strings.Contains(string(data), `"rhs":{"kind":"number","value":"9007199254740993"}`) {
			t.Fatalf("number was normalized or encoded as JSON number: %s", data)
		}
		if query == `where(a="")` && !strings.Contains(string(data), `"rhs":{"kind":"string"}`) {
			t.Fatalf("empty string encoding ambiguous: %s", data)
		}
	}
}

func TestTypedOperandBudgetReservations(t *testing.T) {
	list := parseLEQLValueListForTest(t, `[42,"42"]`)
	for _, tt := range []struct {
		name               string
		call               func(*parseBudget)
		limit, wantCharged int
	}{
		{"scalar", func(b *parseBudget) { valueOperand(nil, "", b) }, 0, 0},
		{"temporary list", func(b *parseBudget) { listOperand(list, b) }, 2, 1},
		{"clone list", func(b *parseBudget) {
			cloneOperand(&Operand{Kind: OperandList, Items: []Operand{{Kind: OperandNumber}, {Kind: OperandString}}}, b)
		}, 2, 1},
		{"semantic clone", func(b *parseBudget) {
			cloneSemanticConditions([]Condition{{RHS: &Operand{Kind: OperandList, Items: []Operand{{Kind: OperandNumber}, {Kind: OperandString}}}}}, b)
		}, 3, 2},
	} {
		t.Run(tt.name, func(t *testing.T) {
			limits := defaultParseLimits()
			limits.maxOutputItems = tt.limit
			b, cancel := newParseBudget(context.Background(), limits)
			defer cancel()
			defer func() {
				p := recover()
				abort, ok := p.(budgetAbort)
				if !ok || !strings.Contains(abort.err.Error(), "output items") {
					t.Fatalf("expected reservation rejection got %v", p)
				}
				if b.output != tt.wantCharged {
					t.Fatalf("reservation partially charged: %d want %d", b.output, tt.wantCharged)
				}
			}()
			tt.call(b)
		})
	}
	// Extraction already pays for the temporary typed list, the owning list copy,
	// and both legacy projections. Semantic copies must use that same budget.
	limits := defaultParseLimits()
	limits.maxOutputItems = 11
	query := `where(a IN [42,"42"])`
	if r := extractConditionsContext(context.Background(), query, limits); len(r.Errors) > 0 || len(r.Conditions) != 1 {
		t.Fatalf("bounded extraction should fit: %+v", r)
	}
	if r := analyzeQueryContext(context.Background(), query, limits); len(r.Errors) != 1 || !strings.Contains(r.Errors[0], "output items") || len(r.Conditions) != 0 || r.Expression != nil {
		t.Fatalf("semantic copy escaped cumulative budget: %+v", r)
	}
	limits.maxOutputItems = 10
	if r := extractConditionsContext(context.Background(), query, limits); len(r.Errors) != 1 || len(r.Conditions) != 0 || r.Expression != nil {
		t.Fatalf("typed output budget escaped failure boundary: %+v", r)
	}
}

func TestTypedOperandHelpersHonorCancellation(t *testing.T) {
	list := parseLEQLValueListForTest(t, `[42,/x/i]`)
	for _, tt := range []struct {
		name string
		call func(*parseBudget)
	}{
		{"scalar", func(b *parseBudget) { valueOperand(nil, "", b) }},
		{"list", func(b *parseBudget) { listOperand(list, b) }},
		{"clone", func(b *parseBudget) {
			cloneOperand(&Operand{Kind: OperandList, Items: []Operand{{Kind: OperandNumber}}}, b)
		}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			b, stop := newParseBudget(ctx, defaultParseLimits())
			defer stop()
			defer func() {
				p := recover()
				abort, ok := p.(budgetAbort)
				if !ok || !strings.Contains(abort.err.Error(), "context canceled") {
					t.Fatalf("canceled helper continued: %v", p)
				}
				if b.output != 0 {
					t.Fatal("canceled helper reserved output")
				}
			}()
			tt.call(b)
		})
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	a, s := ExtractConditionsContext(ctx, `where(a IN [42,/x/i])`), AnalyzeQueryContext(ctx, `where(a IN [42,/x/i])`)
	if len(a.Errors) != 1 || len(s.Errors) != 1 || len(a.Conditions) != 0 || len(s.Conditions) != 0 || a.Expression != nil || s.Expression != nil {
		t.Fatalf("canceled typed parse leaked output: %+v %+v", a, s)
	}
}

// These exact #19 rows now diagnose unsupported key operands, preserving their
// original operator and flags instead of treating them as regex comparisons.
func TestTypedMigratedKeyRegexRows(t *testing.T) {
	for _, tt := range []struct {
		query, operator, flags string
		negated                bool
	}{
		{`where(user==/admin/msxU)`, "==", "msxU", false},
		{`where(user!==/admin/i)`, "!==", "i", false},
		{`where(user NOT ==/admin/i)`, "==", "i", true},
		{`where(NOT user==/admin/i)`, "==", "i", true},
		{`where(NOT user!==/admin/i)`, "!==", "i", true},
		{`where(user NOT !==/admin/i)`, "!==", "i", true},
	} {
		t.Run(tt.query, func(t *testing.T) {
			for name, result := range map[string]struct {
				c []Condition
				e []string
			}{"extract": {ExtractConditions(tt.query).Conditions, ExtractConditions(tt.query).Errors}, "analyze": {AnalyzeQuery(tt.query).Conditions, AnalyzeQuery(tt.query).Errors}} {
				if len(result.e) != 1 || !strings.HasPrefix(result.e[0], "unsupported key-comparison RHS") || len(result.c) != 1 {
					t.Fatalf("%s invalid-key result=%+v", name, result)
				}
				c := result.c[0]
				want := &Operand{Kind: OperandRegex, Value: "admin", RegexFlags: tt.flags}
				if c.Field != "user" || c.Operator != tt.operator || c.Value != "admin" || c.RegexFlags != tt.flags || c.CaseInsensitive || c.Negated != tt.negated || !reflect.DeepEqual(c.RHS, want) {
					t.Fatalf("%s migrated key semantics=%+v RHS=%+v", name, c, c.RHS)
				}
			}
		})
	}
}

func TestTypedInvalidKeysPreserveScopedBooleanTree(t *testing.T) {
	query := `from(source==target) where(NOT (a==/x/i OR ALL(b,c)!==NOCASE(/y/m)), loose)`
	a, s := ExtractConditions(query), AnalyzeQuery(query)
	if len(a.Errors) != 2 || len(s.Errors) != 2 || len(a.Conditions) != 4 || len(s.Conditions) != 4 || literalParseShape(a.Expression) != "and(0,1,or(2,3))" {
		t.Fatalf("invalid key tree/diagnostics changed: %+v %+v shape=%s", a, s, literalParseShape(a.Expression))
	}
	for i, c := range a.Conditions {
		if c.Loose != (i > 0) || c.Negated != (i > 0) || c.CaseInsensitive || c.RHS == nil {
			t.Fatalf("invalid scoped metadata: %+v", c)
		}
		if !reflect.DeepEqual(c, s.Conditions[i]) {
			t.Fatalf("semantic condition changed: %+v %+v", c, s.Conditions[i])
		}
	}
	if s.Expression == nil || len(s.Expression.Children) != 3 || s.Expression.Children[0].Condition != &s.Conditions[0] || s.Expression.Children[1].Condition != &s.Conditions[1] || s.Expression.Children[2].Children[1].Condition != &s.Conditions[3] {
		t.Fatal("best-effort semantic expression references changed")
	}
}
