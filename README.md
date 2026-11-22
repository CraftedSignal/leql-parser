# LEQL parser

The parser extracts conditions and boolean structure from LEQL queries. See
[PARSING_LIMITS.md](PARSING_LIMITS.md) for cooperative limits and cancellation.

## Scoped known-field parsing

Bare words keep keyword-search meaning by default, including dotted words such
as `evil.exe`. Callers with trusted schema knowledge can opt in per call:

```go
result := leql.ExtractConditionsWithOptions(query, leql.Options{
    KnownFields: []string{"process.name", "process.parent_name", "ParentImage"},
})
```

Known names match case-insensitively and are copied into a bounded per-call set.
A matching bare identifier becomes a field-presence predicate. A direct implicit
pair such as `process.name /ok$/i` becomes one fielded regex predicate. The parser
keeps its flags, negation, boolean grouping and loose-WHERE provenance. Unknown
or quoted words never acquire that shorthand. An intervening `AND`/`OR`,
parenthesis or `NOT` prevents pairing; surrounding AND terms retain their order
and boolean structure.

`ExtractConditionsContextWithOptions`, `AnalyzeQueryWithOptions` and
`AnalyzeQueryContextWithOptions` expose the same scoped behavior. Zero options
leave the existing APIs unchanged. Option processing shares the parse's
cancellation, work and output budgets: at most 4,096 supplied names and 128 KiB
of cumulative name bytes are accepted, before normalization or set allocation.
Names also consume output items, leaving the remainder for extraction and
semantic conversion. Rejection and cancellation return errors without partial
conditions; these checks do not add a hard ANTLR preemption guarantee.

## Typed comparison operands

Unambiguous bare-key RHS operands in `<`, `>`, `<=` and `>=` comparisons carry
`OperandField`,
as documented in Rapid7's [Tips and Tricks comparison table](https://docs.rapid7.com/insightidr/tips-and-tricks-for-building-queries/).
Numeric RHS operands keep `OperandNumber`; quoted operands keep `OperandString`.
Digit-leading or sign-followed-digit bare text (such as `100bytes`, `10MB`,
`100ms` and `-10MB`) retains its existing lexical role. The grammar admits these
as identifiers too, so this narrow correction does not universally resolve
digit-leading key names versus values.
The existing `==`/`!==` key-comparison rules and ordinary `=` literal handling
are unchanged. Legacy condition fields, value text and provenance helpers keep
their existing contract. This metadata describes source roles, not backend value
types, coercion, missing/null behavior or query execution. Rapid7's separate
[Components table](https://docs.rapid7.com/insightidr/components-for-building-a-query/)
describes ordering using numeric examples; this narrow bare-key role follows the
explicit key examples in Tips rather than expanding quoted or numeric key roles.
