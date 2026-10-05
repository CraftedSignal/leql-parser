package leql

import "testing"

func TestRegexAndCIDRExpressionPreservesAdvancedSemantics(t *testing.T) {
	result := ExtractConditions(`where((process.name =~ /mimikatz/i OR process.name !~ /benign/) AND source_address IN [IP(10.0.0.0/8), IP("192.168.0.0/16")])`)
	requireNoParseErrors(t, result)
	if len(result.Conditions) != 3 {
		t.Fatalf("expected 3 conditions, got %d: %#v", len(result.Conditions), result.Conditions)
	}

	regex := result.Conditions[0]
	if regex.Field != "process.name" || regex.Operator != "=~" || regex.Value != "mimikatz" || !regex.CaseInsensitive || regex.LogicalOp != "AND" {
		t.Fatalf("case-insensitive regex condition lost semantics: %#v", regex)
	}
	notRegex := result.Conditions[1]
	if notRegex.Field != "process.name" || notRegex.Operator != "!~" || notRegex.Value != "benign" || notRegex.Negated || notRegex.CaseInsensitive || notRegex.LogicalOp != "OR" {
		t.Fatalf("regex-not-match condition lost semantics: %#v", notRegex)
	}
	cidr := result.Conditions[2]
	if cidr.Field != "source_address" || cidr.Operator != "cidr" || cidr.LogicalOp != "AND" || len(cidr.Alternatives) != 2 {
		t.Fatalf("CIDR list condition lost semantics: %#v", cidr)
	}
	if cidr.Alternatives[0] != "10.0.0.0/8" || cidr.Alternatives[1] != "192.168.0.0/16" {
		t.Fatalf("CIDR alternatives = %#v", cidr.Alternatives)
	}
}
