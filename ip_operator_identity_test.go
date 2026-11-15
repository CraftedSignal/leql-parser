package leql

import (
	"context"
	"reflect"
	"strings"
	"testing"
)

func TestIPComparisonPolarityAndExpressionIdentity(t *testing.T) {
	const cidr = "192.168.0.0/24"
	wantRHS := &Operand{Kind: OperandCIDR, Value: cidr, Wrapper: OperandIP}
	for _, tt := range []struct {
		name, query, shape string
		fields             []string
		negated            []bool
	}{
		{"equal", `where(source_address=IP(192.168.0.0/24))`, "0", []string{"source_address"}, []bool{false}},
		{"unequal", `where(source_address!=IP(192.168.0.0/24))`, "0", []string{"source_address"}, []bool{true}},
		{"not equal", `where(NOT (source_address=IP(192.168.0.0/24)))`, "0", []string{"source_address"}, []bool{true}},
		{"not unequal", `where(NOT (source_address!=IP(192.168.0.0/24)))`, "0", []string{"source_address"}, []bool{false}},
		{"field list", `where(source_address,destination_address!=IP(192.168.0.0/24))`, "or(0,1)", []string{"source_address", "destination_address"}, []bool{true, true}},
		{"not field list", `where(NOT (source_address,destination_address!=IP(192.168.0.0/24)))`, "and(0,1)", []string{"source_address", "destination_address"}, []bool{false, false}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			extracted := ExtractConditions(tt.query)
			semantic := AnalyzeQuery(tt.query)
			contextual := AnalyzeQueryContext(context.Background(), tt.query)
			if !reflect.DeepEqual(semantic, contextual) {
				t.Fatalf("ordinary/context semantic mismatch: ordinary=%+v context=%+v", semantic, contextual)
			}
			if len(extracted.Errors) != 0 || len(semantic.Errors) != 0 || len(extracted.Conditions) != len(tt.fields) {
				t.Fatalf("conditions=%+v errors=%v semantic errors=%v", extracted.Conditions, extracted.Errors, semantic.Errors)
			}
			if !reflect.DeepEqual(extracted.Conditions, semantic.Conditions) {
				t.Fatalf("extraction/semantic conditions differ: %+v %+v", extracted.Conditions, semantic.Conditions)
			}
			if extracted.Loose || semantic.Loose || !reflect.DeepEqual(extracted.Commands, []string{"where"}) || !reflect.DeepEqual(semantic.Commands, []string{"where"}) {
				t.Fatalf("WHERE metadata changed: extraction=%+v semantic=%+v", extracted, semantic)
			}
			for i, condition := range extracted.Conditions {
				if condition.Field != tt.fields[i] || condition.Operator != "cidr" || condition.Negated != tt.negated[i] || condition.Value != cidr || !reflect.DeepEqual(condition.RHS, wantRHS) || condition.Loose || condition.PipeStage != 0 {
					t.Errorf("condition %d = %+v; RHS=%+v", i, condition, condition.RHS)
				}
			}
			if got := literalParseShape(extracted.Expression); got != tt.shape {
				t.Errorf("expression shape=%q, want %q", got, tt.shape)
			}
			assertIPExpressionCorrespondence(t, extracted.Expression, semantic.Expression, semantic.Conditions)
		})
	}
}

