package main

import (
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	leql "github.com/craftedsignal/leql-parser"
	sigma "github.com/craftedsignal/sigma-parser"
	"gopkg.in/yaml.v3"
)

type generatedCase struct {
	ID       int64               `json:"id"`
	Seed     int64               `json:"seed"`
	Query    string              `json:"query"`
	Expected []expectedCondition `json:"expected"`
}

type expectedCondition struct {
	Field        string   `json:"field"`
	Operator     string   `json:"operator"`
	Value        string   `json:"value,omitempty"`
	Alternatives []string `json:"alternatives,omitempty"`
	Negated      bool     `json:"negated,omitempty"`
}

type queryEntry struct {
	Source string `json:"source"`
	Name   string `json:"name"`
	Query  string `json:"query"`
}

type runResult struct {
	ID        int64       `json:"id"`
	Query     string      `json:"query,omitempty"`
	SigmaYAML string      `json:"sigma_yaml,omitempty"`
	BackLEQL  string      `json:"back_leql,omitempty"`
	Entry     *queryEntry `json:"entry,omitempty"`
	Failure   *failure    `json:"failure,omitempty"`
}

type failure struct {
	Stage    string   `json:"stage"`
	Reason   string   `json:"reason"`
	Errors   []string `json:"errors,omitempty"`
	Expected []string `json:"expected,omitempty"`
	Actual   []string `json:"actual,omitempty"`
	Missing  []string `json:"missing,omitempty"`
	Extra    []string `json:"extra,omitempty"`
}

type summary struct {
	Total              int64         `json:"total"`
	Passed             int64         `json:"passed"`
	Failed             int64         `json:"failed"`
	DuplicateQueries   int64         `json:"duplicate_queries"`
	LEQLParseErrors    int64         `json:"leql_parse_errors"`
	ExpectedMismatch   int64         `json:"expected_mismatch"`
	SigmaParseErrors   int64         `json:"sigma_parse_errors"`
	SigmaMismatch      int64         `json:"sigma_mismatch"`
	BackLEQLParseError int64         `json:"back_leql_parse_errors"`
	BackLEQLMismatch   int64         `json:"back_leql_mismatch"`
	Duration           time.Duration `json:"duration"`
}

type corpusWriter struct {
	file  *os.File
	enc   *json.Encoder
	first bool
}

type failureWriter struct {
	file *os.File
	enc  *json.Encoder
}

type sigmaRule struct {
	Title     string         `yaml:"title"`
	Status    string         `yaml:"status"`
	Logsource sigmaLogsource `yaml:"logsource"`
	Detection map[string]any `yaml:"detection"`
	Fields    []string       `yaml:"fields,omitempty"`
}

type sigmaLogsource struct {
	Category string `yaml:"category,omitempty"`
	Product  string `yaml:"product,omitempty"`
	Service  string `yaml:"service,omitempty"`
}

type sigmaSelection struct {
	Raw     any
	Negated bool
	Name    string
}

type normCondition struct {
	Field   string
	Op      string
	Value   string
	Negated bool
}

var (
	fields = []string{
		"EventID", "EventCode", "event_code", "status", "result", "action", "severity",
		"user", "username", "account", "source_user", "destination_user", "target_user",
		"source_address", "destination_address", "source_ip", "destination_ip",
		"source_port", "destination_port", "process.name", "process.cmd_line",
		"process.exe_file.description", "process.exe_file.product_name", "file_name",
		"file_path", "domain", "url", "query", "service", "protocol", "connection_status",
		"incoming_bytes", "outgoing_bytes", "risk", "geoip_country_code",
	}
	numericFields = []string{"EventID", "EventCode", "event_code", "status", "source_port", "destination_port", "incoming_bytes", "outgoing_bytes", "risk"}
	stringFields  = []string{"action", "severity", "result", "user", "username", "account", "source_user", "destination_user", "target_user", "source_address", "destination_address", "source_ip", "destination_ip", "process.name", "process.cmd_line", "process.exe_file.description", "process.exe_file.product_name", "file_name", "file_path", "domain", "url", "query", "service", "protocol", "connection_status", "geoip_country_code"}
)

