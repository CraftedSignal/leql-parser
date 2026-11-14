package leql

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/antlr4-go/antlr/v4"
)

func TestExtractContextRejectsCanceledRequest(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	result := ExtractConditionsContext(ctx, `where(x=1 AND y=2)`)
	if result == nil || len(result.Errors) == 0 || !strings.Contains(result.Errors[0], "context canceled") || len(result.Conditions) != 0 {
		t.Fatalf("canceled parse returned conditions or no cancellation: %#v", result)
	}
}

func TestExtractionRejectsTinyBudgetsWithoutPartialConditions(t *testing.T) {
	tests := []struct {
		name, query, want string
		configure         func(*parseLimits)
	}{
		{"input", `where(x=1 AND y=2)`, "input", func(l *parseLimits) { l.maxInputBytes = 4 }},
		{"nesting", `where((x=1))`, "nesting", func(l *parseLimits) { l.maxNestingDepth = 1 }},
		{"tokens", `where(x=1 AND y=2)`, "token", func(l *parseLimits) { l.maxTokens = 2 }},
		{"work", `where(x=1 AND y=2)`, "work", func(l *parseLimits) { l.maxWork = 1 }},
		{"output", `where(x=1 AND y=2)`, "output", func(l *parseLimits) { l.maxOutputItems = 1 }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			limits := defaultParseLimits()
			tt.configure(&limits)
			got := extractConditionsContext(context.Background(), tt.query, limits)
			if got == nil || len(got.Errors) == 0 || !strings.Contains(got.Errors[0], tt.want) || len(got.Conditions) != 0 {
				t.Fatalf("expected %s rejection without partial conditions: %#v", tt.want, got)
			}
		})
	}
}

func TestNestingBudgetIgnoresLiteralAndCommentDelimiters(t *testing.T) {
	limits := defaultParseLimits()
	limits.maxNestingDepth = 1
	got := extractConditionsContext(context.Background(), `where(x="(((")`, limits)
	if got == nil || len(got.Errors) != 0 || len(got.Conditions) != 1 {
		t.Fatalf("literal/comment delimiters counted as syntax nesting: %#v", got)
	}
	if got.Conditions[0].Field != "x" || got.Conditions[0].Value != "(((" {
		t.Fatalf("literal semantics changed: %#v", got.Conditions)
	}
}

type panicDeadlineContext struct{ context.Context }

func (panicDeadlineContext) Deadline() (time.Time, bool) { panic("fixture panic") }

func TestPublicContextBoundaryReturnsFailureOnUnexpectedPanic(t *testing.T) {
	ctx := panicDeadlineContext{context.Background()}
	result := ExtractConditionsContext(ctx, "x=1")
	if len(result.Conditions) != 0 || len(result.Errors) != 1 || !strings.Contains(result.Errors[0], "parser panic") {
		t.Fatalf("unexpected panic did not return failure only: %+v", result)
	}
	semantic := AnalyzeQueryContext(ctx, "x=1")
	if len(semantic.Conditions) != 0 || semantic.Expression != nil || len(semantic.Errors) != 1 || !strings.Contains(semantic.Errors[0], "parser panic") {
		t.Fatalf("unexpected semantic panic did not return failure only: %+v", semantic)
	}
}

func TestAnalyzeContextRejectsCanceledRequest(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	got := AnalyzeQueryContext(ctx, "x=1")
	if got == nil || len(got.Errors) == 0 || !strings.Contains(got.Errors[0], "context canceled") || len(got.Conditions) != 0 {
		t.Fatalf("semantic entrypoint ignored cancellation: %#v", got)
	}
}

func TestAnalyzeSharesBudgetWithSemanticConversion(t *testing.T) {
	limits := defaultParseLimits()
	limits.maxOutputItems = 1
	query := "where(x=1)"
	if result := extractConditionsContext(context.Background(), query, limits); len(result.Errors) != 0 || len(result.Conditions) != 1 {
		t.Fatalf("extraction must fit this budget: %+v", result)
	}
	result := analyzeQueryContext(context.Background(), query, limits)
	if len(result.Conditions) != 0 || result.Expression != nil || len(result.Errors) != 1 || !strings.Contains(result.Errors[0], "output items") {
		t.Fatalf("semantic conversion must share the exhausted budget without partial output: %+v", result)
	}
}

// These small helper fixtures cover the upstream literal-semantic paths before
// their allocations, independently of the public boundary's final size check.
func TestLiteralFieldHelperReservesOutputBeforeAppend(t *testing.T) {
	limits := defaultParseLimits()
	limits.maxOutputItems = 1
	budget, cancel := newParseBudget(context.Background(), limits)
	defer cancel()
	extractor := &conditionExtractor{budget: budget, currentLogicalOp: "AND"}
	defer func() {
		value := recover()
		abort, ok := value.(budgetAbort)
		if !ok || !strings.Contains(abort.err.Error(), "output items") {
			t.Fatalf("expected output rejection before append, got %v", value)
		}
		if len(extractor.conditions) != 1 {
			t.Fatalf("over-budget condition appended: %+v", extractor.conditions)
		}
	}()
	extractor.addFieldValueConditions([]string{"x", "y"}, "=", nil, false, "OR")
}

func TestRegexLiteralValidationSharesWorkAndCancellation(t *testing.T) {
	tree := NewLEQLParser(antlr.NewCommonTokenStream(NewLEQLLexer(antlr.NewInputStream(`where(x > /ok/i)`)), antlr.TokenDefaultChannel)).Query()
	for _, name := range []string{"work", "cancel", "output"} {
		t.Run(name, func(t *testing.T) {
			limits := defaultParseLimits()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			want := ""
			switch name {
			case "work":
				limits.maxWork = 0
				want = "work budget"
			case "cancel":
				cancel()
				want = "context canceled"
			case "output":
				limits.maxOutputItems = 0
				want = "output items"
			}
			budget, stop := newParseBudget(ctx, limits)
			defer stop()
			extractor := &conditionExtractor{budget: budget}
			defer func() {
				value := recover()
				abort, ok := value.(budgetAbort)
				if !ok || !strings.Contains(abort.err.Error(), want) {
					t.Fatalf("expected %s rejection, got %v", want, value)
				}
				if len(extractor.errors) != 0 {
					t.Fatalf("over-budget diagnostic appended: %+v", extractor.errors)
				}
			}()
			extractor.validateRegexLiteralUses(tree)
		})
	}
}
