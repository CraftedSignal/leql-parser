package leql

import (
	"fmt"
	"strings"
	"time"

	"github.com/antlr4-go/antlr/v4"
)

// MaxParseTime is the maximum time allowed for parsing a single query.
// Queries that exceed this are returned with an error.
var MaxParseTime = 5 * time.Second

// Condition represents a field condition extracted from a LEQL query
type Condition struct {
	Field        string   `json:"field"`
	Operator     string   `json:"operator"`
	Value        string   `json:"value"`
	Negated      bool     `json:"negated"`
	PipeStage    int      `json:"pipe_stage"`
	LogicalOp    string   `json:"logical_op"`             // "AND" or "OR" connecting to previous condition
	Alternatives []string `json:"alternatives,omitempty"`  // For OR conditions on same field
	IsComputed   bool     `json:"is_computed,omitempty"`   // True if field was created by a computed expression
	SourceField  string   `json:"source_field,omitempty"`  // Original field before transformation (for computed fields)
}

// ParseResult contains all conditions extracted from the query
type ParseResult struct {
	Conditions     []Condition       `json:"conditions"`
	GroupByFields  []string          `json:"group_by_fields,omitempty"`  // Fields from groupby clause
	ComputedFields map[string]string `json:"computed_fields,omitempty"`  // Map of computed field name -> source field
	Commands       []string          `json:"commands,omitempty"`         // List of commands used in the query (where, groupby, calculate, etc.)
	Joins          []JoinInfo        `json:"joins,omitempty"`            // Extracted join info (LEQL has no joins; kept for API parity)
	Errors         []string          `json:"errors,omitempty"`
}

// FieldProvenance indicates where a field originates relative to a join
type FieldProvenance string

const (
	ProvenanceMain      FieldProvenance = "main"
	ProvenanceJoined    FieldProvenance = "joined"
	ProvenanceJoinKey   FieldProvenance = "join_key"
	ProvenanceAmbiguous FieldProvenance = "ambiguous"
)

// JoinInfo captures the structured decomposition of a JOIN command.
// LEQL does not support joins; this type exists for API parity with KQL/SPL parsers.
type JoinInfo struct {
	Type          string       `json:"type"`
	JoinFields    []string     `json:"join_fields,omitempty"`
	LeftFields    []string     `json:"left_fields,omitempty"`
	RightFields   []string     `json:"right_fields,omitempty"`
	RightTable    string       `json:"right_table,omitempty"`
	Subsearch     *ParseResult `json:"subsearch,omitempty"`
	PipeStage     int          `json:"pipe_stage"`
	ExposedFields []string     `json:"exposed_fields,omitempty"`
}

// conditionExtractor walks the parse tree using the visitor pattern to extract conditions.
type conditionExtractor struct {
	*BaseLEQLParserVisitor
	conditions     []Condition
	groupByFields  []string
	commands       []string
	errors         []string
	currentLogicalOp string
	negated        bool
}

// errorListener collects parse errors
type errorListener struct {
	*antlr.DefaultErrorListener
	errors []string
}

func (l *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.errors = append(l.errors, msg)
}

// ExtractConditions parses a LEQL query and extracts all field conditions.
// This is the main entry point for the LEQL parser.
func ExtractConditions(query string) *ParseResult {
	ch := make(chan *ParseResult, 1)
	go func() {
		ch <- extractConditionsInternal(query)
	}()

	select {
	case result := <-ch:
		return result
	case <-time.After(MaxParseTime):
		return &ParseResult{
			Conditions: []Condition{},
			Commands:   []string{},
			Errors:     []string{fmt.Sprintf("parser timeout: query took longer than %s to parse", MaxParseTime)},
		}
	}
}

func extractConditionsInternal(query string) (result *ParseResult) {
	defer func() {
		if r := recover(); r != nil {
			result = &ParseResult{
				Conditions: []Condition{},
				Commands:   []string{},
				Errors:     []string{fmt.Sprintf("parser panic: %v", r)},
			}
		}
	}()

	input := antlr.NewInputStream(query)
	lexer := NewLEQLLexer(input)

	// Remove default error listener and add our own
	lexer.RemoveErrorListeners()
	lexerErrors := &errorListener{}
	lexer.AddErrorListener(lexerErrors)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewLEQLParser(stream)

	// Remove default error listener and add our own
	parser.RemoveErrorListeners()
	parserErrors := &errorListener{}
	parser.AddErrorListener(parserErrors)

	// Parse the query
	tree := parser.Query()

	// Walk the tree using visitor to extract conditions
	extractor := &conditionExtractor{
		conditions:       make([]Condition, 0),
		groupByFields:    make([]string, 0),
		commands:         make([]string, 0),
		currentLogicalOp: "AND",
	}
	extractor.visit(tree)

	// Combine errors
	allErrors := append(lexerErrors.errors, parserErrors.errors...)
	allErrors = append(allErrors, extractor.errors...)

	return &ParseResult{
		Conditions:     extractor.conditions,
		GroupByFields:  extractor.groupByFields,
		ComputedFields: nil,
		Commands:       extractor.commands,
		Joins:          nil,
		Errors:         allErrors,
	}
}

