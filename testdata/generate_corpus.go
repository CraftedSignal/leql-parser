//go:build ignore

// generate_corpus.go generates a large corpus of valid LEQL queries for parser testing.
// Usage: go run testdata/generate_corpus.go > testdata/corpus_generated.json
package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type entry struct {
	Query       string `json:"query"`
	Description string `json:"description"`
	Source      string `json:"source"`
}

var fields = []string{
	"source_address", "destination_address", "source_ip", "destination_ip",
	"source_user", "destination_user", "user", "username", "account",
	"source_asset", "destination_asset", "asset", "hostname",
	"status", "result", "action", "severity", "risk",
	"service", "protocol", "direction", "connection_status",
	"destination_port", "source_port", "logon_type",
	"geoip_country_code", "geoip_country_name", "geoip_city", "geoip_organization",
	"process.name", "process.cmd_line", "process.exe_file.description",
	"process.exe_file.hashes.md5", "process.exe_file.hashes.sha256",
	"process.exe_file.product_name", "process.username",
	"file_name", "file_path", "file_hash",
	"incoming_bytes", "outgoing_bytes",
	"EventCode", "EventID", "event_code",
	"public_suffix", "query", "url", "domain",
	"source_json.eventCode", "source_json.Category", "source_json.Severity",
	"source_json.Workload", "source_json.AccountType",
	"header.name", "header.severity", "header.device_event_class_id",
	"group", "target_user", "target_account", "source_account",
	"signature", "description", "type", "title",
	"app_protocol_description", "user_agent", "is_blocked",
	"access_types", "env_vars.7.parent_val",
}

var quotedFields = []string{
	"source_address", "destination_address", "geoip_country_code",
	"result", "service", "access_types", "env_vars.7.parent_val",
}

var stringValues = []string{
	"admin", "root", "SYSTEM",
	"SUCCESS", "FAILED", "FAILED_BAD_PASSWORD", "FAILED_BAD_LOGIN", "ACCOUNT_LOCKED_OUT",
	"MEMBER_ADDED_TO_SECURITY_GROUP", "MEMBER_REMOVED_FROM_SECURITY_GROUP",
	"INBOUND", "OUTBOUND", "ACCEPT", "DENY",
	"HIGH", "CRITICAL", "LOW", "MEDIUM", "INFORMATIONAL", "MALICIOUS",
	"powershell.exe", "cmd.exe", "taskkill.exe", "whoami", "net user",
	"United States", "Russia", "China", "Germany", "United Kingdom",
	"error", "warning", "fatal", "exception",
	"o365", "box", "vpn", "sslvpn-session", "kerberos", "krbtgt",
	"Amazon Boardman", "ReadData", "WriteData",
	"SSH", "TLS", "HTTP", "DNS",
}

var numericFields = []string{
	"status", "destination_port", "source_port", "incoming_bytes",
	"outgoing_bytes", "EventCode", "EventID", "event_code", "logon_type",
}

var numbers = []string{
	"0", "1", "3", "10", "22", "53", "80", "200", "300", "400", "403", "404", "443", "500",
	"1024", "3389", "4624", "4625", "4688", "4732", "5000", "8080",
	"50000", "100000000", "50000000",
}

var countryCodes = []string{
	"US", "GB", "DE", "FR", "JP", "CN", "RU", "KP", "IR", "CA", "AU",
	"PL", "UK", "MX", "BR", "IN", "KR", "IL", "NL", "SE",
}

var ipRanges = []string{
	"10.0.0.0/8", "172.16.0.0/12", "172.27.0.0/16", "192.168.0.0/16",
	"169.254.0.0/16", "64.62.128.0/17", "100.64.0.0/10",
}

var compOps = []string{"=", "!=", ">", ">=", "<", "<=", "==", "!=="}
var stringOps = []string{"CONTAINS", "ICONTAINS", "STARTS-WITH", "ISTARTS-WITH"}
var setOps = []string{"IN", "IIN"}
var listStringOps = []string{"CONTAINS-ANY", "ICONTAINS-ANY", "CONTAINS-ALL", "ICONTAINS-ALL", "STARTS-WITH-ANY", "ISTARTS-WITH-ANY"}
var calcFuncs = []string{"count", "COUNT", "bytes", "BYTES"}
var calcFieldFuncs = []string{"sum", "SUM", "average", "AVERAGE", "unique", "UNIQUE", "min", "MIN", "max", "MAX", "standarddeviation", "sd", "SD"}
var sortDirs = []string{"asc", "desc", "ascending", "descending"}

