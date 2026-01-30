package leql

import (
	"testing"
)

// --- Helper ---

func assertConditions(t *testing.T, result *ParseResult, expected []Condition) {
	t.Helper()
	if len(result.Errors) > 0 {
		t.Logf("Parse errors: %v", result.Errors)
	}
	if len(result.Conditions) != len(expected) {
		t.Fatalf("Expected %d conditions, got %d: %+v", len(expected), len(result.Conditions), result.Conditions)
	}
	for i, exp := range expected {
		got := result.Conditions[i]
		if got.Field != exp.Field {
			t.Errorf("Condition %d: expected field %q, got %q", i, exp.Field, got.Field)
		}
		if got.Operator != exp.Operator {
			t.Errorf("Condition %d: expected operator %q, got %q", i, exp.Operator, got.Operator)
		}
		if got.Value != exp.Value {
			t.Errorf("Condition %d: expected value %q, got %q", i, exp.Value, got.Value)
		}
		if got.Negated != exp.Negated {
			t.Errorf("Condition %d: expected negated=%v, got %v", i, exp.Negated, got.Negated)
		}
		if exp.LogicalOp != "" && got.LogicalOp != exp.LogicalOp {
			t.Errorf("Condition %d: expected logicalOp %q, got %q", i, exp.LogicalOp, got.LogicalOp)
		}
		if len(exp.Alternatives) > 0 {
			if len(got.Alternatives) != len(exp.Alternatives) {
				t.Errorf("Condition %d: expected %d alternatives, got %d (%v)", i, len(exp.Alternatives), len(got.Alternatives), got.Alternatives)
			}
		}
	}
}

// --- 1. Simple Comparison ---

func TestExtractConditions_SimpleComparison(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected []Condition
	}{
		{
			name:  "equality",
			query: `where(field=value)`,
			expected: []Condition{
				{Field: "field", Operator: "=", Value: "value", LogicalOp: "AND"},
			},
		},
		{
			name:  "not equal",
			query: `where(status!=404)`,
			expected: []Condition{
				{Field: "status", Operator: "!=", Value: "404", LogicalOp: "AND"},
			},
		},
		{
			name:  "greater than",
			query: `where(response_time>25)`,
			expected: []Condition{
				{Field: "response_time", Operator: ">", Value: "25", LogicalOp: "AND"},
			},
		},
		{
			name:  "greater than or equal",
			query: `where(hits>=10)`,
			expected: []Condition{
				{Field: "hits", Operator: ">=", Value: "10", LogicalOp: "AND"},
			},
		},
		{
			name:  "less than",
			query: `where(level<3)`,
			expected: []Condition{
				{Field: "level", Operator: "<", Value: "3", LogicalOp: "AND"},
			},
		},
		{
			name:  "less than or equal",
			query: `where(age<=65)`,
			expected: []Condition{
				{Field: "age", Operator: "<=", Value: "65", LogicalOp: "AND"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractConditions(tt.query)
			assertConditions(t, result, tt.expected)
		})
	}
}

// --- 2. String Operators ---

func TestExtractConditions_StringOperators(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected []Condition
	}{
		{
			name:  "CONTAINS",
			query: `where(message CONTAINS error)`,
			expected: []Condition{
				{Field: "message", Operator: "contains", Value: "error", LogicalOp: "AND"},
			},
		},
		{
			name:  "ICONTAINS",
			query: `where(user ICONTAINS admin)`,
			expected: []Condition{
				{Field: "user", Operator: "icontains", Value: "admin", LogicalOp: "AND"},
			},
		},
		{
			name:  "STARTS-WITH quoted value",
			query: `where(path STARTS-WITH "/var")`,
			expected: []Condition{
				{Field: "path", Operator: "startswith", Value: "/var", LogicalOp: "AND"},
			},
		},
		{
			name:  "ISTARTS-WITH",
			query: `where(host ISTARTS-WITH web)`,
			expected: []Condition{
				{Field: "host", Operator: "istartswith", Value: "web", LogicalOp: "AND"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractConditions(tt.query)
			assertConditions(t, result, tt.expected)
		})
	}
}