// visit dispatches to the appropriate visitor method based on the node type.
func (e *conditionExtractor) visit(tree antlr.ParseTree) {
	if tree == nil {
		return
	}
	switch ctx := tree.(type) {
	case *QueryContext:
		e.visitQuery(ctx)
	case *FromClauseContext:
		e.visitFromClause(ctx)
	case *WhereClauseContext:
		e.visitWhereClause(ctx)
	case *GroupbyClauseContext:
		e.visitGroupbyClause(ctx)
	case *CalculateClauseContext:
		e.visitCalculateClause(ctx)
	case *SelectClauseContext:
		e.visitSelectClause(ctx)
	case *SortClauseContext:
		e.visitSortClause(ctx)
	case *LimitClauseContext:
		e.visitLimitClause(ctx)
	case *TimesliceClauseContext:
		e.visitTimesliceClause(ctx)
	case *HavingClauseContext:
		e.visitHavingClause(ctx)
	// Boolean expression nodes
	case *AndExprContext:
		e.visitAndExpr(ctx)
	case *OrExprContext:
		e.visitOrExpr(ctx)
	case *NotExprContext:
		e.visitNotExpr(ctx)
	case *ParenExprContext:
		e.visitParenExpr(ctx)
	case *NestedWhereExprContext:
		e.visitNestedWhereExpr(ctx)
	case *ConditionExprContext:
		e.visitConditionExpr(ctx)
	case *KeywordExprContext:
		e.visitKeywordExpr(ctx)
	case *RegexExprContext:
		e.visitRegexExpr(ctx)
	// Condition nodes
	case *ComparisonConditionContext:
		e.visitComparisonCondition(ctx)
	case *StringConditionContext:
		e.visitStringCondition(ctx)
	case *SetConditionContext:
		e.visitSetCondition(ctx)
	case *ListStringConditionContext:
		e.visitListStringCondition(ctx)
	case *NegatedComparisonConditionContext:
		e.visitNegatedComparisonCondition(ctx)
	case *NegatedStringConditionContext:
		e.visitNegatedStringCondition(ctx)
	case *NegatedSetConditionContext:
		e.visitNegatedSetCondition(ctx)
	case *NegatedListStringConditionContext:
		e.visitNegatedListStringCondition(ctx)
	case *PostfixNegatedComparisonConditionContext:
		e.visitPostfixNegatedComparisonCondition(ctx)
	case *PostfixNegatedStringConditionContext:
		e.visitPostfixNegatedStringCondition(ctx)
	case *PostfixNegatedSetConditionContext:
		e.visitPostfixNegatedSetCondition(ctx)
	case *PostfixNegatedListStringConditionContext:
		e.visitPostfixNegatedListStringCondition(ctx)
	case *NocaseConditionContext:
		e.visitNocaseCondition(ctx)
	case *IpConditionContext:
		e.visitIpCondition(ctx)
	case *IpSetConditionContext:
		e.visitIpSetCondition(ctx)
	case *AllFieldsComparisonConditionContext:
		e.visitAllFieldsComparisonCondition(ctx)
	case *AllFieldsSetConditionContext:
		e.visitAllFieldsSetCondition(ctx)
	case *AllFieldsNocaseConditionContext:
		e.visitAllFieldsNocaseCondition(ctx)
	case *FieldExistsConditionContext:
		e.visitFieldExistsCondition(ctx)
	case *NumberKeywordConditionContext:
		e.visitNumberKeywordCondition(ctx)
	default:
		// For unknown nodes, visit children
		if node, ok := tree.(antlr.RuleNode); ok {
			for i := 0; i < node.GetChildCount(); i++ {
				child := node.GetChild(i)
				if pt, ok := child.(antlr.ParseTree); ok {
					e.visit(pt)
				}
			}
		}
	}
}

// visitQuery processes the top-level query rule.
func (e *conditionExtractor) visitQuery(ctx *QueryContext) {
	if ctx.FromClause() != nil {
		e.visit(ctx.FromClause().(antlr.ParseTree))
	}
	if ctx.SelectClause() != nil {
		e.visit(ctx.SelectClause().(antlr.ParseTree))
	}
	if ctx.WhereClause() != nil {
		e.visit(ctx.WhereClause().(antlr.ParseTree))
	}
	if ctx.GroupbyClause() != nil {
		e.visit(ctx.GroupbyClause().(antlr.ParseTree))
	}
	if ctx.HavingClause() != nil {
		e.visit(ctx.HavingClause().(antlr.ParseTree))
	}
	if ctx.CalculateClause() != nil {
		e.visit(ctx.CalculateClause().(antlr.ParseTree))
	}
	if ctx.SortClause() != nil {
		e.visit(ctx.SortClause().(antlr.ParseTree))
	}
	if ctx.LimitClause() != nil {
		e.visit(ctx.LimitClause().(antlr.ParseTree))
	}
	if ctx.TimesliceClause() != nil {
		e.visit(ctx.TimesliceClause().(antlr.ParseTree))
	}
}

