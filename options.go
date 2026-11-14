package leql

import (
	"context"
	"strings"
)

// Options supplies trusted schema knowledge for otherwise ambiguous bare words.
// Zero Options preserves ordinary keyword-search semantics.
type Options struct {
	// KnownFields treats matching bare identifiers as field-presence predicates.
	// Names match case-insensitively; quoted and unknown words remain keywords.
	// A directly adjacent known bare name and /regex/ form one fielded regex
	// predicate; explicit AND/OR and expression boundaries remain separate.
	KnownFields []string
}

// ExtractConditionsWithOptions parses with scoped schema knowledge.
func ExtractConditionsWithOptions(query string, options Options) *ParseResult {
	return ExtractConditionsContextWithOptions(context.Background(), query, options)
}

// ExtractConditionsContextWithOptions shares caller cancellation and parser
// limits with processing of the supplied known-field names.
func ExtractConditionsContextWithOptions(ctx context.Context, query string, options Options) *ParseResult {
	return extractConditionsContextWithOptions(ctx, query, defaultParseLimits(), options)
}

// AnalyzeQueryWithOptions parses and converts with scoped schema knowledge.
func AnalyzeQueryWithOptions(query string, options Options) *SemanticQuery {
	return AnalyzeQueryContextWithOptions(context.Background(), query, options)
}

// AnalyzeQueryContextWithOptions keeps one budget through schema processing,
// parsing and semantic conversion.
func AnalyzeQueryContextWithOptions(ctx context.Context, query string, options Options) *SemanticQuery {
	return analyzeQueryContextWithOptions(ctx, query, defaultParseLimits(), options)
}

// knownFieldSet checks every bound before normalization or map allocation.
// Names consume the same work/output budget as parsing and semantic conversion.
func knownFieldSet(options Options, budget *parseBudget) map[string]struct{} {
	if len(options.KnownFields) == 0 {
		return nil
	}
	budget.check()
	budget.reserveOutput(len(options.KnownFields))
	bytes := 0
	for _, name := range options.KnownFields {
		budget.step()
		if len(name) > budget.limits.maxInputBytes-bytes {
			budget.reject("known-field bytes")
		}
		bytes += len(name)
	}
	fields := make(map[string]struct{}, len(options.KnownFields))
	for _, name := range options.KnownFields {
		budget.step()
		fields[strings.ToLower(name)] = struct{}{}
	}
	return fields
}

func (e *conditionExtractor) isKnownField(name string) bool {
	if len(e.knownFields) == 0 {
		return false
	}
	e.budget.step()
	_, ok := e.knownFields[strings.ToLower(name)]
	return ok
}

// Only adjacent leaves at an implicit-AND boundary select this shorthand.
// The grammar may put surrounding explicit-AND terms inside either child;
// following only their boundary spines leaves OR, grouping and NOT intact.
func (e *conditionExtractor) knownFieldRegexExpression(expressions []IExpressionContext) (*ConditionExpression, bool) {
	if len(e.knownFields) == 0 || len(expressions) != 2 {
		return nil, false
	}
	leftBoundary := e.andBoundary(expressions[0], false)
	rightBoundary := e.andBoundary(expressions[1], true)
	left, ok := leftBoundary.(*ConditionExprContext)
	if !ok || left.Condition() == nil {
		return nil, false
	}
	bare, ok := left.Condition().(*FieldExistsConditionContext)
	if !ok || bare.IDENTIFIER() == nil {
		return nil, false
	}
	right, ok := rightBoundary.(*RegexExprContext)
	if !ok || right.RegexSearch() == nil {
		return nil, false
	}
	literal, ok := right.RegexSearch().(*RegexPatternContext)
	if !ok || literal.REGEX() == nil {
		return nil, false
	}
	if bare.IDENTIFIER().GetSymbol().GetTokenIndex()+1 != literal.REGEX().GetSymbol().GetTokenIndex() {
		return nil, false
	}
	name := bare.IDENTIFIER().GetText()
	if !e.isKnownField(name) {
		return nil, false
	}
	previousLogicalOp := e.currentLogicalOp
	leftExpression := e.buildKnownRegexAndSide(expressions[0], leftBoundary, name, literal)
	e.currentLogicalOp = "AND"
	rightExpression := e.buildKnownRegexAndSide(expressions[1], rightBoundary, "", nil)
	e.currentLogicalOp = previousLogicalOp
	return e.combineKnownRegexAnd(leftExpression, rightExpression), true
}

func (e *conditionExtractor) andBoundary(expression IExpressionContext, left bool) IExpressionContext {
	for {
		e.budget.step()
		and, ok := expression.(*AndExprContext)
		if !ok {
			return expression
		}
		children := and.AllExpression()
		if len(children) != 2 {
			return expression
		}
		if left {
			expression = children[0]
		} else {
			expression = children[1]
		}
	}
}

// Replace the matched bare leaf and omit the matched regex leaf, rebuilding
// only explicit-AND nodes and retaining all other subtrees in their own scope.
func (e *conditionExtractor) buildKnownRegexAndSide(expression, target IExpressionContext, name string, literal *RegexPatternContext) *ConditionExpression {
	e.budget.step()
	if expression == target {
		if literal == nil {
			return nil
		}
		text := literal.REGEX().GetText()
		flags := regexLiteralFlags(text)
		e.budget.reserveOutput(1)
		rhs := &Operand{Kind: OperandRegex, Value: stripRegexSlashes(text), RegexFlags: flags}
		start := len(e.conditions)
		e.addOperandConditions([]string{name}, "=~", rhs, e.negated, "OR", strings.Contains(flags, "i"), flags)
		e.budget.reserveOutput(1)
		return e.conditionRangeExpression(start, len(e.conditions))
	}
	if and, ok := expression.(*AndExprContext); ok {
		children := and.AllExpression()
		if len(children) == 2 {
			left := e.buildKnownRegexAndSide(children[0], target, name, literal)
			previousLogicalOp := e.currentLogicalOp
			e.currentLogicalOp = "AND"
			right := e.buildKnownRegexAndSide(children[1], target, name, literal)
			e.currentLogicalOp = previousLogicalOp
			return e.combineKnownRegexAnd(left, right)
		}
	}
	return e.buildExpression(expression)
}

func (e *conditionExtractor) combineKnownRegexAnd(left, right *ConditionExpression) *ConditionExpression {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	// combineExpressions flattens AND children into a fresh slice. Charge the
	// full allocation before copying, including repeated combinations.
	items := 2
	if left.Operator == BooleanAnd {
		items += len(left.Children) - 1
	}
	if right.Operator == BooleanAnd {
		items += len(right.Children) - 1
	}
	e.budget.step()
	e.budget.reserveOutput(1 + items)
	return combineExpressions(BooleanAnd, left, right)
}
