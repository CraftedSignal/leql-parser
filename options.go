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
// Explicit AND binds tighter than implicit AND, which is left-associative, so
// the terms of one conjunction sit in nested AND and implicit-AND nodes. When a
// pair meets at this boundary, the conjunction is rebuilt term by term in
// source order, forming every adjacent pair; OR, grouping and NOT stay intact
// inside their own terms.
func (e *conditionExtractor) knownFieldRegexExpression(expressions []IExpressionContext) (*ConditionExpression, bool) {
	if len(e.knownFields) == 0 || len(expressions) != 2 {
		return nil, false
	}
	if _, _, ok := e.knownRegexPair(e.andBoundary(expressions[0], false), e.andBoundary(expressions[1], true)); !ok {
		return nil, false
	}
	terms := e.andTerms(expressions[1], e.andTerms(expressions[0], nil))
	previousLogicalOp := e.currentLogicalOp
	var expression *ConditionExpression
	for i := 0; i < len(terms); {
		term, used := e.buildKnownRegexTerm(terms[i:])
		e.currentLogicalOp = "AND"
		expression = e.combineKnownRegexAnd(expression, term)
		i += used
	}
	e.currentLogicalOp = previousLogicalOp
	return expression, true
}

// knownRegexPair reports a known bare name directly followed by a regex literal.
func (e *conditionExtractor) knownRegexPair(leftTerm, rightTerm IExpressionContext) (string, *RegexPatternContext, bool) {
	left, ok := leftTerm.(*ConditionExprContext)
	if !ok || left.Condition() == nil {
		return "", nil, false
	}
	bare, ok := left.Condition().(*FieldExistsConditionContext)
	if !ok || bare.IDENTIFIER() == nil {
		return "", nil, false
	}
	right, ok := rightTerm.(*RegexExprContext)
	if !ok || right.RegexSearch() == nil {
		return "", nil, false
	}
	literal, ok := right.RegexSearch().(*RegexPatternContext)
	if !ok || literal.REGEX() == nil {
		return "", nil, false
	}
	if bare.IDENTIFIER().GetSymbol().GetTokenIndex()+1 != literal.REGEX().GetSymbol().GetTokenIndex() {
		return "", nil, false
	}
	name := bare.IDENTIFIER().GetText()
	if !e.isKnownField(name) {
		return "", nil, false
	}
	return name, literal, true
}

func (e *conditionExtractor) andBoundary(expression IExpressionContext, left bool) IExpressionContext {
	for {
		e.budget.step()
		children := andOperands(expression)
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

// andTerms appends the terms of nested explicit and implicit AND nodes in
// source order.
func (e *conditionExtractor) andTerms(expression IExpressionContext, terms []IExpressionContext) []IExpressionContext {
	e.budget.step()
	children := andOperands(expression)
	if len(children) != 2 {
		return append(terms, expression)
	}
	return e.andTerms(children[1], e.andTerms(children[0], terms))
}

func andOperands(expression IExpressionContext) []IExpressionContext {
	switch and := expression.(type) {
	case *AndExprContext:
		return and.AllExpression()
	case *ImplicitAndExprContext:
		return and.AllExpression()
	}
	return nil
}

// buildKnownRegexTerm builds the first term and reports how many terms it used:
// a known bare name and the regex literal directly after it form one fielded
// regex predicate; any other term keeps its own scope.
func (e *conditionExtractor) buildKnownRegexTerm(terms []IExpressionContext) (*ConditionExpression, int) {
	e.budget.step()
	if len(terms) > 1 {
		if name, literal, ok := e.knownRegexPair(terms[0], terms[1]); ok {
			text := literal.REGEX().GetText()
			flags := regexLiteralFlags(text)
			e.budget.reserveOutput(1)
			rhs := &Operand{Kind: OperandRegex, Value: stripRegexSlashes(text), RegexFlags: flags}
			start := len(e.conditions)
			e.addOperandConditions([]string{name}, "=~", rhs, e.negated, "OR", strings.Contains(flags, "i"), flags)
			e.budget.reserveOutput(1)
			return e.conditionRangeExpression(start, len(e.conditions)), 2
		}
	}
	return e.buildExpression(terms[0]), 1
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