// visitFromClause handles the from(...) clause (log source selection).
func (e *conditionExtractor) visitFromClause(ctx *FromClauseContext) {
	e.commands = append(e.commands, "from")
	if ctx.Expression() != nil {
		e.visit(ctx.Expression().(antlr.ParseTree))
	}
}

// visitWhereClause handles the where(...) clause.
func (e *conditionExtractor) visitWhereClause(ctx *WhereClauseContext) {
	e.commands = append(e.commands, "where")
	if ctx.Expression() != nil {
		e.visit(ctx.Expression().(antlr.ParseTree))
	}
}

// visitGroupbyClause handles the groupby(...) clause.
func (e *conditionExtractor) visitGroupbyClause(ctx *GroupbyClauseContext) {
	e.commands = append(e.commands, "groupby")
	for _, fn := range ctx.AllFieldName() {
		e.groupByFields = append(e.groupByFields, stripQuotes(fn.GetText()))
	}
}

// visitCalculateClause handles the calculate(...) clause.
func (e *conditionExtractor) visitCalculateClause(_ *CalculateClauseContext) {
	e.commands = append(e.commands, "calculate")
}

// visitSelectClause handles the select(...) clause.
func (e *conditionExtractor) visitSelectClause(_ *SelectClauseContext) {
	e.commands = append(e.commands, "select")
}

// visitSortClause handles the sort(...) clause.
func (e *conditionExtractor) visitSortClause(_ *SortClauseContext) {
	e.commands = append(e.commands, "sort")
}

// visitLimitClause handles the limit(...) clause.
func (e *conditionExtractor) visitLimitClause(_ *LimitClauseContext) {
	e.commands = append(e.commands, "limit")
}

// visitTimesliceClause handles the timeslice(...) clause.
func (e *conditionExtractor) visitTimesliceClause(_ *TimesliceClauseContext) {
	e.commands = append(e.commands, "timeslice")
}

// visitHavingClause handles the having(...) clause.
func (e *conditionExtractor) visitHavingClause(_ *HavingClauseContext) {
	e.commands = append(e.commands, "having")
}

// --- Boolean expression visitors ---

func (e *conditionExtractor) visitAndExpr(ctx *AndExprContext) {
	exprs := ctx.AllExpression()
	if len(exprs) >= 1 {
		e.visit(exprs[0].(antlr.ParseTree))
	}
	e.currentLogicalOp = "AND"
	if len(exprs) >= 2 {
		e.visit(exprs[1].(antlr.ParseTree))
	}
}

func (e *conditionExtractor) visitOrExpr(ctx *OrExprContext) {
	exprs := ctx.AllExpression()
	if len(exprs) >= 1 {
		e.visit(exprs[0].(antlr.ParseTree))
	}
	e.currentLogicalOp = "OR"
	if len(exprs) >= 2 {
		e.visit(exprs[1].(antlr.ParseTree))
	}
}

func (e *conditionExtractor) visitNotExpr(ctx *NotExprContext) {
	prevNegated := e.negated
	e.negated = true
	if ctx.Expression() != nil {
		e.visit(ctx.Expression().(antlr.ParseTree))
	}
	e.negated = prevNegated
}

func (e *conditionExtractor) visitParenExpr(ctx *ParenExprContext) {
	if ctx.Expression() != nil {
		e.visit(ctx.Expression().(antlr.ParseTree))
	}
}

func (e *conditionExtractor) visitNestedWhereExpr(ctx *NestedWhereExprContext) {
	if ctx.Expression() != nil {
		e.visit(ctx.Expression().(antlr.ParseTree))
	}
}

func (e *conditionExtractor) visitConditionExpr(ctx *ConditionExprContext) {
	if ctx.Condition() != nil {
		e.visit(ctx.Condition().(antlr.ParseTree))
	}
}

func (e *conditionExtractor) visitKeywordExpr(ctx *KeywordExprContext) {
	if ctx.KeywordSearch() != nil {
		ks := ctx.KeywordSearch()
		switch kw := ks.(type) {
		case *QuotedKeywordContext:
			e.visitQuotedKeyword(kw)
		case *SingleQuotedKeywordContext:
			e.visitSingleQuotedKeyword(kw)
		}
	}
}

func (e *conditionExtractor) visitRegexExpr(ctx *RegexExprContext) {
	if ctx.RegexSearch() != nil {
		rs := ctx.RegexSearch()
		if rp, ok := rs.(*RegexPatternContext); ok {
			e.visitRegexPattern(rp)
		}
	}
}

