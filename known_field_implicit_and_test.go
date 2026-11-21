package leql

import (
	"reflect"
	"testing"
)

// README "Scoped known-field parsing": a direct implicit pair such as
// process.name /ok$/i becomes one fielded regex predicate and "surrounding AND
// terms retain their order and boolean structure". Terms joined to the pair by
// a space are AND terms too, so the pair forms as it does after an explicit AND.
func TestKnownFieldRegexPairAfterImplicitAndTerms(t *testing.T) {
	opts := Options{KnownFields: []string{"process.name", "ParentImage"}}
	r := ExtractConditionsWithOptions(`where(x=1 process.name /ok$/i)`, opts)
	if len(r.Errors) != 0 || len(r.Conditions) != 2 || literalParseShape(r.Expression) != "and(0,1)" {
		t.Fatalf("unexpected extraction: %+v", r)
	}
	if c := r.Conditions[1]; c.Field != "process.name" || c.Operator != "=~" || c.Value != "ok$" || c.RegexFlags != "i" || !c.CaseInsensitive {
		t.Fatalf("pair not formed: %+v", c)
	}
	for _, tt := range []struct{ query, same string }{
		{`where(x=1 process.name /ok$/i)`, `where(x=1 AND process.name /ok$/i)`},
		{`where(foo process.name /ok$/i)`, `where(foo AND process.name /ok$/i)`},
		{`where(a=1 OR b=2 process.name /ok$/i)`, `where(a=1 OR b=2 AND process.name /ok$/i)`},
		{`where(NOT a process.name /ok$/i y=2)`, `where(NOT a AND process.name /ok$/i AND y=2)`},
		{`where(NOT (x=1 process.name /ok$/i), loose)`, `where(NOT (x=1 AND process.name /ok$/i), loose)`},
		{`where(x=1 evil.exe /ok$/i)`, `where(x=1 AND evil.exe /ok$/i)`},
	} {
		t.Run(tt.query, func(t *testing.T) {
			got, want := ExtractConditionsWithOptions(tt.query, opts), ExtractConditionsWithOptions(tt.same, opts)
			if len(got.Errors) != 0 || !reflect.DeepEqual(got, want) {
				t.Fatalf("got %+v\nwant %+v", got, want)
			}
			if got, want := AnalyzeQueryWithOptions(tt.query, opts), AnalyzeQueryWithOptions(tt.same, opts); !reflect.DeepEqual(got, want) {
				t.Fatalf("semantic got %+v\nwant %+v", got, want)
			}
		})
	}
	// Every adjacent pair forms, whether the pairs are joined by a space or by AND.
	for _, query := range []string{`where(process.name /ok$/i ParentImage /x/)`, `where(process.name /ok$/i AND ParentImage /x/)`} {
		t.Run(query, func(t *testing.T) {
			r := ExtractConditionsWithOptions(query, opts)
			if len(r.Errors) != 0 || len(r.Conditions) != 2 || literalParseShape(r.Expression) != "and(0,1)" {
				t.Fatalf("unexpected extraction: %+v", r)
			}
			for i, field := range []string{"process.name", "ParentImage"} {
				if c := r.Conditions[i]; c.Field != field || c.Operator != "=~" {
					t.Fatalf("pair %d not formed: %+v", i, c)
				}
			}
		})
	}
}
