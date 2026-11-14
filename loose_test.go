package leql

import (
	"context"
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

// These tests catch lost WHERE provenance and accidental operator/value lowering.
// JSON decoding keeps the regression executable against the pre-metadata API.
func TestLoosePreservesWhereScopeAcrossAPIs(t *testing.T) {
	tests := []struct {
		query string
		flags []bool
	}{
		{`where(user=admin, loose)`, []bool{true}},
		{`where(admin, LOOSE)`, []bool{true}},
		{`where("Admin Phrase", loose)`, []bool{true}},
		{`where(user,actor=Admin, loose)`, []bool{true, true}},
		{`where(all(user,actor)=Admin, loose)`, []bool{true, true}},
		{`where(NOT (user=Admin OR actor=Root), loose)`, []bool{true, true}},
		{`from(logset=Admin) where(user=admin OR NOT (actor=Root AND n>42), loose)`, []bool{false, true, true, true}},
		{`from(user=admin) where(user=admin, loose)`, []bool{false, true}},
		{`where(where(user=Admin) OR actor=Root, loose)`, []bool{true, true}},
	}
	for _, predicate := range []string{
		`n=42`, `n="42"`, `n=null`, `user=/Admin/`, `user=/Admin/i`,
		`user=NOCASE(Admin)`, `source=IP(10.0.0.0/8)`, `n>42`,
		`user IN [Admin,42]`, `user STARTS-WITH Admin`, `user ICONTAINS Admin`,
		`key1==key2`, `key1!==key2`,
	} {
		tests = append(tests, struct {
			query string
			flags []bool
		}{`where(` + predicate + `, loose)`, []bool{true}})
	}
	apis := []struct {
		name  string
		parse func(string) any
	}{
		{"extract", func(q string) any { return ExtractConditions(q) }},
		{"extract-context", func(q string) any { return ExtractConditionsContext(context.Background(), q) }},
		{"analyze", func(q string) any { return AnalyzeQuery(q) }},
		{"analyze-context", func(q string) any { return AnalyzeQueryContext(context.Background(), q) }},
	}
	for _, tt := range tests {
		for _, api := range apis {
			t.Run(api.name+"/"+tt.query, func(t *testing.T) {
				loose := looseJSON(t, api.parse(tt.query))
				if loose["loose"] != true {
					t.Errorf("query loose metadata missing: %+v", loose)
				}
				if errors, ok := loose["errors"]; ok {
					t.Fatalf("unexpected errors: %v", errors)
				}
				conditions := loose["conditions"].([]any)
				if len(conditions) != len(tt.flags) {
					t.Fatalf("got %d conditions, want %d", len(conditions), len(tt.flags))
				}
				for i, want := range tt.flags {
					flag, present := conditions[i].(map[string]any)["loose"]
					if (want && flag != true) || (!want && present) {
						t.Errorf("condition %d loose = %v (present %v), want %v", i, flag, present, want)
					}
				}
				// Every preexisting field and the entire boolean topology must match
				// the strict query, including semantic expression leaf contents.
				strictQuery := strings.ReplaceAll(strings.ReplaceAll(tt.query, ", loose)", ")"), ", LOOSE)", ")")
				strict := looseJSON(t, api.parse(strictQuery))
				if strings.Contains(string(mustLooseJSON(t, strict)), `"loose"`) {
					t.Error("strict query emitted loose metadata")
				}
				removeLooseJSON(loose)
				if !reflect.DeepEqual(loose, strict) {
					t.Errorf("mode changed existing data:\nloose %+v\nstrict %+v", loose, strict)
				}
			})
		}
	}
}

func TestLooseDedupRetainsStrictFromAndLooseWhere(t *testing.T) {
	result := ExtractConditions(`from(user=admin) where(user=admin, loose)`)
	conditions := DeduplicateConditions(append(result.Conditions, result.Conditions...))
	if len(conditions) != 2 {
		t.Fatalf("dedup erased mode distinction: %+v", conditions)
	}
}

func TestLooseBooleanTopologyAndConditions(t *testing.T) {
	tests := []struct {
		query      string
		conditions []Condition
		topology   string
	}{
		{`where(user=admin, loose)`, []Condition{
			{Loose: true, Field: "user", Operator: "=", Value: "admin", LogicalOp: "AND"},
		}, `{"operator":"condition","condition_index":0}`},
		{`where(user,actor=Admin, loose)`, []Condition{
			{Loose: true, Field: "user", Operator: "=", Value: "Admin", LogicalOp: "AND"},
			{Loose: true, Field: "actor", Operator: "=", Value: "Admin", LogicalOp: "OR"},
		}, `{"operator":"or","condition_index":0,"children":[{"operator":"condition","condition_index":0},{"operator":"condition","condition_index":1}]}`},
		{`where(all(user,actor)=Admin, loose)`, []Condition{
			{Loose: true, Field: "user", Operator: "=", Value: "Admin", LogicalOp: "AND"},
			{Loose: true, Field: "actor", Operator: "=", Value: "Admin", LogicalOp: "AND"},
		}, `{"operator":"and","condition_index":0,"children":[{"operator":"condition","condition_index":0},{"operator":"condition","condition_index":1}]}`},
		{`where(NOT (user=Admin OR actor=Root), loose)`, []Condition{
			{Loose: true, Field: "user", Operator: "=", Value: "Admin", Negated: true, LogicalOp: "AND"},
			{Loose: true, Field: "actor", Operator: "=", Value: "Root", Negated: true, LogicalOp: "OR"},
		}, `{"operator":"and","condition_index":0,"children":[{"operator":"condition","condition_index":0},{"operator":"condition","condition_index":1}]}`},
	}
	for _, tt := range tests {
		t.Run(tt.query, func(t *testing.T) {
			result := ExtractConditions(tt.query)
			if !result.Loose || len(result.Errors) != 0 || !reflect.DeepEqual(result.Conditions, tt.conditions) {
				t.Fatalf("unexpected extraction: %+v", result)
			}
			if got := string(mustLooseJSON(t, result.Expression)); got != tt.topology {
				t.Errorf("topology = %s, want %s", got, tt.topology)
			}
			semantic := AnalyzeQuery(tt.query)
			if !semantic.Loose || len(semantic.Errors) != 0 || !reflect.DeepEqual(semantic.Conditions, tt.conditions) {
				t.Fatalf("unexpected semantic query: %+v", semantic)
			}
			var checkLeaves func(*ConditionExpression, *SemanticExpression)
			checkLeaves = func(extracted *ConditionExpression, converted *SemanticExpression) {
				if converted == nil || extracted.Operator != converted.Operator || len(extracted.Children) != len(converted.Children) {
					t.Fatal("semantic topology differs")
				}
				if extracted.Operator == BooleanCondition && converted.Condition != &semantic.Conditions[extracted.ConditionIndex] {
					t.Fatal("semantic leaf does not point to corresponding flagged condition")
				}
				for i := range extracted.Children {
					checkLeaves(extracted.Children[i], converted.Children[i])
				}
			}
			checkLeaves(result.Expression, semantic.Expression)
		})
	}
}

func TestLooseSemanticCloneAndLeaves(t *testing.T) {
	result := ExtractConditions(`where(user IN [Admin,Root], loose)`)
	semantic := semanticFromExtraction(result, nil)
	if looseJSON(t, semantic)["loose"] != true || looseJSON(t, semantic.Conditions[0])["loose"] != true {
		t.Fatal("semantic clone lost loose mode")
	}
	if semantic.Expression.Condition != &semantic.Conditions[0] {
		t.Fatal("expression does not point at cloned condition")
	}
	// Replace the original condition, including its mode, after conversion.
	result.Conditions[0] = Condition{Field: "mutated"}
	if semantic.Expression.Condition.Field != "user" || looseJSON(t, semantic.Expression.Condition)["loose"] != true {
		t.Fatal("semantic leaf aliases extraction")
	}
}

func looseJSON(t *testing.T, value any) map[string]any {
	t.Helper()
	var decoded map[string]any
	if err := json.Unmarshal(mustLooseJSON(t, value), &decoded); err != nil {
		t.Fatal(err)
	}
	return decoded
}

func mustLooseJSON(t *testing.T, value any) []byte {
	t.Helper()
	data, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func removeLooseJSON(value any) {
	switch value := value.(type) {
	case map[string]any:
		delete(value, "loose")
		for _, child := range value {
			removeLooseJSON(child)
		}
	case []any:
		for _, child := range value {
			removeLooseJSON(child)
		}
	}
}
