// Code generated from LEQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package leql // LEQLParser
import "github.com/antlr4-go/antlr/v4"

type BaseLEQLParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLEQLParserVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitSelectClause(ctx *SelectClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitSelectFieldList(ctx *SelectFieldListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitSelectField(ctx *SelectFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitWhereClause(ctx *WhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitConditionExpr(ctx *ConditionExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitRegexExpr(ctx *RegexExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitNestedWhereExpr(ctx *NestedWhereExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitKeywordExpr(ctx *KeywordExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitComparisonCondition(ctx *ComparisonConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitStringCondition(ctx *StringConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitSetCondition(ctx *SetConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitListStringCondition(ctx *ListStringConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitNegatedComparisonCondition(ctx *NegatedComparisonConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitNegatedStringCondition(ctx *NegatedStringConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitNegatedSetCondition(ctx *NegatedSetConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitNegatedListStringCondition(ctx *NegatedListStringConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitPostfixNegatedComparisonCondition(ctx *PostfixNegatedComparisonConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitPostfixNegatedStringCondition(ctx *PostfixNegatedStringConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitPostfixNegatedSetCondition(ctx *PostfixNegatedSetConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitPostfixNegatedListStringCondition(ctx *PostfixNegatedListStringConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitNocaseCondition(ctx *NocaseConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitIpCondition(ctx *IpConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitIpSetCondition(ctx *IpSetConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitAllFieldsComparisonCondition(ctx *AllFieldsComparisonConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitAllFieldsSetCondition(ctx *AllFieldsSetConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitFieldExistsCondition(ctx *FieldExistsConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitQuotedKeyword(ctx *QuotedKeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitSingleQuotedKeyword(ctx *SingleQuotedKeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitRegexPattern(ctx *RegexPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitFieldList(ctx *FieldListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitAllFieldList(ctx *AllFieldListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitFieldName(ctx *FieldNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitComparisonOp(ctx *ComparisonOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitStringOp(ctx *StringOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitSetOp(ctx *SetOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitListStringOp(ctx *ListStringOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitNocaseValue(ctx *NocaseValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitIpValue(ctx *IpValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitValueList(ctx *ValueListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitValueListItem(ctx *ValueListItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitGroupbyClause(ctx *GroupbyClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitHavingClause(ctx *HavingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitHavingCondition(ctx *HavingConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitCalculateClause(ctx *CalculateClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitCalcExpr(ctx *CalcExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitCalcFunction(ctx *CalcFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitCalcFunctionWithField(ctx *CalcFunctionWithFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitPercentileFunction(ctx *PercentileFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitSortClause(ctx *SortClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitSortSpec(ctx *SortSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitSortDirection(ctx *SortDirectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitLimitClause(ctx *LimitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLEQLParserVisitor) VisitTimesliceClause(ctx *TimesliceClauseContext) interface{} {
	return v.VisitChildren(ctx)
}