// --- 3. Set Operators ---

func TestExtractConditions_SetOperators(t *testing.T) {
	tests := []struct {
		name             string
		query            string
		expectedOperator string
		expectedAltCount int
	}{
		{
			name:             "IN with 3 values",
			query:            `where(status IN [200, 301, 404])`,
			expectedOperator: "in",
			expectedAltCount: 3,
		},
		{
			name:             "IIN with 2 values",
			query:            `where(user IIN ["Admin", "root"])`,
			expectedOperator: "iin",
			expectedAltCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractConditions(tt.query)
			if len(result.Errors) > 0 {
				t.Logf("Parse errors: %v", result.Errors)
			}
			if len(result.Conditions) != 1 {
				t.Fatalf("Expected 1 condition, got %d: %+v", len(result.Conditions), result.Conditions)
			}
			got := result.Conditions[0]
			if got.Operator != tt.expectedOperator {
				t.Errorf("Expected operator %q, got %q", tt.expectedOperator, got.Operator)
			}
			if len(got.Alternatives) != tt.expectedAltCount {
				t.Errorf("Expected %d alternatives, got %d: %v", tt.expectedAltCount, len(got.Alternatives), got.Alternatives)
			}
		})
	}
}

// --- 4. List String Operators ---

func TestExtractConditions_ListStringOperators(t *testing.T) {
	tests := []struct {
		name             string
		query            string
		expectedOperator string
		expectedAltCount int
	}{
		{
			name:             "CONTAINS-ANY",
			query:            `where(message CONTAINS-ANY ["error", "fail"])`,
			expectedOperator: "contains_any",
			expectedAltCount: 2,
		},
		{
			name:             "STARTS-WITH-ANY",
			query:            `where(path STARTS-WITH-ANY ["/var", "/tmp"])`,
			expectedOperator: "startswith_any",
			expectedAltCount: 2,
		},
		{
			name:             "CONTAINS-ALL",
			query:            `where(log CONTAINS-ALL ["error", "critical"])`,
			expectedOperator: "contains_all",
			expectedAltCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractConditions(tt.query)
			if len(result.Errors) > 0 {
				t.Logf("Parse errors: %v", result.Errors)
			}
			if len(result.Conditions) != 1 {
				t.Fatalf("Expected 1 condition, got %d: %+v", len(result.Conditions), result.Conditions)
			}
			got := result.Conditions[0]
			if got.Operator != tt.expectedOperator {
				t.Errorf("Expected operator %q, got %q", tt.expectedOperator, got.Operator)
			}
			if len(got.Alternatives) != tt.expectedAltCount {
				t.Errorf("Expected %d alternatives, got %d: %v", tt.expectedAltCount, len(got.Alternatives), got.Alternatives)
			}
		})
	}
}

// --- 5. Logical Operators ---

func TestExtractConditions_LogicalOperators(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected []Condition
	}{
		{
			name:  "AND two conditions",
			query: `where(a=1 AND b=2)`,
			expected: []Condition{
				{Field: "a", Operator: "=", Value: "1", LogicalOp: "AND"},
				{Field: "b", Operator: "=", Value: "2", LogicalOp: "AND"},
			},
		},
		{
			name:  "OR two conditions",
			query: `where(a=1 OR b=2)`,
			expected: []Condition{
				{Field: "a", Operator: "=", Value: "1", LogicalOp: "AND"},
				{Field: "b", Operator: "=", Value: "2", LogicalOp: "OR"},
			},
		},
		{
			name:  "NOT prefix",
			query: `where(NOT status=404)`,
			expected: []Condition{
				{Field: "status", Operator: "=", Value: "404", Negated: true, LogicalOp: "AND"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractConditions(tt.query)
			assertConditions(t, result, tt.expected)
		})
	}
}

// --- 6. Negation ---

