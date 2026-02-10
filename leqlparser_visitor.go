// Code generated from LEQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package leql // LEQLParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by LEQLParser.
type LEQLParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LEQLParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by LEQLParser#selectClause.
	VisitSelectClause(ctx *SelectClauseContext) interface{}

	// Visit a parse tree produced by LEQLParser#selectFieldList.
	VisitSelectFieldList(ctx *SelectFieldListContext) interface{}

	// Visit a parse tree produced by LEQLParser#selectField.
	VisitSelectField(ctx *SelectFieldContext) interface{}

	// Visit a parse tree produced by LEQLParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by LEQLParser#notExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by LEQLParser#conditionExpr.
	VisitConditionExpr(ctx *ConditionExprContext) interface{}

	// Visit a parse tree produced by LEQLParser#orExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by LEQLParser#regexExpr.
	VisitRegexExpr(ctx *RegexExprContext) interface{}

	// Visit a parse tree produced by LEQLParser#nestedWhereExpr.
	VisitNestedWhereExpr(ctx *NestedWhereExprContext) interface{}

	// Visit a parse tree produced by LEQLParser#keywordExpr.
	VisitKeywordExpr(ctx *KeywordExprContext) interface{}

	// Visit a parse tree produced by LEQLParser#parenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by LEQLParser#andExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by LEQLParser#comparisonCondition.
	VisitComparisonCondition(ctx *ComparisonConditionContext) interface{}

	// Visit a parse tree produced by LEQLParser#stringCondition.
	VisitStringCondition(ctx *StringConditionContext) interface{}

	// Visit a parse tree produced by LEQLParser#setCondition.
	VisitSetCondition(ctx *SetConditionContext) interface{}

	// Visit a parse tree produced by LEQLParser#listStringCondition.
	VisitListStringCondition(ctx *ListStringConditionContext) interface{}

	// Visit a parse tree produced by LEQLParser#negatedComparisonCondition.
	VisitNegatedComparisonCondition(ctx *NegatedComparisonConditionContext) interface{}

	// Visit a parse tree produced by LEQLParser#negatedStringCondition.
	VisitNegatedStringCondition(ctx *NegatedStringConditionContext) interface{}

	// Visit a parse tree produced by LEQLParser#negatedSetCondition.
	VisitNegatedSetCondition(ctx *NegatedSetConditionContext) interface{}

	// Visit a parse tree produced by LEQLParser#negatedListStringCondition.
	VisitNegatedListStringCondition(ctx *NegatedListStringConditionContext) interface{}

	// Visit a parse tree produced by LEQLParser#postfixNegatedComparisonCondition.
	VisitPostfixNegatedComparisonCondition(ctx *PostfixNegatedComparisonConditionContext) interface{}

	// Visit a parse tree produced by LEQLParser#postfixNegatedStringCondition.
	VisitPostfixNegatedStringCondition(ctx *PostfixNegatedStringConditionContext) interface{}

	// Visit a parse tree produced by LEQLParser#postfixNegatedSetCondition.
	VisitPostfixNegatedSetCondition(ctx *PostfixNegatedSetConditionContext) interface{}

	// Visit a parse tree produced by LEQLParser#postfixNegatedListStringCondition.
	VisitPostfixNegatedListStringCondition(ctx *PostfixNegatedListStringConditionContext) interface{}

	// Visit a parse tree produced by LEQLParser#nocaseCondition.
	VisitNocaseCondition(ctx *NocaseConditionContext) interface{}

	// Visit a parse tree produced by LEQLParser#ipCondition.
	VisitIpCondition(ctx *IpConditionContext) interface{}

	// Visit a parse tree produced by LEQLParser#ipSetCondition.
	VisitIpSetCondition(ctx *IpSetConditionContext) interface{}

	// Visit a parse tree produced by LEQLParser#allFieldsComparisonCondition.
	VisitAllFieldsComparisonCondition(ctx *AllFieldsComparisonConditionContext) interface{}

	// Visit a parse tree produced by LEQLParser#allFieldsSetCondition.
	VisitAllFieldsSetCondition(ctx *AllFieldsSetConditionContext) interface{}

	// Visit a parse tree produced by LEQLParser#allFieldsNocaseCondition.
	VisitAllFieldsNocaseCondition(ctx *AllFieldsNocaseConditionContext) interface{}

	// Visit a parse tree produced by LEQLParser#fieldExistsCondition.
	VisitFieldExistsCondition(ctx *FieldExistsConditionContext) interface{}

	// Visit a parse tree produced by LEQLParser#numberKeywordCondition.
	VisitNumberKeywordCondition(ctx *NumberKeywordConditionContext) interface{}

	// Visit a parse tree produced by LEQLParser#quotedKeyword.
	VisitQuotedKeyword(ctx *QuotedKeywordContext) interface{}

	// Visit a parse tree produced by LEQLParser#singleQuotedKeyword.
	VisitSingleQuotedKeyword(ctx *SingleQuotedKeywordContext) interface{}

	// Visit a parse tree produced by LEQLParser#regexPattern.
	VisitRegexPattern(ctx *RegexPatternContext) interface{}

	// Visit a parse tree produced by LEQLParser#fieldList.
	VisitFieldList(ctx *FieldListContext) interface{}

	// Visit a parse tree produced by LEQLParser#allFieldList.
	VisitAllFieldList(ctx *AllFieldListContext) interface{}

	// Visit a parse tree produced by LEQLParser#fieldName.
	VisitFieldName(ctx *FieldNameContext) interface{}

	// Visit a parse tree produced by LEQLParser#comparisonOp.
	VisitComparisonOp(ctx *ComparisonOpContext) interface{}

	// Visit a parse tree produced by LEQLParser#stringOp.
	VisitStringOp(ctx *StringOpContext) interface{}

	// Visit a parse tree produced by LEQLParser#setOp.
	VisitSetOp(ctx *SetOpContext) interface{}

	// Visit a parse tree produced by LEQLParser#listStringOp.
	VisitListStringOp(ctx *ListStringOpContext) interface{}

	// Visit a parse tree produced by LEQLParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by LEQLParser#nocaseValue.
	VisitNocaseValue(ctx *NocaseValueContext) interface{}

	// Visit a parse tree produced by LEQLParser#ipValue.
	VisitIpValue(ctx *IpValueContext) interface{}

	// Visit a parse tree produced by LEQLParser#valueList.
	VisitValueList(ctx *ValueListContext) interface{}

	// Visit a parse tree produced by LEQLParser#valueListItem.
	VisitValueListItem(ctx *ValueListItemContext) interface{}

	// Visit a parse tree produced by LEQLParser#groupbyClause.
	VisitGroupbyClause(ctx *GroupbyClauseContext) interface{}

	// Visit a parse tree produced by LEQLParser#havingClause.
	VisitHavingClause(ctx *HavingClauseContext) interface{}

	// Visit a parse tree produced by LEQLParser#havingCondition.
	VisitHavingCondition(ctx *HavingConditionContext) interface{}

	// Visit a parse tree produced by LEQLParser#calculateClause.
	VisitCalculateClause(ctx *CalculateClauseContext) interface{}

	// Visit a parse tree produced by LEQLParser#calcExpr.
	VisitCalcExpr(ctx *CalcExprContext) interface{}

	// Visit a parse tree produced by LEQLParser#calcFunction.
	VisitCalcFunction(ctx *CalcFunctionContext) interface{}

	// Visit a parse tree produced by LEQLParser#calcFunctionWithField.
	VisitCalcFunctionWithField(ctx *CalcFunctionWithFieldContext) interface{}

	// Visit a parse tree produced by LEQLParser#percentileFunction.
	VisitPercentileFunction(ctx *PercentileFunctionContext) interface{}

	// Visit a parse tree produced by LEQLParser#sortClause.
	VisitSortClause(ctx *SortClauseContext) interface{}

	// Visit a parse tree produced by LEQLParser#sortSpec.
	VisitSortSpec(ctx *SortSpecContext) interface{}

	// Visit a parse tree produced by LEQLParser#sortDirection.
	VisitSortDirection(ctx *SortDirectionContext) interface{}

	// Visit a parse tree produced by LEQLParser#limitClause.
	VisitLimitClause(ctx *LimitClauseContext) interface{}

	// Visit a parse tree produced by LEQLParser#timesliceClause.
	VisitTimesliceClause(ctx *TimesliceClauseContext) interface{}
}