// --- Keyword and regex search visitors ---

func (e *conditionExtractor) visitQuotedKeyword(ctx *QuotedKeywordContext) {
	text := ctx.DOUBLE_STRING().GetText()
	e.addCondition("", "keyword", stripQuotes(text))
}

func (e *conditionExtractor) visitSingleQuotedKeyword(ctx *SingleQuotedKeywordContext) {
	text := ctx.SINGLE_STRING().GetText()
	e.addCondition("", "keyword", stripQuotes(text))
}

func (e *conditionExtractor) visitRegexPattern(ctx *RegexPatternContext) {
	text := ctx.REGEX().GetText()
	pattern := stripRegexSlashes(text)
	e.addCondition("", "regex", pattern)
}

// --- Condition visitors ---

func (e *conditionExtractor) visitComparisonCondition(ctx *ComparisonConditionContext) {
	if ctx.FieldList() == nil || ctx.ComparisonOp() == nil || ctx.Value() == nil {
		return
	}
	fields := extractFieldNames(ctx.FieldList())
	op := extractComparisonOp(ctx.ComparisonOp())
	val := extractValue(ctx.Value())

	if len(fields) == 1 {
		e.addCondition(fields[0], op, val)
	} else {
		// Multiple fields in fieldList = separate conditions with LogicalOp "OR"
		for i, f := range fields {
			logOp := e.currentLogicalOp
			if i > 0 {
				logOp = "OR"
			}
			e.conditions = append(e.conditions, Condition{
				Field:     f,
				Operator:  op,
				Value:     val,
				Negated:   e.negated,
				LogicalOp: logOp,
			})
		}
		// Reset logical op for next condition
		e.currentLogicalOp = "AND"
	}
}

func (e *conditionExtractor) visitStringCondition(ctx *StringConditionContext) {
	if ctx.FieldList() == nil || ctx.StringOp() == nil || ctx.Value() == nil {
		return
	}
	fields := extractFieldNames(ctx.FieldList())
	op := extractStringOp(ctx.StringOp())
	val := extractValue(ctx.Value())

	for i, f := range fields {
		logOp := e.currentLogicalOp
		if i > 0 {
			logOp = "OR"
		}
		e.conditions = append(e.conditions, Condition{
			Field:     f,
			Operator:  op,
			Value:     val,
			Negated:   e.negated,
			LogicalOp: logOp,
		})
	}
	if len(fields) > 0 {
		e.currentLogicalOp = "AND"
	}
}

func (e *conditionExtractor) visitSetCondition(ctx *SetConditionContext) {
	if ctx.FieldList() == nil || ctx.SetOp() == nil || ctx.ValueList() == nil {
		return
	}
	fields := extractFieldNames(ctx.FieldList())
	op := extractSetOp(ctx.SetOp())
	alternatives := extractValueList(ctx.ValueList())

	for _, f := range fields {
		e.conditions = append(e.conditions, Condition{
			Field:        f,
			Operator:     op,
			Value:        strings.Join(alternatives, ", "),
			Negated:      e.negated,
			LogicalOp:    e.currentLogicalOp,
			Alternatives: alternatives,
		})
	}
	e.currentLogicalOp = "AND"
}

func (e *conditionExtractor) visitListStringCondition(ctx *ListStringConditionContext) {
	if ctx.FieldList() == nil || ctx.ListStringOp() == nil || ctx.ValueList() == nil {
		return
	}
	fields := extractFieldNames(ctx.FieldList())
	op := extractListStringOp(ctx.ListStringOp())
	alternatives := extractValueList(ctx.ValueList())

	for _, f := range fields {
		e.conditions = append(e.conditions, Condition{
			Field:        f,
			Operator:     op,
			Value:        strings.Join(alternatives, ", "),
			Negated:      e.negated,
			LogicalOp:    e.currentLogicalOp,
			Alternatives: alternatives,
		})
	}
	e.currentLogicalOp = "AND"
}

func (e *conditionExtractor) visitNegatedComparisonCondition(ctx *NegatedComparisonConditionContext) {
	if ctx.FieldList() == nil || ctx.ComparisonOp() == nil || ctx.Value() == nil {
		return
	}
	fields := extractFieldNames(ctx.FieldList())
	op := extractComparisonOp(ctx.ComparisonOp())
	val := extractValue(ctx.Value())

	for i, f := range fields {
		logOp := e.currentLogicalOp
		if i > 0 {
			logOp = "OR"
		}
		e.conditions = append(e.conditions, Condition{
			Field:     f,
			Operator:  op,
			Value:     val,
			Negated:   true,
			LogicalOp: logOp,
		})
	}
	e.currentLogicalOp = "AND"
}

