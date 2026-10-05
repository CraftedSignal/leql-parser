package leql

import "testing"

func TestExtractConditionsExpression_GroupedNotAppliesDeMorgan(t *testing.T) {
	result := ExtractConditions(`where(NOT ((process.cmd_line ICONTAINS "whoami" AND process.cmd_line ICONTAINS "/priv")))`)
	if len(result.Errors) > 0 {
		t.Fatalf("unexpected parse errors: %v", result.Errors)
	}
	if len(result.Conditions) != 2 {
		t.Fatalf("expected 2 conditions, got %d", len(result.Conditions))
	}
	for index, condition := range result.Conditions {
		if !condition.Negated {
			t.Fatalf("condition %d should be negated: %+v", index, condition)
		}
	}

	assertExpression(t, result.Expression, expressionShape{
		operator: BooleanOr,
		children: []expressionShape{
			{operator: BooleanCondition, conditionIndex: 0},
			{operator: BooleanCondition, conditionIndex: 1},
		},
	})
}

func TestExtractConditionsExpression_MergesFromAndWhereExpressions(t *testing.T) {
	result := ExtractConditions(`from((service=auth OR NOT where(region=us)) AND env=prod) where(status=200)`)
	if len(result.Errors) > 0 {
		t.Fatalf("unexpected parse errors: %v", result.Errors)
	}
	if len(result.Conditions) != 4 {
		t.Fatalf("expected 4 conditions, got %d: %+v", len(result.Conditions), result.Conditions)
	}
	if !result.Conditions[1].Negated {
		t.Fatalf("nested NOT condition should be negated: %+v", result.Conditions[1])
	}

	assertExpression(t, result.Expression, expressionShape{
		operator: BooleanAnd,
		children: []expressionShape{
			{
				operator: BooleanOr,
				children: []expressionShape{
					{operator: BooleanCondition, conditionIndex: 0},
					{operator: BooleanCondition, conditionIndex: 1},
				},
			},
			{operator: BooleanCondition, conditionIndex: 2},
			{operator: BooleanCondition, conditionIndex: 3},
		},
	})
}

func TestExtractConditionsExpression_GroupedNotOrAppliesDeMorgan(t *testing.T) {
	result := ExtractConditions(`where(NOT (process.cmd_line ICONTAINS "whoami" OR process.cmd_line ICONTAINS "/priv"))`)
	if len(result.Errors) > 0 {
		t.Fatalf("unexpected parse errors: %v", result.Errors)
	}
	if len(result.Conditions) != 2 {
		t.Fatalf("expected 2 conditions, got %d", len(result.Conditions))
	}
	for index, condition := range result.Conditions {
		if !condition.Negated {
			t.Fatalf("condition %d should be negated: %+v", index, condition)
		}
	}

	assertExpression(t, result.Expression, expressionShape{
		operator: BooleanAnd,
		children: []expressionShape{
			{operator: BooleanCondition, conditionIndex: 0},
			{operator: BooleanCondition, conditionIndex: 1},
		},
	})
}

func TestExtractConditionsExpression_GroupedNotMultiFieldLeafIsConjoined(t *testing.T) {
	result := ExtractConditions(`where(NOT (process.cmd_line,process.name ICONTAINS "whoami"))`)
	if len(result.Errors) > 0 {
		t.Fatalf("unexpected parse errors: %v", result.Errors)
	}
	if len(result.Conditions) != 2 {
		t.Fatalf("expected 2 conditions, got %d", len(result.Conditions))
	}
	for index, condition := range result.Conditions {
		if !condition.Negated {
			t.Fatalf("condition %d should be negated: %+v", index, condition)
		}
	}

	assertExpression(t, result.Expression, expressionShape{
		operator: BooleanAnd,
		children: []expressionShape{
			{operator: BooleanCondition, conditionIndex: 0},
			{operator: BooleanCondition, conditionIndex: 1},
		},
	})
}

func TestExtractConditionsExpression_SeparateNegationsStayConjoined(t *testing.T) {
	result := ExtractConditions(`where(NOT (process.cmd_line ICONTAINS "whoami") AND NOT process.cmd_line ICONTAINS "/priv")`)
	if len(result.Errors) > 0 {
		t.Fatalf("unexpected parse errors: %v", result.Errors)
	}
	if len(result.Conditions) != 2 {
		t.Fatalf("expected 2 conditions, got %d", len(result.Conditions))
	}

	assertExpression(t, result.Expression, expressionShape{
		operator: BooleanAnd,
		children: []expressionShape{
			{operator: BooleanCondition, conditionIndex: 0},
			{operator: BooleanCondition, conditionIndex: 1},
		},
	})
}

type expressionShape struct {
	operator       BooleanOperator
	conditionIndex int
	children       []expressionShape
}

func assertExpression(t *testing.T, got *ConditionExpression, want expressionShape) {
	t.Helper()
	if got == nil {
		t.Fatalf("expected expression %+v, got nil", want)
	}
	if got.Operator != want.operator {
		t.Fatalf("expected operator %q, got %q in %+v", want.operator, got.Operator, got)
	}
	if got.Operator == BooleanCondition && got.ConditionIndex != want.conditionIndex {
		t.Fatalf("expected condition index %d, got %d", want.conditionIndex, got.ConditionIndex)
	}
	if len(got.Children) != len(want.children) {
		t.Fatalf("expected %d children, got %d in %+v", len(want.children), len(got.Children), got)
	}
	for index := range want.children {
		assertExpression(t, got.Children[index], want.children[index])
	}
}
