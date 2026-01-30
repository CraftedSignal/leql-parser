// Code generated from LEQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package leql // LEQLParser
import "github.com/antlr4-go/antlr/v4"

// BaseLEQLParserListener is a complete listener for a parse tree produced by LEQLParser.
type BaseLEQLParserListener struct{}

var _ LEQLParserListener = &BaseLEQLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLEQLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLEQLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLEQLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLEQLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseLEQLParserListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseLEQLParserListener) ExitQuery(ctx *QueryContext) {}

// EnterSelectClause is called when production selectClause is entered.
func (s *BaseLEQLParserListener) EnterSelectClause(ctx *SelectClauseContext) {}

// ExitSelectClause is called when production selectClause is exited.
func (s *BaseLEQLParserListener) ExitSelectClause(ctx *SelectClauseContext) {}

// EnterSelectFieldList is called when production selectFieldList is entered.
func (s *BaseLEQLParserListener) EnterSelectFieldList(ctx *SelectFieldListContext) {}

// ExitSelectFieldList is called when production selectFieldList is exited.
func (s *BaseLEQLParserListener) ExitSelectFieldList(ctx *SelectFieldListContext) {}

// EnterSelectField is called when production selectField is entered.
func (s *BaseLEQLParserListener) EnterSelectField(ctx *SelectFieldContext) {}