func main() {
	rng := rand.New(rand.NewSource(42)) // deterministic
	var queries []entry
	seen := map[string]bool{}

	add := func(q, desc string) {
		if seen[q] {
			return
		}
		seen[q] = true
		queries = append(queries, entry{Query: q, Description: desc, Source: "generated"})
	}

	// 1. Simple comparisons: field op value (all combinations)
	for _, f := range fields {
		for _, op := range compOps {
			v := pick(rng, stringValues)
			add(fmt.Sprintf(`where(%s%s"%s")`, f, op, v), fmt.Sprintf("%s %s %s", f, op, v))
		}
	}

	// 2. Numeric comparisons
	for _, f := range numericFields {
		for _, op := range compOps {
			v := pick(rng, numbers)
			add(fmt.Sprintf(`where(%s%s%s)`, f, op, v), fmt.Sprintf("%s %s %s", f, op, v))
		}
	}

	// 3. String operators
	for _, f := range fields[:20] {
		for _, op := range stringOps {
			v := pick(rng, stringValues)
			add(fmt.Sprintf(`where(%s %s "%s")`, f, op, v), fmt.Sprintf("%s %s %s", f, op, v))
			// Unquoted value
			add(fmt.Sprintf(`where(%s %s %s)`, f, op, pickIdent(rng)), fmt.Sprintf("%s %s ident", f, op))
		}
	}

	// 4. Set operators (IN/IIN)
	for _, f := range fields[:25] {
		for _, op := range setOps {
			vals := pickN(rng, stringValues, 2+rng.Intn(5))
			quotedVals := make([]string, len(vals))
			for i, v := range vals {
				quotedVals[i] = fmt.Sprintf(`"%s"`, v)
			}
			add(fmt.Sprintf(`where(%s %s [%s])`, f, op, strings.Join(quotedVals, ", ")),
				fmt.Sprintf("%s %s %d values", f, op, len(vals)))
		}
		// Numeric IN
		if contains(numericFields, f) {
			nums := pickN(rng, numbers, 2+rng.Intn(4))
			add(fmt.Sprintf(`where(%s IN [%s])`, f, strings.Join(nums, ", ")),
				fmt.Sprintf("%s IN %d numbers", f, len(nums)))
		}
	}

	// 5. List string operators
	for _, f := range fields[:15] {
		for _, op := range listStringOps {
			vals := pickN(rng, stringValues, 2+rng.Intn(4))
			quotedVals := make([]string, len(vals))
			for i, v := range vals {
				quotedVals[i] = fmt.Sprintf(`"%s"`, v)
			}
			add(fmt.Sprintf(`where(%s %s [%s])`, f, op, strings.Join(quotedVals, ", ")),
				fmt.Sprintf("%s %s %d values", f, op, len(vals)))
		}
	}

	// 6. Negation: NOT field op value (prefix)
	for i := 0; i < 200; i++ {
		f := pick(rng, fields)
		op := pick(rng, compOps[:2]) // = or !=
		v := pick(rng, stringValues)
		add(fmt.Sprintf(`where(NOT %s%s"%s")`, f, op, v), fmt.Sprintf("NOT %s %s %s", f, op, v))
	}

	// 7. Negation: field NOT op value (postfix)
	for i := 0; i < 200; i++ {
		f := pick(rng, fields)
		op := pick(rng, setOps)
		vals := pickN(rng, stringValues, 2+rng.Intn(4))
		quotedVals := make([]string, len(vals))
		for j, v := range vals {
			quotedVals[j] = fmt.Sprintf(`"%s"`, v)
		}
		add(fmt.Sprintf(`where(%s NOT %s [%s])`, f, op, strings.Join(quotedVals, ", ")),
			fmt.Sprintf("%s NOT %s %d values", f, op, len(vals)))
	}
	for i := 0; i < 200; i++ {
		f := pick(rng, fields)
		op := pick(rng, stringOps)
		v := pick(rng, stringValues)
		add(fmt.Sprintf(`where(%s NOT %s "%s")`, f, op, v), fmt.Sprintf("%s NOT %s %s", f, op, v))
	}

	// 8. NOCASE conditions
	for i := 0; i < 100; i++ {
		f := pick(rng, fields)
		v := pick(rng, stringValues)
		add(fmt.Sprintf(`where(%s=NOCASE("%s"))`, f, v), fmt.Sprintf("NOCASE %s = %s", f, v))
		add(fmt.Sprintf(`where(%s=NOCASE(%s))`, f, pickIdent(rng)), fmt.Sprintf("NOCASE %s = ident", f))
	}

	// 9. IP conditions
	for _, f := range []string{"source_address", "destination_address", "source_ip", "destination_ip"} {
		for _, ip := range ipRanges {
			add(fmt.Sprintf(`where(%s=IP(%s))`, f, ip), fmt.Sprintf("IP %s = %s", f, ip))
		}
		// IP in set
		ips := pickN(rng, ipRanges, 2+rng.Intn(3))
		ipVals := make([]string, len(ips))
		for i, ip := range ips {
			ipVals[i] = fmt.Sprintf("IP(%s)", ip)
		}
		add(fmt.Sprintf(`where(%s IN [%s])`, f, strings.Join(ipVals, ", ")),
			fmt.Sprintf("%s IN IP ranges", f))
		add(fmt.Sprintf(`where(%s NOT IN [%s])`, f, strings.Join(ipVals, ", ")),
			fmt.Sprintf("%s NOT IN IP ranges", f))
	}

	// 10. AND/OR combinations
	for i := 0; i < 2000; i++ {
		n := 2 + rng.Intn(4) // 2-5 conditions
		conds := make([]string, n)
		for j := 0; j < n; j++ {
			conds[j] = randCondition(rng)
		}
		ops := make([]string, n-1)
		for j := range ops {
			if rng.Intn(3) == 0 {
				ops[j] = " OR "
			} else {
				ops[j] = " AND "
			}
		}
		var q strings.Builder
		q.WriteString(conds[0])
		for j := 0; j < len(ops); j++ {
			q.WriteString(ops[j])
			q.WriteString(conds[j+1])
		}
		add(fmt.Sprintf("where(%s)", q.String()), fmt.Sprintf("%d conditions combined", n))
	}

	// 11. Parenthesized expressions
	for i := 0; i < 500; i++ {
		c1, c2, c3 := randCondition(rng), randCondition(rng), randCondition(rng)
		add(fmt.Sprintf("where((%s OR %s) AND %s)", c1, c2, c3), "paren OR AND")
		add(fmt.Sprintf("where(%s AND (%s OR %s))", c1, c2, c3), "AND paren OR")
	}

	// 12. Nested where()
	for i := 0; i < 300; i++ {
		c := randCondition(rng)
		add(fmt.Sprintf("where(where(%s))", c), "nested where")
	}
	for i := 0; i < 100; i++ {
		c1, c2 := randCondition(rng), randCondition(rng)
		add(fmt.Sprintf("where(where(%s) AND %s)", c1, c2), "nested where AND")
		add(fmt.Sprintf("where(where(%s AND %s))", c1, c2), "nested where inner AND")
	}

	// 13. Quoted field names
	for i := 0; i < 300; i++ {
		f := pick(rng, quotedFields)
		op := pick(rng, compOps[:2])
		v := pick(rng, stringValues)
		add(fmt.Sprintf(`where("%s"%s"%s")`, f, op, v), fmt.Sprintf("quoted field %s", f))
	}
	for i := 0; i < 200; i++ {
		f := pick(rng, quotedFields)
		op := pick(rng, setOps)
		vals := pickN(rng, stringValues, 2+rng.Intn(3))
		quotedVals := make([]string, len(vals))
		for j, v := range vals {
			quotedVals[j] = fmt.Sprintf(`"%s"`, v)
		}
		add(fmt.Sprintf(`where("%s" %s [%s])`, f, op, strings.Join(quotedVals, ", ")),
			fmt.Sprintf("quoted %s %s", f, op))
		add(fmt.Sprintf(`where("%s" NOT %s [%s])`, f, op, strings.Join(quotedVals, ", ")),
			fmt.Sprintf("quoted %s NOT %s", f, op))
	}

	// 14. Keyword searches
	for _, v := range stringValues {
		add(fmt.Sprintf(`where("%s")`, v), fmt.Sprintf("keyword %s", v))
	}

	// 15. Regex searches
	regexPatterns := []string{
		"/error/", "/error/i", "/.*password.*/", "/pass{2,5}word/",
		"/w{3}/", "/467./", `/192\.168\.1\.\d+/`, `/\\broot\\b/i`,
		"/failed.*login/i", "/cmd\\.exe|powershell\\.exe/i",
		`/total sale (?P<saleValue>\d*)/`,
		`/eventCode\":(?P<EVID>\d{4})/`,
	}
	for _, r := range regexPatterns {
		add(fmt.Sprintf("where(%s)", r), "regex search")
	}
	for _, f := range fields[:10] {
		for _, r := range regexPatterns[:4] {
			add(fmt.Sprintf("where(%s=%s)", f, r), fmt.Sprintf("field regex %s", f))
		}
	}

	// 16. ALL() conditions
	for i := 0; i < 100; i++ {
		n := 2 + rng.Intn(3)
		fs := pickN(rng, fields, n)
		op := pick(rng, compOps[:2])
		v := pick(rng, stringValues)
		add(fmt.Sprintf(`where(ALL(%s)%s"%s")`, strings.Join(fs, ", "), op, v),
			fmt.Sprintf("ALL %d fields", n))
	}

	// 17. groupby clause combinations
	for i := 0; i < 500; i++ {
		n := 1 + rng.Intn(4)
		fs := pickN(rng, fields, n)
		cond := randCondition(rng)
		add(fmt.Sprintf("where(%s) groupby(%s)", cond, strings.Join(fs, ", ")),
			fmt.Sprintf("groupby %d fields", n))
	}
	for i := 0; i < 200; i++ {
		n := 1 + rng.Intn(3)
		fs := pickN(rng, fields, n)
		add(fmt.Sprintf("groupby(%s)", strings.Join(fs, ", ")),
			fmt.Sprintf("bare groupby %d fields", n))
	}

	// 18. calculate clause
	for _, calc := range calcFuncs {
		add(fmt.Sprintf("calculate(%s)", calc), fmt.Sprintf("calculate %s", calc))
		for i := 0; i < 20; i++ {
			cond := randCondition(rng)
			add(fmt.Sprintf("where(%s) calculate(%s)", cond, calc),
				fmt.Sprintf("where calculate %s", calc))
		}
	}
	for _, calc := range calcFieldFuncs {
		f := pick(rng, fields)
		add(fmt.Sprintf("calculate(%s:%s)", calc, f), fmt.Sprintf("calculate %s:%s", calc, f))
		for i := 0; i < 10; i++ {
			cond := randCondition(rng)
			f2 := pick(rng, fields)
			add(fmt.Sprintf("where(%s) calculate(%s:%s)", cond, calc, f2),
				fmt.Sprintf("where calculate %s:field", calc))
		}
	}

	// 19. having clause (calculate must come before having in LEQL)
	for i := 0; i < 200; i++ {
		cond := randCondition(rng)
		f := pick(rng, fields)
		n := pick(rng, numbers)
		op := pick(rng, compOps[:6])
		gf := pick(rng, fields)
		calcType := pick(rng, []string{"count", "unique:" + f, "sum:" + f, "average:" + f})
		add(fmt.Sprintf("where(%s) groupby(%s) calculate(count) having(%s%s%s)", cond, gf, calcType, op, n),
			"having clause")
	}

	// 20. sort clause
	for _, dir := range sortDirs {
		for i := 0; i < 30; i++ {
			cond := randCondition(rng)
			gf := pick(rng, fields)
			add(fmt.Sprintf("where(%s) groupby(%s) sort(%s)", cond, gf, dir),
				fmt.Sprintf("sort %s", dir))
		}
	}

	// 21. limit clause
	for i := 0; i < 200; i++ {
		cond := randCondition(rng)
		gf := pick(rng, fields)
		n1 := 1 + rng.Intn(1000)
		add(fmt.Sprintf("where(%s) groupby(%s) limit(%d)", cond, gf, n1),
			fmt.Sprintf("limit %d", n1))
		if rng.Intn(2) == 0 {
			n2 := 1 + rng.Intn(50)
			add(fmt.Sprintf("where(%s) groupby(%s) limit(%d, %d)", cond, gf, n1, n2),
				fmt.Sprintf("limit %d, %d", n1, n2))
		}
	}

	// 22. timeslice clause
	for _, ts := range []string{"30", "60", "30s", "5m", "1h", "30m", "1d", "7d"} {
		for i := 0; i < 20; i++ {
			cond := randCondition(rng)
			add(fmt.Sprintf("where(%s) calculate(count) timeslice(%s)", cond, ts),
				fmt.Sprintf("timeslice %s", ts))
		}
	}

	// 23. select clause
	for i := 0; i < 100; i++ {
		n := 1 + rng.Intn(4)
		fs := pickN(rng, fields, n)
		cond := randCondition(rng)
		add(fmt.Sprintf("select(%s) where(%s)", strings.Join(fs, ", "), cond),
			fmt.Sprintf("select %d fields", n))
	}

	// 24. Full pipeline combinations
	for i := 0; i < 3000; i++ {
		var parts []string

		// where clause (80% chance)
		if rng.Intn(5) > 0 {
			n := 1 + rng.Intn(3)
			conds := make([]string, n)
			for j := 0; j < n; j++ {
				conds[j] = randCondition(rng)
			}
			ops := []string{" AND ", " OR "}
			var w strings.Builder
			w.WriteString(conds[0])
			for j := 1; j < n; j++ {
				w.WriteString(ops[rng.Intn(2)])
				w.WriteString(conds[j])
			}
			parts = append(parts, fmt.Sprintf("where(%s)", w.String()))
		}

		// groupby (60% chance)
		if rng.Intn(5) > 1 {
			n := 1 + rng.Intn(3)
			fs := pickN(rng, fields, n)
			parts = append(parts, fmt.Sprintf("groupby(%s)", strings.Join(fs, ", ")))
		}

		// calculate (40% chance) — must come before having in LEQL
		if rng.Intn(5) > 2 {
			if rng.Intn(2) == 0 {
				parts = append(parts, fmt.Sprintf("calculate(%s)", pick(rng, calcFuncs)))
			} else {
				parts = append(parts, fmt.Sprintf("calculate(%s:%s)",
					pick(rng, calcFieldFuncs), pick(rng, fields)))
			}
		}

		// having (20% chance)
		if rng.Intn(5) == 0 {
			op := pick(rng, compOps[:6])
			n := pick(rng, numbers)
			parts = append(parts, fmt.Sprintf("having(count%s%s)", op, n))
		}

		// sort (30% chance)
		if rng.Intn(3) == 0 {
			parts = append(parts, fmt.Sprintf("sort(%s)", pick(rng, sortDirs)))
		}

		// limit (20% chance)
		if rng.Intn(5) == 0 {
			parts = append(parts, fmt.Sprintf("limit(%d)", 1+rng.Intn(500)))
		}

		// timeslice (10% chance)
		if rng.Intn(10) == 0 {
			ts := pick(rng, []string{"30s", "5m", "30m", "1h", "1d"})
			parts = append(parts, fmt.Sprintf("timeslice(%s)", ts))
		}

		if len(parts) > 0 {
			q := strings.Join(parts, " ")
			add(q, fmt.Sprintf("pipeline %d parts", len(parts)))
		}
	}

	// 25. Edge cases
	// Field exists
	for _, f := range fields[:20] {
		add(fmt.Sprintf("where(%s)", f), fmt.Sprintf("field exists %s", f))
	}
	// Multi-field OR (comma-separated)
	for i := 0; i < 100; i++ {
		n := 2 + rng.Intn(3)
		fs := pickN(rng, fields, n)
		op := pick(rng, compOps[:2])
		v := pick(rng, stringValues)
		add(fmt.Sprintf(`where(%s%s"%s")`, strings.Join(fs, ","), op, v),
			fmt.Sprintf("multi-field %s", op))
	}

	// 26. Percentile functions
	for _, pctl := range []string{"percentile", "pctl", "PERCENTILE", "PCTL"} {
		for _, n := range []string{"50", "90", "95", "99"} {
			f := pick(rng, numericFields)
			add(fmt.Sprintf("calculate(%s(%s):%s)", pctl, n, f),
				fmt.Sprintf("%s(%s):%s", pctl, n, f))
		}
	}

	// 27. sort with key qualifier
	for i := 0; i < 50; i++ {
		gf := pick(rng, fields)
		add(fmt.Sprintf("groupby(%s) calculate(count) sort(asc#key)", gf), "sort asc by key")
		add(fmt.Sprintf("groupby(%s) calculate(count) sort(desc#key)", gf), "sort desc by key")
	}

	// 28. Double-nested where with quoted fields and postfix negation
	for i := 0; i < 200; i++ {
		f := pick(rng, quotedFields)
		op := pick(rng, setOps)
		vals := pickN(rng, stringValues, 2+rng.Intn(3))
		qv := make([]string, len(vals))
		for j, v := range vals {
			qv[j] = fmt.Sprintf(`"%s"`, v)
		}
		c2f := pick(rng, quotedFields)
		c2v := pick(rng, stringValues)
		add(fmt.Sprintf(`where(where("%s" NOT %s [%s] and "%s" = "%s"))`, f, op, strings.Join(qv, ", "), c2f, c2v),
			"nested where quoted postfix NOT")
	}

	// 29. Mixed case operators
	mixedCaseOps := []string{"Contains", "contains", "CONTAINS", "Icontains", "icontains",
		"Starts-With", "starts-with", "Istarts-With"}
	for _, op := range mixedCaseOps {
		f := pick(rng, fields)
		v := pick(rng, stringValues)
		add(fmt.Sprintf(`where(%s %s "%s")`, f, op, v), fmt.Sprintf("mixed case %s", op))
	}

	// 30. Triple quoted strings
	for i := 0; i < 50; i++ {
		f := pick(rng, fields)
		v := pick(rng, stringValues)
		add(fmt.Sprintf(`where(%s='''%s''')`, f, v), "triple single quote")
		add(fmt.Sprintf(`where(%s="""%s""")`, f, v), "triple double quote")
	}

	// 31. Bare keyword search (unquoted)
	for i := 0; i < 50; i++ {
		f := pickIdent(rng)
		add(fmt.Sprintf("where(%s)", f), fmt.Sprintf("bare keyword %s", f))
	}

	// 32. NOT expressions
	for i := 0; i < 200; i++ {
		c := randCondition(rng)
		add(fmt.Sprintf("where(NOT (%s))", c), "NOT parenthesized")
	}
	for i := 0; i < 100; i++ {
		c1, c2 := randCondition(rng), randCondition(rng)
		add(fmt.Sprintf("where(NOT %s AND %s)", c1, c2), "NOT first AND second")
		add(fmt.Sprintf("where(%s AND NOT %s)", c1, c2), "first AND NOT second")
	}

	// 33. Deeply nested parentheses
	for i := 0; i < 100; i++ {
		c := randCondition(rng)
		add(fmt.Sprintf("where((((%s))))", c), "deeply nested parens")
		c2 := randCondition(rng)
		add(fmt.Sprintf("where(((%s) AND (%s)))", c, c2), "double paren AND")
	}

	// 34. More pipeline combos to reach target
	for i := 0; i < 40000; i++ {
		q := randFullQuery(rng)
		add(q, "random pipeline")
	}

	// Shuffle for randomness
	rng.Shuffle(len(queries), func(i, j int) {
		queries[i], queries[j] = queries[j], queries[i]
	})

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(queries); err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "Generated %d unique queries\n", len(queries))
}

