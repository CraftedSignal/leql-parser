package leql

import (
	"reflect"
	"testing"
)

// Postfix NOT is Rapid7's documented exclusion form (where(key1, key2 NOT
// STARTS-WITH "aaa")); negating ICONTAINS, ISTARTS-WITH or IENDS-WITH keeps
// the operator case-insensitive in every NOT form.
func TestNegatedCaseInsensitiveStringOperatorsKeepFlag(t *testing.T) {
	for _, op := range []string{"ICONTAINS", "ISTARTS-WITH", "IENDS-WITH", "CONTAINS", "STARTS-WITH", "ENDS-WITH"} {
		want := op[0] == 'I'
		for _, query := range []string{"where(a NOT " + op + " x)", "where(NOT a " + op + " x)"} {
			r := ExtractConditions(query)
			if len(r.Errors) != 0 || len(r.Conditions) != 1 || !r.Conditions[0].Negated || r.Conditions[0].CaseInsensitive != want {
				t.Errorf("%s: conditions=%+v errors=%q, want negated with CaseInsensitive=%v", query, r.Conditions, r.Errors, want)
			}
		}
		// The condition rule's own prefix-NOT alternative.
		r := extractConditionRuleForTest(t, "NOT a "+op+" x")
		if len(r.Errors) != 0 || len(r.Conditions) != 1 || !r.Conditions[0].Negated || r.Conditions[0].CaseInsensitive != want {
			t.Errorf("condition NOT a %s x: conditions=%+v errors=%q, want CaseInsensitive=%v", op, r.Conditions, r.Errors, want)
		}
	}
}

// Rapid7: IIN compares "text values ... in case-insensitive way", and the
// I-prefixed list operators match "case-insensitively". IN, whose name also
// starts with I, stays case-sensitive.
func TestCaseInsensitiveListOperatorsSetFlag(t *testing.T) {
	for _, op := range []string{"IIN", "ICONTAINS-ANY", "ICONTAINS-ALL", "ISTARTS-WITH-ANY", "IENDS-WITH-ANY", "IN", "CONTAINS-ANY", "CONTAINS-ALL", "STARTS-WITH-ANY", "ENDS-WITH-ANY"} {
		want := op != "IN" && op[0] == 'I'
		queries := []string{"where(a " + op + " [x, y])", "where(a,b " + op + " [x, y])", "where(a NOT " + op + " [x, y])", "where(NOT a " + op + " [x, y])"}
		if op == "IIN" || op == "IN" {
			queries = append(queries, "where(ALL(a,b) "+op+" [x, y])")
		}
		for _, query := range queries {
			r := ExtractConditions(query)
			if len(r.Errors) != 0 || len(r.Conditions) == 0 {
				t.Errorf("%s: conditions=%+v errors=%q", query, r.Conditions, r.Errors)
				continue
			}
			for _, c := range r.Conditions {
				if c.CaseInsensitive != want {
					t.Errorf("%s: %+v, want CaseInsensitive=%v", query, c, want)
				}
			}
		}
		// The condition rule's own prefix-NOT alternative.
		r := extractConditionRuleForTest(t, "NOT a "+op+" [x, y]")
		if len(r.Errors) != 0 || len(r.Conditions) != 1 || r.Conditions[0].CaseInsensitive != want {
			t.Errorf("condition NOT a %s [x, y]: conditions=%+v errors=%q, want CaseInsensitive=%v", op, r.Conditions, r.Errors, want)
		}
	}
}

// Rapid7's "Use variables in queries" example excludes private ranges with
// source_ip NOT IN [IP(10.0.0.0/8), ...]. A list of only IP() items is a CIDR
// match in every IN form, as it already was without NOT and with a leading NOT.
func TestIPListsMatchAsCIDRInEveryInForm(t *testing.T) {
	list := `[IP(10.0.0.0/8), IP(192.168.0.0/16)]`
	for _, tt := range []struct {
		query   string
		negated bool
	}{
		{`where(source_ip IN ` + list + `)`, false},
		{`where(NOT source_ip IN ` + list + `)`, true},
		{`where(source_ip NOT IN ` + list + `)`, true},
		{`where(source_ip, destination_ip NOT IN ` + list + `)`, true},
		{`where(ALL(source_ip, destination_ip) IN ` + list + `)`, false},
	} {
		r := ExtractConditions(tt.query)
		if len(r.Errors) != 0 || len(r.Conditions) == 0 || !HasComplexWhereConditions(r) {
			t.Errorf("%s: conditions=%+v errors=%q", tt.query, r.Conditions, r.Errors)
			continue
		}
		for _, c := range r.Conditions {
			if c.Operator != "cidr" || c.Negated != tt.negated {
				t.Errorf("%s: %+v, want operator cidr with negated=%v", tt.query, c, tt.negated)
			}
		}
	}
	if got, want := ExtractConditions(`where(source_ip NOT IN `+list+`)`), ExtractConditions(`where(NOT source_ip IN `+list+`)`); !reflect.DeepEqual(got, want) {
		t.Errorf("postfix NOT IN differs from prefix NOT:\n got %+v\nwant %+v", got, want)
	}
	// The condition rule's own prefix-NOT alternative.
	if r := extractConditionRuleForTest(t, "NOT source_ip IN "+list); len(r.Errors) != 0 || len(r.Conditions) != 1 || r.Conditions[0].Operator != "cidr" || !r.Conditions[0].Negated {
		t.Errorf("condition NOT source_ip IN [IP..]: conditions=%+v errors=%q", r.Conditions, r.Errors)
	}
	// A list that mixes IP() and plain items stays a set match, as without NOT.
	if r := ExtractConditions(`where(source_ip NOT IN [IP(10.0.0.0/8), "x"])`); len(r.Conditions) != 1 || r.Conditions[0].Operator != "in" {
		t.Errorf("mixed list: conditions=%+v errors=%q", r.Conditions, r.Errors)
	}
}
