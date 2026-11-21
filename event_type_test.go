package leql

import "testing"

// A rule that excludes an event ID is not a rule for that event. A rule that
// matches several event IDs has no single event type: Rapid7 documents IN as
// "a shortcut for multiple ORs", and EventID IN [4688, 4624] already gives "".
func TestEventTypeOnlyFromPositiveEventIDMatches(t *testing.T) {
	for _, tt := range []struct{ query, want string }{
		{`where(EventID=4688)`, "windows_4688"},
		{`where(EventID IN [4688])`, "windows_4688"},
		{`where(EventCode=4624 OR EventID=4624)`, "windows_4624"},
		{`where(EventID=4688 AND NOT EventID=4624)`, "windows_4688"},
		{`where(EventID!=4688)`, ""},
		{`where(NOT EventID=4688)`, ""},
		{`where(EventID NOT IN [4688])`, ""},
		{`where(EventID>4688)`, ""},
		{`where(EventID==4688)`, ""},
		{`where(EventID IN [4688, 4624])`, ""},
		{`where(EventID=4688 OR EventID=4624)`, ""},
	} {
		if got := GetEventTypeFromConditions(ExtractConditions(tt.query)); got != tt.want {
			t.Errorf("%s: event type %q, want %q", tt.query, got, tt.want)
		}
	}
}