func pick(rng *rand.Rand, s []string) string {
	return s[rng.Intn(len(s))]
}

func pickN(rng *rand.Rand, s []string, n int) []string {
	if n > len(s) {
		n = len(s)
	}
	perm := rng.Perm(len(s))
	result := make([]string, n)
	for i := 0; i < n; i++ {
		result[i] = s[perm[i]]
	}
	return result
}

func pickIdent(rng *rand.Rand) string {
	idents := []string{"admin", "error", "warning", "success", "failed", "powershell", "ssh", "root", "vpn", "alert", "blocked", "true", "false"}
	return idents[rng.Intn(len(idents))]
}

func contains(s []string, v string) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}

func randCondition(rng *rand.Rand) string {
	f := pick(rng, fields)
	switch rng.Intn(8) {
	case 0: // comparison with string
		op := pick(rng, compOps[:2])
		v := pick(rng, stringValues)
		return fmt.Sprintf(`%s%s"%s"`, f, op, v)
	case 1: // comparison with number
		nf := pick(rng, numericFields)
		op := pick(rng, compOps)
		n := pick(rng, numbers)
		return fmt.Sprintf(`%s%s%s`, nf, op, n)
	case 2: // string op
		op := pick(rng, stringOps)
		v := pick(rng, stringValues)
		return fmt.Sprintf(`%s %s "%s"`, f, op, v)
	case 3: // set op
		op := pick(rng, setOps)
		vals := pickN(rng, stringValues, 2+rng.Intn(3))
		qv := make([]string, len(vals))
		for i, v := range vals {
			qv[i] = fmt.Sprintf(`"%s"`, v)
		}
		return fmt.Sprintf(`%s %s [%s]`, f, op, strings.Join(qv, ", "))
	case 4: // NOCASE
		v := pick(rng, stringValues)
		return fmt.Sprintf(`%s=NOCASE("%s")`, f, v)
	case 5: // IP
		ipf := pick(rng, []string{"source_address", "destination_address", "source_ip"})
		ip := pick(rng, ipRanges)
		return fmt.Sprintf(`%s=IP(%s)`, ipf, ip)
	case 6: // quoted field
		qf := pick(rng, quotedFields)
		op := pick(rng, compOps[:2])
		v := pick(rng, stringValues)
		return fmt.Sprintf(`"%s"%s"%s"`, qf, op, v)
	default: // regex
		return fmt.Sprintf(`%s=/error.*/i`, f)
	}
}