func TestExtractConditions_Negation(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected []Condition
	}{
		{
			name:  "NOT IN",
			query: `where(NOT key IN ["aaa", "bbb"])`,
			expected: []Condition{
				{Field: "key", Operator: "in", Value: "aaa, bbb", Negated: true, LogicalOp: "AND"},
			},
		},
		{
			name:  "NOT CONTAINS",
			query: `where(NOT key CONTAINS error)`,
			expected: []Condition{
				{Field: "key", Operator: "contains", Value: "error", Negated: true, LogicalOp: "AND"},
			},
		},
		{
			name:  "not equal operator",
			query: `where(key!=value)`,
			expected: []Condition{
				{Field: "key", Operator: "!=", Value: "value", LogicalOp: "AND"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractConditions(tt.query)
			assertConditions(t, result, tt.expected)
		})
	}
}

// --- 7. Quoted Values ---

func TestExtractConditions_QuotedValues(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected []Condition
	}{
		{
			name:  "double quoted value",
			query: `where(name="John Doe")`,
			expected: []Condition{
				{Field: "name", Operator: "=", Value: "John Doe", LogicalOp: "AND"},
			},
		},
		{
			name:  "single quoted value",
			query: `where(name='John Doe')`,
			expected: []Condition{
				{Field: "name", Operator: "=", Value: "John Doe", LogicalOp: "AND"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractConditions(tt.query)
			assertConditions(t, result, tt.expected)
		})
	}
}

// --- 8. Regex ---

func TestExtractConditions_Regex(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected []Condition
	}{
		{
			name:  "regex keyword search with flags",
			query: `where(/error/i)`,
			expected: []Condition{
				{Field: "", Operator: "regex", Value: "error", LogicalOp: "AND"},
			},
		},
		{
			name:  "field regex comparison",
			query: `where(field=/pattern/)`,
			expected: []Condition{
				{Field: "field", Operator: "=", Value: "pattern", LogicalOp: "AND"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractConditions(tt.query)
			assertConditions(t, result, tt.expected)
		})
	}
}

// --- 9. Nested Keys ---

func TestExtractConditions_NestedKeys(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected []Condition
	}{
		{
			name:  "dotted JSON field",
			query: `where(source_json.Workload=SharePoint)`,
			expected: []Condition{
				{Field: "source_json.Workload", Operator: "=", Value: "SharePoint", LogicalOp: "AND"},
			},
		},
		{
			name:  "dotted process field",
			query: `where(process.name=powershell.exe)`,
			expected: []Condition{
				{Field: "process.name", Operator: "=", Value: "powershell.exe", LogicalOp: "AND"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractConditions(tt.query)
			assertConditions(t, result, tt.expected)
		})
	}
}

// --- 10. Keyword Search ---

func TestExtractConditions_KeywordSearch(t *testing.T) {
	result := ExtractConditions(`where("Amazon Boardman")`)
	if len(result.Errors) > 0 {
		t.Logf("Parse errors: %v", result.Errors)
	}
	if len(result.Conditions) != 1 {
		t.Fatalf("Expected 1 condition, got %d: %+v", len(result.Conditions), result.Conditions)
	}
	got := result.Conditions[0]
	if got.Operator != "keyword" {
		t.Errorf("Expected operator %q, got %q", "keyword", got.Operator)
	}
	if got.Value != "Amazon Boardman" {
		t.Errorf("Expected value %q, got %q", "Amazon Boardman", got.Value)
	}
}

// --- 11. IsStatisticalQuery ---

func TestIsStatisticalQuery(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected bool
	}{
		{
			name:     "where + calculate is statistical",
			query:    `where(status=200) calculate(count)`,
			expected: true,
		},
		{
			name:     "where + groupby + calculate is statistical",
			query:    `where(status=200) groupby(user) calculate(sum:bytes)`,
			expected: true,
		},
		{
			name:     "where only is not statistical",
			query:    `where(status=200)`,
			expected: false,
		},
		{
			name:     "where + groupby without calculate is not statistical",
			query:    `where(status=200) groupby(user)`,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractConditions(tt.query)
			got := IsStatisticalQuery(result)
			if got != tt.expected {
				t.Errorf("IsStatisticalQuery() = %v, expected %v (commands: %v)", got, tt.expected, result.Commands)
			}
		})
	}
}