func main() {
	var (
		total       = flag.Int64("n", 1_000_000, "number of generated LEQL queries to test")
		seed        = flag.Int64("seed", 4242, "base random seed; use a negative value for crypto-random choices")
		workers     = flag.Int("workers", runtime.NumCPU(), "parallel worker count")
		corpusPath  = flag.String("corpus", "testdata/generated/leql_sigma_roundtrip_corpus.json", "verified generated corpus JSON array path; empty disables writing")
		failPath    = flag.String("failures", "testdata/generated/leql_sigma_roundtrip_failures.jsonl", "failure JSONL path; empty disables writing")
		failLimit   = flag.Int("failure-limit", 1000, "maximum failure records to write")
		progress    = flag.Int64("progress", 10000, "print progress every N completed cases")
		strict      = flag.Bool("strict", false, "exit non-zero if any generated case fails")
		stopOnFirst = flag.Bool("stop-on-first", false, "stop scheduling new work after first failure")
		unique      = flag.Bool("unique", false, "write only unique query strings; keep generating until n unique verified queries are written")
	)
	flag.Parse()

	if *total < 0 {
		fatalf("-n must be non-negative")
	}
	if *workers <= 0 {
		*workers = 1
	}

	cw, err := newCorpusWriter(*corpusPath)
	if err != nil {
		fatalf("open corpus writer: %v", err)
	}
	defer cw.close()

	fw, err := newFailureWriter(*failPath)
	if err != nil {
		fatalf("open failure writer: %v", err)
	}
	defer fw.close()

	start := time.Now()
	s, failureRecords := runGenerated(*seed, *total, *workers, *failLimit, *progress, *stopOnFirst, *unique, cw, fw, start)
	printSummary(s, *corpusPath, *failPath, failureRecords)
	if *strict && s.Failed > 0 {
		os.Exit(1)
	}
}

func runGenerated(seed, target int64, workers, failLimit int, progress int64, stopOnFirst, unique bool, cw *corpusWriter, fw *failureWriter, start time.Time) (summary, int) {
	if unique {
		return runGeneratedUnique(seed, target, workers, failLimit, progress, stopOnFirst, cw, fw, start)
	}
	return runGeneratedFixed(seed, target, workers, failLimit, progress, stopOnFirst, cw, fw, start)
}

func runGeneratedFixed(seed, total int64, workers, failLimit int, progress int64, stopOnFirst bool, cw *corpusWriter, fw *failureWriter, start time.Time) (summary, int) {
	jobs := make(chan int64, workers*2)
	results := make(chan runResult, workers*2)
	done := make(chan struct{})
	var doneOnce sync.Once
	stop := func() { doneOnce.Do(func() { close(done) }) }

	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for id := range jobs {
				select {
				case <-done:
					return
				default:
				}
				results <- processID(seed, id)
			}
		}()
	}

	go func() {
		defer close(jobs)
		for i := int64(0); i < total; i++ {
			select {
			case <-done:
				return
			case jobs <- i:
			}
		}
	}()
	go func() {
		wg.Wait()
		close(results)
	}()

	var s summary
	var failureRecords int
	for result := range results {
		acceptResult(result, &s, &failureRecords, failLimit, cw, fw)
		if result.Failure != nil && stopOnFirst {
			stop()
		}
		if progress > 0 && s.Total%progress == 0 {
			printProgress(s, total, start)
		}
		if stopOnFirst && s.Failed > 0 {
			break
		}
	}
	s.Duration = time.Since(start)
	return s, failureRecords
}

func runGeneratedUnique(seed, target int64, workers, failLimit int, progress int64, stopOnFirst bool, cw *corpusWriter, fw *failureWriter, start time.Time) (summary, int) {
	jobs := make(chan int64, workers*2)
	results := make(chan runResult, workers*2)
	seen := make(map[[32]byte]struct{}, mapCapacityHint(target))

	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for id := range jobs {
				results <- processID(seed, id)
			}
		}()
	}

	var s summary
	var failureRecords int
	nextID := int64(0)
	active := 0
	maxActive := workers * 2
	jobsClosed := false
	closeJobs := func() {
		if !jobsClosed {
			close(jobs)
			jobsClosed = true
		}
	}
	schedule := func() {
		for !jobsClosed && active < maxActive && s.Passed+int64(active) < target {
			jobs <- nextID
			nextID++
			active++
		}
	}

	schedule()
	for active > 0 {
		result := <-results
		active--
		s.Total++
		if result.Failure == nil {
			hash := sha256.Sum256([]byte(result.Query))
			if _, ok := seen[hash]; ok {
				s.DuplicateQueries++
			} else {
				seen[hash] = struct{}{}
				s.Passed++
				if result.Entry != nil {
					if err := cw.write(*result.Entry); err != nil {
						closeJobs()
						fatalf("write corpus: %v", err)
					}
				}
			}
		} else {
			recordFailure(result, &s, &failureRecords, failLimit, fw)
			if stopOnFirst {
				closeJobs()
			}
		}

		if !jobsClosed {
			schedule()
		}
		if s.Passed == target {
			closeJobs()
		}
		if progress > 0 && (s.Total%progress == 0 || s.Passed == target) {
			printProgress(s, target, start)
		}
	}
	closeJobs()
	wg.Wait()
	s.Duration = time.Since(start)
	return s, failureRecords
}

