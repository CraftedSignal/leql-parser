package leql

import (
	"context"
	"reflect"
	"strings"
	"testing"
)

func TestWave4gOrderingFieldOperands(t *testing.T) {
	for _, op := range []string{"<", ">", "<=", ">="} {
		for _, tc := range []struct {
			source, value string
			kind          OperandKind
		}{
			{"KeyB", "KeyB", OperandField},
			{"3", "3", OperandNumber},
			{`"KeyB"`, "KeyB", OperandString},
			{`'KeyB'`, "KeyB", OperandString},
		} {
			t.Run(op+tc.source, func(t *testing.T) {
				query := "where(KeyA" + op + tc.source + ")"
				cs := requireTypedConditions(t, query, []Operand{{Kind: tc.kind, Value: tc.value}}, "0")
				if cs[0].Field != "KeyA" || cs[0].Operator != op || cs[0].Value != tc.value || cs[0].SourceField != "" || cs[0].Negated || cs[0].CaseInsensitive || cs[0].IsComputed {
					t.Fatalf("legacy comparison projection changed: %+v", cs[0])
				}
				if ClassifyFieldProvenance(ExtractConditions(query), "KeyB") != ProvenanceMain {
					t.Fatal("legacy field provenance contract changed")
				}
			})
		}
		requireTypedConditions(t, "where(KeyA"+op+"KeyB OR KeyA"+op+`"KeyB")`, []Operand{{Kind: OperandField, Value: "KeyB"}, {Kind: OperandString, Value: "KeyB"}}, "or(0,1)")
	}
	for _, tc := range []struct {
		op   string
		kind OperandKind
	}{
		{"==", OperandField}, {"!==", OperandField}, {"=", OperandString}, {"!=", OperandString},
	} {
		requireTypedConditions(t, "where(KeyA"+tc.op+"KeyB)", []Operand{{Kind: tc.kind, Value: "KeyB"}}, "0")
	}
}

func TestWave4gOrderingExpansionScopeAndClone(t *testing.T) {
	for _, tc := range []struct {
		query, shape string
		count        int
		negated      bool
	}{
		{`where(KeyA,Other<KeyB)`, "or(0,1)", 2, false},
		{`where(ALL(KeyA,Other)>=KeyB)`, "and(0,1)", 2, false},
		{`where(NOT KeyA<KeyB)`, "0", 1, true},
		{`where(KeyA NOT <=KeyB)`, "0", 1, true},
	} {
		t.Run(tc.query, func(t *testing.T) {
			rhs := make([]Operand, tc.count)
			for i := range rhs {
				rhs[i] = Operand{Kind: OperandField, Value: "KeyB"}
			}
			cs := requireTypedConditions(t, tc.query, rhs, tc.shape)
			for _, c := range cs {
				if c.Negated != tc.negated || c.SourceField != "" {
					t.Fatalf("negation/source projection changed: %+v", c)
				}
			}
		})
	}
	cs := requireTypedConditions(t, `from(KeyA<KeyB) where(KeyC>=KeyD, loose)`, []Operand{{Kind: OperandField, Value: "KeyB"}, {Kind: OperandField, Value: "KeyD"}}, "and(0,1)")
	if cs[0].Loose || !cs[1].Loose || cs[0].Field != "KeyA" || cs[1].Field != "KeyC" || cs[0].Operator != "<" || cs[1].Operator != ">=" {
		t.Fatalf("FROM/WHERE scope changed: %+v", cs)
	}
	p := ExtractConditions(`where(KeyA,Other<KeyB)`)
	s := semanticFromExtraction(p, nil)
	want := Operand{Kind: OperandField, Value: "KeyB"}
	p.Conditions[0].RHS.Kind, p.Conditions[0].RHS.Value = OperandString, "mutated"
	if !reflect.DeepEqual(p.Conditions[1].RHS, &want) || !reflect.DeepEqual(s.Conditions[0].RHS, &want) || !reflect.DeepEqual(s.Conditions[1].RHS, &want) {
		t.Fatal("typed RHS aliases siblings or semantic clone")
	}
	if s.Expression.Children[0].Condition != &s.Conditions[0] || s.Expression.Children[1].Condition != &s.Conditions[1] {
		t.Fatal("semantic leaves do not own cloned conditions")
	}
}

func TestWave4gOrderingSharedBudgetsAndCancellation(t *testing.T) {
	const query = `from(KeyA<KeyB) where(ALL(KeyC,KeyD)>=KeyE, loose)`
	for _, semantic := range []bool{false, true} {
		b, cancel := newParseBudget(context.Background(), defaultParseLimits())
		p := extractConditionsInternalWithOptions(query, b, Options{})
		if semantic {
			_ = semanticFromExtraction(p, b)
		}
		used := *b
		cancel()
		if len(p.Errors) != 0 || len(p.Conditions) != 3 {
			t.Fatalf("budget control: %+v", p)
		}
		for _, resource := range []struct {
			name string
			used int
			set  func(*parseLimits, int)
		}{
			{"work budget", used.work, func(l *parseLimits, n int) { l.maxWork = n }},
			{"output items", used.output, func(l *parseLimits, n int) { l.maxOutputItems = n }},
		} {
			for _, delta := range []int{0, -1} {
				l := defaultParseLimits()
				resource.set(&l, resource.used+delta)
				var cs []Condition
				var errs []string
				var tree bool
				if semantic {
					r := analyzeQueryContext(context.Background(), query, l)
					cs, errs, tree = r.Conditions, r.Errors, r.Expression != nil
				} else {
					r := extractConditionsContext(context.Background(), query, l)
					cs, errs, tree = r.Conditions, r.Errors, r.Expression != nil
				}
				if delta == 0 && (len(errs) > 0 || len(cs) != 3 || !tree) || delta < 0 && (len(cs) > 0 || tree || !strings.Contains(strings.Join(errs, " "), resource.name)) {
					t.Fatalf("%s exact/one-short semantic=%v delta=%d: %+v %v tree=%v", resource.name, semantic, delta, cs, errs, tree)
				}
			}
		}
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	p, s := ExtractConditionsContext(ctx, query), AnalyzeQueryContext(ctx, query)
	if len(p.Conditions) != 0 || len(s.Conditions) != 0 || p.Expression != nil || s.Expression != nil || !strings.Contains(strings.Join(p.Errors, " "), "context canceled") || !strings.Contains(strings.Join(s.Errors, " "), "context canceled") {
		t.Fatalf("cancellation leaked output: %+v / %+v", p, s)
	}
}

func TestWave4gOrderingNumericLookingBoundary(t *testing.T) {
	for _, op := range []string{"<", ">", "<=", ">="} {
		for _, tc := range []struct {
			source, value string
			kind          OperandKind
		}{
			{"KeyB", "KeyB", OperandField},
			{"object.KeyB", "object.KeyB", OperandField},
			{"100bytes", "100bytes", OperandString},
			{"10MB", "10MB", OperandString},
			{"100ms", "100ms", OperandString},
			{"-10MB", "-10MB", OperandString},
			{"-100ms", "-100ms", OperandString},
			{"-3", "-3", OperandNumber},
			{`"100ms"`, "100ms", OperandString},
		} {
			t.Run(op+tc.source, func(t *testing.T) {
				requireTypedConditions(t, "where(KeyA"+op+tc.source+")", []Operand{{Kind: tc.kind, Value: tc.value}}, "0")
			})
		}
	}
}
