package leql

import "testing"

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