func (e *conditionExtractor) visitNegatedStringCondition(ctx *NegatedStringConditionContext) {
	if ctx.FieldList() == nil || ctx.StringOp() == nil || ctx.Value() == nil {
		return
	}
	fields := extractFieldNames(ctx.FieldList())
	op := extractStringOp(ctx.StringOp())
	val := extractValue(ctx.Value())

	for i, f := range fields {
		logOp := e.currentLogicalOp
		if i > 0 {
			logOp = "OR"
		}
		e.conditions = append(e.conditions, Condition{
			Field:     f,
			Operator:  op,
			Value:     val,
			Negated:   true,
			LogicalOp: logOp,
		})
	}
	e.currentLogicalOp = "AND"
}

func (e *conditionExtractor) visitNegatedSetCondition(ctx *NegatedSetConditionContext) {
	if ctx.FieldList() == nil || ctx.SetOp() == nil || ctx.ValueList() == nil {
		return
	}
	fields := extractFieldNames(ctx.FieldList())
	op := extractSetOp(ctx.SetOp())
	alternatives := extractValueList(ctx.ValueList())

	for _, f := range fields {
		e.conditions = append(e.conditions, Condition{
			Field:        f,
			Operator:     op,
			Value:        strings.Join(alternatives, ", "),
			Negated:      true,
			LogicalOp:    e.currentLogicalOp,
			Alternatives: alternatives,
		})
	}
	e.currentLogicalOp = "AND"
}

func (e *conditionExtractor) visitNegatedListStringCondition(ctx *NegatedListStringConditionContext) {
	if ctx.FieldList() == nil || ctx.ListStringOp() == nil || ctx.ValueList() == nil {
		return
	}
	fields := extractFieldNames(ctx.FieldList())
	op := extractListStringOp(ctx.ListStringOp())
	alternatives := extractValueList(ctx.ValueList())

	for _, f := range fields {
		e.conditions = append(e.conditions, Condition{
			Field:        f,
			Operator:     op,
			Value:        strings.Join(alternatives, ", "),
			Negated:      true,
			LogicalOp:    e.currentLogicalOp,
			Alternatives: alternatives,
		})
	}
	e.currentLogicalOp = "AND"
}

// Postfix negation: field NOT op value (e.g., field NOT IN [...])

func (e *conditionExtractor) visitPostfixNegatedComparisonCondition(ctx *PostfixNegatedComparisonConditionContext) {
	if ctx.FieldList() == nil || ctx.ComparisonOp() == nil || ctx.Value() == nil {
		return
	}
	fields := extractFieldNames(ctx.FieldList())
	op := extractComparisonOp(ctx.ComparisonOp())
	val := extractValue(ctx.Value())

	for i, f := range fields {
		logOp := e.currentLogicalOp
		if i > 0 {
			logOp = "OR"
		}
		e.conditions = append(e.conditions, Condition{
			Field:     f,
			Operator:  op,
			Value:     val,
			Negated:   true,
			LogicalOp: logOp,
		})
	}
	e.currentLogicalOp = "AND"
}

func (e *conditionExtractor) visitPostfixNegatedStringCondition(ctx *PostfixNegatedStringConditionContext) {
	if ctx.FieldList() == nil || ctx.StringOp() == nil || ctx.Value() == nil {
		return
	}
	fields := extractFieldNames(ctx.FieldList())
	op := extractStringOp(ctx.StringOp())
	val := extractValue(ctx.Value())

	for i, f := range fields {
		logOp := e.currentLogicalOp
		if i > 0 {
			logOp = "OR"
		}
		e.conditions = append(e.conditions, Condition{
			Field:     f,
			Operator:  op,
			Value:     val,
			Negated:   true,
			LogicalOp: logOp,
		})
	}
	e.currentLogicalOp = "AND"
}

func (e *conditionExtractor) visitPostfixNegatedSetCondition(ctx *PostfixNegatedSetConditionContext) {
	if ctx.FieldList() == nil || ctx.SetOp() == nil || ctx.ValueList() == nil {
		return
	}
	fields := extractFieldNames(ctx.FieldList())
	op := extractSetOp(ctx.SetOp())
	alternatives := extractValueList(ctx.ValueList())

	for _, f := range fields {
		e.conditions = append(e.conditions, Condition{
			Field:        f,
			Operator:     op,
			Value:        strings.Join(alternatives, ", "),
			Negated:      true,
			LogicalOp:    e.currentLogicalOp,
			Alternatives: alternatives,
		})
	}
	e.currentLogicalOp = "AND"
}

func (e *conditionExtractor) visitPostfixNegatedListStringCondition(ctx *PostfixNegatedListStringConditionContext) {
	if ctx.FieldList() == nil || ctx.ListStringOp() == nil || ctx.ValueList() == nil {
		return
	}
	fields := extractFieldNames(ctx.FieldList())
	op := extractListStringOp(ctx.ListStringOp())
	alternatives := extractValueList(ctx.ValueList())

	for _, f := range fields {
		e.conditions = append(e.conditions, Condition{
			Field:        f,
			Operator:     op,
			Value:        strings.Join(alternatives, ", "),
			Negated:      true,
			LogicalOp:    e.currentLogicalOp,
			Alternatives: alternatives,
		})
	}
	e.currentLogicalOp = "AND"
}