func acceptResult(result runResult, s *summary, failureRecords *int, failLimit int, cw *corpusWriter, fw *failureWriter) {
	s.Total++
	if result.Failure == nil {
		s.Passed++
		if result.Entry != nil {
			if err := cw.write(*result.Entry); err != nil {
				fatalf("write corpus: %v", err)
			}
		}
		return
	}
	recordFailure(result, s, failureRecords, failLimit, fw)
}

func recordFailure(result runResult, s *summary, failureRecords *int, failLimit int, fw *failureWriter) {
	s.Failed++
	countFailureStage(s, result.Failure.Stage)
	if *failureRecords < failLimit {
		if err := fw.write(result); err != nil {
			fatalf("write failure: %v", err)
		}
		(*failureRecords)++
	}
}

func processCase(tc generatedCase) runResult {
	result := runResult{ID: tc.ID, Query: tc.Query}

	leqlResult := leql.ExtractConditions(tc.Query)
	if leqlResult == nil {
		result.Failure = &failure{Stage: "leql_parse", Reason: "ExtractConditions returned nil"}
		return result
	}
	if len(leqlResult.Errors) > 0 {
		result.Failure = &failure{Stage: "leql_parse", Reason: "LEQL parser returned errors", Errors: leqlResult.Errors}
		return result
	}

	expected := normalizeExpected(tc.Expected)
	actualLEQL := normalizeLEQLConditions(flattenLEQLConditions(leqlResult))
	if missing, extra := compareConditionSets(expected, actualLEQL); len(missing) > 0 || len(extra) > 0 {
		result.Failure = &failure{
			Stage:    "expected_mismatch",
			Reason:   "LEQL parser extraction differed from generated semantic oracle",
			Expected: conditionKeys(expected),
			Actual:   conditionKeys(actualLEQL),
			Missing:  conditionKeys(missing),
			Extra:    conditionKeys(extra),
		}
		return result
	}

	sigmaYAML, err := leqlResultToSigmaYAML(tc.ID, leqlResult)
	if err != nil {
		result.Failure = &failure{Stage: "sigma_build", Reason: err.Error()}
		return result
	}
	result.SigmaYAML = sigmaYAML

	sigmaResult := sigma.ExtractConditions(sigmaYAML)
	if sigmaResult == nil {
		result.Failure = &failure{Stage: "sigma_parse", Reason: "Sigma ExtractConditions returned nil"}
		return result
	}
	if len(sigmaResult.Errors) > 0 {
		result.Failure = &failure{Stage: "sigma_parse", Reason: "Sigma parser returned errors", Errors: sigmaResult.Errors}
		return result
	}

	actualSigma := normalizeSigmaConditions(sigmaResult.Conditions)
	if missing, extra := compareConditionSets(actualLEQL, actualSigma); len(missing) > 0 || len(extra) > 0 {
		result.Failure = &failure{
			Stage:    "sigma_mismatch",
			Reason:   "LEQL parse result and Sigma parser result differ",
			Expected: conditionKeys(actualLEQL),
			Actual:   conditionKeys(actualSigma),
			Missing:  conditionKeys(missing),
			Extra:    conditionKeys(extra),
		}
		return result
	}

	backLEQL, err := sigmaResultToLEQL(sigmaResult)
	if err != nil {
		result.Failure = &failure{Stage: "back_leql_build", Reason: err.Error()}
		return result
	}
	result.BackLEQL = backLEQL

	backResult := leql.ExtractConditions(backLEQL)
	if backResult == nil {
		result.Failure = &failure{Stage: "back_leql_parse", Reason: "LEQL parser returned nil for back-converted LEQL"}
		return result
	}
	if len(backResult.Errors) > 0 {
		result.Failure = &failure{Stage: "back_leql_parse", Reason: "LEQL parser returned errors for back-converted LEQL", Errors: backResult.Errors}
		return result
	}
	actualBack := normalizeLEQLConditions(flattenLEQLConditions(backResult))
	if missing, extra := compareConditionSets(actualSigma, actualBack); len(missing) > 0 || len(extra) > 0 {
		result.Failure = &failure{
			Stage:    "back_leql_mismatch",
			Reason:   "Sigma parser result and back-converted LEQL parser result differ",
			Expected: conditionKeys(actualSigma),
			Actual:   conditionKeys(actualBack),
			Missing:  conditionKeys(missing),
			Extra:    conditionKeys(extra),
		}
		return result
	}

	result.Entry = &queryEntry{
		Source: "generated_leql_sigma_roundtrip",
		Name:   fmt.Sprintf("generated_%09d", tc.ID),
		Query:  tc.Query,
	}
	return result
}

