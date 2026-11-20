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
