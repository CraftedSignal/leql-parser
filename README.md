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
