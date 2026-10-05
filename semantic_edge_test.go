package leql

import (
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

type wrappedFieldListContext struct{ *FieldListContext }
type wrappedAllFieldListContext struct{ *AllFieldListContext }
type wrappedValueContext struct{ *ValueContext }
type wrappedValueListContext struct{ *ValueListContext }
type wrappedValueListItemContext struct{ *ValueListItemContext }
type wrappedIPValueContext struct{ *IpValueContext }

func parseLEQLValueListForTest(t *testing.T, text string) IValueListContext {
	t.Helper()
	lexer := NewLEQLLexer(antlr.NewInputStream(text))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewLEQLParser(stream)
	return parser.ValueList()
}

func TestMultiFieldCompoundConditionsAreOrLinked(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected []Condition
	}{
		{
			name:  "regex",
			query: `where(source,destination =~ /private/)`,
			expected: []Condition{
				{Field: "source", Operator: "=~", Value: "private", LogicalOp: "AND"},
				{Field: "destination", Operator: "=~", Value: "private", LogicalOp: "OR"},
			},
		},
		{
			name:  "string operator",
			query: `where(message,detail ENDS-WITH ".exe")`,
			expected: []Condition{
				{Field: "message", Operator: "endswith", Value: ".exe", LogicalOp: "AND"},
				{Field: "detail", Operator: "endswith", Value: ".exe", LogicalOp: "OR"},
			},
		},
		{
			name:  "set",
			query: `where(status,result IN [200, 201])`,
			expected: []Condition{
				{Field: "status", Operator: "in", Value: "200, 201", LogicalOp: "AND", Alternatives: []string{"200", "201"}},
				{Field: "result", Operator: "in", Value: "200, 201", LogicalOp: "OR", Alternatives: []string{"200", "201"}},
			},
		},
		{
			name:  "list string",
			query: `where(message,detail CONTAINS-ANY ["error", "fail"])`,
			expected: []Condition{
				{Field: "message", Operator: "contains_any", Value: "error, fail", LogicalOp: "AND", Alternatives: []string{"error", "fail"}},
				{Field: "detail", Operator: "contains_any", Value: "error, fail", LogicalOp: "OR", Alternatives: []string{"error", "fail"}},
			},
		},
		{
			name:  "postfix set",
			query: `where(status,result NOT IN [200, 201])`,
			expected: []Condition{
				{Field: "status", Operator: "in", Value: "200, 201", Negated: true, LogicalOp: "AND", Alternatives: []string{"200", "201"}},
				{Field: "result", Operator: "in", Value: "200, 201", Negated: true, LogicalOp: "OR", Alternatives: []string{"200", "201"}},
			},
		},
		{
			name:  "postfix list string",
			query: `where(message,detail NOT CONTAINS-ANY ["debug", "trace"])`,
			expected: []Condition{
				{Field: "message", Operator: "contains_any", Value: "debug, trace", Negated: true, LogicalOp: "AND", Alternatives: []string{"debug", "trace"}},
				{Field: "detail", Operator: "contains_any", Value: "debug, trace", Negated: true, LogicalOp: "OR", Alternatives: []string{"debug", "trace"}},
			},
		},
		{
			name:  "ip set",
			query: `where(source_address,destination_address IN [IP(10.0.0.0/8), IP(192.168.0.0/16)])`,
			expected: []Condition{
				{Field: "source_address", Operator: "cidr", Value: "10.0.0.0/8, 192.168.0.0/16", LogicalOp: "AND", Alternatives: []string{"10.0.0.0/8", "192.168.0.0/16"}},
				{Field: "destination_address", Operator: "cidr", Value: "10.0.0.0/8, 192.168.0.0/16", LogicalOp: "OR", Alternatives: []string{"10.0.0.0/8", "192.168.0.0/16"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requireExactConditions(t, ExtractConditions(tt.query), tt.expected)
		})
	}
}

func TestConditionRuleMultiFieldPrefixNegationIsOrLinked(t *testing.T) {
	tests := []struct {
		name      string
		condition string
		expected  []Condition
	}{
		{
			name:      "regex",
			condition: `NOT source,destination =~ /private/`,
			expected: []Condition{
				{Field: "source", Operator: "=~", Value: "private", Negated: true, LogicalOp: "AND"},
				{Field: "destination", Operator: "=~", Value: "private", Negated: true, LogicalOp: "OR"},
			},
		},
		{
			name:      "comparison",
			condition: `NOT status,result=404`,
			expected: []Condition{
				{Field: "status", Operator: "=", Value: "404", Negated: true, LogicalOp: "AND"},
				{Field: "result", Operator: "=", Value: "404", Negated: true, LogicalOp: "OR"},
			},
		},
		{
			name:      "string",
			condition: `NOT message,detail CONTAINS error`,
			expected: []Condition{
				{Field: "message", Operator: "contains", Value: "error", Negated: true, LogicalOp: "AND"},
				{Field: "detail", Operator: "contains", Value: "error", Negated: true, LogicalOp: "OR"},
			},
		},
		{
			name:      "set",
			condition: `NOT status,result IN [200, 404]`,
			expected: []Condition{
				{Field: "status", Operator: "in", Value: "200, 404", Negated: true, LogicalOp: "AND", Alternatives: []string{"200", "404"}},
				{Field: "result", Operator: "in", Value: "200, 404", Negated: true, LogicalOp: "OR", Alternatives: []string{"200", "404"}},
			},
		},
		{
			name:      "list string",
			condition: `NOT message,detail CONTAINS-ANY ["error", "fail"]`,
			expected: []Condition{
				{Field: "message", Operator: "contains_any", Value: "error, fail", Negated: true, LogicalOp: "AND", Alternatives: []string{"error", "fail"}},
				{Field: "detail", Operator: "contains_any", Value: "error, fail", Negated: true, LogicalOp: "OR", Alternatives: []string{"error", "fail"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requireExactConditions(t, extractConditionRuleForTest(t, tt.condition), tt.expected)
		})
	}
}

func TestClauseCommandsKeywordAndParserDiagnostics(t *testing.T) {
	result := ExtractConditions(`select(hostname) where(status=200) groupby(user) calculate(count) having(count>5) sort(desc) limit(10) timeslice(30m)`)
	requireNoParseErrors(t, result)
	expectedCommands := []string{"select", "where", "groupby", "calculate", "having", "sort", "limit", "timeslice"}
	if strings.Join(result.Commands, ",") != strings.Join(expectedCommands, ",") {
		t.Fatalf("commands mismatch: got %v want %v", result.Commands, expectedCommands)
	}

	requireExactConditions(t, ExtractConditions(`where('single quoted keyword')`), []Condition{{
		Operator:  "keyword",
		Value:     "single quoted keyword",
		LogicalOp: "AND",
	}})
	requireExactConditions(t, ExtractConditions(`where(username)`), []Condition{{
		Field:     "username",
		Operator:  "exists",
		Value:     "*",
		LogicalOp: "AND",
	}})
	if result := ExtractConditions(`where(status=)`); result == nil || len(result.Errors) == 0 {
		t.Fatalf("expected syntax diagnostics for malformed query, got %#v", result)
	}

	listener := &errorListener{}
	listener.SyntaxError(nil, nil, 1, 1, "manual diagnostic", nil)
	if len(listener.errors) != 1 || listener.errors[0] != "manual diagnostic" {
		t.Fatalf("manual syntax diagnostic was not recorded: %v", listener.errors)
	}
}

func TestAdditionalHelperAndVisitorGuardEdges(t *testing.T) {
	if got := DeduplicateConditions(nil); got != nil {
		t.Fatalf("nil conditions should dedupe to nil, got %v", got)
	}

	extractor := &conditionExtractor{currentLogicalOp: "AND"}
	parent := NewConditionContext(nil, nil, -1)
	extractor.visit(nil)
	extractor.visit(NewIpSetConditionContext(nil, parent))
	extractor.visitComparisonCondition(NewComparisonConditionContext(nil, parent))
	extractor.visitRegexMatchCondition(NewRegexMatchConditionContext(nil, parent))
	extractor.visitNegatedRegexMatchCondition(NewNegatedRegexMatchConditionContext(nil, parent))
	extractor.visitPostfixNegatedRegexMatchCondition(NewPostfixNegatedRegexMatchConditionContext(nil, parent))
	extractor.visitStringCondition(NewStringConditionContext(nil, parent))
	extractor.visitSetCondition(NewSetConditionContext(nil, parent))
	extractor.visitListStringCondition(NewListStringConditionContext(nil, parent))
	extractor.visitNegatedComparisonCondition(NewNegatedComparisonConditionContext(nil, parent))
	extractor.visitNegatedStringCondition(NewNegatedStringConditionContext(nil, parent))
	extractor.visitNegatedSetCondition(NewNegatedSetConditionContext(nil, parent))
	extractor.visitNegatedListStringCondition(NewNegatedListStringConditionContext(nil, parent))
	extractor.visitPostfixNegatedComparisonCondition(NewPostfixNegatedComparisonConditionContext(nil, parent))
	extractor.visitPostfixNegatedStringCondition(NewPostfixNegatedStringConditionContext(nil, parent))
	extractor.visitPostfixNegatedSetCondition(NewPostfixNegatedSetConditionContext(nil, parent))
	extractor.visitPostfixNegatedListStringCondition(NewPostfixNegatedListStringConditionContext(nil, parent))
	extractor.visitNocaseCondition(NewNocaseConditionContext(nil, parent))
	extractor.visitIpCondition(NewIpConditionContext(nil, parent))
	extractor.visitIpSetCondition(NewIpSetConditionContext(nil, parent))
	extractor.visitAllFieldsComparisonCondition(NewAllFieldsComparisonConditionContext(nil, parent))
	extractor.visitAllFieldsSetCondition(NewAllFieldsSetConditionContext(nil, parent))
	extractor.visitAllFieldsNocaseCondition(NewAllFieldsNocaseConditionContext(nil, parent))
	extractor.visitFieldExistsCondition(NewFieldExistsConditionContext(nil, parent))
	extractor.visitNumberKeywordCondition(NewNumberKeywordConditionContext(nil, parent))
	extractor.visitVariableCondition(NewVariableConditionContext(nil, parent))

	if len(extractor.conditions) != 0 {
		t.Fatalf("nil-guard visitor calls should not emit conditions, got %+v", extractor.conditions)
	}
}

func TestClauseVisitorMethodsAppendCommands(t *testing.T) {
	extractor := &conditionExtractor{}
	extractor.visitCalculateClause(nil)
	extractor.visitSelectClause(nil)
	extractor.visitSortClause(nil)
	extractor.visitLimitClause(nil)
	extractor.visitTimesliceClause(nil)
	extractor.visitHavingClause(nil)

	expected := []string{"calculate", "select", "sort", "limit", "timeslice", "having"}
	if strings.Join(extractor.commands, ",") != strings.Join(expected, ",") {
		t.Fatalf("commands mismatch: got %v want %v", extractor.commands, expected)
	}

	listener := &errorListener{}
	listener.SyntaxError(nil, nil, 0, 0, "direct diagnostic", nil)
	if len(listener.errors) != 1 || listener.errors[0] != "direct diagnostic" {
		t.Fatalf("syntax error listener did not record diagnostic: %v", listener.errors)
	}
}

func TestValueAndUtilityEdgeBranches(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected []Condition
	}{
		{
			name:  "triple single quoted value",
			query: `where(message='''multi word''')`,
			expected: []Condition{{
				Field:     "message",
				Operator:  "=",
				Value:     "multi word",
				LogicalOp: "AND",
			}},
		},
		{
			name:  "triple double quoted value",
			query: `where(message="""multi word""")`,
			expected: []Condition{{
				Field:     "message",
				Operator:  "=",
				Value:     "multi word",
				LogicalOp: "AND",
			}},
		},
		{
			name:  "case insensitive ends with",
			query: `where(path IENDS-WITH ".EXE")`,
			expected: []Condition{{
				Field:     "path",
				Operator:  "iendswith",
				Value:     ".EXE",
				LogicalOp: "AND",
			}},
		},
		{
			name:  "ends with any",
			query: `where(path ENDS-WITH-ANY [".exe", ".dll"])`,
			expected: []Condition{{
				Field:        "path",
				Operator:     "endswith_any",
				Value:        ".exe, .dll",
				LogicalOp:    "AND",
				Alternatives: []string{".exe", ".dll"},
			}},
		},
		{
			name:  "case insensitive ends with any",
			query: `where(path IENDS-WITH-ANY [".EXE", ".DLL"])`,
			expected: []Condition{{
				Field:        "path",
				Operator:     "iendswith_any",
				Value:        ".EXE, .DLL",
				LogicalOp:    "AND",
				Alternatives: []string{".EXE", ".DLL"},
			}},
		},
		{
			name:  "time unit value",
			query: `where(duration=30m)`,
			expected: []Condition{{
				Field:     "duration",
				Operator:  "=",
				Value:     "30m",
				LogicalOp: "AND",
			}},
		},
		{
			name:  "variable value",
			query: `where(user=${target_user})`,
			expected: []Condition{{
				Field:     "user",
				Operator:  "=",
				Value:     "${target_user}",
				LogicalOp: "AND",
			}},
		},
		{
			name:  "IP value wrapper",
			query: `where(source_address=IP("10.0.0.1"))`,
			expected: []Condition{{
				Field:     "source_address",
				Operator:  "cidr",
				Value:     "10.0.0.1",
				LogicalOp: "AND",
			}},
		},
		{
			name:  "multi field IP value wrapper",
			query: `where(source_address,destination_address=IP("10.0.0.1"))`,
			expected: []Condition{
				{Field: "source_address", Operator: "cidr", Value: "10.0.0.1", LogicalOp: "AND"},
				{Field: "destination_address", Operator: "cidr", Value: "10.0.0.1", LogicalOp: "OR"},
			},
		},
		{
			name:  "field list nocase",
			query: `where(user,actor=NOCASE(admin))`,
			expected: []Condition{
				{Field: "user", Operator: "=", Value: "admin", LogicalOp: "AND"},
				{Field: "actor", Operator: "=", Value: "admin", LogicalOp: "OR"},
			},
		},
		{
			name:  "ALL nocase",
			query: `where(ALL(user,actor)=NOCASE(admin))`,
			expected: []Condition{
				{Field: "user", Operator: "=", Value: "admin", LogicalOp: "AND"},
				{Field: "actor", Operator: "=", Value: "admin", LogicalOp: "AND"},
			},
		},
		{
			name:  "ALL set",
			query: `where(ALL(user,actor) IN [admin, root])`,
			expected: []Condition{
				{Field: "user", Operator: "in", Value: "admin, root", LogicalOp: "AND", Alternatives: []string{"admin", "root"}},
				{Field: "actor", Operator: "in", Value: "admin, root", LogicalOp: "AND", Alternatives: []string{"admin", "root"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requireExactConditions(t, ExtractConditions(tt.query), tt.expected)
		})
	}

	if got := extractIPValue(NewEmptyIpValueContext()); got != "" {
		t.Fatalf("empty IP context returned %q", got)
	}
	if got := extractFieldNames(&wrappedFieldListContext{FieldListContext: NewEmptyFieldListContext()}); got != nil {
		t.Fatalf("wrapped field list returned %v", got)
	}
	if got := extractAllFieldNames(&wrappedAllFieldListContext{AllFieldListContext: NewEmptyAllFieldListContext()}); got != nil {
		t.Fatalf("wrapped all-field list returned %v", got)
	}
	if got := extractStringOp(NewEmptyStringOpContext()); got != "" {
		t.Fatalf("empty string op returned %q", got)
	}
	if got := extractListStringOp(NewEmptyListStringOpContext()); got != "" {
		t.Fatalf("empty list string op returned %q", got)
	}
	if got := extractValue(&wrappedValueContext{ValueContext: NewEmptyValueContext()}); got != "" {
		t.Fatalf("wrapped value context returned %q", got)
	}
	if got := extractValueList(&wrappedValueListContext{ValueListContext: NewEmptyValueListContext()}); got != nil {
		t.Fatalf("wrapped value list returned %v", got)
	}
	valueList := NewEmptyValueListContext()
	valueList.AddChild(&wrappedValueListItemContext{ValueListItemContext: NewEmptyValueListItemContext()})
	if got := extractValueList(valueList); len(got) != 0 {
		t.Fatalf("wrapped value list item returned %v", got)
	}
	if valueListItemsAllIP(nil) {
		t.Fatal("nil value list reported all IP values")
	}
	if valueListItemsAllIP(&wrappedValueListContext{ValueListContext: NewEmptyValueListContext()}) {
		t.Fatal("wrapped value list reported all IP values")
	}
	if valueListItemsAllIP(NewEmptyValueListContext()) {
		t.Fatal("empty value list reported all IP values")
	}
	mixed := NewEmptyValueListContext()
	mixed.AddChild(&wrappedValueListItemContext{ValueListItemContext: NewEmptyValueListItemContext()})
	if valueListItemsAllIP(mixed) {
		t.Fatal("mixed/empty value list item reported all IP values")
	}
	allIP := parseLEQLValueListForTest(t, `[IP(10.0.0.0/8), IP("192.168.0.0/16")]`)
	if !valueListItemsAllIP(allIP) {
		t.Fatal("IP-only value list was not recognized")
	}
	if got := extractIPValue(&wrappedIPValueContext{IpValueContext: NewEmptyIpValueContext()}); got != "" {
		t.Fatalf("wrapped IP context returned %q", got)
	}
	if got := HasUnmappedComputedFields(&ParseResult{}); got {
		t.Fatal("LEQL should not report unmapped computed fields")
	}
	if got := IsStatisticalQuery(nil); got {
		t.Fatal("nil result should not be statistical")
	}
	if got := HasComplexWhereConditions(nil); got {
		t.Fatal("nil result should not be complex")
	}
	if got := GetEventTypeFromConditions(nil); got != "" {
		t.Fatalf("nil result event type = %q", got)
	}

	for _, tt := range []struct {
		query    string
		expected string
	}{
		{`where(EventCode=4720)`, "windows_4720"},
		{`where(EventCode=1)`, "sysmon_1"},
		{`where(EventCode=3)`, "sysmon_3"},
		{`where(EventCode=9999)`, ""},
	} {
		if got := GetEventTypeFromConditions(ExtractConditions(tt.query)); got != tt.expected {
			t.Fatalf("%s event type = %q, want %q", tt.query, got, tt.expected)
		}
	}
}
