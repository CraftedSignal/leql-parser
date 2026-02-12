package leql

import (
	"strings"
	"testing"
)

func FuzzLEQLParser(f *testing.F) {
	seeds := []string{
		// --- Basic comparison operators ---
		`where(field=value)`,
		`where(field!=value)`,
		`where(field>10)`,
		`where(field>=10)`,
		`where(field<10)`,
		`where(field<=10)`,

		// --- Numeric field names (the grammar supports NUMBER in fieldName) ---
		`where(1=1)`,
		`where(3=3)`,
		`where(0=false)`,
		`where(status=200 AND 1=1)`,
		`where(1=1 AND 3=3 AND field=value)`,
		`where(1=1 OR 2=2)`,
		`where(NOT 1=1)`,
		`where(42>10)`,

		// --- String operators ---
		`where(field CONTAINS error)`,
		`where(field ICONTAINS error)`,
		`where(field STARTS-WITH prefix)`,
		`where(field ISTARTS-WITH prefix)`,
		`where(message CONTAINS "critical error")`,
		`where(path STARTS-WITH "/usr/local/bin")`,
		`where(host ICONTAINS "PROD")`,
		`where(log ISTARTS-WITH "WARN")`,

		// --- Set operators ---
		`where(field IN [a, b, c])`,
		`where(field IIN ["A", "B"])`,
		`where(status IN [200, 301, 404, 500])`,
		`where(user IIN ["Admin", "root", "system"])`,
		`where(field NOT IN [a, b])`,
		`where(NOT field IN [x, y, z])`,

		// --- List string operators ---
		`where(field CONTAINS-ANY ["error", "fail"])`,
		`where(field ICONTAINS-ANY ["error", "fail"])`,
		`where(field CONTAINS-ALL ["error", "critical"])`,
		`where(field ICONTAINS-ALL ["error", "critical"])`,
		`where(field STARTS-WITH-ANY ["/var", "/tmp"])`,
		`where(field ISTARTS-WITH-ANY ["/var", "/tmp"])`,
		`where(message CONTAINS-ANY ["login", "logout", "timeout"])`,
		`where(path STARTS-WITH-ANY ["/etc", "/usr", "/var", "/tmp"])`,
		`where(log ICONTAINS-ALL ["error", "database", "timeout"])`,

		// --- Boolean logic ---
		`where(a=1 AND b=2)`,
		`where(a=1 OR b=2)`,
		`where(NOT field=value)`,
		`where((a=1 OR b=2) AND c=3)`,
		`where(a=1 AND b=2 AND c=3)`,
		`where(a=1 OR b=2 OR c=3)`,
		`where((a=1 AND b=2) OR (c=3 AND d=4))`,
		`where(NOT (a=1 OR b=2))`,
		`where(a=1 AND NOT b=2)`,
		`where(NOT a=1 AND NOT b=2)`,
		`where((NOT a=1) OR b=2)`,
		`where(a=1 AND (b=2 OR c=3) AND d=4)`,

		// --- Implicit AND (space-separated) ---
		`where(a=1 b=2)`,
		`where(a=1 b=2 c=3)`,
		`where(status=200 user=admin)`,

		// --- Complex boolean logic ---
		`where((status=200 OR status=301) AND user!=anonymous AND NOT path STARTS-WITH "/health")`,
		`where(severity="HIGH" AND (action=block OR action=drop) AND NOT source_ip=IP(10.0.0.0/8))`,

		// --- Nested where clauses ---
		`where(where(field=value))`,
		`where(where(a=1 AND b=2))`,
		`where(where(where(deep=nested)))`,
		`where(where(status=200) AND user=admin)`,
		`where(where("access_types" = "ReadData"))`,
		`where(where("geoip_country_code" NOT IN ["PL", "UK"] and "result" = "SUCCESS"))`,

		// --- NOCASE wrapper ---
		`where(field=NOCASE(value))`,
		`where(user=NOCASE(admin))`,
		`where(hostname=NOCASE("web-server-01"))`,
		`where(status=NOCASE(success))`,

		// --- IP/CIDR matching ---
		`where(ip=IP(10.0.0.0/8))`,
		`where(source_address=IP(192.168.1.0/24))`,
		`where(destination_ip=IP(172.16.0.0/12))`,
		`where(source_address=IP(10.0.0.1))`,

		// --- Regex patterns ---
		`where(/regex/i)`,
		`where(field=/pattern/)`,
		`where(/error|warning|critical/i)`,
		`where(field=/^admin.*/)`,
		`where(/\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}/)`,
		`where(user=/^(admin|root|system)$/)`,
		`where(/powershell|cmd\.exe|wscript/i)`,

		// --- Dotted field names ---
		`where(source_json.Workload=SharePoint)`,
		`where(process.name=powershell.exe)`,
		`where(source_json.Operation=FileUploaded)`,
		`where(event.category=authentication)`,
		`where(source_json.ClientIP STARTS-WITH "10.")`,
		`where(source_json.properties.status.errorCode!=0)`,
		`where("env_vars.7.parent_val" = "C:\ProgramData")`,

		// --- Quoted field names ---
		`where("source_address" = "111.161.118.40")`,
		`where("geoip_country_code" != "US")`,
		`where("result" = "SUCCESS")`,

		// --- Keyword search ---
		`where("Amazon Boardman")`,
		`where('single quoted keyword')`,

		// --- Multiple clauses ---
		`groupby(field) calculate(count)`,
		`groupby(user, status) calculate(count)`,
		`where(x=1) groupby(y) calculate(sum:z) sort(desc) limit(10)`,
		`where(result=FAILED) groupby(user) having(count>5)`,
		`calculate(count) timeslice(30m)`,
		`calculate(count) timeslice(1h)`,
		`select(hostname, process.name) where(status=200)`,
		`where(status>=300) groupby(status) calculate(count) sort(desc) limit(350)`,
		`where(severity="HIGH") groupby(asset) calculate(unique:signature)`,
		`where(direction=OUTBOUND) groupby(destination_address) calculate(sum:outgoing_bytes)`,
		`where(connection_status=DENY) groupby(source_address)`,
		`where(geoip_country_code!="US") groupby(geoip_country_name)`,
		`where(result=SUCCESS) groupby(service, geoip_country_name, user_agent)`,
		`where(status>=400) groupby(url) calculate(count) sort(desc) limit(50)`,
		`where(action=ALLOW) groupby(source_address, destination_port) calculate(count) sort(desc) limit(100)`,
		`where(result=FAILED) groupby(user) calculate(count) having(count>10) sort(desc) limit(20)`,
		`calculate(bytes) timeslice(5m)`,
		`groupby(user) calculate(unique:source_address)`,
		`where(severity IN ["HIGH", "CRITICAL"]) groupby(signature) calculate(count) sort(desc)`,

		// --- Real InsightIDR detection patterns ---
		`where(source_json.Workload=SharePoint AND source_json.Operation=FileUploaded) groupby(source_json.UserId) calculate(count)`,
		`where(geoip_country_code!="US" AND result=SUCCESS) groupby(user, geoip_country_name, service)`,
		`where(connection_status=DENY AND destination_port=445) groupby(source_address) calculate(count) sort(desc)`,
		`where(process.name CONTAINS-ANY ["powershell.exe", "cmd.exe", "wscript.exe", "cscript.exe"])`,
		`where(source_json.Operation IN ["Add member to role.", "Add member to group."] AND source_json.Workload=AzureActiveDirectory)`,
		`where(result=FAILED) groupby(user) calculate(count) having(count>5) sort(desc) limit(10)`,
		`where(NOT geoip_country_code IN ["US", "CA", "GB"] AND result=SUCCESS AND service=sslvpn)`,
		`where(destination_port IN [4444, 5555, 6666, 8080, 8443]) groupby(source_address, destination_address)`,
		`where(source_json.properties.status.errorCode!=0 AND source_json.category=SignInLogs)`,
		`where(EventCode=4688 AND process.name ICONTAINS "mimikatz")`,
		`where(EventCode=4625) groupby(user) calculate(count) having(count>10)`,
		`where(source_json.Operation CONTAINS-ANY ["MailItemsAccessed", "Send", "SearchQueryInitiatedExchange"])`,

		// --- Postfix negation ---
		`where(field NOT IN [a, b, c])`,
		`where(user NOT CONTAINS admin)`,
		`where(status NOT IN [200, 301])`,
		`where(path NOT STARTS-WITH "/health")`,

		// --- Edge cases ---
		``,
		`where()`,
		`groupby()`,
		`calculate()`,
		`where(=)`,
		`where(field=)`,
		`where(=value)`,
		`where(field field)`,
		`where(field=value AND)`,
		`where(AND field=value)`,
		`where(OR)`,
		`where(NOT)`,
		`where(NOT NOT field=value)`,
		`where((()))`,
		`where(field IN [])`,
		`where(field CONTAINS "")`,
		`where(field="")`,
		`where(  field  =  value  )`,
		`sort(asc)`,
		`sort(desc)`,
		`limit(0)`,
		`limit(999999)`,
		`timeslice(1s)`,
		`timeslice(24h)`,
		`from(logsource) where(field=value)`,
		`calculate(average:response_time)`,
		`calculate(min:latency)`,
		`calculate(max:bytes)`,
		`calculate(standarddeviation:response_time)`,
		`calculate(sd:latency)`,
		`calculate(unique:user)`,

		// --- Variable conditions ---
		`where($variable=value)`,

		// --- Triple-quoted strings ---
		`where(field="""triple quoted""")`,
		`where(field='''triple single''')`,

		// --- Strict equality ---
		`where(field==value)`,
		`where(field!==value)`,

		// --- ALL() field list ---
		`where(ALL(field1, field2)=value)`,
		`where(ALL(a, b, c) IN [x, y])`,

		// --- Multi-field conditions ---
		`where(field1, field2 = value)`,
		`where(src_ip, dst_ip = "10.0.0.1")`,
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, query string) {
		// Must not panic -- that's the only requirement
		result := ExtractConditions(query)
		_ = result
	})
}