func (e *conditionExtractor) visitNocaseCondition(ctx *NocaseConditionContext) {
	if ctx.FieldList() == nil || ctx.ComparisonOp() == nil || ctx.NocaseValue() == nil {
		return
	}
	fields := extractFieldNames(ctx.FieldList())
	op := extractComparisonOp(ctx.ComparisonOp())

	// Unwrap NOCASE(value) to get the inner value
	val := ""
	nocaseCtx := ctx.NocaseValue()
	if nocaseCtx != nil && nocaseCtx.Value() != nil {
		val = extractValue(nocaseCtx.Value())
	}

	for i, f := range fields {
		logOp := e.currentLogicalOp
		if i > 0 {
			logOp = "OR"
		}
		e.conditions = append(e.conditions, Condition{
			Field:     f,
			Operator:  op,
			Value:     val,
			Negated:   e.negated,
			LogicalOp: logOp,
		})
	}
	e.currentLogicalOp = "AND"
}

func (e *conditionExtractor) visitIpCondition(ctx *IpConditionContext) {
	if ctx.FieldList() == nil || ctx.IpValue() == nil {
		return
	}
	fields := extractFieldNames(ctx.FieldList())

	// Unwrap IP(cidr) to get the CIDR value
	val := extractIPValue(ctx.IpValue())

	for i, f := range fields {
		logOp := e.currentLogicalOp
		if i > 0 {
			logOp = "OR"
		}
		e.conditions = append(e.conditions, Condition{
			Field:     f,
			Operator:  "cidr",
			Value:     val,
			Negated:   e.negated,
			LogicalOp: logOp,
		})
	}
	e.currentLogicalOp = "AND"
}

func (e *conditionExtractor) visitIpSetCondition(ctx *IpSetConditionContext) {
	if ctx.FieldList() == nil || ctx.ValueList() == nil {
		return
	}
	fields := extractFieldNames(ctx.FieldList())
	alternatives := extractValueListWithIP(ctx.ValueList())

	for _, f := range fields {
		e.conditions = append(e.conditions, Condition{
			Field:        f,
			Operator:     "in",
			Value:        strings.Join(alternatives, ", "),
			Negated:      e.negated,
			LogicalOp:    e.currentLogicalOp,
			Alternatives: alternatives,
		})
	}
	e.currentLogicalOp = "AND"
}

func (e *conditionExtractor) visitAllFieldsComparisonCondition(ctx *AllFieldsComparisonConditionContext) {
	if ctx.AllFieldList() == nil || ctx.ComparisonOp() == nil || ctx.Value() == nil {
		return
	}
	fields := extractAllFieldNames(ctx.AllFieldList())
	op := extractComparisonOp(ctx.ComparisonOp())
	val := extractValue(ctx.Value())

	for i, f := range fields {
		logOp := e.currentLogicalOp
		if i > 0 {
			logOp = "AND"
		}
		e.conditions = append(e.conditions, Condition{
			Field:     f,
			Operator:  op,
			Value:     val,
			Negated:   e.negated,
			LogicalOp: logOp,
		})
	}
	e.currentLogicalOp = "AND"
}

func (e *conditionExtractor) visitAllFieldsSetCondition(ctx *AllFieldsSetConditionContext) {
	if ctx.AllFieldList() == nil || ctx.SetOp() == nil || ctx.ValueList() == nil {
		return
	}
	fields := extractAllFieldNames(ctx.AllFieldList())
	op := extractSetOp(ctx.SetOp())
	alternatives := extractValueList(ctx.ValueList())

	for i, f := range fields {
		logOp := e.currentLogicalOp
		if i > 0 {
			logOp = "AND"
		}
		e.conditions = append(e.conditions, Condition{
			Field:        f,
			Operator:     op,
			Value:        strings.Join(alternatives, ", "),
			Negated:      e.negated,
			LogicalOp:    logOp,
			Alternatives: alternatives,
		})
	}
	e.currentLogicalOp = "AND"
}

func (e *conditionExtractor) visitAllFieldsNocaseCondition(ctx *AllFieldsNocaseConditionContext) {
	if ctx.AllFieldList() == nil || ctx.ComparisonOp() == nil || ctx.NocaseValue() == nil {
		return
	}
	fields := extractAllFieldNames(ctx.AllFieldList())
	op := extractComparisonOp(ctx.ComparisonOp())
	val := ""
	if ctx.NocaseValue() != nil && ctx.NocaseValue().Value() != nil {
		val = extractValue(ctx.NocaseValue().Value())
	}

	for i, f := range fields {
		logOp := e.currentLogicalOp
		if i > 0 {
			logOp = "AND"
		}
		e.conditions = append(e.conditions, Condition{
			Field:     f,
			Operator:  op,
			Value:     val,
			Negated:   e.negated,
			LogicalOp: logOp,
		})
	}
	e.currentLogicalOp = "AND"
}

