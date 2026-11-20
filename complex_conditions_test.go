package leql

import "testing"

// HasComplexWhereConditions reports whether "any condition uses regex or cidr
// operators"; a regex on a field (=~ and !~, also written = /re/ and != /re/)
// is a regex condition like the whole-event /re/.
func TestHasComplexWhereConditionsCountsFieldRegex(t *testing.T) {
	for _, tt := range []struct {
		query string
		want  bool
	}{
		{`where(/adm.*/i)`, true},
		{`where(user=/adm.*/i)`, true},
		{`where(user!=/adm.*/)`, true},
		{`where(user =~ /adm.*/)`, true},
		{`where(user !~ /adm.*/)`, true},
		{`where(NOT user=/adm.*/)`, true},
		{`where(user=admin)`, false},
		{`where(user!=admin)`, false},
	} {
		r := ExtractConditions(tt.query)
		if len(r.Errors) != 0 {
			t.Errorf("%s: errors %q", tt.query, r.Errors)
		}
		if got := HasComplexWhereConditions(r); got != tt.want {
			t.Errorf("%s: HasComplexWhereConditions=%v, want %v (conditions %+v)", tt.query, got, tt.want, r.Conditions)
		}
	}
}
