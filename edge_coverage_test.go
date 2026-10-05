package leql

import (
	"strings"
	"testing"
	"time"

	"github.com/antlr4-go/antlr/v4"
)

func extractConditionRuleForTest(t *testing.T, condition string) *ParseResult {
	t.Helper()

	input := antlr.NewInputStream(condition)
	lexer := NewLEQLLexer(input)
	lexer.RemoveErrorListeners()
	lexerErrors := &errorListener{}
	lexer.AddErrorListener(lexerErrors)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewLEQLParser(stream)
	parser.RemoveErrorListeners()
	parserErrors := &errorListener{}
	parser.AddErrorListener(parserErrors)

	extractor := &conditionExtractor{
		conditions:       make([]Condition, 0),
		groupByFields:    make([]string, 0),
		commands:         make([]string, 0),
		currentLogicalOp: "AND",
	}
	extractor.visit(parser.Condition())

	errs := append([]string{}, lexerErrors.errors...)
	errs = append(errs, parserErrors.errors...)
	errs = append(errs, extractor.errors...)
	return &ParseResult{
		Conditions: extractor.conditions,
		Commands:   extractor.commands,
		Errors:     errs,
	}
}

func extractExpressionForTest(t *testing.T, expression string) *ParseResult {
	t.Helper()

	input := antlr.NewInputStream(expression)
	lexer := NewLEQLLexer(input)
	lexer.RemoveErrorListeners()
	lexerErrors := &errorListener{}
	lexer.AddErrorListener(lexerErrors)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewLEQLParser(stream)
	parser.RemoveErrorListeners()
	parserErrors := &errorListener{}
	parser.AddErrorListener(parserErrors)

	extractor := &conditionExtractor{
		conditions:       make([]Condition, 0),
		groupByFields:    make([]string, 0),
		commands:         make([]string, 0),
		currentLogicalOp: "AND",
	}
	extractor.visit(parser.Expression())

	errs := append([]string{}, lexerErrors.errors...)
	errs = append(errs, parserErrors.errors...)
	errs = append(errs, extractor.errors...)
	return &ParseResult{
		Conditions: extractor.conditions,
		Commands:   extractor.commands,
		Errors:     errs,
	}
}

func requireNoParseErrors(t *testing.T, result *ParseResult) {
	t.Helper()
	if result == nil {
		t.Fatal("parse result is nil")
	}
	if len(result.Errors) != 0 {
		t.Fatalf("unexpected parse errors: %v", result.Errors)
	}
}

func TestExpressionVisitorFallbackPreservesBooleanSemantics(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   []Condition
	}{
		{
			name:       "explicit and parens",
			expression: `(a=1 AND b=2)`,
			expected: []Condition{
				{Field: "a", Operator: "=", Value: "1", LogicalOp: "AND"},
				{Field: "b", Operator: "=", Value: "2", LogicalOp: "AND"},
			},
		},
		{
			name:       "implicit and",
			expression: `a=1 b=2`,
			expected: []Condition{
				{Field: "a", Operator: "=", Value: "1", LogicalOp: "AND"},
				{Field: "b", Operator: "=", Value: "2", LogicalOp: "AND"},
			},
		},
		{
			name:       "or",
			expression: `a=1 OR b=2`,
			expected: []Condition{
				{Field: "a", Operator: "=", Value: "1", LogicalOp: "AND"},
				{Field: "b", Operator: "=", Value: "2", LogicalOp: "OR"},
			},
		},
		{
			name:       "not nested where",
			expression: `NOT where(b=2)`,
			expected: []Condition{
				{Field: "b", Operator: "=", Value: "2", Negated: true, LogicalOp: "AND"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requireExactConditions(t, extractExpressionForTest(t, tt.expression), tt.expected)
		})
	}
}

func requireExactConditions(t *testing.T, result *ParseResult, expected []Condition) {
	t.Helper()
	requireNoParseErrors(t, result)
	if len(result.Conditions) != len(expected) {
		t.Fatalf("expected %d conditions, got %d: %+v", len(expected), len(result.Conditions), result.Conditions)
	}
	for i, want := range expected {
		got := result.Conditions[i]
		if got.Field != want.Field ||
			got.Operator != want.Operator ||
			got.Value != want.Value ||
			got.Negated != want.Negated ||
			got.LogicalOp != want.LogicalOp {
			t.Fatalf("condition %d mismatch:\n got: %+v\nwant: %+v", i, got, want)
		}
		if strings.Join(got.Alternatives, "\x00") != strings.Join(want.Alternatives, "\x00") {
			t.Fatalf("condition %d alternatives mismatch: got %v want %v", i, got.Alternatives, want.Alternatives)
		}
	}
}

