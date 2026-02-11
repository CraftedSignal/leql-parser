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

var variables = []string{
	"${private_ip}", "${na_region}", "${trusted_ips}", "${monitored_users}",
	"${threshold}", "${alert_level}", "${allowed_countries}", "${blocked_ports}",
	"${service_accounts}", "${admin_group}", "${vip_users}", "${baseline_count}",
	"${risk_threshold}", "${exclude_ips}", "${watchlist}", "${scan_target}",
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

	// 34. Complex real-world-style queries
	// Long regex patterns (like Windows event code lists from PCI packs)
	eventCodes := []string{"4624", "4625", "4648", "4672", "4688", "4720", "4722", "4724", "4725", "4726",
		"4728", "4732", "4735", "4740", "4756", "4767", "4781", "5136", "5137", "5141",
		"7036", "1102", "4719", "4697", "4698", "4699"}
	for i := 0; i < 200; i++ {
		n := 3 + rng.Intn(15) // 3-17 codes
		codes := pickN(rng, eventCodes, n)
		regexPat := strings.Join(codes, "|")
		f := pick(rng, []string{"data.eventCode", "source_json.eventCode", "xml.system.eventid", "EventCode"})
		gf := pick(rng, []string{"data.computerName", "source_json.computerName", "destination_account", "user"})
		add(fmt.Sprintf("where(%s=/%s/)groupby(%s) limit(10000)", f, regexPat, gf),
			fmt.Sprintf("event code regex %d codes", n))
		add(fmt.Sprintf("where(%s=/%s/) groupby(%s) calculate(count) limit(10000)", f, regexPat, gf),
			fmt.Sprintf("event code regex+calc %d codes", n))
	}

	// Multi-OR regex chains (separate OR'd field conditions)
	for i := 0; i < 200; i++ {
		n := 2 + rng.Intn(4)
		parts := make([]string, n)
		for j := 0; j < n; j++ {
			codes := pickN(rng, eventCodes, 2+rng.Intn(8))
			f := pick(rng, []string{"data.eventCode", "source_json.eventCode"})
			parts[j] = fmt.Sprintf("%s=/%s/", f, strings.Join(codes, "|"))
		}
		gf := pick(rng, fields)
		add(fmt.Sprintf("where(%s)groupby(%s) limit(10000)", strings.Join(parts, " OR "), gf),
			fmt.Sprintf("multi-regex OR %d parts", n))
	}

	// Deep boolean combinations (3-6 levels)
	for i := 0; i < 500; i++ {
		q := randDeepBooleanQuery(rng, 3+rng.Intn(4))
		add(fmt.Sprintf("where(%s)", q), "deep boolean")
	}

	// Complex multi-clause pipelines with all features
	for i := 0; i < 500; i++ {
		// Build a complex where clause
		n := 3 + rng.Intn(5)
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
		// Build full pipeline with many clauses
		var parts []string
		parts = append(parts, fmt.Sprintf("where(%s)", w.String()))
		gfs := pickN(rng, fields, 1+rng.Intn(3))
		parts = append(parts, fmt.Sprintf("groupby(%s)", strings.Join(gfs, ", ")))
		if rng.Intn(2) == 0 {
			parts = append(parts, fmt.Sprintf("calculate(%s:%s)", pick(rng, calcFieldFuncs), pick(rng, fields)))
		} else {
			parts = append(parts, fmt.Sprintf("calculate(%s)", pick(rng, calcFuncs)))
		}
		if rng.Intn(3) == 0 {
			parts = append(parts, fmt.Sprintf("having(count%s%s)", pick(rng, compOps[:6]), pick(rng, numbers)))
		}
		if rng.Intn(2) == 0 {
			if rng.Intn(3) == 0 {
				parts = append(parts, fmt.Sprintf("sort(%s#key)", pick(rng, sortDirs[:2])))
			} else {
				parts = append(parts, fmt.Sprintf("sort(%s)", pick(rng, sortDirs)))
			}
		}
		parts = append(parts, fmt.Sprintf("limit(%d)", 1+rng.Intn(10000)))
		if rng.Intn(3) == 0 {
			ts := pick(rng, []string{"30s", "5m", "10m", "30m", "1h", "1d", "7d"})
			parts = append(parts, fmt.Sprintf("timeslice(%s)", ts))
		}
		add(strings.Join(parts, " "), "complex pipeline")
	}

	// Select + complex pipeline
	for i := 0; i < 200; i++ {
		sf := pickN(rng, fields, 2+rng.Intn(4))
		cond := randCondition(rng)
		gfs := pickN(rng, fields, 1+rng.Intn(2))
		parts := []string{
			fmt.Sprintf("select(%s)", strings.Join(sf, ", ")),
			fmt.Sprintf("where(%s)", cond),
			fmt.Sprintf("groupby(%s)", strings.Join(gfs, ", ")),
			fmt.Sprintf("calculate(%s)", pick(rng, calcFuncs)),
			fmt.Sprintf("sort(%s)", pick(rng, sortDirs)),
			fmt.Sprintf("limit(%d)", 1+rng.Intn(100)),
		}
		add(strings.Join(parts, " "), "select pipeline")
	}

	// Nested where with complex conditions
	for i := 0; i < 300; i++ {
		n := 2 + rng.Intn(4)
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
		gf := pick(rng, fields)
		add(fmt.Sprintf("where(where(%s)) groupby(%s) calculate(count)", w.String(), gf),
			"nested where complex")
	}

	// Parenthesized groups with mixed operators
	for i := 0; i < 500; i++ {
		c := make([]string, 4+rng.Intn(4))
		for j := range c {
			c[j] = randCondition(rng)
		}
		// Build: (a AND b) OR (c AND d) AND (e OR f)
		groups := []string{}
		idx := 0
		for idx < len(c) {
			sz := 2 + rng.Intn(2)
			if idx+sz > len(c) {
				sz = len(c) - idx
			}
			gc := c[idx : idx+sz]
			innerOp := " AND "
			if rng.Intn(3) == 0 {
				innerOp = " OR "
			}
			if sz > 1 {
				groups = append(groups, fmt.Sprintf("(%s)", strings.Join(gc, innerOp)))
			} else {
				groups = append(groups, gc[0])
			}
			idx += sz
		}
		outerOp := " OR "
		if rng.Intn(2) == 0 {
			outerOp = " AND "
		}
		add(fmt.Sprintf("where(%s)", strings.Join(groups, outerOp)),
			fmt.Sprintf("grouped boolean %d terms", len(c)))
	}

	// IP set conditions (multiple IP ranges in lists)
	for i := 0; i < 100; i++ {
		f := pick(rng, []string{"source_address", "destination_address", "source_ip"})
		n := 2 + rng.Intn(5)
		ips := pickN(rng, ipRanges, n)
		ipVals := make([]string, len(ips))
		for j, ip := range ips {
			ipVals[j] = fmt.Sprintf("IP(%s)", ip)
		}
		add(fmt.Sprintf("where(%s NOT IN [%s])", f, strings.Join(ipVals, ", ")),
			fmt.Sprintf("IP NOT IN %d ranges", n))
		add(fmt.Sprintf("where(%s IN [%s])", f, strings.Join(ipVals, ", ")),
			fmt.Sprintf("IP IN %d ranges", n))
		// Complex: IP filter + other conditions
		v := pick(rng, stringValues)
		add(fmt.Sprintf(`where(%s NOT IN [%s] AND user!="%s" AND connection_status = DENY)`, f, strings.Join(ipVals, ", "), v),
			"IP filter complex")
	}

	// Number keyword searches
	for _, n := range numbers {
		add(fmt.Sprintf("where(%s)", n), fmt.Sprintf("number keyword %s", n))
	}

	// ALL() with NOCASE
	for i := 0; i < 50; i++ {
		fs := pickN(rng, fields, 2+rng.Intn(2))
		v := pick(rng, stringValues)
		add(fmt.Sprintf(`where(ALL(%s)=NOCASE("%s"))`, strings.Join(fs, ", "), v),
			"ALL NOCASE")
	}

	// COUNT/BYTES with field
	for i := 0; i < 100; i++ {
		f := pick(rng, fields)
		gf := pick(rng, fields)
		calc := pick(rng, []string{"COUNT", "count", "BYTES", "bytes"})
		add(fmt.Sprintf("groupby(%s) calculate(%s:%s)", gf, calc, f),
			fmt.Sprintf("calc %s:field", calc))
	}

	// Percentile with full pipeline
	for _, pctl := range []string{"percentile", "pctl", "PERCENTILE", "PCTL"} {
		for _, p := range []string{"50", "90", "95", "99", "99.9"} {
			f := pick(rng, numericFields)
			cond := randCondition(rng)
			gf := pick(rng, fields)
			add(fmt.Sprintf("where(%s) groupby(%s) calculate(%s(%s):%s) sort(desc) limit(100)", cond, gf, pctl, p, f),
				fmt.Sprintf("percentile pipeline %s(%s)", pctl, p))
		}
	}

	// Docker/monitoring-style deep field paths
	deepFields := []string{
		"stats.memory_stats.usage", "stats.cpu_stats.cpu_usage.total_usage",
		"stats.networks.eth0.rx_bytes", "stats.networks.eth0.tx_bytes",
		"stats.networks.eth0.rx_packets", "stats.networks.eth0.tx_packets",
		"stats.networks.eth0.rx_errors", "stats.networks.eth0.tx_errors",
		"source_json.properties.mfaDetail.authMethod",
		"source_json.properties.riskLevelDuringSignIn",
		"process.exe_file.hashes.md5", "process.exe_file.hashes.sha256",
		"process.exe_file.product_name", "process.exe_file.description",
		"xml.eventdata.targetusername", "xml.system.eventid",
	}
	for i := 0; i < 200; i++ {
		f := pick(rng, deepFields)
		op := pick(rng, compOps[:4])
		v := pick(rng, stringValues)
		add(fmt.Sprintf(`where(%s%s"%s")`, f, op, v), fmt.Sprintf("deep field %s", f))
		add(fmt.Sprintf(`where(%s!=null) calculate(average:%s)`, f, f), fmt.Sprintf("null check %s", f))
		gf := pick(rng, fields)
		add(fmt.Sprintf(`where(%s!=null) groupby(%s) calculate(average:%s) sort(desc)`, f, gf, f),
			fmt.Sprintf("deep field pipeline %s", f))
	}

	// 35. from() clause
	logTypes := []string{"asset_auth", "vpn", "firewall", "dns", "endpoint", "cloud", "syslog", "windows_event"}
	for i := 0; i < 200; i++ {
		lt := pick(rng, logTypes)
		cond := randCondition(rng)
		add(fmt.Sprintf(`from(event_type = "%s") where(%s)`, lt, cond),
			fmt.Sprintf("from %s", lt))
		gf := pick(rng, fields)
		add(fmt.Sprintf(`from(log_type = "%s") where(%s) groupby(%s) calculate(count)`, lt, cond, gf),
			fmt.Sprintf("from %s pipeline", lt))
	}

	// 36. loose modifier
	for i := 0; i < 200; i++ {
		cond := randCondition(rng)
		add(fmt.Sprintf("where(%s, loose)", cond), "loose condition")
		v := pick(rng, stringValues)
		add(fmt.Sprintf(`where("%s", loose)`, v), "loose keyword")
		gf := pick(rng, fields)
		add(fmt.Sprintf(`where("%s", loose) groupby(%s)`, v, gf), "loose keyword groupby")
	}

	// 37. Variable references ${var_name}
	for i := 0; i < 200; i++ {
		f := pick(rng, fields)
		v := pick(rng, variables)
		op := pick(rng, compOps[:2])
		add(fmt.Sprintf("where(%s%s%s)", f, op, v), fmt.Sprintf("variable %s", v))
	}
	for i := 0; i < 100; i++ {
		f := pick(rng, fields)
		v := pick(rng, variables)
		add(fmt.Sprintf("where(%s NOT IN [%s])", f, v), fmt.Sprintf("NOT IN variable %s", v))
		add(fmt.Sprintf("where(%s IN [%s])", f, v), fmt.Sprintf("IN variable %s", v))
	}
	for i := 0; i < 100; i++ {
		f := pick(rng, fields)
		v1 := pick(rng, variables)
		v2 := pick(rng, stringValues)
		add(fmt.Sprintf(`where(%s IN [%s, "%s"])`, f, v1, v2), "variable in mixed list")
		add(fmt.Sprintf(`where(%s NOT IN [%s, "%s"])`, f, v1, v2), "NOT IN variable mixed list")
	}
	// Variable with full pipeline
	for i := 0; i < 100; i++ {
		f := pick(rng, fields)
		v := pick(rng, variables)
		gf := pick(rng, fields)
		add(fmt.Sprintf("where(%s NOT IN [%s]) groupby(%s) calculate(count)", f, v, gf),
			"variable pipeline")
	}
	// Variable with from() and loose
	for i := 0; i < 50; i++ {
		lt := pick(rng, logTypes)
		f := pick(rng, fields)
		v := pick(rng, variables)
		add(fmt.Sprintf(`from(event_type = "%s") where(%s NOT IN [%s])`, lt, f, v),
			"from with variable")
	}

	// 38. More pipeline combos to reach target
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
	switch rng.Intn(10) {
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
	case 7: // regex
		return fmt.Sprintf(`%s=/error.*/i`, f)
	case 8: // variable comparison
		v := pick(rng, variables)
		op := pick(rng, compOps[:2])
		return fmt.Sprintf(`%s%s%s`, f, op, v)
	default: // variable in set
		v := pick(rng, variables)
		return fmt.Sprintf(`%s NOT IN [%s]`, f, v)
	}
}

func randDeepBooleanQuery(rng *rand.Rand, depth int) string {
	if depth <= 0 {
		return randCondition(rng)
	}
	switch rng.Intn(5) {
	case 0: // AND group
		left := randDeepBooleanQuery(rng, depth-1)
		right := randDeepBooleanQuery(rng, depth-1)
		return fmt.Sprintf("(%s AND %s)", left, right)
	case 1: // OR group
		left := randDeepBooleanQuery(rng, depth-1)
		right := randDeepBooleanQuery(rng, depth-1)
		return fmt.Sprintf("(%s OR %s)", left, right)
	case 2: // NOT
		inner := randDeepBooleanQuery(rng, depth-1)
		return fmt.Sprintf("NOT (%s)", inner)
	case 3: // nested where
		inner := randDeepBooleanQuery(rng, depth-1)
		return fmt.Sprintf("where(%s)", inner)
	default: // base condition
		return randCondition(rng)
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
