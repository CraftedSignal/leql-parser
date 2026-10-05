package leql

import "testing"

func TestAnalyzeQueryPreservesLEQLAdvancedSemantics(t *testing.T) {
	semantic := AnalyzeQuery(`where((process.name =~ /mimikatz/i OR process.name !~ /benign/) AND source_address IN [IP(10.0.0.0/8), IP("192.168.0.0/16")]) groupby(user) calculate(count) sort(desc) limit(10)`)
	if semantic == nil {
		t.Fatal("expected semantic query")
	}
	if len(semantic.Errors) != 0 {
		t.Fatalf("unexpected errors: %v", semantic.Errors)
	}
	if len(semantic.Conditions) != 3 {
		t.Fatalf("expected 3 conditions, got %+v", semantic.Conditions)
	}
	if semantic.Conditions[0].Field != "process.name" || semantic.Conditions[0].Operator != "=~" || !semantic.Conditions[0].CaseInsensitive {
		t.Fatalf("case-insensitive regex semantics lost: %+v", semantic.Conditions[0])
	}
	if semantic.Conditions[1].Operator != "!~" || semantic.Conditions[1].Negated {
		t.Fatalf("regex-not-match semantics lost: %+v", semantic.Conditions[1])
	}
	if semantic.Conditions[2].Operator != "cidr" || len(semantic.Conditions[2].Alternatives) != 2 {
		t.Fatalf("CIDR alternatives lost: %+v", semantic.Conditions[2])
	}
	if len(semantic.GroupByFields) != 1 || semantic.GroupByFields[0] != "user" {
		t.Fatalf("groupby fields not preserved: %+v", semantic.GroupByFields)
	}
	if semantic.ComputedFields["calculate"] != "count()" {
		t.Fatalf("calculate command not preserved: %+v", semantic.ComputedFields)
	}
	if got := semantic.Commands; len(got) != 5 || got[0] != "where" || got[1] != "groupby" || got[2] != "calculate" || got[3] != "sort" || got[4] != "limit" {
		t.Fatalf("commands not preserved: %+v", got)
	}

	expression := semantic.Expression
	if expression == nil || expression.Operator != BooleanAnd || len(expression.Children) != 2 {
		t.Fatalf("expected top-level AND expression, got %#v", expression)
	}
	if expression.Children[0].Operator != BooleanOr || len(expression.Children[0].Children) != 2 {
		t.Fatalf("expected regex OR group, got %#v", expression.Children[0])
	}
	if expression.Children[1].Condition == nil || expression.Children[1].Condition.Field != "source_address" {
		t.Fatalf("expected CIDR leaf, got %#v", expression.Children[1])
	}
}

func TestAnalyzeQueryBuildsExpressionFromLogicalOps(t *testing.T) {
	semantic := AnalyzeQuery(`where(status=200 OR result=OK)`)
	if len(semantic.Errors) != 0 {
		t.Fatalf("unexpected errors: %v", semantic.Errors)
	}
	if semantic.Expression == nil || semantic.Expression.Operator != BooleanOr || len(semantic.Expression.Children) != 2 {
		t.Fatalf("expected OR expression from parser tree/logical ops, got %#v", semantic.Expression)
	}
}

