package leql

import "github.com/antlr4-go/antlr/v4"

// OperandKind describes lexical syntax and the field-reference role of key
// comparisons. It does not describe the runtime type of an event value.
type OperandKind string

const (
	OperandString   OperandKind = "string"
	OperandNumber   OperandKind = "number"
	OperandField    OperandKind = "field"
	OperandRegex    OperandKind = "regex"
	OperandVariable OperandKind = "variable"
	OperandCIDR     OperandKind = "cidr"
	OperandTimeUnit OperandKind = "time_unit"
	OperandList     OperandKind = "list"
)

// OperandWrapper preserves a source wrapper independently of its inner kind.
type OperandWrapper string

const (
	OperandNoCase OperandWrapper = "nocase"
	OperandIP     OperandWrapper = "ip"
)

// Operand preserves a predicate's RHS syntax. Numbers retain their exact text;
// strings use the parser's existing quote stripping without extra unescaping.
// Only list roots have Items, in source order. Nil Condition.RHS means unknown.
type Operand struct {
	Kind       OperandKind    `json:"kind"`
	Value      string         `json:"value,omitempty"`
	RegexFlags string         `json:"regex_flags,omitempty"`
	Wrapper    OperandWrapper `json:"wrapper,omitempty"`
	Items      []Operand      `json:"items,omitempty"`
}

func isKeyComparison(operator string) bool { return operator == "==" || operator == "!==" }

func directFieldOperand(value IValueContext) bool {
	literal, ok := value.(*ValueContext)
	return ok && (literal.IDENTIFIER() != nil || literal.NUMBER() != nil || literal.SINGLE_STRING() != nil || literal.DOUBLE_STRING() != nil)
}

func scalarOperand(value IValueContext, budget *parseBudget) Operand {
	budget.step()
	operand := Operand{Kind: OperandString, Value: extractValue(value)}
	if literal, ok := value.(*ValueContext); ok {
		switch {
		case literal.NUMBER() != nil:
			operand.Kind = OperandNumber
		case literal.REGEX() != nil:
			operand.Kind = OperandRegex
			operand.RegexFlags = regexLiteralFlags(literal.REGEX().GetText())
		case literal.VARIABLE() != nil:
			operand.Kind = OperandVariable
		case literal.IP_CIDR() != nil:
			operand.Kind = OperandCIDR
		case literal.TIME_UNIT() != nil:
			operand.Kind = OperandTimeUnit
		}
	}
	return operand
}

func valueOperand(value IValueContext, operator string, budget *parseBudget) *Operand {
	// Reserve before declaring the operand whose address escapes to the caller.
	budget.reserveOutput(1)
	operand := scalarOperand(value, budget)
	if isKeyComparison(operator) && directFieldOperand(value) {
		operand.Kind = OperandField
	}
	return &operand
}

func ipOperand(value IIpValueContext, budget *parseBudget) Operand {
	budget.step()
	operand := Operand{Kind: OperandString, Wrapper: OperandIP}
	if value != nil {
		if literal := value.IP_CIDR(); literal != nil {
			operand.Kind = OperandCIDR
			operand.Value = literal.GetText()
		} else if value.Value() != nil {
			operand = scalarOperand(value.Value(), budget)
			operand.Wrapper = OperandIP
		}
	}
	return operand
}

func listOperand(value IValueListContext, budget *parseBudget) *Operand {
	budget.step()
	budget.reserveOutput(1)
	operand := &Operand{Kind: OperandList}
	if value == nil {
		return operand
	}
	// Count before allocating; generated AllValueListItem builds its own slice.
	count := 0
	for _, child := range value.GetChildren() {
		budget.step()
		if _, ok := child.(IValueListItemContext); ok {
			count++
		}
	}
	budget.reserveOutput(count)
	operand.Items = make([]Operand, 0, count)
	for _, child := range value.GetChildren() {
		budget.step()
		item, ok := child.(IValueListItemContext)
		if !ok {
			continue
		}
		if item.IpValue() != nil {
			operand.Items = append(operand.Items, ipOperand(item.IpValue(), budget))
		} else if item.Value() != nil {
			operand.Items = append(operand.Items, scalarOperand(item.Value(), budget))
		}
	}
	return operand
}

func cloneOperand(operand *Operand, budget *parseBudget) *Operand {
	budget.step()
	if operand == nil {
		return nil
	}
	budget.reserveOutput(1)
	clone := *operand
	clone.Items = cloneOperandItems(operand.Items, budget)
	return &clone
}

func cloneOperandItems(items []Operand, budget *parseBudget) []Operand {
	budget.step()
	if items == nil {
		return nil
	}
	budget.reserveOutput(len(items))
	out := make([]Operand, len(items))
	for i, item := range items {
		budget.step()
		out[i] = item
		out[i].Items = cloneOperandItems(item.Items, budget)
	}
	return out
}

func operandAlternatives(operand *Operand, budget *parseBudget) []string {
	budget.step()
	if operand == nil || operand.Kind != OperandList {
		return nil
	}
	budget.reserveOutput(len(operand.Items))
	out := make([]string, len(operand.Items))
	for i, item := range operand.Items {
		budget.step()
		out[i] = item.Value
	}
	return out
}

// A comparison predicate owns its entire operand, including nested regexes.
// Diagnose unsupported key syntax once before the generic regex walk descends.
func keyPredicateOperand(tree antlr.Tree, budget *parseBudget) (string, string, bool, bool) {
	comparison, ok := tree.(interface{ ComparisonOp() IComparisonOpContext })
	if !ok || comparison.ComparisonOp() == nil {
		return "", "", false, false
	}
	switch tree.(type) {
	case interface{ FieldList() IFieldListContext }, interface{ AllFieldList() IAllFieldListContext }:
	default:
		return "", "", false, false
	}
	op := comparison.ComparisonOp().GetText()
	if !isKeyComparison(op) {
		return "", "", false, false
	}
	if direct, ok := tree.(interface{ Value() IValueContext }); ok && direct.Value() != nil {
		value := direct.Value()
		kind := string(scalarOperand(value, budget).Kind)
		if value.TRIPLE_SINGLE_STRING() != nil || value.TRIPLE_DOUBLE_STRING() != nil {
			kind = "triple-quoted string"
		}
		return op, kind, directFieldOperand(value), true
	}
	if wrapped, ok := tree.(interface{ NocaseValue() INocaseValueContext }); ok && wrapped.NocaseValue() != nil {
		return op, "nocase wrapper", false, true
	}
	if wrapped, ok := tree.(interface{ IpValue() IIpValueContext }); ok && wrapped.IpValue() != nil {
		return op, "ip wrapper", false, true
	}
	return op, "recovered value", false, true
}
