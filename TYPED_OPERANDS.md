# Typed right-hand operands

Field predicates returned by `ExtractConditions`, `AnalyzeQuery`, and their
context variants include `Condition.RHS *Operand`. Existing `Value`,
`Alternatives`, operator, negation, and Loose metadata remain available. Bare
keyword, number, variable, and regex searches have no syntactic RHS and leave
this pointer nil. Manually constructed legacy conditions may also leave it nil;
nil means unknown operand metadata, not a string value.

`Operand.Kind` records source syntax and, for key comparisons, the reference
role. It does not assert an event's runtime value type. The kinds are `string`,
`number`, `field`, `regex`, `variable`, `cidr`, `time_unit`, and `list`.

| Source RHS | Typed representation |
| --- | --- |
| `42`, `-1e-03`, `9007199254740993` | `number`, exact lexical text in `Value` |
| `"42"`, `'42'`, bare `null`, `true`, `false` | `string`, existing extracted text |
| `b` or `"b"` after `==` / `!==` | `field`, field name `b` |
| `/x/mi` | `regex`, pattern `x`, `RegexFlags: "mi"` |
| `${name}` | `variable`, unchanged `${name}` |
| `10.0.0.0/8`, `30m` | `cidr` and `time_unit`, unchanged lexical text |
| `IP("10.0.0.0/8")` | `string`, `Wrapper: "ip"` |
| `IP(10.0.0.0/8)` | `cidr`, `Wrapper: "ip"` |
| `NOCASE(42)` | `number`, `Wrapper: "nocase"` |
| `[42,"42",/x/i]` | `list`, ordered typed `Items` |

Numbers are never converted through floating point or normalized. There are no
null or boolean operand kinds. Strings retain the existing quote stripping and
backslash content; this API does not add unescaping, quote-style preservation,
case folding, numeric range handling, or an evaluator. List roots have only
`Kind` and `Items`; each item's `Value` matches its legacy `Alternatives` entry.
Only the grammar's existing one-level lists are parsed. Direct regex list items
are supported under IN/IIN, including negation and ALL-field expansion. Flags
remain on each item; IIN does not add `i` to a regex. Other regex/operator and
regex-wrapper restrictions remain in force.

## Key comparisons and diagnostics

`==` and `!==` compare keys. Direct identifiers, NUMBER tokens, and single- or
double-quoted RHS tokens become `field` operands. In particular, `a==42` and
`a=="42"` both reference a field named `42`; `a=42` is a numeric literal and
`a="42"` is a string literal. Existing acceptance of quoted and numeric names is
a parser grammar compatibility choice. Bare keywords absent from the existing
value rule, such as `loose`, are not newly accepted.

Other grammar-accepted RHS forms produce one diagnostic per source predicate
starting with `unsupported key-comparison RHS for ==: expected field name` (or
`!==`). Best-effort conditions and boolean trees preserve the key operator,
negation, Loose scope, original typed operand, and legacy inner value. Direct
regex flags remain in both RHS and the legacy `RegexFlags` field. Invalid key
operands do not derive `CaseInsensitive`; flags nested inside wrappers remain
only in RHS. They are not lowered to regex or CIDR operators.

These diagnostics describe this parser's supported key model. They do not make
a claim about all undocumented dynamic-field forms. Consumers must treat
`Errors` as non-executable input; best-effort extraction is not a guarantee for
consumers that ignore diagnostics. Normal syntax recovery remains unchanged.
The grammar has no postfix-NOT wrapper production or ALL-fields IP wrapper
production, so those combinations continue to return syntax diagnostics.
Cancellation, limit exhaustion, and unexpected panic return failure-only results
without conditions or expressions.

## Compatibility and downstream adoption

The Go field and JSON `rhs` property are additive. Nil RHS is omitted from JSON,
but parser-produced field predicates now include it, so their JSON is no longer
byte-identical to older strict-query output. For example:

```json
{"kind":"number","value":"9007199254740993"}
```

The numeric value is a JSON string. An empty string is encoded as
`{"kind":"string"}`. External unkeyed `Condition` literals can require changes
when this new struct field is added; keyed literals remain compatible.
Semantic conversion deep-copies RHS and its items, and expanded conditions own
separate mutable operands. Deduplication includes lexical kind, reference role,
wrapper, ordered items, flags, and nil-versus-known metadata.

Parser metadata alone does not implement downstream translation or evaluation.
Consumers that currently read only Value/Alternatives must adopt RHS explicitly:
preserve numeric spelling until a supported target conversion, keep quoted
numbers as strings, translate `field` references with correct `==`/`!==`
semantics, retain mixed list kinds and Loose behavior, and collect RHS field
references where appropriate. Targets that cannot represent an operation must
report that limitation. Do not store RHS references in `SourceField`, which
retains its computed/original-field provenance meaning. Existing platform
adapters and dependency releases are separate adoption work.

See [parsing limits](PARSING_LIMITS.md) for cumulative metadata accounting and
[Loose mode](LOOSE_MODE.md) for scope and evaluation limits.
