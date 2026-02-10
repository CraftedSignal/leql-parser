// Code generated from LEQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package leql // LEQLParser
import "github.com/antlr4-go/antlr/v4"

// LEQLParserListener is a complete listener for a parse tree produced by LEQLParser.
type LEQLParserListener interface {
	antlr.ParseTreeListener

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterSelectClause is called when entering the selectClause production.
	EnterSelectClause(c *SelectClauseContext)

	// EnterSelectFieldList is called when entering the selectFieldList production.
	EnterSelectFieldList(c *SelectFieldListContext)

	// EnterSelectField is called when entering the selectField production.
	EnterSelectField(c *SelectFieldContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterNotExpr is called when entering the notExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterConditionExpr is called when entering the conditionExpr production.
	EnterConditionExpr(c *ConditionExprContext)

	// EnterOrExpr is called when entering the orExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterRegexExpr is called when entering the regexExpr production.
	EnterRegexExpr(c *RegexExprContext)

	// EnterNestedWhereExpr is called when entering the nestedWhereExpr production.
	EnterNestedWhereExpr(c *NestedWhereExprContext)

	// EnterKeywordExpr is called when entering the keywordExpr production.
	EnterKeywordExpr(c *KeywordExprContext)

	// EnterParenExpr is called when entering the parenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterAndExpr is called when entering the andExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterComparisonCondition is called when entering the comparisonCondition production.
	EnterComparisonCondition(c *ComparisonConditionContext)

	// EnterStringCondition is called when entering the stringCondition production.
	EnterStringCondition(c *StringConditionContext)

	// EnterSetCondition is called when entering the setCondition production.
	EnterSetCondition(c *SetConditionContext)

	// EnterListStringCondition is called when entering the listStringCondition production.
	EnterListStringCondition(c *ListStringConditionContext)

	// EnterNegatedComparisonCondition is called when entering the negatedComparisonCondition production.
	EnterNegatedComparisonCondition(c *NegatedComparisonConditionContext)

	// EnterNegatedStringCondition is called when entering the negatedStringCondition production.
	EnterNegatedStringCondition(c *NegatedStringConditionContext)

	// EnterNegatedSetCondition is called when entering the negatedSetCondition production.
	EnterNegatedSetCondition(c *NegatedSetConditionContext)

	// EnterNegatedListStringCondition is called when entering the negatedListStringCondition production.
	EnterNegatedListStringCondition(c *NegatedListStringConditionContext)

	// EnterPostfixNegatedComparisonCondition is called when entering the postfixNegatedComparisonCondition production.
	EnterPostfixNegatedComparisonCondition(c *PostfixNegatedComparisonConditionContext)

	// EnterPostfixNegatedStringCondition is called when entering the postfixNegatedStringCondition production.
	EnterPostfixNegatedStringCondition(c *PostfixNegatedStringConditionContext)

	// EnterPostfixNegatedSetCondition is called when entering the postfixNegatedSetCondition production.
	EnterPostfixNegatedSetCondition(c *PostfixNegatedSetConditionContext)

	// EnterPostfixNegatedListStringCondition is called when entering the postfixNegatedListStringCondition production.
	EnterPostfixNegatedListStringCondition(c *PostfixNegatedListStringConditionContext)

	// EnterNocaseCondition is called when entering the nocaseCondition production.
	EnterNocaseCondition(c *NocaseConditionContext)

	// EnterIpCondition is called when entering the ipCondition production.
	EnterIpCondition(c *IpConditionContext)

	// EnterIpSetCondition is called when entering the ipSetCondition production.
	EnterIpSetCondition(c *IpSetConditionContext)

	// EnterAllFieldsComparisonCondition is called when entering the allFieldsComparisonCondition production.
	EnterAllFieldsComparisonCondition(c *AllFieldsComparisonConditionContext)

	// EnterAllFieldsSetCondition is called when entering the allFieldsSetCondition production.
	EnterAllFieldsSetCondition(c *AllFieldsSetConditionContext)

	// EnterFieldExistsCondition is called when entering the fieldExistsCondition production.
	EnterFieldExistsCondition(c *FieldExistsConditionContext)

	// EnterQuotedKeyword is called when entering the quotedKeyword production.
	EnterQuotedKeyword(c *QuotedKeywordContext)

	// EnterSingleQuotedKeyword is called when entering the singleQuotedKeyword production.
	EnterSingleQuotedKeyword(c *SingleQuotedKeywordContext)

	// EnterRegexPattern is called when entering the regexPattern production.
	EnterRegexPattern(c *RegexPatternContext)

	// EnterFieldList is called when entering the fieldList production.
	EnterFieldList(c *FieldListContext)

	// EnterAllFieldList is called when entering the allFieldList production.
	EnterAllFieldList(c *AllFieldListContext)

	// EnterFieldName is called when entering the fieldName production.
	EnterFieldName(c *FieldNameContext)

	// EnterComparisonOp is called when entering the comparisonOp production.
	EnterComparisonOp(c *ComparisonOpContext)

	// EnterStringOp is called when entering the stringOp production.
	EnterStringOp(c *StringOpContext)

	// EnterSetOp is called when entering the setOp production.
	EnterSetOp(c *SetOpContext)

	// EnterListStringOp is called when entering the listStringOp production.
	EnterListStringOp(c *ListStringOpContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterNocaseValue is called when entering the nocaseValue production.
	EnterNocaseValue(c *NocaseValueContext)

	// EnterIpValue is called when entering the ipValue production.
	EnterIpValue(c *IpValueContext)

	// EnterValueList is called when entering the valueList production.
	EnterValueList(c *ValueListContext)

	// EnterValueListItem is called when entering the valueListItem production.
	EnterValueListItem(c *ValueListItemContext)

	// EnterGroupbyClause is called when entering the groupbyClause production.
	EnterGroupbyClause(c *GroupbyClauseContext)

	// EnterHavingClause is called when entering the havingClause production.
	EnterHavingClause(c *HavingClauseContext)

	// EnterHavingCondition is called when entering the havingCondition production.
	EnterHavingCondition(c *HavingConditionContext)

	// EnterCalculateClause is called when entering the calculateClause production.
	EnterCalculateClause(c *CalculateClauseContext)

	// EnterCalcExpr is called when entering the calcExpr production.
	EnterCalcExpr(c *CalcExprContext)

	// EnterCalcFunction is called when entering the calcFunction production.
	EnterCalcFunction(c *CalcFunctionContext)

	// EnterCalcFunctionWithField is called when entering the calcFunctionWithField production.
	EnterCalcFunctionWithField(c *CalcFunctionWithFieldContext)

	// EnterPercentileFunction is called when entering the percentileFunction production.
	EnterPercentileFunction(c *PercentileFunctionContext)

	// EnterSortClause is called when entering the sortClause production.
	EnterSortClause(c *SortClauseContext)

	// EnterSortSpec is called when entering the sortSpec production.
	EnterSortSpec(c *SortSpecContext)

	// EnterSortDirection is called when entering the sortDirection production.
	EnterSortDirection(c *SortDirectionContext)

	// EnterLimitClause is called when entering the limitClause production.
	EnterLimitClause(c *LimitClauseContext)

	// EnterTimesliceClause is called when entering the timesliceClause production.
	EnterTimesliceClause(c *TimesliceClauseContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitSelectClause is called when exiting the selectClause production.
	ExitSelectClause(c *SelectClauseContext)

	// ExitSelectFieldList is called when exiting the selectFieldList production.
	ExitSelectFieldList(c *SelectFieldListContext)

	// ExitSelectField is called when exiting the selectField production.
	ExitSelectField(c *SelectFieldContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitNotExpr is called when exiting the notExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitConditionExpr is called when exiting the conditionExpr production.
	ExitConditionExpr(c *ConditionExprContext)

	// ExitOrExpr is called when exiting the orExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitRegexExpr is called when exiting the regexExpr production.
	ExitRegexExpr(c *RegexExprContext)

	// ExitNestedWhereExpr is called when exiting the nestedWhereExpr production.
	ExitNestedWhereExpr(c *NestedWhereExprContext)

	// ExitKeywordExpr is called when exiting the keywordExpr production.
	ExitKeywordExpr(c *KeywordExprContext)

	// ExitParenExpr is called when exiting the parenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitAndExpr is called when exiting the andExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitComparisonCondition is called when exiting the comparisonCondition production.
	ExitComparisonCondition(c *ComparisonConditionContext)

	// ExitStringCondition is called when exiting the stringCondition production.
	ExitStringCondition(c *StringConditionContext)

	// ExitSetCondition is called when exiting the setCondition production.
	ExitSetCondition(c *SetConditionContext)

	// ExitListStringCondition is called when exiting the listStringCondition production.
	ExitListStringCondition(c *ListStringConditionContext)

	// ExitNegatedComparisonCondition is called when exiting the negatedComparisonCondition production.
	ExitNegatedComparisonCondition(c *NegatedComparisonConditionContext)

	// ExitNegatedStringCondition is called when exiting the negatedStringCondition production.
	ExitNegatedStringCondition(c *NegatedStringConditionContext)

	// ExitNegatedSetCondition is called when exiting the negatedSetCondition production.
	ExitNegatedSetCondition(c *NegatedSetConditionContext)

	// ExitNegatedListStringCondition is called when exiting the negatedListStringCondition production.
	ExitNegatedListStringCondition(c *NegatedListStringConditionContext)

	// ExitPostfixNegatedComparisonCondition is called when exiting the postfixNegatedComparisonCondition production.
	ExitPostfixNegatedComparisonCondition(c *PostfixNegatedComparisonConditionContext)

	// ExitPostfixNegatedStringCondition is called when exiting the postfixNegatedStringCondition production.
	ExitPostfixNegatedStringCondition(c *PostfixNegatedStringConditionContext)

	// ExitPostfixNegatedSetCondition is called when exiting the postfixNegatedSetCondition production.
	ExitPostfixNegatedSetCondition(c *PostfixNegatedSetConditionContext)

	// ExitPostfixNegatedListStringCondition is called when exiting the postfixNegatedListStringCondition production.
	ExitPostfixNegatedListStringCondition(c *PostfixNegatedListStringConditionContext)

	// ExitNocaseCondition is called when exiting the nocaseCondition production.
	ExitNocaseCondition(c *NocaseConditionContext)

	// ExitIpCondition is called when exiting the ipCondition production.
	ExitIpCondition(c *IpConditionContext)

	// ExitIpSetCondition is called when exiting the ipSetCondition production.
	ExitIpSetCondition(c *IpSetConditionContext)

	// ExitAllFieldsComparisonCondition is called when exiting the allFieldsComparisonCondition production.
	ExitAllFieldsComparisonCondition(c *AllFieldsComparisonConditionContext)

	// ExitAllFieldsSetCondition is called when exiting the allFieldsSetCondition production.
	ExitAllFieldsSetCondition(c *AllFieldsSetConditionContext)

	// ExitFieldExistsCondition is called when exiting the fieldExistsCondition production.
	ExitFieldExistsCondition(c *FieldExistsConditionContext)

	// ExitQuotedKeyword is called when exiting the quotedKeyword production.
	ExitQuotedKeyword(c *QuotedKeywordContext)

	// ExitSingleQuotedKeyword is called when exiting the singleQuotedKeyword production.
	ExitSingleQuotedKeyword(c *SingleQuotedKeywordContext)

	// ExitRegexPattern is called when exiting the regexPattern production.
	ExitRegexPattern(c *RegexPatternContext)

	// ExitFieldList is called when exiting the fieldList production.
	ExitFieldList(c *FieldListContext)

	// ExitAllFieldList is called when exiting the allFieldList production.
	ExitAllFieldList(c *AllFieldListContext)

	// ExitFieldName is called when exiting the fieldName production.
	ExitFieldName(c *FieldNameContext)

	// ExitComparisonOp is called when exiting the comparisonOp production.
	ExitComparisonOp(c *ComparisonOpContext)

	// ExitStringOp is called when exiting the stringOp production.
	ExitStringOp(c *StringOpContext)

	// ExitSetOp is called when exiting the setOp production.
	ExitSetOp(c *SetOpContext)

	// ExitListStringOp is called when exiting the listStringOp production.
	ExitListStringOp(c *ListStringOpContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitNocaseValue is called when exiting the nocaseValue production.
	ExitNocaseValue(c *NocaseValueContext)

	// ExitIpValue is called when exiting the ipValue production.
	ExitIpValue(c *IpValueContext)

	// ExitValueList is called when exiting the valueList production.
	ExitValueList(c *ValueListContext)

	// ExitValueListItem is called when exiting the valueListItem production.
	ExitValueListItem(c *ValueListItemContext)

	// ExitGroupbyClause is called when exiting the groupbyClause production.
	ExitGroupbyClause(c *GroupbyClauseContext)

	// ExitHavingClause is called when exiting the havingClause production.
	ExitHavingClause(c *HavingClauseContext)

	// ExitHavingCondition is called when exiting the havingCondition production.
	ExitHavingCondition(c *HavingConditionContext)

	// ExitCalculateClause is called when exiting the calculateClause production.
	ExitCalculateClause(c *CalculateClauseContext)

	// ExitCalcExpr is called when exiting the calcExpr production.
	ExitCalcExpr(c *CalcExprContext)

	// ExitCalcFunction is called when exiting the calcFunction production.
	ExitCalcFunction(c *CalcFunctionContext)

	// ExitCalcFunctionWithField is called when exiting the calcFunctionWithField production.
	ExitCalcFunctionWithField(c *CalcFunctionWithFieldContext)

	// ExitPercentileFunction is called when exiting the percentileFunction production.
	ExitPercentileFunction(c *PercentileFunctionContext)

	// ExitSortClause is called when exiting the sortClause production.
	ExitSortClause(c *SortClauseContext)

	// ExitSortSpec is called when exiting the sortSpec production.
	ExitSortSpec(c *SortSpecContext)

	// ExitSortDirection is called when exiting the sortDirection production.
	ExitSortDirection(c *SortDirectionContext)

	// ExitLimitClause is called when exiting the limitClause production.
	ExitLimitClause(c *LimitClauseContext)

	// ExitTimesliceClause is called when exiting the timesliceClause production.
	ExitTimesliceClause(c *TimesliceClauseContext)
}