func TestPostfixNegationVariantsPreserveSemantics(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected []Condition
	}{
		{
			name:  "comparison",
			query: `where(status NOT = 404)`,
			expected: []Condition{{
				Field:     "status",
				Operator:  "=",
				Value:     "404",
				Negated:   true,
				LogicalOp: "AND",
			}},
		},
		{
			name:  "multi field comparison",
			query: `where(status,result NOT = 404)`,
			expected: []Condition{
				{Field: "status", Operator: "=", Value: "404", Negated: true, LogicalOp: "AND"},
				{Field: "result", Operator: "=", Value: "404", Negated: true, LogicalOp: "OR"},
			},
		},
		{
			name:  "regex match",
			query: `where(process.name NOT =~ /benign/i)`,
			expected: []Condition{{
				Field:     "process.name",
				Operator:  "=~",
				Value:     "benign",
				Negated:   true,
				LogicalOp: "AND",
			}},
		},
		{
			name:  "multi field regex match",
			query: `where(source,destination NOT =~ /private/i)`,
			expected: []Condition{
				{Field: "source", Operator: "=~", Value: "private", Negated: true, LogicalOp: "AND"},
				{Field: "destination", Operator: "=~", Value: "private", Negated: true, LogicalOp: "OR"},
			},
		},
		{
			name:  "list string",
			query: `where(message NOT CONTAINS-ANY ["error", "fail"])`,
			expected: []Condition{{
				Field:        "message",
				Operator:     "contains_any",
				Value:        "error, fail",
				Negated:      true,
				LogicalOp:    "AND",
				Alternatives: []string{"error", "fail"},
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requireExactConditions(t, ExtractConditions(tt.query), tt.expected)
		})
	}
}