func processID(seed, id int64) (result runResult) {
	result = runResult{ID: id}
	defer func() {
		if r := recover(); r != nil {
			result.Failure = &failure{Stage: "generator_panic", Reason: fmt.Sprintf("panic while generating or processing case: %v", r)}
		}
	}()
	return processCase(generateCase(seed, id))
}

func generateCase(seed, id int64) generatedCase {
	return generateBoundedSemanticCase(seed, id)
}

func generateSimpleSearch(rng *rand.Rand) (string, []expectedCondition) {
	conds := make([]expectedCondition, 0, rng.Intn(4)+2)
	for i := 0; i < cap(conds); i++ {
		conds = append(conds, randomCondition(rng))
	}
	return fmt.Sprintf("where(%s)", joinExpected(conds, "AND")), conds
}

func generateBooleanSearch(rng *rand.Rand) (string, []expectedCondition) {
	field := oneOf(rng, stringFields...)
	values := uniqueValues(rng, valuesForField(field), rng.Intn(3)+2)
	orConds := make([]expectedCondition, 0, len(values))
	for _, value := range values {
		orConds = append(orConds, expectedCondition{Field: field, Operator: "=", Value: value})
	}
	left := "(" + joinExpected(orConds, "OR") + ")"

	right := randomCondition(rng)
	expected := append([]expectedCondition(nil), orConds...)
	expected = append(expected, right)

	if rng.Intn(4) == 0 {
		for i := range orConds {
			orConds[i].Negated = true
		}
		expected = append([]expectedCondition(nil), orConds...)
		expected = append(expected, right)
		left = "NOT " + left
	}

	query := fmt.Sprintf("where(%s AND %s)", left, renderCondition(right))
	return query, expected
}

func generatePipelineSearch(rng *rand.Rand) (string, []expectedCondition) {
	first := randomCondition(rng)
	second := randomCondition(rng)
	third := randomCondition(rng)
	query := fmt.Sprintf("select(%s, %s, %s) where(%s AND %s AND %s) sort(desc) limit(%d)",
		first.Field, second.Field, third.Field, renderCondition(first), renderCondition(second), renderCondition(third), rng.Intn(500)+1)
	return query, []expectedCondition{first, second, third}
}

func generateStringOperatorSearch(rng *rand.Rand) (string, []expectedCondition) {
	field := oneOf(rng, "process.cmd_line", "process.name", "file_name", "file_path", "url", "domain", "service")
	a := expectedCondition{Field: field, Operator: oneOf(rng, "contains", "icontains", "startswith", "istartswith"), Value: randomValueForField(rng, field)}
	b := expectedCondition{Field: field, Operator: oneOf(rng, "contains", "icontains", "startswith", "istartswith"), Value: randomValueForField(rng, field), Negated: rng.Intn(2) == 0}
	query := fmt.Sprintf("where(%s OR %s)", renderCondition(a), renderCondition(b))
	return query, []expectedCondition{a, b}
}

func generateAggregationSearch(rng *rand.Rand) (string, []expectedCondition) {
	pre := randomCondition(rng)
	groupField := oneOf(rng, "user", "source_address", "process.name", "destination_port", "service")
	query := fmt.Sprintf("where(%s) groupby(%s) calculate(count) sort(desc) limit(%d)",
		renderCondition(pre), groupField, rng.Intn(500)+1)
	return query, []expectedCondition{pre}
}

func generateFormattedSearch(rng *rand.Rand) (string, []expectedCondition) {
	query, expected := generatePipelineSearch(rng)
	query = strings.ReplaceAll(query, " AND ", "\n    AND ")
	query = strings.ReplaceAll(query, " OR ", "\n    OR ")
	query = strings.ReplaceAll(query, ") ", ")\n")
	return query, expected
}

func randomCondition(rng *rand.Rand) expectedCondition {
	field := oneOf(rng, fields...)
	if contains(numericFields, field) {
		return expectedCondition{Field: field, Operator: oneOf(rng, "=", "!=", ">", "<", ">=", "<="), Value: strconv.Itoa(rng.Intn(9000) + 1)}
	}
	if rng.Intn(5) == 0 {
		values := uniqueValues(rng, valuesForField(field), rng.Intn(4)+2)
		return expectedCondition{Field: field, Operator: "in", Value: values[0], Alternatives: values, Negated: rng.Intn(4) == 0}
	}
	op := oneOf(rng, "=", "!=", "contains", "icontains", "startswith", "istartswith")
	return expectedCondition{Field: field, Operator: op, Value: randomValueForField(rng, field)}
}