func (e *conditionExtractor) visitFieldExistsCondition(ctx *FieldExistsConditionContext) {
	if ctx.IDENTIFIER() == nil {
		return
	}
	field := ctx.IDENTIFIER().GetText()
	e.addCondition(field, "exists", "*")
}

func (e *conditionExtractor) visitNumberKeywordCondition(ctx *NumberKeywordConditionContext) {
	if ctx.NUMBER() == nil {
		return
	}
	e.addCondition("_keyword", "search", ctx.NUMBER().GetText())
}

// addCondition appends a new condition using current extractor state.
func (e *conditionExtractor) addCondition(field, operator, value string) {
	e.conditions = append(e.conditions, Condition{
		Field:     field,
		Operator:  operator,
		Value:     value,
		Negated:   e.negated,
		LogicalOp: e.currentLogicalOp,
	})
	e.currentLogicalOp = "AND"
}

// --- Field and value extraction helpers ---

// extractFieldNames returns the field names from a fieldList context.
func extractFieldNames(ctx IFieldListContext) []string {
	if ctx == nil {
		return nil
	}
	flCtx, ok := ctx.(*FieldListContext)
	if !ok {
		return nil
	}
	var fields []string
	for _, fn := range flCtx.AllFieldName() {
		fields = append(fields, stripQuotes(fn.GetText()))
	}
	return fields
}

// extractAllFieldNames returns the field names from an allFieldList context (ALL(...)).
func extractAllFieldNames(ctx IAllFieldListContext) []string {
	if ctx == nil {
		return nil
	}
	aflCtx, ok := ctx.(*AllFieldListContext)
	if !ok {
		return nil
	}
	var fields []string
	for _, fn := range aflCtx.AllFieldName() {
		fields = append(fields, stripQuotes(fn.GetText()))
	}
	return fields
}

// extractComparisonOp returns the operator string from a comparisonOp context.
func extractComparisonOp(ctx IComparisonOpContext) string {
	if ctx == nil {
		return ""
	}
	return ctx.GetText()
}

// extractStringOp returns the lowercase operator name from a stringOp context.
func extractStringOp(ctx IStringOpContext) string {
	if ctx == nil {
		return ""
	}
	text := strings.ToLower(ctx.GetText())
	// Normalize hyphenated operators to underscore-free lowercase names
	switch text {
	case "contains":
		return "contains"
	case "icontains":
		return "icontains"
	case "starts-with":
		return "startswith"
	case "istarts-with":
		return "istartswith"
	default:
		return strings.ReplaceAll(text, "-", "")
	}
}

// extractSetOp returns the lowercase operator name from a setOp context.
func extractSetOp(ctx ISetOpContext) string {
	if ctx == nil {
		return ""
	}
	return strings.ToLower(ctx.GetText())
}

// extractListStringOp returns the operator name from a listStringOp context.
func extractListStringOp(ctx IListStringOpContext) string {
	if ctx == nil {
		return ""
	}
	text := strings.ToLower(ctx.GetText())
	// Normalize hyphenated LEQL operators to canonical names
	switch text {
	case "contains-any":
		return "contains_any"
	case "icontains-any":
		return "icontains_any"
	case "contains-all":
		return "contains_all"
	case "icontains-all":
		return "icontains_all"
	case "starts-with-any":
		return "startswith_any"
	case "istarts-with-any":
		return "istartswith_any"
	default:
		return strings.ReplaceAll(text, "-", "_")
	}
}

// extractValue returns the string value from a value context, stripping quotes as needed.
func extractValue(ctx IValueContext) string {
	if ctx == nil {
		return ""
	}
	valCtx, ok := ctx.(*ValueContext)
	if !ok {
		return ctx.GetText()
	}

	if valCtx.DOUBLE_STRING() != nil {
		return stripQuotes(valCtx.DOUBLE_STRING().GetText())
	}
	if valCtx.SINGLE_STRING() != nil {
		return stripQuotes(valCtx.SINGLE_STRING().GetText())
	}
	if valCtx.TRIPLE_SINGLE_STRING() != nil {
		return stripTripleQuotes(valCtx.TRIPLE_SINGLE_STRING().GetText())
	}
	if valCtx.TRIPLE_DOUBLE_STRING() != nil {
		return stripTripleQuotes(valCtx.TRIPLE_DOUBLE_STRING().GetText())
	}
	if valCtx.REGEX() != nil {
		return stripRegexSlashes(valCtx.REGEX().GetText())
	}
	// NUMBER, TIME_UNIT, IDENTIFIER: return raw text
	return valCtx.GetText()
}