func TestConditionRulePrefixNegationVariantsPreserveSemantics(t *testing.T) {
	tests := []struct {
		name      string
		condition string
		expected  []Condition
	}{
		{
			name:      "comparison",
			condition: `NOT status=404`,
			expected: []Condition{{
				Field:     "status",
				Operator:  "=",
				Value:     "404",
				Negated:   true,
				LogicalOp: "AND",
			}},
		},
		{
			name:      "regex match",
			condition: `NOT process.name =~ /benign/i`,
			expected: []Condition{{
				Field:     "process.name",
				Operator:  "=~",
				Value:     "benign",
				Negated:   true,
				LogicalOp: "AND",
			}},
		},
		{
			name:      "string operator",
			condition: `NOT message CONTAINS error`,
			expected: []Condition{{
				Field:     "message",
				Operator:  "contains",
				Value:     "error",
				Negated:   true,
				LogicalOp: "AND",
			}},
		},
		{
			name:      "set",
			condition: `NOT status IN [200, 404]`,
			expected: []Condition{{
				Field:        "status",
				Operator:     "in",
				Value:        "200, 404",
				Negated:      true,
				LogicalOp:    "AND",
				Alternatives: []string{"200", "404"},
			}},
		},
		{
			name:      "list string",
			condition: `NOT message CONTAINS-ANY ["error", "fail"]`,
			expected: []Condition{{
				Field:        "message",
				Operator:     "contains_any",
				Value:        "error, fail",
				Negated:      true,
				LogicalOp:    "AND",
				Alternatives: []string{"error", "fail"},
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requireExactConditions(t, extractConditionRuleForTest(t, tt.condition), tt.expected)
		})
	}
}

func TestIPSetValuesPreserveCIDRWrappers(t *testing.T) {
	result := ExtractConditions(`where(source_address IN [IP(10.0.0.0/8), IP("192.168.0.0/16")])`)
	requireExactConditions(t, result, []Condition{{
		Field:        "source_address",
		Operator:     "cidr",
		Value:        "10.0.0.0/8, 192.168.0.0/16",
		LogicalOp:    "AND",
		Alternatives: []string{"10.0.0.0/8", "192.168.0.0/16"},
	}})
}

func TestIpSetVisitorBodyWithConstructedContext(t *testing.T) {
	input := antlr.NewInputStream(`source_address,destination_address IN [IP(10.0.0.0/8), IP("192.168.0.0/16")]`)
	lexer := NewLEQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewLEQLParser(stream)
	condition := parser.Condition()

	setCtx, ok := condition.(*SetConditionContext)
	if !ok {
		t.Fatalf("expected generic set context from parser, got %T", condition)
	}

	ipCtx := NewIpSetConditionContext(nil, NewConditionContext(nil, nil, -1))
	ipCtx.AddChild(setCtx.FieldList())
	ipCtx.AddChild(setCtx.ValueList())

	extractor := &conditionExtractor{currentLogicalOp: "AND"}
	extractor.visitIpSetCondition(ipCtx)

	requireExactConditions(t, &ParseResult{Conditions: extractor.conditions}, []Condition{
		{Field: "source_address", Operator: "cidr", Value: "10.0.0.0/8, 192.168.0.0/16", LogicalOp: "AND", Alternatives: []string{"10.0.0.0/8", "192.168.0.0/16"}},
		{Field: "destination_address", Operator: "cidr", Value: "10.0.0.0/8, 192.168.0.0/16", LogicalOp: "OR", Alternatives: []string{"10.0.0.0/8", "192.168.0.0/16"}},
	})
}

func TestExtractorRecoveryAndTimeoutEdges(t *testing.T) {
	oldHook := extractHook
	oldMaxParseTime := MaxParseTime
	defer func() {
		extractHook = oldHook
		MaxParseTime = oldMaxParseTime
	}()

	extractHook = func(string) { panic("forced panic") }
	result := ExtractConditions(`where(status=200)`)
	if result == nil || len(result.Errors) == 0 || !strings.Contains(result.Errors[0], "parser panic: forced panic") {
		t.Fatalf("expected panic recovery error, got %#v", result)
	}

	block := make(chan struct{})
	extractHook = func(string) { <-block }
	MaxParseTime = time.Nanosecond
	result = ExtractConditions(`where(status=200)`)
	close(block)
	if result == nil || len(result.Errors) == 0 || !strings.Contains(result.Errors[0], "parser timeout") {
		t.Fatalf("expected timeout recovery error, got %#v", result)
	}
}

func TestHelperGuardEdges(t *testing.T) {
	if got := extractFieldNames(nil); got != nil {
		t.Fatalf("nil field list returned %v", got)
	}
	if got := extractAllFieldNames(nil); got != nil {
		t.Fatalf("nil all-field list returned %v", got)
	}
	if got := extractComparisonOp(nil); got != "" {
		t.Fatalf("nil comparison op returned %q", got)
	}
	if got := extractStringOp(nil); got != "" {
		t.Fatalf("nil string op returned %q", got)
	}
	if got := extractSetOp(nil); got != "" {
		t.Fatalf("nil set op returned %q", got)
	}
	if got := extractListStringOp(nil); got != "" {
		t.Fatalf("nil list string op returned %q", got)
	}
	if got := extractValue(nil); got != "" {
		t.Fatalf("nil value returned %q", got)
	}
	if got := extractValueList(nil); got != nil {
		t.Fatalf("nil value list returned %v", got)
	}
	if got := extractValueListWithIP(nil); got != nil {
		t.Fatalf("nil IP value list returned %v", got)
	}
	if got := extractIPValue(nil); got != "" {
		t.Fatalf("nil IP value returned %q", got)
	}
	if got := stripQuotes("bare"); got != "bare" {
		t.Fatalf("bare string stripped to %q", got)
	}
	if got := stripTripleQuotes("short"); got != "short" {
		t.Fatalf("short triple-quote input stripped to %q", got)
	}
	if got := stripRegexSlashes("literal"); got != "literal" {
		t.Fatalf("non-regex input stripped to %q", got)
	}
	if got := stripRegexSlashes("/"); got != "/" {
		t.Fatalf("single slash stripped to %q", got)
	}
	if got := stripRegexSlashes("/unterminated"); got != "/unterminated" {
		t.Fatalf("unterminated regex stripped to %q", got)
	}
	if got := ClassifyFieldProvenance(nil, "anything"); got != ProvenanceMain {
		t.Fatalf("LEQL provenance should always be main, got %q", got)
	}
	if boolKey(false) != "0" || boolKey(true) != "1" {
		t.Fatalf("boolKey returned unexpected values")
	}
}

func expressionContextForTest(t *testing.T, expression string) IExpressionContext {
	t.Helper()

	input := antlr.NewInputStream(expression)
	lexer := NewLEQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewLEQLParser(stream)

	return parser.Expression()
}

func TestExpressionBuilderHelperEdges(t *testing.T) {
	extractor := &conditionExtractor{currentLogicalOp: "AND"}

	if got := extractor.buildExpression(nil); got != nil {
		t.Fatalf("nil expression tree returned %#v", got)
	}
	if got := extractor.buildBinaryExpression(nil, BooleanAnd); got != nil {
		t.Fatalf("empty binary expression returned %#v", got)
	}

	expr := expressionContextForTest(t, `status=200`)
	if got := extractor.buildBinaryExpression([]IExpressionContext{expr}, BooleanAnd); got == nil || got.Operator != BooleanCondition {
		t.Fatalf("single binary expression did not collapse to condition, got %#v", got)
	}

	if got := invertExpression(nil); got != nil {
		t.Fatalf("nil inverted expression returned %#v", got)
	}

	leaf := conditionIndexExpression(0)
	if got := invertExpressionChildren(BooleanOr, []*ConditionExpression{leaf}); got != leaf {
		t.Fatalf("single inverted child was not returned directly, got %#v", got)
	}

	if regexValueHasCaseInsensitiveFlag(`plain`) {
		t.Fatalf("plain literal reported a regex flag")
	}
	if regexValueHasCaseInsensitiveFlag(`/cmd.*/`) {
		t.Fatalf("regex without flags reported case-insensitive flag")
	}
	if !regexValueHasCaseInsensitiveFlag(`/cmd.*/i`) {
		t.Fatalf("regex with i flag did not report case-insensitive flag")
	}
}