// ExitSelectField is called when production selectField is exited.
func (s *BaseLEQLParserListener) ExitSelectField(ctx *SelectFieldContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseLEQLParserListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseLEQLParserListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterNotExpr is called when production notExpr is entered.
func (s *BaseLEQLParserListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production notExpr is exited.
func (s *BaseLEQLParserListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterConditionExpr is called when production conditionExpr is entered.
func (s *BaseLEQLParserListener) EnterConditionExpr(ctx *ConditionExprContext) {}

// ExitConditionExpr is called when production conditionExpr is exited.
func (s *BaseLEQLParserListener) ExitConditionExpr(ctx *ConditionExprContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BaseLEQLParserListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BaseLEQLParserListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterRegexExpr is called when production regexExpr is entered.
func (s *BaseLEQLParserListener) EnterRegexExpr(ctx *RegexExprContext) {}

// ExitRegexExpr is called when production regexExpr is exited.
func (s *BaseLEQLParserListener) ExitRegexExpr(ctx *RegexExprContext) {}

// EnterKeywordExpr is called when production keywordExpr is entered.
func (s *BaseLEQLParserListener) EnterKeywordExpr(ctx *KeywordExprContext) {}

// ExitKeywordExpr is called when production keywordExpr is exited.
func (s *BaseLEQLParserListener) ExitKeywordExpr(ctx *KeywordExprContext) {}

// EnterParenExpr is called when production parenExpr is entered.
func (s *BaseLEQLParserListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production parenExpr is exited.
func (s *BaseLEQLParserListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BaseLEQLParserListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BaseLEQLParserListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterComparisonCondition is called when production comparisonCondition is entered.
func (s *BaseLEQLParserListener) EnterComparisonCondition(ctx *ComparisonConditionContext) {}

// ExitComparisonCondition is called when production comparisonCondition is exited.
func (s *BaseLEQLParserListener) ExitComparisonCondition(ctx *ComparisonConditionContext) {}

// EnterStringCondition is called when production stringCondition is entered.
func (s *BaseLEQLParserListener) EnterStringCondition(ctx *StringConditionContext) {}

// ExitStringCondition is called when production stringCondition is exited.
func (s *BaseLEQLParserListener) ExitStringCondition(ctx *StringConditionContext) {}

// EnterSetCondition is called when production setCondition is entered.
func (s *BaseLEQLParserListener) EnterSetCondition(ctx *SetConditionContext) {}

// ExitSetCondition is called when production setCondition is exited.
func (s *BaseLEQLParserListener) ExitSetCondition(ctx *SetConditionContext) {}

// EnterListStringCondition is called when production listStringCondition is entered.
func (s *BaseLEQLParserListener) EnterListStringCondition(ctx *ListStringConditionContext) {}

// ExitListStringCondition is called when production listStringCondition is exited.
func (s *BaseLEQLParserListener) ExitListStringCondition(ctx *ListStringConditionContext) {}

// EnterNegatedComparisonCondition is called when production negatedComparisonCondition is entered.
func (s *BaseLEQLParserListener) EnterNegatedComparisonCondition(ctx *NegatedComparisonConditionContext) {
}

// ExitNegatedComparisonCondition is called when production negatedComparisonCondition is exited.
func (s *BaseLEQLParserListener) ExitNegatedComparisonCondition(ctx *NegatedComparisonConditionContext) {
}

// EnterNegatedStringCondition is called when production negatedStringCondition is entered.
func (s *BaseLEQLParserListener) EnterNegatedStringCondition(ctx *NegatedStringConditionContext) {}

// ExitNegatedStringCondition is called when production negatedStringCondition is exited.
func (s *BaseLEQLParserListener) ExitNegatedStringCondition(ctx *NegatedStringConditionContext) {}

// EnterNegatedSetCondition is called when production negatedSetCondition is entered.
func (s *BaseLEQLParserListener) EnterNegatedSetCondition(ctx *NegatedSetConditionContext) {}

// ExitNegatedSetCondition is called when production negatedSetCondition is exited.
func (s *BaseLEQLParserListener) ExitNegatedSetCondition(ctx *NegatedSetConditionContext) {}

// EnterNegatedListStringCondition is called when production negatedListStringCondition is entered.
func (s *BaseLEQLParserListener) EnterNegatedListStringCondition(ctx *NegatedListStringConditionContext) {
}

// ExitNegatedListStringCondition is called when production negatedListStringCondition is exited.
func (s *BaseLEQLParserListener) ExitNegatedListStringCondition(ctx *NegatedListStringConditionContext) {
}

// EnterNocaseCondition is called when production nocaseCondition is entered.
func (s *BaseLEQLParserListener) EnterNocaseCondition(ctx *NocaseConditionContext) {}

// ExitNocaseCondition is called when production nocaseCondition is exited.
func (s *BaseLEQLParserListener) ExitNocaseCondition(ctx *NocaseConditionContext) {}

// EnterIpCondition is called when production ipCondition is entered.
func (s *BaseLEQLParserListener) EnterIpCondition(ctx *IpConditionContext) {}

// ExitIpCondition is called when production ipCondition is exited.
func (s *BaseLEQLParserListener) ExitIpCondition(ctx *IpConditionContext) {}

// EnterIpSetCondition is called when production ipSetCondition is entered.
func (s *BaseLEQLParserListener) EnterIpSetCondition(ctx *IpSetConditionContext) {}

// ExitIpSetCondition is called when production ipSetCondition is exited.
func (s *BaseLEQLParserListener) ExitIpSetCondition(ctx *IpSetConditionContext) {}

// EnterAllFieldsComparisonCondition is called when production allFieldsComparisonCondition is entered.
func (s *BaseLEQLParserListener) EnterAllFieldsComparisonCondition(ctx *AllFieldsComparisonConditionContext) {
}

// ExitAllFieldsComparisonCondition is called when production allFieldsComparisonCondition is exited.
func (s *BaseLEQLParserListener) ExitAllFieldsComparisonCondition(ctx *AllFieldsComparisonConditionContext) {
}

// EnterAllFieldsSetCondition is called when production allFieldsSetCondition is entered.
func (s *BaseLEQLParserListener) EnterAllFieldsSetCondition(ctx *AllFieldsSetConditionContext) {}

// ExitAllFieldsSetCondition is called when production allFieldsSetCondition is exited.
func (s *BaseLEQLParserListener) ExitAllFieldsSetCondition(ctx *AllFieldsSetConditionContext) {}

// EnterFieldExistsCondition is called when production fieldExistsCondition is entered.
func (s *BaseLEQLParserListener) EnterFieldExistsCondition(ctx *FieldExistsConditionContext) {}

// ExitFieldExistsCondition is called when production fieldExistsCondition is exited.
func (s *BaseLEQLParserListener) ExitFieldExistsCondition(ctx *FieldExistsConditionContext) {}

// EnterQuotedKeyword is called when production quotedKeyword is entered.
func (s *BaseLEQLParserListener) EnterQuotedKeyword(ctx *QuotedKeywordContext) {}

// ExitQuotedKeyword is called when production quotedKeyword is exited.
func (s *BaseLEQLParserListener) ExitQuotedKeyword(ctx *QuotedKeywordContext) {}

// EnterSingleQuotedKeyword is called when production singleQuotedKeyword is entered.
func (s *BaseLEQLParserListener) EnterSingleQuotedKeyword(ctx *SingleQuotedKeywordContext) {}

// ExitSingleQuotedKeyword is called when production singleQuotedKeyword is exited.
func (s *BaseLEQLParserListener) ExitSingleQuotedKeyword(ctx *SingleQuotedKeywordContext) {}

// EnterRegexPattern is called when production regexPattern is entered.
func (s *BaseLEQLParserListener) EnterRegexPattern(ctx *RegexPatternContext) {}

// ExitRegexPattern is called when production regexPattern is exited.
func (s *BaseLEQLParserListener) ExitRegexPattern(ctx *RegexPatternContext) {}

// EnterFieldList is called when production fieldList is entered.
func (s *BaseLEQLParserListener) EnterFieldList(ctx *FieldListContext) {}

// ExitFieldList is called when production fieldList is exited.
func (s *BaseLEQLParserListener) ExitFieldList(ctx *FieldListContext) {}

// EnterAllFieldList is called when production allFieldList is entered.
func (s *BaseLEQLParserListener) EnterAllFieldList(ctx *AllFieldListContext) {}

// ExitAllFieldList is called when production allFieldList is exited.
func (s *BaseLEQLParserListener) ExitAllFieldList(ctx *AllFieldListContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseLEQLParserListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseLEQLParserListener) ExitFieldName(ctx *FieldNameContext) {}

// EnterComparisonOp is called when production comparisonOp is entered.
func (s *BaseLEQLParserListener) EnterComparisonOp(ctx *ComparisonOpContext) {}

// ExitComparisonOp is called when production comparisonOp is exited.
func (s *BaseLEQLParserListener) ExitComparisonOp(ctx *ComparisonOpContext) {}

// EnterStringOp is called when production stringOp is entered.
func (s *BaseLEQLParserListener) EnterStringOp(ctx *StringOpContext) {}

// ExitStringOp is called when production stringOp is exited.
func (s *BaseLEQLParserListener) ExitStringOp(ctx *StringOpContext) {}

// EnterSetOp is called when production setOp is entered.
func (s *BaseLEQLParserListener) EnterSetOp(ctx *SetOpContext) {}

// ExitSetOp is called when production setOp is exited.
func (s *BaseLEQLParserListener) ExitSetOp(ctx *SetOpContext) {}

// EnterListStringOp is called when production listStringOp is entered.
func (s *BaseLEQLParserListener) EnterListStringOp(ctx *ListStringOpContext) {}

// ExitListStringOp is called when production listStringOp is exited.
func (s *BaseLEQLParserListener) ExitListStringOp(ctx *ListStringOpContext) {}

// EnterValue is called when production value is entered.
func (s *BaseLEQLParserListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseLEQLParserListener) ExitValue(ctx *ValueContext) {}

// EnterNocaseValue is called when production nocaseValue is entered.
func (s *BaseLEQLParserListener) EnterNocaseValue(ctx *NocaseValueContext) {}

// ExitNocaseValue is called when production nocaseValue is exited.
func (s *BaseLEQLParserListener) ExitNocaseValue(ctx *NocaseValueContext) {}

// EnterIpValue is called when production ipValue is entered.
func (s *BaseLEQLParserListener) EnterIpValue(ctx *IpValueContext) {}

// ExitIpValue is called when production ipValue is exited.
func (s *BaseLEQLParserListener) ExitIpValue(ctx *IpValueContext) {}

// EnterValueList is called when production valueList is entered.
func (s *BaseLEQLParserListener) EnterValueList(ctx *ValueListContext) {}

// ExitValueList is called when production valueList is exited.
func (s *BaseLEQLParserListener) ExitValueList(ctx *ValueListContext) {}

// EnterValueListItem is called when production valueListItem is entered.
func (s *BaseLEQLParserListener) EnterValueListItem(ctx *ValueListItemContext) {}

// ExitValueListItem is called when production valueListItem is exited.
func (s *BaseLEQLParserListener) ExitValueListItem(ctx *ValueListItemContext) {}

// EnterGroupbyClause is called when production groupbyClause is entered.
func (s *BaseLEQLParserListener) EnterGroupbyClause(ctx *GroupbyClauseContext) {}

// ExitGroupbyClause is called when production groupbyClause is exited.
func (s *BaseLEQLParserListener) ExitGroupbyClause(ctx *GroupbyClauseContext) {}

// EnterHavingClause is called when production havingClause is entered.
func (s *BaseLEQLParserListener) EnterHavingClause(ctx *HavingClauseContext) {}

// ExitHavingClause is called when production havingClause is exited.
func (s *BaseLEQLParserListener) ExitHavingClause(ctx *HavingClauseContext) {}

// EnterHavingCondition is called when production havingCondition is entered.
func (s *BaseLEQLParserListener) EnterHavingCondition(ctx *HavingConditionContext) {}

// ExitHavingCondition is called when production havingCondition is exited.
func (s *BaseLEQLParserListener) ExitHavingCondition(ctx *HavingConditionContext) {}

// EnterCalculateClause is called when production calculateClause is entered.
func (s *BaseLEQLParserListener) EnterCalculateClause(ctx *CalculateClauseContext) {}

// ExitCalculateClause is called when production calculateClause is exited.
func (s *BaseLEQLParserListener) ExitCalculateClause(ctx *CalculateClauseContext) {}

// EnterCalcExpr is called when production calcExpr is entered.
func (s *BaseLEQLParserListener) EnterCalcExpr(ctx *CalcExprContext) {}

// ExitCalcExpr is called when production calcExpr is exited.
func (s *BaseLEQLParserListener) ExitCalcExpr(ctx *CalcExprContext) {}

// EnterCalcFunction is called when production calcFunction is entered.
func (s *BaseLEQLParserListener) EnterCalcFunction(ctx *CalcFunctionContext) {}

// ExitCalcFunction is called when production calcFunction is exited.
func (s *BaseLEQLParserListener) ExitCalcFunction(ctx *CalcFunctionContext) {}

// EnterCalcFunctionWithField is called when production calcFunctionWithField is entered.
func (s *BaseLEQLParserListener) EnterCalcFunctionWithField(ctx *CalcFunctionWithFieldContext) {}

// ExitCalcFunctionWithField is called when production calcFunctionWithField is exited.
func (s *BaseLEQLParserListener) ExitCalcFunctionWithField(ctx *CalcFunctionWithFieldContext) {}

// EnterPercentileFunction is called when production percentileFunction is entered.
func (s *BaseLEQLParserListener) EnterPercentileFunction(ctx *PercentileFunctionContext) {}

// ExitPercentileFunction is called when production percentileFunction is exited.
func (s *BaseLEQLParserListener) ExitPercentileFunction(ctx *PercentileFunctionContext) {}

// EnterSortClause is called when production sortClause is entered.
func (s *BaseLEQLParserListener) EnterSortClause(ctx *SortClauseContext) {}

// ExitSortClause is called when production sortClause is exited.
func (s *BaseLEQLParserListener) ExitSortClause(ctx *SortClauseContext) {}

// EnterSortDirection is called when production sortDirection is entered.
func (s *BaseLEQLParserListener) EnterSortDirection(ctx *SortDirectionContext) {}

// ExitSortDirection is called when production sortDirection is exited.
func (s *BaseLEQLParserListener) ExitSortDirection(ctx *SortDirectionContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BaseLEQLParserListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BaseLEQLParserListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterTimesliceClause is called when production timesliceClause is entered.
func (s *BaseLEQLParserListener) EnterTimesliceClause(ctx *TimesliceClauseContext) {}

// ExitTimesliceClause is called when production timesliceClause is exited.
func (s *BaseLEQLParserListener) ExitTimesliceClause(ctx *TimesliceClauseContext) {}
