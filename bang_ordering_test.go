package leql

import (
	"context"
	"reflect"
	"strings"
	"testing"
)

func TestBangOrderingTypedPredicates(t *testing.T) {
	for _, tc := range []struct {
		atom, operator string
		rhs            Operand
		negated        bool
		count          int
	}{
		{`a!>100`, ">", Operand{Kind: OperandNumber, Value: "100"}, true, 1},
		{`a!>=100`, ">=", Operand{Kind: OperandNumber, Value: "100"}, true, 1},
		{`a!<100`, "<", Operand{Kind: OperandNumber, Value: "100"}, true, 1},
		{`a!<=100`, "<=", Operand{Kind: OperandNumber, Value: "100"}, true, 1},
		{`NOT a!>100`, ">", Operand{Kind: OperandNumber, Value: "100"}, false, 1},
		{`a NOT !>100`, ">", Operand{Kind: OperandNumber, Value: "100"}, false, 1},
		{`a!>NOCASE("z")`, ">", Operand{Kind: OperandString, Value: "z", Wrapper: OperandNoCase}, true, 1},
		{`ALL(a,b)!>100`, ">", Operand{Kind: OperandNumber, Value: "100"}, true, 2},
		{`ALL(a,b)!>NOCASE("z")`, ">", Operand{Kind: OperandString, Value: "z", Wrapper: OperandNoCase}, true, 2},
		{`a!>Target`, ">", Operand{Kind: OperandField, Value: "Target"}, true, 1},
		{`a!>"100"`, ">", Operand{Kind: OperandString, Value: "100"}, true, 1},
		{`a!>10MB`, ">", Operand{Kind: OperandString, Value: "10MB"}, true, 1},
	} {
		t.Run(tc.atom, func(t *testing.T) {
			operands, shape := []Operand{tc.rhs}, "0"
			if tc.count == 2 {
				operands, shape = append(operands, tc.rhs), "and(0,1)"
			}
			for _, c := range requireTypedConditions(t, "where("+tc.atom+")", operands, shape) {
				if c.Operator != tc.operator || c.Negated != tc.negated || c.Loose || c.IsComputed || c.SourceField != "" {
					t.Errorf("bang predicate changed projection: %+v", c)
				}
			}
		})
	}
}

func TestBangOrderingBooleanComposition(t *testing.T) {
	for _, tc := range []struct {
		atom, shape string
		negated     []bool
	}{
		{`a,b!>100`, "or(0,1)", []bool{true, true}},
		{`NOT (a,b!>100)`, "and(0,1)", []bool{false, false}},
		{`ALL(a,b)!>100`, "and(0,1)", []bool{true, true}},
		{`NOT (ALL(a,b)!>100)`, "or(0,1)", []bool{false, false}},
		{`NOT NOT (a,b!>100)`, "or(0,1)", []bool{true, true}},
		{`NOT (a,b!>100 OR c=ok)`, "and(0,1,2)", []bool{false, false, true}},
	} {
		t.Run(tc.atom, func(t *testing.T) {
			rhs := []Operand{{Kind: OperandNumber, Value: "100"}, {Kind: OperandNumber, Value: "100"}}
			if len(tc.negated) == 3 {
				rhs = append(rhs, Operand{Kind: OperandString, Value: "ok"})
			}
			cs := requireTypedConditions(t, "where("+tc.atom+")", rhs, tc.shape)
			for i, c := range cs {
				if c.Negated != tc.negated[i] {
					t.Errorf("leaf %d polarity: %+v", i, c)
				}
			}
		})
	}
	cs := requireTypedConditions(t, `from(a!>1) where(b!<2, loose)`, []Operand{{Kind: OperandNumber, Value: "1"}, {Kind: OperandNumber, Value: "2"}}, "and(0,1)")
	if cs[0].Loose || !cs[1].Loose || !cs[0].Negated || !cs[1].Negated || cs[0].Operator != ">" || cs[1].Operator != "<" {
		t.Errorf("scope projection changed: %+v", cs)
	}
	p := ExtractConditions(`where(a,b!>Target)`)
	s := semanticFromExtraction(p, nil)
	if len(p.Conditions) != 2 || len(s.Conditions) != 2 {
		t.Fatalf("missing clone fixture conditions: %+v", p)
	}
	p.Conditions[0].RHS.Value = "mutated"
	want := &Operand{Kind: OperandField, Value: "Target"}
	if !reflect.DeepEqual(p.Conditions[1].RHS, want) || !reflect.DeepEqual(s.Conditions[0].RHS, want) || !reflect.DeepEqual(s.Conditions[1].RHS, want) {
		t.Fatal("bang operands alias siblings or semantic clones")
	}
}

func TestBangOrderingBoundaries(t *testing.T) {
	for _, tc := range []struct {
		atom, operator string
		rhs            Operand
	}{
		{`a>1`, ">", Operand{Kind: OperandNumber, Value: "1"}},
		{`a!=b`, "!=", Operand{Kind: OperandString, Value: "b"}},
		{`a!==b`, "!==", Operand{Kind: OperandField, Value: "b"}},
		{`a!~/x/`, "!~", Operand{Kind: OperandRegex, Value: "x"}},
	} {
		cs := requireTypedConditions(t, "where("+tc.atom+")", []Operand{tc.rhs}, "0")
		if cs[0].Operator != tc.operator || cs[0].Negated {
			t.Errorf("legacy operator changed: %+v", cs)
		}
	}
	for _, query := range []string{`where(a!>IP(10.0.0.0/8))`, `where(a!>/x/i)`, `calculate(count) having(count!>1)`, `where(a!!>1)`, `where(a!=>1)`} {
		p, s := ExtractConditions(query), AnalyzeQuery(query)
		if len(p.Errors) == 0 || len(s.Errors) == 0 {
			t.Errorf("unsupported boundary accepted: %s: %+v / %+v", query, p, s)
		}
		if strings.Contains(query, "IP(") && (len(p.Conditions) != 1 || p.Conditions[0].Operator != ">" || !p.Conditions[0].Negated || !strings.Contains(strings.Join(p.Errors, " "), "unsupported IP comparison")) {
			t.Errorf("IP bang polarity/diagnostic lost: %+v", p)
		}
	}
}

func TestBangOrderingAtomicResources(t *testing.T) {
	const query = `where(a,b!>Target)`
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if p := ExtractConditionsContext(ctx, query); len(p.Errors) == 0 || len(p.Conditions) != 0 || p.Expression != nil {
		t.Fatalf("canceled bang parse returned a model: %+v", p)
	}
	for _, resource := range []string{"work", "output"} {
		limits := defaultParseLimits()
		if resource == "work" {
			limits.maxWork = 1
		} else {
			limits.maxOutputItems = 1
		}
		p := extractConditionsContext(context.Background(), query, limits)
		s := analyzeQueryContext(context.Background(), query, limits)
		if len(p.Errors) == 0 || len(s.Errors) == 0 || len(p.Conditions) != 0 || len(s.Conditions) != 0 || p.Expression != nil || s.Expression != nil {
			t.Errorf("%s failure returned partial bang conditions: %+v / %+v", resource, p, s)
		}
	}
}