func renderCondition(cond expectedCondition) string {
	if len(cond.Alternatives) > 0 {
		if cond.Negated {
			return fmt.Sprintf("%s NOT IN [%s]", cond.Field, joinLEQLValues(cond.Alternatives))
		}
		return fmt.Sprintf("%s IN [%s]", cond.Field, joinLEQLValues(cond.Alternatives))
	}
	body := ""
	switch strings.ToLower(cond.Operator) {
	case "contains", "icontains":
		body = fmt.Sprintf("%s %s %s", cond.Field, strings.ToUpper(cond.Operator), formatLEQLValue(cond.Value))
	case "startswith", "istartswith":
		op := "STARTS-WITH"
		if strings.EqualFold(cond.Operator, "istartswith") {
			op = "ISTARTS-WITH"
		}
		body = fmt.Sprintf("%s %s %s", cond.Field, op, formatLEQLValue(cond.Value))
	default:
		body = fmt.Sprintf("%s %s %s", cond.Field, cond.Operator, formatLEQLValue(cond.Value))
	}
	if cond.Negated {
		return "NOT " + body
	}
	return body
}

func joinExpected(conditions []expectedCondition, op string) string {
	parts := make([]string, 0, len(conditions))
	for _, cond := range conditions {
		parts = append(parts, renderCondition(cond))
	}
	return strings.Join(parts, " "+op+" ")
}

