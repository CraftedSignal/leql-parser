package leql

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// Rapid7 documents unquoted keyword searches and /pattern/i literals:
// https://docs.rapid7.com/insightidr/components-for-building-a-query/
// These cases catch field-existence misclassification, regex/scalar conflation,
// dropped flags, and lost boolean grouping at both public entry points.
func TestLiteralSemanticsKeywords(t *testing.T) {
	tests := []struct {
		query  string
		leaves []literalLeaf
		shape  string
	}{
		{`where(run)`, []literalLeaf{{Operator: "keyword", Value: "run", LogicalOp: "AND"}}, "0"},
		{`where("run")`, []literalLeaf{{Operator: "keyword", Value: "run", LogicalOp: "AND"}}, "0"},
		{`where(Amazon AND Boardman)`, []literalLeaf{{Operator: "keyword", Value: "Amazon", LogicalOp: "AND"}, {Operator: "keyword", Value: "Boardman", LogicalOp: "AND"}}, "and(0,1)"},
		{`where(Amazon Boardman)`, []literalLeaf{{Operator: "keyword", Value: "Amazon", LogicalOp: "AND"}, {Operator: "keyword", Value: "Boardman", LogicalOp: "AND"}}, "and(0,1)"},
		{`where(Amazon OR NOT Boardman)`, []literalLeaf{{Operator: "keyword", Value: "Amazon", LogicalOp: "AND"}, {Operator: "keyword", Value: "Boardman", Negated: true, LogicalOp: "OR"}}, "or(0,1)"},
	}
	for _, tt := range tests {
		t.Run(tt.query, func(t *testing.T) { assertLiteralSemantics(t, tt.query, tt.leaves, tt.shape) })
	}
}

func TestLiteralSemanticsRegex(t *testing.T) {
	tests := []struct {
		query, field, op, value, flags string
		negated, caseInsensitive       bool
	}{
		{`where(/error/)`, "", "regex", "error", "", false, false},
		{`where(/error/i)`, "", "regex", "error", "i", false, true},
		{`where(/error/msxU)`, "", "regex", "error", "msxU", false, false},
		{`where(/a\/b/im)`, "", "regex", `a\/b`, "im", false, true},
		{`where(NOT /error/i)`, "", "regex", "error", "i", true, true},
		{`where(user=/.*admin.*/i)`, "user", "=~", ".*admin.*", "i", false, true},
		{`where(user=/.*admin.*/)`, "user", "=~", ".*admin.*", "", false, false},
		{`where(user!=/admin/i)`, "user", "!~", "admin", "i", false, true},
		{`where(NOT user=/admin/i)`, "user", "=~", "admin", "i", true, true},
		{`where(user NOT =/admin/i)`, "user", "=~", "admin", "i", true, true},
		{`where(user NOT !=/admin/i)`, "user", "!~", "admin", "i", true, true},
		{`where(user=/a\/b/mi)`, "user", "=~", `a\/b`, "mi", false, true},
		{`where(user='.*admin.*')`, "user", "=", ".*admin.*", "", false, false},
		{`where('run')`, "", "keyword", "run", "", false, false},
		{`where(NOT user!=/admin/i)`, "user", "!~", "admin", "i", true, true},
		{`where(user=~/admin/im)`, "user", "=~", "admin", "im", false, true},
		{`where(NOT user!~/admin/s)`, "user", "!~", "admin", "s", true, false},
		{`where(user NOT =~/admin/xU)`, "user", "=~", "admin", "xU", true, false},
		{`where(user=".*admin.*")`, "user", "=", ".*admin.*", "", false, false},
		{`where(user="/admin/i")`, "user", "=", "/admin/i", "", false, false},
		{`where("/admin/i")`, "", "keyword", "/admin/i", "", false, false},
		{`where(user=~"/admin/i")`, "user", "=~", "/admin/i", "", false, false},
		{`where(user>3)`, "user", ">", "3", "", false, false},
	}
	for _, tt := range tests {
		t.Run(tt.query, func(t *testing.T) {
			assertLiteralSemantics(t, tt.query, []literalLeaf{{Field: tt.field, Operator: tt.op, Value: tt.value, RegexFlags: tt.flags, CaseInsensitive: tt.caseInsensitive, Negated: tt.negated, LogicalOp: "AND"}}, "0")
		})
	}
}

func TestLiteralSemanticsMultiFieldRegex(t *testing.T) {
	tests := []struct {
		query, shape, secondOp string
		negated                bool
	}{
		{`where(user,actor=/admin/i)`, "or(0,1)", "OR", false},
		{`where(ALL(user,actor)=/admin/i)`, "and(0,1)", "AND", false},
		{`where(NOT (user,actor=/admin/i))`, "and(0,1)", "OR", true},
		{`where(NOT (ALL(user,actor)=/admin/i))`, "or(0,1)", "AND", true},
		{`where(user,actor NOT =/admin/i)`, "or(0,1)", "OR", true},
	}
	for _, tt := range tests {
		t.Run(tt.query, func(t *testing.T) {
			assertLiteralSemantics(t, tt.query, []literalLeaf{
				{Field: "user", Operator: "=~", Value: "admin", RegexFlags: "i", CaseInsensitive: true, Negated: tt.negated, LogicalOp: "AND"},
				{Field: "actor", Operator: "=~", Value: "admin", RegexFlags: "i", CaseInsensitive: true, Negated: tt.negated, LogicalOp: tt.secondOp},
			}, tt.shape)
		})
	}
}

