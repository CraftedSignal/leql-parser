package leql

import (
	"context"
	"reflect"
	"strings"
	"testing"
)

func TestKnownFieldsPreserveScopedDetectionPredicates(t *testing.T) {
	opts := Options{KnownFields: []string{"process.name", "process.parent_name", "ParentImage"}}
	tests := []struct {
		name, query               string
		fields, operators, values []string
		negated                   []bool
		shape                     string
	}{
		{"negated known field", `where(NOT PROCESS.PARENT_NAME OR ParentImage)`, []string{"PROCESS.PARENT_NAME", "ParentImage"}, []string{"exists", "exists"}, []string{"", ""}, []bool{true, false}, "or(0,1)"},
		{"implicit known field regex", `where(process.name /ok$/i)`, []string{"process.name"}, []string{"=~"}, []string{"ok$"}, []bool{false}, "0"},
		{"prefix AND", `where(x=1 AND process.name /ok$/i)`, []string{"x", "process.name"}, []string{"=", "=~"}, []string{"1", "ok$"}, []bool{false, false}, "and(0,1)"},
		{"suffix AND", `where(process.name /ok$/i AND y=2)`, []string{"process.name", "y"}, []string{"=~", "="}, []string{"ok$", "2"}, []bool{false, false}, "and(0,1)"},
		{"both AND sides", `where(x=1 AND process.name /ok$/i AND y=2)`, []string{"x", "process.name", "y"}, []string{"=", "=~", "="}, []string{"1", "ok$", "2"}, []bool{false, false, false}, "and(0,1,2)"},
		{"AND before regex stays separate", `where(x=1 AND process.name AND /ok$/i AND y=2)`, []string{"x", "process.name", "", "y"}, []string{"=", "exists", "regex", "="}, []string{"1", "", "ok$", "2"}, []bool{false, false, false, false}, "and(0,1,2,3)"},
		{"NOT before regex stays separate", `where(process.name NOT /ok$/i)`, []string{"process.name", ""}, []string{"exists", "regex"}, []string{"", "ok$"}, []bool{false, true}, "and(0,1)"},
		{"negated pair", `where(NOT (process.name /ok$/im))`, []string{"process.name"}, []string{"=~"}, []string{"ok$"}, []bool{true}, "0"},
		{"explicit AND stays separate", `where(process.name AND /ok$/i)`, []string{"process.name", ""}, []string{"exists", "regex"}, []string{"", "ok$"}, []bool{false, false}, "and(0,1)"},
		{"OR stays separate", `where(process.name OR /ok$/i)`, []string{"process.name", ""}, []string{"exists", "regex"}, []string{"", "ok$"}, []bool{false, false}, "or(0,1)"},
		{"unknown stays keyword", `where(evil.exe /ok$/i)`, []string{"", ""}, []string{"keyword", "regex"}, []string{"evil.exe", "ok$"}, []bool{false, false}, "and(0,1)"},
		{"quoted stays keyword", `where("process.name" /ok$/i)`, []string{"", ""}, []string{"keyword", "regex"}, []string{"process.name", "ok$"}, []bool{false, false}, "and(0,1)"},
		{"group boundary stays separate", `where((process.name) /ok$/i)`, []string{"process.name", ""}, []string{"exists", "regex"}, []string{"", "ok$"}, []bool{false, false}, "and(0,1)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractConditionsWithOptions(tt.query, opts)
			if len(result.Errors) != 0 || len(result.Conditions) != len(tt.fields) {
				t.Fatalf("unexpected extraction: %+v", result)
			}
			for i, c := range result.Conditions {
				if c.Field != tt.fields[i] || c.Operator != tt.operators[i] || c.Value != tt.values[i] || c.Negated != tt.negated[i] {
					t.Fatalf("condition %d lost scoped semantics: %+v", i, c)
				}
			}
			if got := literalParseShape(result.Expression); got != tt.shape {
				t.Fatalf("expression=%s want=%s", got, tt.shape)
			}
			semantic := AnalyzeQueryWithOptions(tt.query, opts)
			if len(semantic.Errors) != 0 || !reflect.DeepEqual(semantic.Conditions, result.Conditions) {
				t.Fatalf("semantic extraction differs: %+v", semantic)
			}
			assertKnownFieldExpression(t, result.Expression, semantic.Expression, semantic.Conditions)
			if tt.name == "implicit known field regex" || tt.name == "negated pair" {
				c := result.Conditions[0]
				if c.RHS == nil || c.RHS.Kind != OperandRegex || c.RHS.RegexFlags != c.RegexFlags || !c.CaseInsensitive {
					t.Fatalf("regex flags/type lost: %+v", c)
				}
			}
		})
	}
}