func randFullQuery(rng *rand.Rand) string {
	var parts []string

	// select (5%)
	if rng.Intn(20) == 0 {
		n := 1 + rng.Intn(3)
		fs := pickN(rng, fields, n)
		parts = append(parts, fmt.Sprintf("select(%s)", strings.Join(fs, ", ")))
	}

	// where (85%)
	if rng.Intn(7) > 0 {
		n := 1 + rng.Intn(4)
		conds := make([]string, n)
		for j := 0; j < n; j++ {
			conds[j] = randCondition(rng)
		}
		var w strings.Builder
		w.WriteString(conds[0])
		for j := 1; j < n; j++ {
			if rng.Intn(3) == 0 {
				w.WriteString(" OR ")
			} else {
				w.WriteString(" AND ")
			}
			w.WriteString(conds[j])
		}
		// Occasionally wrap in nested where
		if rng.Intn(10) == 0 {
			parts = append(parts, fmt.Sprintf("where(where(%s))", w.String()))
		} else {
			parts = append(parts, fmt.Sprintf("where(%s)", w.String()))
		}
	}

	// groupby (50%)
	if rng.Intn(2) == 0 {
		n := 1 + rng.Intn(3)
		fs := pickN(rng, fields, n)
		parts = append(parts, fmt.Sprintf("groupby(%s)", strings.Join(fs, ", ")))
	}

	// calculate (30%) — must come before having in LEQL
	if rng.Intn(3) == 0 {
		if rng.Intn(2) == 0 {
			parts = append(parts, fmt.Sprintf("calculate(%s)", pick(rng, calcFuncs)))
		} else {
			parts = append(parts, fmt.Sprintf("calculate(%s:%s)", pick(rng, calcFieldFuncs), pick(rng, fields)))
		}
	}

	// having (10%)
	if rng.Intn(10) == 0 {
		op := pick(rng, compOps[:6])
		n := pick(rng, numbers)
		parts = append(parts, fmt.Sprintf("having(count%s%s)", op, n))
	}

	// sort (15%)
	if rng.Intn(7) == 0 {
		if rng.Intn(3) == 0 {
			parts = append(parts, fmt.Sprintf("sort(%s#key)", pick(rng, sortDirs[:2])))
		} else {
			parts = append(parts, fmt.Sprintf("sort(%s)", pick(rng, sortDirs)))
		}
	}

	// limit (15%)
	if rng.Intn(7) == 0 {
		n1 := 1 + rng.Intn(500)
		if rng.Intn(3) == 0 {
			n2 := 1 + rng.Intn(50)
			parts = append(parts, fmt.Sprintf("limit(%d, %d)", n1, n2))
		} else {
			parts = append(parts, fmt.Sprintf("limit(%d)", n1))
		}
	}

	// timeslice (10%)
	if rng.Intn(10) == 0 {
		ts := pick(rng, []string{"30", "60", "30s", "5m", "30m", "1h", "1d", "7d"})
		parts = append(parts, fmt.Sprintf("timeslice(%s)", ts))
	}

	if len(parts) == 0 {
		return "calculate(count)"
	}
	return strings.Join(parts, " ")
}