// --- 12. GetEventTypeFromConditions ---

func TestGetEventTypeFromConditions(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected string
	}{
		{
			name:     "EventCode 4688",
			query:    `where(EventCode=4688)`,
			expected: "windows_4688",
		},
		{
			name:     "EventCode 4624",
			query:    `where(EventCode=4624)`,
			expected: "windows_4624",
		},
		{
			name:     "EventID 4625",
			query:    `where(EventID=4625)`,
			expected: "windows_4625",
		},
		{
			name:     "non-event field",
			query:    `where(status=200)`,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractConditions(tt.query)
			got := GetEventTypeFromConditions(result)
			if got != tt.expected {
				t.Errorf("GetEventTypeFromConditions() = %q, expected %q", got, tt.expected)
			}
		})
	}
}

// --- 13. DeduplicateConditions ---

func TestDeduplicateConditions(t *testing.T) {
	conditions := []Condition{
		{Field: "status", Operator: "=", Value: "200"},
		{Field: "user", Operator: "=", Value: "admin"},
		{Field: "status", Operator: "=", Value: "200"}, // duplicate
	}

	result := DeduplicateConditions(conditions)
	if len(result) != 2 {
		t.Errorf("Expected 2 conditions after dedup, got %d: %+v", len(result), result)
	}
}

// --- 14. ClassifyFieldProvenance ---

func TestClassifyFieldProvenance(t *testing.T) {
	result := &ParseResult{
		Conditions: []Condition{
			{Field: "status", Operator: "=", Value: "200"},
		},
	}
	prov := ClassifyFieldProvenance(result, "status")
	if prov != ProvenanceMain {
		t.Errorf("Expected ProvenanceMain, got %q", prov)
	}
}

// --- 15. HasUnmappedComputedFields ---

func TestHasUnmappedComputedFields(t *testing.T) {
	result := ExtractConditions(`where(field=value)`)
	if HasUnmappedComputedFields(result) {
		t.Error("Expected HasUnmappedComputedFields to return false for LEQL")
	}
}

// --- 16. HasComplexWhereConditions ---

func TestHasComplexWhereConditions(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected bool
	}{
		{
			name:     "regex is complex",
			query:    `where(/pattern/)`,
			expected: true,
		},
		{
			name:     "simple equality is not complex",
			query:    `where(field=value)`,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractConditions(tt.query)
			got := HasComplexWhereConditions(result)
			if got != tt.expected {
				t.Errorf("HasComplexWhereConditions() = %v, expected %v (conditions: %+v)", got, tt.expected, result.Conditions)
			}
		})
	}
}

// --- 17. ParseResult Commands ---

func TestParseResult_Commands(t *testing.T) {
	result := ExtractConditions(`where(x=1) groupby(y) calculate(count)`)
	if len(result.Errors) > 0 {
		t.Logf("Parse errors: %v", result.Errors)
	}

	expectedCmds := []string{"where", "groupby", "calculate"}
	if len(result.Commands) != len(expectedCmds) {
		t.Fatalf("Expected commands %v, got %v", expectedCmds, result.Commands)
	}
	for i, exp := range expectedCmds {
		if result.Commands[i] != exp {
			t.Errorf("Command %d: expected %q, got %q", i, exp, result.Commands[i])
		}
	}
}

// --- 18. ParseResult GroupByFields ---

func TestParseResult_GroupByFields(t *testing.T) {
	result := ExtractConditions(`groupby(user, status) calculate(count)`)
	if len(result.Errors) > 0 {
		t.Logf("Parse errors: %v", result.Errors)
	}

	expectedFields := []string{"user", "status"}
	if len(result.GroupByFields) != len(expectedFields) {
		t.Fatalf("Expected group-by fields %v, got %v", expectedFields, result.GroupByFields)
	}
	for i, exp := range expectedFields {
		if result.GroupByFields[i] != exp {
			t.Errorf("GroupByField %d: expected %q, got %q", i, exp, result.GroupByFields[i])
		}
	}
}