func TestSemanticInternalExtractionClonesAndPreservesJoinParity(t *testing.T) {
	result := &ParseResult{
		Conditions: []Condition{{
			Field:        "status",
			Operator:     "in",
			Value:        "200, 201",
			Alternatives: []string{"200", "201"},
		}},
		GroupByFields:  []string{"user"},
		ComputedFields: map[string]string{"calculate": "count()"},
		Commands:       []string{"where", "groupby", "calculate"},
		Joins: []JoinInfo{{
			Type:          "unused",
			JoinFields:    []string{"user"},
			LeftFields:    []string{"left_user"},
			RightFields:   []string{"right_user"},
			RightTable:    "right",
			Subsearch:     &ParseResult{Conditions: []Condition{{Field: "child", Operator: "=", Value: "1"}}},
			PipeStage:     1,
			ExposedFields: []string{"child"},
		}},
		Errors: []string{"warning"},
	}

	semantic := semanticFromExtraction(result)
	if semantic == nil {
		t.Fatal("expected semantic query")
	}

	result.Conditions[0].Alternatives[0] = "mutated"
	result.GroupByFields[0] = "mutated"
	result.ComputedFields["calculate"] = "mutated"
	result.Commands[0] = "mutated"
	result.Joins[0].JoinFields[0] = "mutated"
	result.Joins[0].LeftFields[0] = "mutated"
	result.Joins[0].RightFields[0] = "mutated"
	result.Joins[0].ExposedFields[0] = "mutated"
	result.Errors[0] = "mutated"

	if semantic.Conditions[0].Alternatives[0] != "200" || semantic.GroupByFields[0] != "user" || semantic.ComputedFields["calculate"] != "count()" {
		t.Fatalf("semantic data was aliased: %+v", semantic)
	}
	if semantic.Commands[0] != "where" || semantic.Errors[0] != "warning" {
		t.Fatalf("commands/errors were aliased: %+v", semantic)
	}
	join := semantic.Joins[0]
	if join.JoinFields[0] != "user" || join.LeftFields[0] != "left_user" || join.RightFields[0] != "right_user" || join.ExposedFields[0] != "child" {
		t.Fatalf("join metadata was aliased: %+v", join)
	}
	if join.Subquery == nil || len(join.Subquery.Conditions) != 1 || join.Subquery.Conditions[0].Field != "child" {
		t.Fatalf("join subquery not preserved: %+v", join.Subquery)
	}
}

func TestSemanticInternalExtractionEdges(t *testing.T) {
	if got := semanticFromExtraction(nil); got != nil {
		t.Fatalf("nil parse result converted to %+v", got)
	}
	empty := semanticFromExtraction(&ParseResult{})
	if empty == nil {
		t.Fatal("expected empty semantic query")
	}
	if empty.Expression != nil || empty.Conditions != nil || empty.Joins != nil {
		t.Fatalf("empty parse result should not synthesize semantics: %+v", empty)
	}

	conditions := []Condition{{Field: "a"}, {Field: "b", LogicalOp: "OR"}}
	if got := semanticExpressionFromParseExpression(&ConditionExpression{Operator: BooleanCondition, ConditionIndex: -1}, conditions); got != nil {
		t.Fatalf("negative condition index converted to %#v", got)
	}
	if got := semanticExpressionFromParseExpression(&ConditionExpression{Operator: BooleanCondition, ConditionIndex: 2}, conditions); got != nil {
		t.Fatalf("out-of-range condition index converted to %#v", got)
	}
	if got := semanticExpressionFromParseExpression(&ConditionExpression{Operator: "unknown"}, conditions); got != nil {
		t.Fatalf("unknown operator converted to %#v", got)
	}
	if got := semanticExpressionFromParseExpression(&ConditionExpression{Operator: BooleanAnd}, conditions); got != nil {
		t.Fatalf("empty AND converted to %#v", got)
	}
	if got := semanticExpressionFromParseExpression(&ConditionExpression{
		Operator: BooleanAnd,
		Children: []*ConditionExpression{{Operator: BooleanCondition, ConditionIndex: 0}, {Operator: BooleanCondition, ConditionIndex: 99}},
	}, conditions); got != nil {
		t.Fatalf("invalid child expression converted to %#v", got)
	}

	left := &SemanticExpression{Operator: BooleanAnd, Children: []*SemanticExpression{{Operator: BooleanCondition, Condition: &conditions[0]}}}
	right := &SemanticExpression{Operator: BooleanAnd, Children: []*SemanticExpression{{Operator: BooleanCondition, Condition: &conditions[1]}}}
	combined := combineSemanticExpressions(BooleanAnd, left, right)
	if combined == nil || len(combined.Children) != 2 {
		t.Fatalf("same-operator combine did not flatten children: %#v", combined)
	}
	fromLogicalOps := semanticExpressionFromConditions([]Condition{
		{Field: "first"},
		{Field: "second"},
		{Field: "third", LogicalOp: "OR"},
	})
	if fromLogicalOps == nil || fromLogicalOps.Operator != BooleanOr || len(fromLogicalOps.Children) != 2 {
		t.Fatalf("logical-op expression not rebuilt correctly: %#v", fromLogicalOps)
	}
	if got := combineSemanticExpressions(BooleanAnd, nil, right); got != right {
		t.Fatalf("nil left combine = %#v", got)
	}
	if got := combineSemanticExpressions(BooleanAnd, left, nil); got != left {
		t.Fatalf("nil right combine = %#v", got)
	}
}
