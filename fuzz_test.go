package leql

import "testing"

func FuzzLEQLParser(f *testing.F) {
	seeds := []string{
		`where(field=value)`,
		`where(field!=value)`,
		`where(field>10)`,
		`where(field>=10)`,
		`where(field<10)`,
		`where(field<=10)`,
		`where(field CONTAINS error)`,
		`where(field ICONTAINS error)`,
		`where(field STARTS-WITH prefix)`,
		`where(field ISTARTS-WITH prefix)`,
		`where(field IN [a, b, c])`,
		`where(field IIN ["A", "B"])`,
		`where(field CONTAINS-ANY ["error", "fail"])`,
		`where(field ICONTAINS-ANY ["error", "fail"])`,
		`where(field CONTAINS-ALL ["error", "critical"])`,
		`where(field ICONTAINS-ALL ["error", "critical"])`,
		`where(field STARTS-WITH-ANY ["/var", "/tmp"])`,
		`where(field ISTARTS-WITH-ANY ["/var", "/tmp"])`,
		`where(a=1 AND b=2)`,
		`where(a=1 OR b=2)`,
		`where(NOT field=value)`,
		`where((a=1 OR b=2) AND c=3)`,
		`where(/regex/i)`,
		`where(field=/pattern/)`,
		`where(field=NOCASE(value))`,
		`where(ip=IP(10.0.0.0/8))`,
		`where(name="John Doe")`,
		`where(name='John Doe')`,
		`where(source_json.Workload=SharePoint)`,
		`where(process.name=powershell.exe)`,
		`where("Amazon Boardman")`,
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
		``,
		`where()`,
		`groupby()`,
		`calculate()`,
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, query string) {
		// Must not panic — that's the only requirement
		result := ExtractConditions(query)
		_ = result
	})
}