// extractValueList extracts string values from a valueList context.
func extractValueList(ctx IValueListContext) []string {
	if ctx == nil {
		return nil
	}
	vlCtx, ok := ctx.(*ValueListContext)
	if !ok {
		return nil
	}
	var values []string
	for _, item := range vlCtx.AllValueListItem() {
		itemCtx, ok := item.(*ValueListItemContext)
		if !ok {
			continue
		}
		if itemCtx.IpValue() != nil {
			values = append(values, extractIPValue(itemCtx.IpValue()))
		} else if itemCtx.Value() != nil {
			values = append(values, extractValue(itemCtx.Value()))
		}
	}
	return values
}

// extractValueListWithIP is similar to extractValueList but specifically for IP set conditions.
// It unwraps IP() wrappers from value list items.
func extractValueListWithIP(ctx IValueListContext) []string {
	return extractValueList(ctx)
}

// extractIPValue unwraps IP(cidr) or IP(value) to get the inner value string.
func extractIPValue(ctx IIpValueContext) string {
	if ctx == nil {
		return ""
	}
	ipCtx, ok := ctx.(*IpValueContext)
	if !ok {
		return ""
	}
	// IP_CIDR token takes priority
	if ipCtx.IP_CIDR() != nil {
		return ipCtx.IP_CIDR().GetText()
	}
	if ipCtx.Value() != nil {
		return extractValue(ipCtx.Value())
	}
	return ""
}

// --- String processing helpers ---

// stripQuotes removes surrounding single or double quotes from a string.
func stripQuotes(s string) string {
	if len(s) >= 2 {
		if (s[0] == '"' && s[len(s)-1] == '"') || (s[0] == '\'' && s[len(s)-1] == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

// stripTripleQuotes removes surrounding triple quotes (''' or """) from a string.
func stripTripleQuotes(s string) string {
	if len(s) >= 6 {
		if strings.HasPrefix(s, `"""`) && strings.HasSuffix(s, `"""`) {
			return s[3 : len(s)-3]
		}
		if strings.HasPrefix(s, "'''") && strings.HasSuffix(s, "'''") {
			return s[3 : len(s)-3]
		}
	}
	return s
}

// stripRegexSlashes removes surrounding /.../ and optional flags from a regex literal.
func stripRegexSlashes(s string) string {
	if len(s) < 2 || s[0] != '/' {
		return s
	}
	// Find the closing slash
	lastSlash := strings.LastIndex(s, "/")
	if lastSlash <= 0 {
		return s
	}
	// Return the pattern between the slashes (flags are discarded for the value)
	return s[1:lastSlash]
}

// --- Exported utility functions ---

// ClassifyFieldProvenance determines where a field originates relative to joins.
// LEQL does not support joins, so this always returns ProvenanceMain.
func ClassifyFieldProvenance(_ *ParseResult, _ string) FieldProvenance {
	return ProvenanceMain
}

// DeduplicateConditions removes duplicate conditions based on Field+Operator+Value.
func DeduplicateConditions(conditions []Condition) []Condition {
	if len(conditions) == 0 {
		return conditions
	}

	seen := make(map[string]bool)
	result := make([]Condition, 0, len(conditions))

	for _, cond := range conditions {
		key := strings.ToLower(cond.Field) + "|" + cond.Operator + "|" + cond.Value
		if !seen[key] {
			seen[key] = true
			result = append(result, cond)
		}
	}

	return result
}

// IsStatisticalQuery returns true if the query contains a calculate clause.
func IsStatisticalQuery(result *ParseResult) bool {
	for _, cmd := range result.Commands {
		if cmd == "calculate" {
			return true
		}
	}
	return false
}

// HasUnmappedComputedFields returns whether any computed field could not be traced back
// to a source field. LEQL does not have eval/extend, so this always returns false.
func HasUnmappedComputedFields(_ *ParseResult) bool {
	return false
}

// HasComplexWhereConditions returns true if any condition uses regex or cidr operators.
func HasComplexWhereConditions(result *ParseResult) bool {
	complexOperators := map[string]bool{
		"regex": true,
		"cidr":  true,
	}

	for _, cond := range result.Conditions {
		if complexOperators[cond.Operator] {
			return true
		}
	}
	return false
}

// GetEventTypeFromConditions detects Windows Event types based on EventID/EventCode conditions.
// Returns event type strings like "windows_4688", "sysmon_1", etc.
func GetEventTypeFromConditions(result *ParseResult) string {
	var eventID string

	for _, cond := range result.Conditions {
		fieldLower := strings.ToLower(cond.Field)

		// Check for EventID or EventCode
		if fieldLower == "eventid" || fieldLower == "eventcode" {
			eventID = cond.Value
		}
	}

	if eventID == "" {
		return ""
	}

	// Map event IDs to event types
	switch eventID {
	case "4688":
		return "windows_4688"
	case "4624":
		return "windows_4624"
	case "4625":
		return "windows_4625"
	case "4720":
		return "windows_4720"
	case "1":
		return "sysmon_1"
	case "3":
		return "sysmon_3"
	}

	return ""
}
