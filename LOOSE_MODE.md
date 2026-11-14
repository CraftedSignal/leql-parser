# Loose WHERE mode

`ExtractConditions`, `ExtractConditionsContext`, `AnalyzeQuery`, and
`AnalyzeQueryContext` preserve the trailing `where(..., loose)` modifier as
`Loose` on `ParseResult` and `SemanticQuery`. Each `Condition` produced inside
that WHERE scope also has `Loose: true`, including conditions expanded from
multi-field predicates and nested expressions. FROM conditions keep `Loose: false`.
Deduplication retains strict and loose variants of the same predicate.

The JSON property is `"loose": true`; false values are omitted. The Loose fields
are additive for JSON consumers and keyed Go struct literals. External unkeyed
literals may need updating. Typed RHS metadata separately adds JSON properties
to strict-query field predicates as described below.

The condition flag records scope provenance. Operators, values, alternatives,
regex flags, case sensitivity flags, negation, and boolean grouping retain their
existing representation. For example, `where(user=admin, loose)` still has
operator `=` and value `admin`, together with the new mode metadata. Consumers
must honor that metadata when interpreting loose matching.

[Rapid7's comparison documentation](https://docs.rapid7.com/insightidr/components-for-building-a-query/)
describes partial, case-insensitive string matching in loose mode and directional
containment for key comparisons. This parser preserves the modifier rather than
implementing universal executable loose semantics. The documentation does not
establish how loose affects every numeric, regex, list, wrapper, or explicit
string operator. Those predicates receive provenance without operator lowering.

Typed RHS metadata distinguishes quoted numeric strings from numbers and records
key-reference roles, alongside the legacy Value/Alternatives text projections.
Consumers that discard `Loose` still lose the user's matching intent.

Typed RHS syntax and key-reference roles are preserved alongside Loose metadata.
See [typed operands](TYPED_OPERANDS.md) for the API and downstream adoption work.
