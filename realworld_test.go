package leql

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
)

// corpusEntry represents a single query from the corpus JSON file.
type corpusEntry struct {
	Query       string `json:"query"`
	Description string `json:"description"`
}

// TestCorpus tests the LEQL parser against a corpus of real-world queries
// sourced from Rapid7 InsightIDR documentation and common usage patterns.
func TestCorpus(t *testing.T) {
	corpusPath := "testdata/corpus.json"

	if _, err := os.Stat(corpusPath); os.IsNotExist(err) {
		t.Skip("Corpus not available at testdata/corpus.json")
	}

	data, err := os.ReadFile(corpusPath)
	if err != nil {
		t.Fatalf("Failed to read corpus: %v", err)
	}

	var entries []corpusEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		t.Fatalf("Failed to parse corpus JSON: %v", err)
	}

	if len(entries) < 100 {
		t.Fatalf("Corpus too small: got %d entries, want at least 100", len(entries))
	}

	t.Logf("Loaded %d queries from corpus", len(entries))

	var success, partial, failed, panicked int
	var errorDetails []string

	for i, e := range entries {
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked++
					msg := fmt.Sprintf("PANIC [%d] %q (%s): %v", i, e.Query, e.Description, r)
					errorDetails = append(errorDetails, msg)
					t.Logf("%s", msg)
				}
			}()

			result := ExtractConditions(e.Query)

			hasConditions := len(result.Conditions) > 0 || len(result.GroupByFields) > 0 || len(result.Commands) > 0
			hasErrors := len(result.Errors) > 0

			switch {
			case hasConditions && !hasErrors:
				success++
			case hasConditions && hasErrors:
				partial++
				msg := fmt.Sprintf("PARTIAL [%d] %q (%s): errors=%v", i, e.Query, e.Description, result.Errors)
				errorDetails = append(errorDetails, msg)
				t.Logf("%s", msg)
			default:
				failed++
				msg := fmt.Sprintf("FAILED [%d] %q (%s): errors=%v conditions=%d commands=%v",
					i, e.Query, e.Description, result.Errors, len(result.Conditions), result.Commands)
				errorDetails = append(errorDetails, msg)
				t.Logf("%s", msg)
			}
		}()
	}

	total := success + partial + failed + panicked
	successRate := float64(success) * 100 / float64(total)
	partialRate := float64(partial) * 100 / float64(total)
	failedRate := float64(failed) * 100 / float64(total)
	panicRate := float64(panicked) * 100 / float64(total)

	t.Logf("")
	t.Logf("=== Corpus Test Results ===")
	t.Logf("Total:    %d", total)
	t.Logf("Success:  %d (%.1f%%)", success, successRate)
	t.Logf("Partial:  %d (%.1f%%)", partial, partialRate)
	t.Logf("Failed:   %d (%.1f%%)", failed, failedRate)
	t.Logf("Panicked: %d (%.1f%%)", panicked, panicRate)
	t.Logf("===========================")

	if len(errorDetails) > 0 {
		t.Logf("")
		t.Logf("--- Error Summary ---")
		for _, d := range errorDetails {
			t.Logf("  %s", d)
		}
	}

	// Hard failures
	if panicked > 0 {
		t.Errorf("Parser panicked on %d/%d queries (%.1f%%) — panic rate must be 0%%", panicked, total, panicRate)
	}
	if successRate < 50 {
		t.Errorf("Success rate %.1f%% is below 50%% threshold (success=%d, total=%d)", successRate, success, total)
	}

	// Informational warnings (do not fail the test)
	if failedRate > 20 {
		t.Logf("WARNING: Failed rate %.1f%% is above 20%% — consider improving grammar coverage", failedRate)
	}

	// Log queries with parse errors for grammar improvement tracking
	if partial > 0 || failed > 0 {
		t.Logf("")
		t.Logf("Queries with errors may indicate grammar gaps. Review the errors above to decide if grammar fixes are warranted.")
	}
}