func TestKnownFieldsKeepDefaultKeywordsAndLooseProvenance(t *testing.T) {
	query := `where(run OR evil.exe)`
	if got, want := ExtractConditionsWithOptions(query, Options{}), ExtractConditions(query); !reflect.DeepEqual(got, want) {
		t.Fatalf("zero options changed defaults: got=%+v want=%+v", got, want)
	}
	for _, c := range ExtractConditionsWithOptions(query, Options{KnownFields: []string{"process.name"}}).Conditions {
		if c.Field != "" || c.Operator != "keyword" {
			t.Fatalf("unknown word became field: %+v", c)
		}
	}
	opts := Options{KnownFields: []string{"process.name"}}
	got := ExtractConditionsWithOptions(`where(process.name /ok/i, loose)`, opts)
	semantic := AnalyzeQueryWithOptions(`where(process.name /ok/i, loose)`, opts)
	if len(got.Errors) != 0 || len(got.Conditions) != 1 || !got.Loose || !got.Conditions[0].Loose || !semantic.Loose || !semantic.Conditions[0].Loose {
		t.Fatalf("loose metadata lost: %+v %+v", got, semantic)
	}
}

func TestKnownFieldOptionsRespectSharedBoundsAndCancellation(t *testing.T) {
	tests := []struct {
		name  string
		setup func(*parseLimits)
		opts  Options
		want  string
	}{
		{"entry budget", func(l *parseLimits) { l.maxOutputItems = 1 }, Options{KnownFields: []string{"one", "two"}}, "output"},
		{"name bytes", func(l *parseLimits) { l.maxInputBytes = 16 }, Options{KnownFields: []string{"0123456789", "0123456789"}}, "known-field bytes"},
		{"option work", func(l *parseLimits) { l.maxWork = 1 }, Options{KnownFields: []string{"one", "two"}}, "work budget"},
		{"shared output", func(l *parseLimits) { l.maxOutputItems = 1 }, Options{KnownFields: []string{"x"}}, "output"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			limits := defaultParseLimits()
			tt.setup(&limits)
			got := extractConditionsContextWithOptions(context.Background(), `where(x)`, limits, tt.opts)
			if len(got.Errors) != 1 || !strings.Contains(got.Errors[0], tt.want) || len(got.Conditions) != 0 || got.Expression != nil {
				t.Fatalf("limit returned partial extraction: %+v", got)
			}
			semantic := analyzeQueryContextWithOptions(context.Background(), `where(x)`, limits, tt.opts)
			if len(semantic.Errors) != 1 || !strings.Contains(semantic.Errors[0], tt.want) || len(semantic.Conditions) != 0 || semantic.Expression != nil {
				t.Fatalf("limit returned partial semantic result: %+v", semantic)
			}
		})
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	opts := Options{KnownFields: []string{"x"}}
	if got := ExtractConditionsContextWithOptions(ctx, `where(x)`, opts); len(got.Errors) != 1 || !strings.Contains(got.Errors[0], "context canceled") || len(got.Conditions) != 0 {
		t.Fatalf("cancellation ignored: %+v", got)
	}
	if got := AnalyzeQueryContextWithOptions(ctx, `where(x)`, opts); len(got.Errors) != 1 || !strings.Contains(got.Errors[0], "context canceled") || len(got.Conditions) != 0 {
		t.Fatalf("semantic cancellation ignored: %+v", got)
	}
}

func TestKnownFieldConsumerRegressionFixtures(t *testing.T) {
	opts := Options{KnownFields: []string{"process.name", "process.parent_name"}}
	regex := ExtractConditionsWithOptions(`where(process.name /(\\cmd\.exe|\\powershell\.exe)$/i AND process.cmd_line ICONTAINS-ANY ["whoami", "hostname"])`, opts)
	if len(regex.Errors) != 0 || len(regex.Conditions) != 2 || regex.Conditions[0].Field != "process.name" || regex.Conditions[0].Operator != "=~" || regex.Conditions[0].RegexFlags != "i" || regex.Conditions[1].Field != "process.cmd_line" || literalParseShape(regex.Expression) != "and(0,1)" {
		t.Fatalf("consumer regex/alternative predicates changed: %+v", regex)
	}
	presence := ExtractConditionsWithOptions(`where(process.name ISTARTS-WITH "C:\\Windows\\System32\\" AND NOT process.parent_name)`, opts)
	if len(presence.Errors) != 0 || len(presence.Conditions) != 2 || presence.Conditions[1].Field != "process.parent_name" || presence.Conditions[1].Operator != "exists" || !presence.Conditions[1].Negated || literalParseShape(presence.Expression) != "and(0,1)" {
		t.Fatalf("consumer escaped-string/negated-presence predicates changed: %+v", presence)
	}
}

func assertKnownFieldExpression(t *testing.T, parsed *ConditionExpression, semantic *SemanticExpression, conditions []Condition) {
	t.Helper()
	if parsed == nil || semantic == nil || parsed.Operator != semantic.Operator || len(parsed.Children) != len(semantic.Children) {
		t.Fatal("semantic expression changed scoped topology")
	}
	if parsed.Operator == BooleanCondition && semantic.Condition != &conditions[parsed.ConditionIndex] {
		t.Fatal("semantic leaf does not reference its cloned condition")
	}
	for i, child := range parsed.Children {
		assertKnownFieldExpression(t, child, semantic.Children[i], conditions)
	}
}
