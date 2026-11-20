package leql

import (
	"reflect"
	"testing"
)

// Rapid7's comparison operators "support values that include numbers with
// units" (value<100bytes). Unquoted values that start with a digit stay one
// string operand instead of splitting into a number and an implicit keyword.
func TestDigitLeadingBareValuesStayWhole(t *testing.T) {
	for _, tt := range []struct{ query, field, op, value string }{
		{`where(value<100bytes)`, "value", "<", "100bytes"},
		{`where(size>=10MB)`, "size", ">=", "10MB"},
		{`where(duration<100ms)`, "duration", "<", "100ms"},
		{`where(md5=0cc175b9c0f1b6a831c399e269772661)`, "md5", "=", "0cc175b9c0f1b6a831c399e269772661"},
		{`where(date=2023-01-01)`, "date", "=", "2023-01-01"},
	} {
		t.Run(tt.query, func(t *testing.T) {
			r := ExtractConditions(tt.query)
			if len(r.Errors) != 0 || len(r.Conditions) != 1 {
				t.Fatalf("conditions=%+v errors=%q", r.Conditions, r.Errors)
			}
			c := r.Conditions[0]
			if c.Field != tt.field || c.Operator != tt.op || c.Value != tt.value || !reflect.DeepEqual(c.RHS, &Operand{Kind: OperandString, Value: tt.value}) {
				t.Fatalf("got %+v rhs=%+v", c, c.RHS)
			}
		})
	}
	// Numbers and time units that the same text could also spell keep their kinds.
	for _, tt := range []struct {
		source string
		kind   OperandKind
	}{{"1e3", OperandNumber}, {"3.1415", OperandNumber}, {"-5", OperandNumber}, {"30m", OperandTimeUnit}} {
		t.Run(tt.source, func(t *testing.T) {
			requireTypedConditions(t, "where(n="+tt.source+")", []Operand{{Kind: tt.kind, Value: tt.source}}, "0")
		})
	}
}

// A bare dotted IPv4 address stays one value, and the rest of the WHERE is kept.
func TestBareIPv4ValueStaysWhole(t *testing.T) {
	r := ExtractConditions(`where(source_ip=192.168.1.1 AND user=admin)`)
	if len(r.Errors) != 0 || len(r.Conditions) != 2 {
		t.Fatalf("conditions=%+v errors=%q", r.Conditions, r.Errors)
	}
	ip, user := r.Conditions[0], r.Conditions[1]
	if ip.Field != "source_ip" || ip.Operator != "=" || ip.Value != "192.168.1.1" || !reflect.DeepEqual(ip.RHS, &Operand{Kind: OperandString, Value: "192.168.1.1"}) {
		t.Fatalf("address condition=%+v rhs=%+v", ip, ip.RHS)
	}
	if user.Field != "user" || user.Operator != "=" || user.Value != "admin" || user.LogicalOp != "AND" {
		t.Fatalf("second condition=%+v", user)
	}
	if got := literalParseShape(r.Expression); got != "and(0,1)" {
		t.Fatalf("expression=%s", got)
	}
}