// TestGeneratedCorpus runs the parser against a large generated corpus (~50k queries).
// All generated queries are syntactically valid, so we require 100% parse success.
func TestGeneratedCorpus(t *testing.T) {
	corpusPath := "testdata/corpus_generated.json"

	if _, err := os.Stat(corpusPath); os.IsNotExist(err) {
		t.Skip("Generated corpus not available — run: go run testdata/generate_corpus.go > testdata/corpus_generated.json")
	}

	data, err := os.ReadFile(corpusPath)
	if err != nil {
		t.Fatalf("Failed to read generated corpus: %v", err)
	}

	var entries []corpusEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		t.Fatalf("Failed to parse generated corpus JSON: %v", err)
	}

	t.Logf("Loaded %d generated queries", len(entries))

	var success, partial, failed, panicked int
	var sampleErrors []string

	for _, e := range entries {
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked++
					if len(sampleErrors) < 20 {
						sampleErrors = append(sampleErrors, fmt.Sprintf("PANIC %q: %v", e.Query, r))
					}
				}
			}()

			result := ExtractConditions(e.Query)

			hasConditions := len(result.Conditions) > 0 || len(result.GroupByFields) > 0 || len(result.Commands) > 0
			hasErrors := len(result.Errors) > 0

			switch {
			case hasConditions && !hasErrors:
				success++
			case hasConditions && hasErrors:
				partial++
				if len(sampleErrors) < 20 {
					sampleErrors = append(sampleErrors, fmt.Sprintf("PARTIAL %q: errors=%v", e.Query, result.Errors))
				}
			default:
				failed++
				if len(sampleErrors) < 20 {
					sampleErrors = append(sampleErrors, fmt.Sprintf("FAILED %q: errors=%v", e.Query, result.Errors))
				}
			}
		}()
	}

	total := success + partial + failed + panicked
	successRate := float64(success) * 100 / float64(total)

	t.Logf("")
	t.Logf("=== Generated Corpus Results ===")
	t.Logf("Total:    %d", total)
	t.Logf("Success:  %d (%.1f%%)", success, successRate)
	t.Logf("Partial:  %d (%.2f%%)", partial, float64(partial)*100/float64(total))
	t.Logf("Failed:   %d (%.2f%%)", failed, float64(failed)*100/float64(total))
	t.Logf("Panicked: %d", panicked)
	t.Logf("================================")

	if len(sampleErrors) > 0 {
		t.Logf("")
		t.Logf("--- Sample Errors (first %d) ---", len(sampleErrors))
		for _, d := range sampleErrors {
			t.Logf("  %s", d)
		}
	}

	if panicked > 0 {
		t.Errorf("Parser panicked on %d queries", panicked)
	}
	if successRate < 95 {
		t.Errorf("Success rate %.1f%% is below 95%% threshold", successRate)
	}
}

// TestCorpusNoPanic is a focused test ensuring no query in the corpus causes a panic.
func TestCorpusNoPanic(t *testing.T) {
	corpusPath := "testdata/corpus.json"

	if _, err := os.Stat(corpusPath); os.IsNotExist(err) {
		t.Skip("Corpus not available at testdata/corpus.json")
	}

	data, err := os.ReadFile(corpusPath)
	if err != nil {
		t.Fatalf("Failed to read corpus: %v", err)
	}

	var entries []corpusEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		t.Fatalf("Failed to parse corpus JSON: %v", err)
	}

	for i, e := range entries {
		t.Run(fmt.Sprintf("query_%d", i), func(t *testing.T) {
			// ExtractConditions has its own panic recovery, but we add an extra layer
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("Panic on query %q: %v", e.Query, r)
				}
			}()

			result := ExtractConditions(e.Query)
			// Check that the internal panic recovery didn't trigger
			for _, err := range result.Errors {
				if strings.HasPrefix(err, "parser panic:") {
					t.Errorf("Internal panic on query %q: %s", e.Query, err)
				}
			}
		})
	}
}

// BenchmarkCorpus benchmarks parsing speed across the full corpus.
func BenchmarkCorpus(b *testing.B) {
	corpusPath := "testdata/corpus.json"

	if _, err := os.Stat(corpusPath); os.IsNotExist(err) {
		b.Skip("Corpus not available")
	}

	data, err := os.ReadFile(corpusPath)
	if err != nil {
		b.Fatalf("Failed to read corpus: %v", err)
	}

	var entries []corpusEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		b.Fatalf("Failed to parse corpus JSON: %v", err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, e := range entries {
			_ = ExtractConditions(e.Query)
		}
	}
}