func TestIPComparisonGroupedSiblingAndWhereMode(t *testing.T) {
	query := `where((source_address!=IP(192.168.0.0/24) AND status=ok) OR destination_address=IP(192.168.0.0/24), loose)`
	extracted := ExtractConditions(query)
	semantic := AnalyzeQuery(query)
	contextual := AnalyzeQueryContext(context.Background(), query)
	if !reflect.DeepEqual(semantic, contextual) || !reflect.DeepEqual(extracted.Conditions, semantic.Conditions) {
		t.Fatalf("API parity: extraction=%+v ordinary=%+v context=%+v", extracted, semantic, contextual)
	}
	if len(extracted.Errors) != 0 || len(extracted.Conditions) != 3 || !extracted.Loose || !semantic.Loose || !reflect.DeepEqual(extracted.Commands, []string{"where"}) {
		t.Fatalf("unexpected extraction: %+v", extracted)
	}
	if got := literalParseShape(extracted.Expression); got != "or(and(0,1),2)" {
		t.Errorf("shape=%q", got)
	}
	for i, want := range []struct {
		field, operator string
		negated         bool
	}{
		{"source_address", "cidr", true},
		{"status", "=", false},
		{"destination_address", "cidr", false},
	} {
		condition := extracted.Conditions[i]
		if condition.Field != want.field || condition.Operator != want.operator || condition.Negated != want.negated || !condition.Loose || condition.PipeStage != 0 {
			t.Errorf("sibling %d = %+v", i, condition)
		}
		if i != 1 && !reflect.DeepEqual(condition.RHS, &Operand{Kind: OperandCIDR, Value: "192.168.0.0/24", Wrapper: OperandIP}) {
			t.Errorf("IP sibling %d RHS=%+v", i, condition.RHS)
		}
	}
	assertIPExpressionCorrespondence(t, extracted.Expression, semantic.Expression, semantic.Conditions)
}

func TestIPComparisonUnsupportedOperatorAndKeyControls(t *testing.T) {
	for _, operator := range []string{">", ">=", "<", "<="} {
		t.Run(operator, func(t *testing.T) {
			query := "where(source_address" + operator + "IP(192.168.0.0/24))"
			extracted, semantic := ExtractConditions(query), AnalyzeQuery(query)
			if len(extracted.Conditions) != 1 || len(semantic.Conditions) != 1 || extracted.Conditions[0].Operator != operator || semantic.Conditions[0].Operator != operator {
				t.Fatalf("operator identity lost: %+v %+v", extracted, semantic)
			}
			if len(extracted.Errors) != 1 || !strings.Contains(extracted.Errors[0], "unsupported IP comparison operator") || !strings.Contains(extracted.Errors[0], operator) || !reflect.DeepEqual(extracted.Errors, semantic.Errors) {
				t.Fatalf("diagnostic missing: extraction=%v semantic=%v", extracted.Errors, semantic.Errors)
			}
			assertIPExpressionCorrespondence(t, extracted.Expression, semantic.Expression, semantic.Conditions)
		})
	}
	for _, operator := range []string{"==", "!=="} {
		t.Run(operator, func(t *testing.T) {
			query := "where(source_address" + operator + "IP(192.168.0.0/24))"
			extracted, semantic := ExtractConditions(query), AnalyzeQuery(query)
			if len(extracted.Conditions) != 1 || len(extracted.Errors) != 1 || !strings.Contains(extracted.Errors[0], "unsupported key-comparison RHS") || extracted.Conditions[0].Operator != operator || !reflect.DeepEqual(extracted.Conditions, semantic.Conditions) || !reflect.DeepEqual(extracted.Errors, semantic.Errors) {
				t.Fatalf("key comparison changed: extraction=%+v semantic=%+v", extracted, semantic)
			}
		})
	}
}

func assertIPExpressionCorrespondence(t *testing.T, parsed *ConditionExpression, semantic *SemanticExpression, conditions []Condition) {
	t.Helper()
	if parsed == nil || semantic == nil || parsed.Operator != semantic.Operator || len(parsed.Children) != len(semantic.Children) {
		t.Fatalf("expression topology differs: parsed=%+v semantic=%+v", parsed, semantic)
	}
	if parsed.Operator == BooleanCondition {
		if parsed.ConditionIndex < 0 || parsed.ConditionIndex >= len(conditions) || semantic.Condition != &conditions[parsed.ConditionIndex] {
			t.Fatalf("expression leaf does not bind condition: parsed=%+v semantic=%+v", parsed, semantic)
		}
	}
	for i := range parsed.Children {
		assertIPExpressionCorrespondence(t, parsed.Children[i], semantic.Children[i], conditions)
	}
}