func TestLiteralSemanticsUnsupportedRegexOperands(t *testing.T) {
	for _, query := range []string{`where(user>/admin/i)`, `where(user CONTAINS /admin/i)`, `where(user CONTAINS-ANY [/admin/m])`, `where(user=NOCASE(/admin/i))`, `where(run) calculate(count) having(count=/admin/i)`} {
		t.Run(query, func(t *testing.T) {
			for name, errs := range map[string][]string{"extract": ExtractConditions(query).Errors, "analyze": AnalyzeQuery(query).Errors} {
				if !strings.Contains(strings.Join(errs, " "), "regex literal") {
					t.Errorf("%s silently accepts unsupported regex operand: %v", name, errs)
				}
			}
		})
	}
}

type literalLeaf struct {
	Field           string `json:"field"`
	Operator        string `json:"operator"`
	Value           string `json:"value"`
	Negated         bool   `json:"negated"`
	CaseInsensitive bool   `json:"case_insensitive"`
	RegexFlags      string `json:"regex_flags"`
	LogicalOp       string `json:"logical_op"`
}

func assertLiteralSemantics(t *testing.T, query string, want []literalLeaf, shape string) {
	t.Helper()
	extracted, semantic := ExtractConditions(query), AnalyzeQuery(query)
	for name, result := range map[string]struct {
		conditions []Condition
		errors     []string
	}{"extract": {extracted.Conditions, extracted.Errors}, "analyze": {semantic.Conditions, semantic.Errors}} {
		if len(result.errors) > 0 {
			t.Fatalf("%s errors: %v", name, result.errors)
		}
		data, err := json.Marshal(result.conditions)
		if err != nil {
			t.Fatal(err)
		}
		var got []literalLeaf
		if err := json.Unmarshal(data, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s leaves = %+v; want %+v", name, got, want)
		}
	}
	if got := literalParseShape(extracted.Expression); got != shape {
		t.Errorf("extract shape=%s want %s", got, shape)
	}
	var semanticShape func(*SemanticExpression) string
	semanticShape = func(expr *SemanticExpression) string {
		if expr == nil {
			return "nil"
		}
		if expr.Operator == BooleanCondition {
			for i := range semantic.Conditions {
				if expr.Condition == &semantic.Conditions[i] {
					return fmt.Sprint(i)
				}
			}
			t.Errorf("semantic expression leaf is not a semantic condition: %+v", expr.Condition)
			return "missing"
		}
		var children []string
		for _, child := range expr.Children {
			children = append(children, semanticShape(child))
		}
		return string(expr.Operator) + "(" + strings.Join(children, ",") + ")"
	}
	if got := semanticShape(semantic.Expression); got != shape {
		t.Errorf("analyze shape=%s want %s", got, shape)
	}
}

func literalParseShape(expr *ConditionExpression) string {
	if expr == nil {
		return "nil"
	}
	if expr.Operator == BooleanCondition {
		return fmt.Sprint(expr.ConditionIndex)
	}
	var children []string
	for _, child := range expr.Children {
		children = append(children, literalParseShape(child))
	}
	return string(expr.Operator) + "(" + strings.Join(children, ",") + ")"
}

func TestLiteralSemanticsDedupKeepsDistinctRegexFlags(t *testing.T) {
	result := ExtractConditions(`where(user=/admin/ OR user=/admin/i OR user=/admin/m OR user=/admin/i)`)
	conditions := DeduplicateConditions(result.Conditions)
	if len(conditions) != 3 {
		t.Fatalf("dedup merged distinct regex flags: %+v", conditions)
	}
}

func TestLiteralSemanticsConditionRulePrefixNot(t *testing.T) {
	// Parse the condition rule directly to cover the generated prefix-NOT contexts;
	// the whole-query parser can choose an expression-level NOT instead.
	for _, tt := range []struct{ query, op string }{
		{`NOT user,actor=/admin/im`, "=~"},
		{`NOT user,actor!=/admin/im`, "!~"},
		{`NOT user,actor=~/admin/im`, "=~"},
	} {
		query := tt.query
		t.Run(query, func(t *testing.T) {
			result := extractConditionRuleForTest(t, query)
			requireNoParseErrors(t, result)
			if len(result.Conditions) != 2 {
				t.Fatalf("conditions=%+v", result.Conditions)
			}
			for _, condition := range result.Conditions {
				data, _ := json.Marshal(condition)
				var got literalLeaf
				_ = json.Unmarshal(data, &got)
				if got.Operator != tt.op || got.RegexFlags != "im" || !got.CaseInsensitive || !got.Negated {
					t.Errorf("prefix NOT lost regex semantics: %+v", got)
				}
			}
		})
	}
}