// TestConditionCompleteness verifies that ExtractConditions extracts all expected
// field names from a variety of LEQL query patterns. This is a deterministic test
// that catches regressions in field extraction across all supported syntax forms.
func TestConditionCompleteness(t *testing.T) {
	tests := []struct {
		name           string
		query          string
		expectedFields []string
	}{
		{
			name:           "simple equality",
			query:          `where(status=200)`,
			expectedFields: []string{"status"},
		},
		{
			name:           "numeric field name standalone",
			query:          `where(1=1)`,
			expectedFields: []string{"1"},
		},
		{
			name:           "numeric field names with surrounding conditions",
			query:          `where(status=200 AND 1=1 AND user=admin)`,
			expectedFields: []string{"status", "1", "user"},
		},
		{
			name:           "multiple AND conditions",
			query:          `where(action=block AND severity=HIGH AND protocol=TCP)`,
			expectedFields: []string{"action", "severity", "protocol"},
		},
		{
			name:           "multiple OR conditions",
			query:          `where(status=200 OR status=301 OR status=404)`,
			expectedFields: []string{"status"},
		},
		{
			name:           "mixed AND/OR conditions",
			query:          `where((status=200 OR status=404) AND user=admin)`,
			expectedFields: []string{"status", "user"},
		},
		{
			name:           "NOT conditions",
			query:          `where(NOT status=404 AND user=admin)`,
			expectedFields: []string{"status", "user"},
		},
		{
			name:           "CONTAINS operator",
			query:          `where(message CONTAINS error AND severity=HIGH)`,
			expectedFields: []string{"message", "severity"},
		},
		{
			name:           "STARTS-WITH operator",
			query:          `where(path STARTS-WITH "/var" AND host=webserver)`,
			expectedFields: []string{"path", "host"},
		},
		{
			name:           "ICONTAINS operator",
			query:          `where(user ICONTAINS admin)`,
			expectedFields: []string{"user"},
		},
		{
			name:           "IN set operator",
			query:          `where(status IN [200, 301, 404])`,
			expectedFields: []string{"status"},
		},
		{
			name:           "IIN set operator",
			query:          `where(user IIN ["Admin", "root"])`,
			expectedFields: []string{"user"},
		},
		{
			name:           "CONTAINS-ANY list operator",
			query:          `where(message CONTAINS-ANY ["error", "fail", "critical"])`,
			expectedFields: []string{"message"},
		},
		{
			name:           "CONTAINS-ALL list operator",
			query:          `where(log CONTAINS-ALL ["error", "database"])`,
			expectedFields: []string{"log"},
		},
		{
			name:           "NOCASE wrapper",
			query:          `where(hostname=NOCASE(webserver) AND status=200)`,
			expectedFields: []string{"hostname", "status"},
		},
		{
			name:           "IP CIDR matching",
			query:          `where(source_address=IP(10.0.0.0/8) AND destination_port=443)`,
			expectedFields: []string{"source_address", "destination_port"},
		},
		{
			name:           "dotted field names",
			query:          `where(source_json.Workload=SharePoint AND source_json.Operation=FileUploaded)`,
			expectedFields: []string{"source_json.Workload", "source_json.Operation"},
		},
		{
			name:           "regex field condition",
			query:          `where(field=/pattern/ AND status=200)`,
			expectedFields: []string{"field", "status"},
		},
		{
			name:           "implicit AND (space-separated)",
			query:          `where(status=200 user=admin)`,
			expectedFields: []string{"status", "user"},
		},
		{
			name:           "postfix NOT IN",
			query:          `where(geoip_country_code NOT IN ["US", "CA"] AND result=SUCCESS)`,
			expectedFields: []string{"geoip_country_code", "result"},
		},
		{
			name:           "real InsightIDR brute force detection",
			query:          `where(result=FAILED) groupby(user) calculate(count)`,
			expectedFields: []string{"result"},
		},
		{
			name:           "real InsightIDR geo anomaly",
			query:          `where(geoip_country_code!="US" AND result=SUCCESS AND service=sslvpn)`,
			expectedFields: []string{"geoip_country_code", "result", "service"},
		},
		{
			name:           "real InsightIDR process monitoring",
			query:          `where(process.name CONTAINS-ANY ["powershell.exe", "cmd.exe", "wscript.exe"])`,
			expectedFields: []string{"process.name"},
		},
		{
			name:           "complex nested boolean with NOT and OR",
			query:          `where((action=block OR action=drop) AND NOT source_address=IP(10.0.0.0/8) AND severity=HIGH)`,
			expectedFields: []string{"action", "source_address", "severity"},
		},
		{
			name:           "quoted field names",
			query:          `where("source_address" = "10.0.0.1" AND "result" = "SUCCESS")`,
			expectedFields: []string{"source_address", "result"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractConditions(tt.query)

			extractedFields := make(map[string]bool)
			for _, c := range result.Conditions {
				extractedFields[strings.ToLower(c.Field)] = true
			}

			for _, expected := range tt.expectedFields {
				if !extractedFields[strings.ToLower(expected)] {
					t.Errorf("field %q not extracted\n  query: %s\n  extracted: %v\n  errors: %v",
						expected, tt.query, extractedFields, result.Errors)
				}
			}
		})
	}
}