func leqlResultToSigmaYAML(id int64, result *leql.ParseResult) (string, error) {
	conditions := flattenLEQLConditions(result)
	if len(conditions) == 0 {
		return "", fmt.Errorf("no LEQL conditions to convert")
	}
	rule := sigmaRule{
		Title:     fmt.Sprintf("Generated LEQL Roundtrip %d", id),
		Status:    "test",
		Logsource: inferLogsource(conditions),
		Detection: make(map[string]any, len(conditions)+1),
	}
	conditionParts := make([]string, 0, len(conditions)*2)
	fieldSet := make(map[string]bool)
	for i, cond := range conditions {
		selection := leqlConditionToSigmaSelection(cond, fmt.Sprintf("selection_%06d", i))
		rule.Detection[selection.Name] = selection.Raw
		if len(conditionParts) > 0 {
			op := strings.ToLower(cond.LogicalOp)
			if op != "or" {
				op = "and"
			}
			conditionParts = append(conditionParts, op)
		}
		ref := selection.Name
		if selection.Negated {
			ref = "not " + ref
		}
		conditionParts = append(conditionParts, ref)
		if cond.Field != "" {
			fieldSet[cond.Field] = true
		}
	}
	rule.Detection["condition"] = strings.Join(conditionParts, " ")
	for field := range fieldSet {
		rule.Fields = append(rule.Fields, field)
	}
	sort.Strings(rule.Fields)
	data, err := yaml.Marshal(rule)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func leqlConditionToSigmaSelection(cond leql.Condition, name string) sigmaSelection {
	field := cond.Field
	op, negated := canonicalLEQLOperator(cond.Operator, cond.Negated)
	values := conditionValues(cond.Value, cond.Alternatives)
	if field == "" || op == "keyword" {
		if len(values) > 1 {
			return sigmaSelection{Name: name, Raw: values, Negated: negated}
		}
		return sigmaSelection{Name: name, Raw: first(values), Negated: negated}
	}
	key := field
	var value any = sigmaValue(values)
	switch op {
	case "=":
	case "contains":
		key += "|contains"
	case "startswith":
		key += "|startswith"
	case "endswith":
		key += "|endswith"
	case ">":
		key += "|gt"
	case ">=":
		key += "|gte"
	case "<":
		key += "|lt"
	case "<=":
		key += "|lte"
	default:
		key += "|" + op
	}
	return sigmaSelection{Name: name, Raw: map[string]any{key: value}, Negated: negated}
}

func sigmaResultToLEQL(result *sigma.ParseResult) (string, error) {
	if result == nil || len(result.Conditions) == 0 {
		return "", fmt.Errorf("no Sigma conditions to convert")
	}
	parts := make([]string, 0, len(result.Conditions)*2)
	for i, cond := range result.Conditions {
		if i > 0 {
			logical := strings.ToLower(cond.LogicalOp)
			if logical != "or" {
				logical = "and"
			}
			parts = append(parts, logical)
		}
		expr := sigmaConditionToLEQL(cond)
		if cond.Negated {
			expr = "NOT " + expr
		}
		parts = append(parts, expr)
	}
	return "where(" + strings.Join(parts, " ") + ")", nil
}

func sigmaConditionToLEQL(cond sigma.Condition) string {
	values := conditionValues(cond.Value, cond.Alternatives)
	if cond.Field == "" || cond.Operator == "keyword" {
		return joinAlternativeExpressions(values, func(v string) string {
			return formatLEQLValue(v)
		})
	}
	switch cond.Operator {
	case "=":
		if len(values) > 1 {
			return fmt.Sprintf("%s IN [%s]", cond.Field, joinLEQLValues(values))
		}
		return fmt.Sprintf("%s = %s", cond.Field, formatLEQLValue(first(values)))
	case "contains":
		return joinAlternativeExpressions(values, func(v string) string {
			return fmt.Sprintf("%s CONTAINS %s", cond.Field, formatLEQLValue(v))
		})
	case "startswith":
		return joinAlternativeExpressions(values, func(v string) string {
			return fmt.Sprintf("%s STARTS-WITH %s", cond.Field, formatLEQLValue(v))
		})
	case "endswith":
		return joinAlternativeExpressions(values, func(v string) string {
			return fmt.Sprintf("%s ENDS-WITH %s", cond.Field, formatLEQLValue(v))
		})
	case ">":
		return joinAlternativeExpressions(values, func(v string) string {
			return fmt.Sprintf("%s > %s", cond.Field, formatLEQLValue(v))
		})
	case ">=":
		return joinAlternativeExpressions(values, func(v string) string {
			return fmt.Sprintf("%s >= %s", cond.Field, formatLEQLValue(v))
		})
	case "<":
		return joinAlternativeExpressions(values, func(v string) string {
			return fmt.Sprintf("%s < %s", cond.Field, formatLEQLValue(v))
		})
	case "<=":
		return joinAlternativeExpressions(values, func(v string) string {
			return fmt.Sprintf("%s <= %s", cond.Field, formatLEQLValue(v))
		})
	default:
		return fmt.Sprintf("%s = %s", cond.Field, formatLEQLValue(first(values)))
	}
}

func flattenLEQLConditions(result *leql.ParseResult) []leql.Condition {
	if result == nil {
		return nil
	}
	conditions := append([]leql.Condition(nil), result.Conditions...)
	return conditions
}

func normalizeExpected(conditions []expectedCondition) []normCondition {
	normalized := make([]normCondition, 0, len(conditions))
	for _, cond := range conditions {
		op, negated := canonicalLEQLOperator(cond.Operator, cond.Negated)
		for _, value := range conditionValues(cond.Value, cond.Alternatives) {
			normalized = append(normalized, normalizeParts(cond.Field, op, value, negated))
		}
	}
	return dedupeNorm(normalized)
}

func normalizeLEQLConditions(conditions []leql.Condition) []normCondition {
	normalized := make([]normCondition, 0, len(conditions))
	for _, cond := range conditions {
		op, negated := canonicalLEQLOperator(cond.Operator, cond.Negated)
		for _, value := range conditionValues(cond.Value, cond.Alternatives) {
			normalized = append(normalized, normalizeParts(cond.Field, op, value, negated))
		}
	}
	return dedupeNorm(normalized)
}

func normalizeSigmaConditions(conditions []sigma.Condition) []normCondition {
	normalized := make([]normCondition, 0, len(conditions))
	for _, cond := range conditions {
		op := cond.Operator
		if op == "keyword" {
			op = "contains"
		}
		for _, value := range conditionValues(cond.Value, cond.Alternatives) {
			normalized = append(normalized, normalizeParts(cond.Field, op, value, cond.Negated))
		}
	}
	return dedupeNorm(normalized)
}

func canonicalLEQLOperator(op string, negated bool) (string, bool) {
	switch strings.ToLower(op) {
	case "==", "=~", "=":
		return "=", negated
	case "!=", "!~", "!==":
		return "=", !negated
	case "contains", "icontains", "contains_any", "icontains_any", "contains_all", "icontains_all":
		return "contains", negated
	case "!contains", "!icontains":
		return "contains", !negated
	case "startswith", "istartswith", "startswith_any", "istartswith_any":
		return "startswith", negated
	case "!startswith", "!istartswith":
		return "startswith", !negated
	case "endswith", "iendswith":
		return "endswith", negated
	case "!endswith", "!iendswith":
		return "endswith", !negated
	case "in", "iin":
		return "=", negated
	default:
		return op, negated
	}
}

func normalizeParts(field, op, value string, negated bool) normCondition {
	return normCondition{Field: strings.ToLower(field), Op: op, Value: canonicalValue(value), Negated: negated}
}

func canonicalValue(value string) string {
	for strings.Contains(value, `\\`) {
		value = strings.ReplaceAll(value, `\\`, `\`)
	}
	return value
}

func dedupeNorm(conditions []normCondition) []normCondition {
	seen := make(map[normCondition]bool, len(conditions))
	result := make([]normCondition, 0, len(conditions))
	for _, cond := range conditions {
		if !seen[cond] {
			seen[cond] = true
			result = append(result, cond)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return conditionKey(result[i]) < conditionKey(result[j])
	})
	return result
}

func compareConditionSets(expected, actual []normCondition) ([]normCondition, []normCondition) {
	actualSet := make(map[normCondition]bool, len(actual))
	for _, cond := range actual {
		actualSet[cond] = true
	}
	var missing []normCondition
	for _, cond := range expected {
		if !actualSet[cond] {
			missing = append(missing, cond)
		}
	}
	expectedSet := make(map[normCondition]bool, len(expected))
	for _, cond := range expected {
		expectedSet[cond] = true
	}
	var extra []normCondition
	for _, cond := range actual {
		if !expectedSet[cond] {
			extra = append(extra, cond)
		}
	}
	return missing, extra
}

func missingConditions(expected, actual []normCondition) []normCondition {
	missing, _ := compareConditionSets(expected, actual)
	return missing
}

func conditionKeys(conditions []normCondition) []string {
	keys := make([]string, len(conditions))
	for i, cond := range conditions {
		keys[i] = conditionKey(cond)
	}
	sort.Strings(keys)
	return keys
}

func conditionKey(cond normCondition) string {
	prefix := ""
	if cond.Negated {
		prefix = "!"
	}
	return prefix + cond.Field + "|" + cond.Op + "|" + cond.Value
}

func conditionValues(value string, alternatives []string) []string {
	if len(alternatives) > 0 {
		out := append([]string(nil), alternatives...)
		sort.Strings(out)
		return out
	}
	return []string{value}
}

func sigmaValue(values []string) any {
	if len(values) > 1 {
		return values
	}
	return first(values)
}

func first(values []string) string {
	if len(values) == 0 {
		return ""
	}
	return values[0]
}

func joinAlternativeExpressions(values []string, render func(string) string) string {
	if len(values) == 0 {
		return render("")
	}
	if len(values) == 1 {
		return render(values[0])
	}
	parts := make([]string, 0, len(values))
	for _, value := range values {
		parts = append(parts, render(value))
	}
	return "(" + strings.Join(parts, " OR ") + ")"
}

func formatLEQLValue(value string) string {
	if isNumeric(value) || strings.EqualFold(value, "true") || strings.EqualFold(value, "false") {
		return value
	}
	return strconv.Quote(value)
}

func joinLEQLValues(values []string) string {
	out := make([]string, 0, len(values))
	for _, value := range values {
		out = append(out, formatLEQLValue(value))
	}
	return strings.Join(out, ", ")
}

func inferLogsource(conditions []leql.Condition) sigmaLogsource {
	for _, cond := range conditions {
		field := strings.ToLower(cond.Field)
		if strings.Contains(field, "process") || strings.EqualFold(cond.Field, "file_name") || strings.EqualFold(cond.Field, "file_path") {
			return sigmaLogsource{Category: "process_creation", Product: "windows"}
		}
		if strings.Contains(field, "ip") || strings.Contains(field, "port") {
			return sigmaLogsource{Category: "network_connection", Product: "windows"}
		}
	}
	return sigmaLogsource{Category: "generic", Product: "windows"}
}

func valuesForField(field string) []string {
	switch field {
	case "file_name", "process.name":
		return []string{"cmd.exe", "powershell.exe", "pwsh.exe", "rundll32.exe", "mshta.exe", "svchost.exe"}
	case "process.cmd_line":
		return []string{"-enc", "downloadstring", "whoami", "net user", " /c ", "Invoke-WebRequest"}
	case "source_address", "destination_address", "source_ip", "destination_ip":
		return []string{"10.0.0.5", "192.168.1.10", "172.16.1.20", "8.8.8.8", "203.0.113.10"}
	case "action":
		return []string{"allow", "deny", "create", "delete", "login", "logout"}
	case "status", "result", "connection_status":
		return []string{"Success", "Failure", "Failed", "0", "50074", "Denied"}
	case "file_path":
		return []string{`C:\Windows\System32`, `C:\Users\Public`, `C:\ProgramData`, `C:\Temp`}
	case "severity":
		return []string{"LOW", "MEDIUM", "HIGH", "CRITICAL", "INFORMATIONAL"}
	case "service", "protocol":
		return []string{"http", "dns", "ssh", "smb", "kerberos", "vpn"}
	case "url", "domain", "query":
		return []string{"example.com", "update.microsoft.com", "pastebin.com", "/admin", "/login", "malware.test"}
	case "process.exe_file.description", "process.exe_file.product_name":
		return []string{"error", "critical", "failed password", "mfa required", "admin login"}
	default:
		return []string{"admin", "root", "SYSTEM", "svc_app", "web01", "test"}
	}
}

func randomValueForField(rng *rand.Rand, field string) string {
	values := valuesForField(field)
	return values[rng.Intn(len(values))]
}

func uniqueValues(rng *rand.Rand, pool []string, count int) []string {
	if count > len(pool) {
		count = len(pool)
	}
	perm := rng.Perm(len(pool))
	values := make([]string, 0, count)
	for i := 0; i < count; i++ {
		values = append(values, pool[perm[i]])
	}
	sort.Strings(values)
	return values
}

func oneOf[T any](rng *rand.Rand, values ...T) T {
	return values[rng.Intn(len(values))]
}

func contains(values []string, needle string) bool {
	for _, value := range values {
		if value == needle {
			return true
		}
	}
	return false
}

func isNumeric(value string) bool {
	if value == "" {
		return false
	}
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}

func seedForCase(seed, id int64) int64 {
	x := uint64(seed) + 0x9e3779b97f4a7c15 + uint64(id)*0xbf58476d1ce4e5b9
	x = (x ^ (x >> 30)) * 0xbf58476d1ce4e5b9
	x = (x ^ (x >> 27)) * 0x94d049bb133111eb
	return int64(x ^ (x >> 31))
}

func newCorpusWriter(path string) (*corpusWriter, error) {
	if path == "" {
		return &corpusWriter{}, nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	if _, err := file.WriteString("[\n"); err != nil {
		file.Close()
		return nil, err
	}
	return &corpusWriter{file: file, enc: json.NewEncoder(file), first: true}, nil
}

func (w *corpusWriter) write(entry queryEntry) error {
	if w.file == nil {
		return nil
	}
	if !w.first {
		if _, err := w.file.WriteString(",\n"); err != nil {
			return err
		}
	}
	w.first = false
	return w.enc.Encode(entry)
}

func (w *corpusWriter) close() error {
	if w.file == nil {
		return nil
	}
	if _, err := w.file.WriteString("]\n"); err != nil {
		w.file.Close()
		return err
	}
	return w.file.Close()
}

func newFailureWriter(path string) (*failureWriter, error) {
	if path == "" {
		return &failureWriter{}, nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return &failureWriter{file: file, enc: json.NewEncoder(file)}, nil
}

func (w *failureWriter) write(result runResult) error {
	if w.file == nil {
		return nil
	}
	return w.enc.Encode(result)
}

func (w *failureWriter) close() error {
	if w.file == nil {
		return nil
	}
	return w.file.Close()
}

func countFailureStage(s *summary, stage string) {
	switch stage {
	case "leql_parse":
		s.LEQLParseErrors++
	case "expected_mismatch":
		s.ExpectedMismatch++
	case "sigma_parse":
		s.SigmaParseErrors++
	case "sigma_mismatch":
		s.SigmaMismatch++
	case "back_leql_parse":
		s.BackLEQLParseError++
	case "back_leql_mismatch":
		s.BackLEQLMismatch++
	}
}

func printProgress(s summary, total int64, start time.Time) {
	elapsed := time.Since(start)
	rate := float64(s.Total) / elapsed.Seconds()
	remaining := ""
	if rate > 0 && total > s.Total {
		remaining = fmt.Sprintf(", eta=%s", time.Duration(float64(total-s.Total)/rate)*time.Second)
	}
	duplicates := ""
	if s.DuplicateQueries > 0 {
		duplicates = fmt.Sprintf(" duplicates=%d", s.DuplicateQueries)
	}
	fmt.Fprintf(os.Stderr, "processed=%d/%d passed=%d failed=%d%s rate=%.0f/s%s\n", s.Total, total, s.Passed, s.Failed, duplicates, rate, remaining)
}

func printSummary(s summary, corpusPath, failPath string, failureRecords int) {
	fmt.Printf("generated=%d passed=%d failed=%d duplicates=%d duration=%s\n", s.Total, s.Passed, s.Failed, s.DuplicateQueries, s.Duration)
	fmt.Printf("failures: leql_parse=%d expected_mismatch=%d sigma_parse=%d sigma_mismatch=%d back_leql_parse=%d back_leql_mismatch=%d\n",
		s.LEQLParseErrors, s.ExpectedMismatch, s.SigmaParseErrors, s.SigmaMismatch, s.BackLEQLParseError, s.BackLEQLMismatch)
	if corpusPath != "" {
		fmt.Printf("verified corpus: %s\n", corpusPath)
	}
	if failPath != "" {
		fmt.Printf("failure records: %s (%d written)\n", failPath, failureRecords)
	}
}

func mapCapacityHint(target int64) int {
	if target <= 0 {
		return 0
	}
	if int64(int(target)) != target {
		return 0
	}
	return int(target)
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
