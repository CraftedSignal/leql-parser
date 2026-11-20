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

// Rapid7 ("Exclude results with comparison operators"): "You must use ! before
// these operators to exclude data from your results", e.g. key != ["aaa", "bbb"].
// The list exclusion extracts what the equivalent IN form extracts.
func TestNotEqualsListExcludesEveryValue(t *testing.T) {
	r := ExtractConditions(`where(key != ["aaa", "bbb"])`)
	if len(r.Errors) != 0 || len(r.Conditions) != 1 {
		t.Fatalf("conditions=%+v errors=%q", r.Conditions, r.Errors)
	}
	if c := r.Conditions[0]; c.Field != "key" || c.Operator != "in" || !c.Negated || c.Value != "aaa, bbb" || !reflect.DeepEqual(c.Alternatives, []string{"aaa", "bbb"}) {
		t.Fatalf("got %+v", c)
	}
	for _, tt := range []struct{ query, same string }{
		{`where(key != ["aaa", "bbb"])`, `where(key NOT IN ["aaa", "bbb"])`},
		{`where(key1, key2 != ["aaa", "bbb"])`, `where(key1, key2 NOT IN ["aaa", "bbb"])`},
		{`where(NOT key != [aaa, 42])`, `where(key IN [aaa, 42])`},
		{`where(user=admin AND key != [aaa] OR x=1)`, `where(user=admin AND key NOT IN [aaa] OR x=1)`},
		{`where(source_ip != [IP(10.0.0.0/8), IP(192.168.0.0/16)])`, `where(NOT source_ip IN [IP(10.0.0.0/8), IP(192.168.0.0/16)])`},
	} {
		t.Run(tt.query, func(t *testing.T) {
			got, want := ExtractConditions(tt.query), ExtractConditions(tt.same)
			if len(got.Errors) != 0 || len(got.Conditions) == 0 || !reflect.DeepEqual(got, want) {
				t.Fatalf("got %+v\nwant %+v", got, want)
			}
			if got, want := AnalyzeQuery(tt.query), AnalyzeQuery(tt.same); !reflect.DeepEqual(got, want) {
				t.Fatalf("semantic got %+v\nwant %+v", got, want)
			}
		})
	}
}

// Rapid7's example query "Group changes made to a certain group" compares an
// unquoted value with a leading hyphen.
func TestLeadingHyphenBareValuesStayWhole(t *testing.T) {
	r := ExtractConditions(`where(action IN [MEMBER_ADDED_TO_SECURITY_GROUP, MEMBER_REMOVED_FROM_SECURITY_GROUP]AND group CONTAINS -job-admins)`)
	if len(r.Errors) != 0 || len(r.Conditions) != 2 {
		t.Fatalf("conditions=%+v errors=%q", r.Conditions, r.Errors)
	}
	if c := r.Conditions[1]; c.Field != "group" || c.Operator != "contains" || c.Value != "-job-admins" || !reflect.DeepEqual(c.RHS, &Operand{Kind: OperandString, Value: "-job-admins"}) {
		t.Fatalf("got %+v rhs=%+v", c, c.RHS)
	}
	for _, tt := range []struct {
		source string
		kind   OperandKind
	}{{"--no-sandbox", OperandString}, {"-5MB", OperandString}, {"-5", OperandNumber}, {"-1e-03", OperandNumber}} {
		t.Run(tt.source, func(t *testing.T) {
			requireTypedConditions(t, "where(n="+tt.source+")", []Operand{{Kind: tt.kind, Value: tt.source}}, "0")
		})
	}
}
