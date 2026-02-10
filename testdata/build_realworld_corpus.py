#!/usr/bin/env python3
"""Build a comprehensive real-world LEQL corpus from multiple sources."""
import json
import sys

queries = []
seen = set()

def add(query, description, source="real-world"):
    q = query.strip()
    if not q or q in seen:
        return
    seen.add(q)
    queries.append({"query": q, "description": description, "source": source})

# ============================================================
# Source 1: Rapid7 Example Queries (official docs)
# ============================================================
rapid7_examples = [
    ('where(result AND result != SUCCESS) groupby(destination_user) calculate(count)', 'Failed authentication by user'),
    ('where(result ISTARTS-WITH "FAILED") groupby(destination_user) calculate(count)', 'Failed auth alternative'),
    ('groupby(source_user)', 'Admin actions by user'),
    ('groupby(action)', 'All admin actions'),
    ('where(source_user="Arnold Holt")', 'Activity by specific user'),
    ('where(source_user=NOCASE("arnold holt"))', 'Case-insensitive user search'),
    ('where(source_user="Tina Gonzales (Admin)")', 'User with parens in name'),
    ('where(source_user=NOCASE("tina gonzales (admin)"))', 'NOCASE with parens'),
    ('where(source_user="rrn:uba:us:14f8eba8-46c8-474b-a982-29476e7a8bd8:user:JA5G9PI3PC9M")', 'UBA RRN user search'),
    ('where(source_user ICONTAINS admin)groupby(source_user)', 'Users with admin in name'),
    ('where(source_user ICONTAINS admin)groupby(action)', 'Admin actions grouped'),
    ('where(source_user ICONTAINS admin AND action=MEMBER_ADDED_TO_SECURITY_GROUP) groupby(group)', 'Admin group additions'),
    ('where(action="MEMBER_ADDED_TO_SECURITY_GROUP" AND group="vpn-users")groupby(target_user)', 'Users added to VPN group'),
    ('where(action="MEMBER_ADDED_TO_SECURITY_GROUP")groupby(source_user)', 'Accounts adding users to groups'),
    ('where(target_account NOT ICONTAINS "$" AND action=PRIVILEGE_ESCALATION)groupby(target_account, group, source_user, timestamp)', 'Privilege escalation tracking'),
    ('where(action IN [MEMBER_ADDED_TO_SECURITY_GROUP, MEMBER_REMOVED_FROM_SECURITY_GROUP] AND group CONTAINS "-job-admins")', 'Group changes tracking'),
    ('where(/:\\d{2} (?P<host>\\w+)./ AND /4732 EVENT/ OR /\\s636 EVENT/) groupby(host)', 'Windows security group changes'),
    ('where(/:\\d{2} (?P<host>\\w+)./ AND /4740 EVENT/ OR /\\s644 EVENT/) groupby(host)', 'Account lockout events'),
    ('where(/:\\d{2} (?P<host>\\w+)./ AND /1102 EVENT/ OR /\\s517 EVENT/) groupby(host)', 'Audit log cleared events'),
    ('where(/4719 EVENT/ OR /\\s612 EVENT/)', 'System audit policy changed'),
    ('where(/eventCode\\":\\d{4}/) groupby(EVID)', 'Group by event code'),
    ('where(/eventCode\\":\\"(?P<EVID>\\d{4})/)groupby(EVID)', 'Named capture event code'),
    ('where(/computerName\\":\\"(?P<HostName>[\\w\\d\\-]*)/)groupby(HostName)', 'Group by computer name'),
    ('where(severity ICONTAINS "MALICIOUS")groupby(asset)having(count>=5)', 'Malicious alerts by asset'),
    ('where(result AND result != SUCCESS) groupby(destination_user) calculate(count)', 'Failed auth counts'),
    ('where(result starts-with "failed") groupby(destination_user) calculate(count)', 'Failed auth starts-with'),
    ('where(/(?P<ip>\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3})/) groupby(ip) calculate(count)', 'IP address extraction'),
    ('where(service NOT IIN ["krbtgt", "kerberos"] AND result ISTARTS-WITH "failed") groupby(destination_account)', 'Non-Kerberos failed logins'),
    ('where(service NOT IN [krbtgt, kerberos]AND result="FAILED_BAD_PASSWORD")groupby("destination_asset")', 'Non-Kerberos password failures'),
    ('where(result ISTARTS-WITH "FAILED")groupby(destination_user)', 'Failed auth accounts'),
    ('where(/4625 EVENT/ OR /\\s529 EVENT/)', 'Failed login events'),
    ('where(/:\\d{2} (?P<host>\\w+)./ AND /4625 EVENT/ OR /\\s529 EVENT/) groupby(host)', 'Failed logins by host'),
    ('where(file_event="delete" OR "write" OR "modify") calculate(count)', 'File events count'),
    ('where(incoming_bytes>0 AND geoip_country_code NOT IN [US, IE, GB, DE, JP, CA, AU])groupby(geoip_country_code)', 'Downloads from non-US countries'),
    ('where(geoip_country_name!="United States")groupby(geoip_country_name)', 'Traffic from non-US'),
    ('calculate(count)', 'Simple count'),
    ('where(direction=OUTBOUND AND geoip_country_code!=US)groupby(destination_address)calculate(sum:outgoing_bytes)limit(10)', 'Top outbound by country'),
    ('where(direction=INBOUND)groupby(destination_address)calculate(sum:incoming_bytes)limit(10)', 'Top inbound destinations'),
    ('where(direction="OUTBOUND" AND destination_address="52.205.169.150")groupby(user)', 'Outbound to specific IP'),
    ('where(connection_status="DENY") calculate(count)', 'Denied connections count'),
    ('where(direction="OUTBOUND" AND connection_status="DENY")calculate(count)', 'Outbound denied count'),
    ('where(connection_status="ACCEPT" AND direction="OUTBOUND" AND destination_port NOT IN ["443", "80", "53"]) groupby(destination_port)', 'Non-standard outbound ports'),
    ('where(direction=OUTBOUND)groupby(destination_address)calculate(sum:outgoing_bytes)', 'Outbound bytes by destination'),
    ('where(direction=INBOUND)groupby(source_address)calculate(sum:incoming_bytes)', 'Inbound bytes by source'),
    ('where(geoip_country_name IN [Czechia, Russia, "Hong Kong"] AND connection_status = ACCEPT AND direction=INBOUND)groupby(geoip_country_name)', 'Inbound from risky countries'),
    ('where(connection_status = DENY AND source_address NOT IN [IP(10.0.0.0/8),IP(172.27.0.0/16),IP(169.254.0.0/16),IP(192.168.0.0/16),IP(172.16.0.0/16)])', 'Denied non-RFC1918 sources'),
    ('where(user!="unknown" AND connection_status = DENY AND source_address NOT IN [IP(10.0.0.0/8), IP(172.27.0.0/16), IP(169.254.0.0/16), IP(192.168.0.0/16), IP(172.16.0.0/16)])', 'Denied known users non-private'),
    ('where(connection_status="DENY" AND geoip_country_name!="United States") groupby(geoip_country_name) calculate(count)', 'Denied by country'),
    ('where(connection_status=DENY AND direction=INBOUND AND geoip_country_name!="United States") groupby(geoip_country_name) calculate(count)', 'Inbound denied by country'),
    ('where(direction="OUTBOUND" AND outgoing_bytes>50000000 AND geoip_organization="Box.com")', 'Large outbound to Box'),
    ('where(stats.networks.eth0.rx_bytes!=null) calculate(average:stats.networks.eth0.rx_bytes)', 'Network receive bytes avg'),
    ('where(action="OBTAIN")groupby(asset)calculate(unique:ip)having(unique:ip>=5)', 'Assets with multiple IPs'),
    ('where(severity="INFORMATIONAL")groupby(signature)limit(1000)', 'Informational alerts'),
    ('where(severity="INFORMATIONAL" AND signature=/.*ssh.*/i)groupby(signature)', 'SSH informational alerts'),
    ('where(severity="LOW")groupby(description)', 'Low severity alerts'),
    ('groupby(severity)calculate(unique:signature)', 'Unique signatures by severity'),
    ('where(severity="HIGH")groupby(asset)calculate(unique:signature)', 'High severity unique per asset'),
    ('where(severity="CRITICAL")groupby(asset)calculate(unique:signature)', 'Critical alerts per asset'),
    ('where(geoip_country_name="United States")calculate(count)', 'US traffic count'),
    ('where(geoip_country_name="United States")groupby(user)calculate(count)', 'US traffic by user'),
    ('where(geoip_city="San Jose")groupby(user)', 'San Jose users'),
    ('where(geoip_city=/Providence|Framingham|Dallas|Minneapolis|Appleton|Phoenix|Omaha|Melbourne|Tuzla|Leeds|Zurich|Singapore|Toronto/i)groupby(geoip_city)', 'Multi-city regex'),
    ('where(geoip_country_name="Russia")', 'Russian traffic'),
    ('where(service="box")groupby(user)', 'Box service users'),
    ('where(service="o365")groupby(user)', 'O365 users'),
    ('where(geoip_country_code!="US") groupby(geoip_country_name)', 'Non-US by country'),
    ('where(geoip_country_name AND geoip_country_name!=/United States|Canada|Mexico/i AND result=SUCCESS)groupby(geoip_country_name)limit(100)', 'Successful non-NA logins'),
    ('where(source_json.ApplicationId="[insert application ID]")groupby(account)', 'Azure AD app logins'),
    ('where(result=SUCCESS)groupby(service, geoip_country_name, user_agent)', 'Successful logins detail'),
    ('where(source_json.operationName="Sign-in activity" AND result=SUCCESS)groupby(source_json.properties.appDisplayName)', 'Azure successful sign-ins by app'),
    ('where(source_json.operationName="Sign-in activity" AND result=FAILED_ACCOUNT_DISABLED)groupby(user)', 'Disabled account sign-in attempts'),
    ('where(source_json.properties.riskLevelDuringSignIn!=none)groupby("source_json.properties.riskLevelDuringSignIn")', 'Azure risk level sign-ins'),
    ('groupby(source_json.properties.mfaDetail.authMethod)', 'MFA methods used'),
    ('where(source_json.properties.mfaDetail.authMethod=/Text message/i)groupby(result)', 'SMS MFA results'),
    ('where(source_json.properties.mfaDetail.authMethod=/Mobile app verification code/i)groupby(result)', 'Mobile MFA results'),
    ('where(source_json.properties.mfaDetail.authMethod=/Phone call \\(Authentication phone\\)/i)groupby(result)', 'Phone MFA results'),
    ('where(source_json.properties.mfaDetail.authMethod="OATH verification code")groupby(result)', 'OATH MFA results'),
    ('where(direction="OUTBOUND" AND app_protocol_description ISTARTS-WITH "TLS")groupby(app_protocol_description)', 'Outbound TLS protocols'),
    ('where(source_json.outcome.result=SUCCESS)groupby(service)', 'Successful outcomes by service'),
    ('where(severity="High")groupby(type, title)', 'High severity by type'),
    ('groupby(header.name)', 'CEF header names'),
    ('groupby(header.device_event_class_id)', 'CEF device event classes'),
    ('groupby(header.severity)', 'CEF severity levels'),
    ('groupby(source_json.eventCode)', 'Event codes'),
    ('where(source_json.eventCode=1116)groupby(user, asset, file_path)', 'Defender detection events'),
    ('groupby(risk)', 'Risk levels'),
    ('calculate(count)timeslice(1d)', 'Daily log count'),
    ('groupby(source_json.Category, source_json.Severity)', 'Categories by severity'),
    ('where(is_blocked=true)calculate(count)', 'Blocked email count'),
    ('where(is_blocked=true)groupby(source_json.sender)calculate(unique:source_json.sender)', 'Blocked senders'),
    ('where(data.logName ICONTAINS "SECURITY")groupby(data.eventCode)', 'Security event codes'),
    ('where(process.name=/powershell.exe/i AND process.cmd_line NOT IN ["null", /.*ps1.*/i])groupby(process.cmd_line)', 'PowerShell without scripts'),
    ('groupby(process.exe_file.description)', 'Process descriptions'),
    ('groupby(process.hash_reputation.reputation)', 'Hash reputations'),
    ('where(process.name=NOCASE(taskkill.exe)) groupby(process.cmd_line)', 'Taskkill usage'),
    ('where(process.exe_file.description icontains-any ["ssh", "telnet"]) groupby(process.name, hostname, process.username)', 'SSH/Telnet processes'),
    ('where(user="Pete Coors")groupby(file_name)', 'Files accessed by user'),
    ('where(file_name="audit.csv")groupby(user)', 'Users accessing file'),
    ('where(query="safebrowsing.google.com" AND user!="unknown")groupby(user)limit(20)', 'Safe browsing queries'),
    ('groupby(source_json.Workload, action)', 'Workloads and actions'),
    ('groupby(source_user) calculate(count) limit(10)', 'Top 10 users'),
    ('where(public_suffix AND public_suffix NOT IN [com, net, org]) groupby(public_suffix)', 'Websites outside common domains'),
    ('where(public_suffix="ru") groupby(query)', 'Russian websites visited'),
]
for q, d in rapid7_examples:
    add(q, d, "rapid7-docs")

# ============================================================
# Source 2: GitHub Cheat Sheet (Outs1d3r-Net)
# ============================================================
cheat_sheet_1 = [
    ('where(source_user="grp adm" AND action="ACCOUNT_DISABLED")', 'Account disabled by user'),
    ('where(source_user="grp adm" AND action="ACCOUNT_DISABLED") calculate(count)', 'Account disabled count'),
    ('where(asset="iphone_of_siem")', 'Asset search'),
    ('where(4725)', 'Event code search'),
    ('where(client_mac="11:22:33:aa:bb:cc")', 'MAC address search'),
    ('where("192.168.1.1")', 'IP keyword search'),
    ('where(source_ip="192.168.1.1")', 'Source IP filter'),
    ('where(source_address="192.168.1.1")', 'Source address filter'),
    ('where(destination_ip="192.168.1.1")', 'Destination IP filter'),
    ('where(destination_address="192.168.1.1")', 'Destination address filter'),
    ('where(source_port="22")', 'Source port SSH'),
    ('where(destination_port="443")', 'Destination port HTTPS'),
    ('where(http_method="get")', 'HTTP GET requests'),
    ('where(http_method="post")', 'HTTP POST requests'),
    ('where(url="https://google.com/")', 'URL search'),
    ('where(url_host="google.com")', 'URL host search'),
    ('where(direction="OUTBOUND")', 'Outbound direction'),
    ('where(direction="INTERNAL")', 'Internal direction'),
    ('where(/Alice (?P<VARIABLE>\\s*)/)', 'Named capture regex'),
    ('where(/@domain.com/)', 'Domain regex search'),
]
for q, d in cheat_sheet_1:
    add(q, d, "github-outs1d3r-cheatsheet")

# ============================================================
# Source 3: GitHub Cheat Sheet (pawan-regoti)
# ============================================================
cheat_sheet_2 = [
    ('where(connection_status=DENY)groupby(source_address)', 'Denied connections by source'),
    ('where(source_ip=/192.168.1./)', 'Source IP regex'),
    ('where(result=FAILED_BAD_LOGIN)calculate(UNIQUE:source_ip)', 'Unique IPs with bad login'),
    ('where(service=0365 AND result=FAILED_BAD_LOGIN)groupby(source_ip)', 'O365 failed logins by IP'),
    ('where(/facebook/ AND user!="unknown")groupby(user)', 'Facebook access by user'),  # note: uses group_by in original
    ('where(geoip_country_code!=US)groupby(geoip_country_name)sort(desc)', 'Non-US traffic sorted'),
]
for q, d in cheat_sheet_2:
    add(q, d, "github-pawan-cheatsheet")

# ============================================================
# Source 4: Rapid7 r7insight_analytic_packs - Windows PCI
# ============================================================
windows_pci = [
    ('where(data.eventCode=/4692|4693|4694|4695|4688|4696|4689|5712|4928|4929|4930|4931|4934|4935|4936|4937|4662|5169|5137|5138|5141|5136|5139|4932|4933|4668|4665|4666|4667|4818|4875|4876|4868|4869|4870|4871|4872|4873|4874|4877|4878|4879|4880|4881|4882|4883|4884|4885|4886|4887|4888|4889|4890|4891|4892|4893|4894|4895|4896|4897|4898|4899|4900|5120|5121|5122|5123|5124|5125|5126|5127|5145|5142|5144/ OR data.eventCode=/5140|5143|5168|4664|4985|5051|5031|5150|5151|5154|5155|5156|5157|5158|5159|5146|5147|5152|5153|4656|4658|4690|4660|4659|4661|4663|4698|4699|5889|5890|4671|4691|4700|4701|4702|5148|5149|5888|4657|5039|4660|4659|4661|4663|5478|4960|4961|4962|4963|4965|5479|5480|5483|5484|5485|5024|5025|5030|5027|5028|5029|5032|5033|5034|5035|5037|5050|5058|5059|5071|6400|6401|6402|6403|6404|6405|6406|6407|6408|6409|6420|6421|6422|6423|6424|6416|6419|4608|4609|4616|4621|4697|4610|4611|4614|4622|4618|5038|5057|4612|4615|4816|5056|5060|5061|5062|6281|6410|6417|6418/)groupby(data.computerName) limit(10000)', 'Windows PCI object access events'),
    ('where(source_json.eventCode=/4783|4784|4785|4786|4787|4788|4789|4790|4791|4792|4742|4743|4746|4751|4744|4745|4747|4748|4749|4750|4752|4753|4759|4760|4761|4762|4782|4793|4729|4733|4757|4727|4728|4730|4731|4732|4734|4735|4737|4754|4755|4756|4758|4764|4799|4720|4722|4723|4724|4725|4726|4767|4781|4797|4738|4740|4765|4766|4780|4794|4798|5376|5377/)groupby(source_json.computerName)limit(10000)', 'Windows account management events'),
    ('where(data.eventCode=/4783|4784|4785|4786|4787|4788|4789|4790|4791|4792|4742|4743|4746|4751|4744|4745|4747|4748|4749|4750|4752|4753|4759|4760|4761|4762|4782|4793|4729|4733|4757|4727|4728|4730|4731|4732|4734|4735|4737|4754|4755|4756|4758|4764|4799|4720|4722|4723|4724|4725|4726|4767|4781|4797|4738|4740|4765|4766|4780|4794|4798|5376|5377/)groupby(data.computerName) limit(10000)', 'Windows account management data format'),
    ('where(data.eventCode=4624) groupby(xml.eventdata.targetusername) calculate(count) limit(10000)', 'Successful logins by user'),
    ('where(data.eventCode=4624) groupby(xml.eventdata.targetusername)', 'Successful logins list'),
    ('where(data.eventCode=7036 AND "Windows Event Log") groupby(data.computerName) limit(10000)', 'Service changes'),
    ('where(source_json.eventCode=4624) groupby(destination_account) calculate(count) limit(10000)', 'Successful logins IDR format'),
    ('where(source_json.eventCode=4624) groupby(source_json.eventCode)', 'Login event codes'),
    ('groupby(destination_user)', 'Destination users'),
    ('where(xml.system.eventid=/4715|4719|4902|4904|4905|4906|4907|4908|4912|4817|4739|4865|4866|4867|4713|4716|4717|4718|4864|4707|4913|4703|4704|4705|4706|4714|4911|4709|4710|4712|4711/)groupby(xml.system.eventid)', 'Windows policy change events'),
    ('where(destination_port = 80) groupby(source_address)', 'HTTP traffic by source'),
    ('where(data.eventCode=4625) groupby(xml.eventdata.targetusername) calculate(count) limit(10000)', 'Failed logins by user'),
    ('where(data.eventCode=4625) groupby(xml.eventdata.targetusername)', 'Failed logins list'),
    ('where(source_json.eventCode=4625) groupby(destination_account) calculate(count) limit(10000)', 'Failed logins IDR format'),
    ('where(source_json.eventCode=4625) groupby(source_json.eventCode)', 'Failed login event codes'),
    ('where(result=FAILED_BAD_PASSWORD OR result=FAILED_BAD_LOGIN OR result = FAILED_OTHER) calculate(count) timeslice(60)', 'Failed auth trend'),
    ('where(data.eventCode=4672) groupby(data.computerName) limit(10000)', 'Special privilege logons'),
    ('where(source_json.eventCode=4672)groupby(source_json.computerName) limit(10000)', 'Special privilege IDR format'),
    ('where(xml.eventdata.param2=stopped AND xml.system.eventid=7036)groupby(xml.eventdata.param1)', 'Stopped services'),
    ('where(asset_os_family=windows) groupby(asset)', 'Windows assets'),
    ('where(asset_os_family=windows) groupby(file_event)', 'Windows file events'),
    ('where(asset_os_family=windows) groupby(user)', 'Windows users'),
    ('where(asset_os_family=windows) groupby(process)', 'Windows processes'),
    ('where(asset_os_family=windows) groupby(file_name) limit(100)', 'Windows file names'),
    ('where(asset_os_family=windows) calculate(COUNT)timeslice(1440m)', 'Windows daily activity'),
]
for q, d in windows_pci:
    add(q, d, "r7insight-windows-pci")

# ============================================================
# Source 5: Rapid7 r7insight_analytic_packs - Linux PCI
# ============================================================
linux_pci = [
    ('where(type=ANOM_ROOT_TRANS)calculate(count)', 'Root transition anomalies'),
    ('where(type=ANOM_LOGIN_ACCT OR type=ANOM_LOGIN_FAILURES OR type=ANOM_LOGIN_LOCATION OR type=ANOM_LOGIN_SESSIONS OR type=ANOM_LOGIN_TIME OR type=CRYPTO_KEY_USER OR type=USER_ACCT OR type=USER_AUTH OR type=USER_CMD OR type=USER_LOGIN OR type=USER_ERR OR type=USER_LOGOUT)groupby(type)', 'Linux audit event types'),
    ('where(/DAEMON/ AND /success/)groupby(pid)', 'Successful daemon operations'),
    ('where(/EXECVE/)calculate(count)', 'Command execution count'),
    ('where(type=SYSCALL)groupby(comm)', 'System calls by command'),
    ('where(AuditD)groupby(key)', 'AuditD by key'),
    ('where(asset_os_family=linux) groupby(asset)', 'Linux assets'),
    ('where(asset_os_family=linux) groupby(file_event)', 'Linux file events'),
    ('where(asset_os_family=linux) groupby(user)', 'Linux users'),
    ('where(asset_os_family=linux) groupby(process)', 'Linux processes'),
    ('where(asset_os_family=linux) groupby(file_name) limit(100)', 'Linux file names'),
    ('where(asset_os_family=linux) calculate(COUNT)timeslice(1440m)', 'Linux daily activity'),
]
for q, d in linux_pci:
    add(q, d, "r7insight-linux-pci")

# ============================================================
# Source 6: Rapid7 r7insight_analytic_packs - Apache
# ============================================================
apache = [
    ('groupby(http.status) calculate(COUNT:http.status)', 'HTTP status codes'),
    ('groupby(http.path) calculate(COUNT:http.path)', 'HTTP paths'),
    ('calculate(AVERAGE:http.bytes)', 'Average HTTP bytes'),
    ('where(http.status=401) groupby(http.path) calculate(COUNT:http.status)', 'Unauthorized requests'),
    ('where(http.status=200) groupby(http.path) calculate(COUNT:http.status)', 'Successful requests'),
    ('where(http.status=404) groupby(http.path) calculate(COUNT:http.status)', 'Not found requests'),
    ('calculate(unique:http.addr)', 'Unique HTTP addresses'),
    ('where(http.status=500) groupby(http.path) calculate(COUNT:http.status)', 'Server errors'),
]
for q, d in apache:
    add(q, d, "r7insight-apache")

# ============================================================
# Source 7: Rapid7 r7insight_analytic_packs - Nginx
# ============================================================
nginx = [
    ('where(/\\"GET .{1,}\\.php/) calculate(COUNT)', 'PHP requests'),
    ('where(/\\"GET .{1,}\\.jpg/) calculate(COUNT)', 'JPEG requests'),
    ('where(http.status) groupby(http.referer) calculate(COUNT:http.referer)', 'Referrers by status'),
    ('groupby(http.agent) calculate(COUNT:http.agent)', 'User agents'),
    ('where(/\\"GET .{1,}\\.php/i) calculate(COUNT)', 'PHP requests case-insensitive'),
    ('where(/\\"GET .{1,}\\.jpg/i) calculate(COUNT)', 'JPEG requests case-insensitive'),
]
for q, d in nginx:
    add(q, d, "r7insight-nginx")

# ============================================================
# Source 8: Rapid7 r7insight_analytic_packs - Salesforce
# ============================================================
salesforce = [
    ('where(LOGIN_EVENT) calculate(COUNT)', 'Salesforce logins'),
    ('where(LOGOUT_EVENT) calculate(COUNT)', 'Salesforce logouts'),
    ('where(LOGIN_AS_EVENT) calculate(COUNT)', 'Salesforce login-as events'),
]
for q, d in salesforce:
    add(q, d, "r7insight-salesforce")

# ============================================================
# Source 9: Rapid7 r7insight_analytic_packs - Docker
# ============================================================
docker = [
    ('groupby(image) calculate(average:stats.memory_stats.usage) sort(desc)', 'Docker memory by image'),
    ('groupby(image) calculate(average:stats.cpu_stats.cpu_usage.total_usage) sort(desc)', 'Docker CPU by image'),
    ('groupby(name) calculate(average:stats.memory_stats.usage) sort(desc)', 'Docker memory by container'),
    ('groupby(name) calculate(average:stats.cpu_stats.cpu_usage.total_usage) sort(desc)', 'Docker CPU by container'),
    ('calculate(average:stats.memory_stats.usage) timeslice(10m)', 'Docker memory trend'),  # note: timeSlice in original
    ('calculate(average:stats.cpu_stats.cpu_usage.total_usage) timeslice(10m)', 'Docker CPU trend'),
    ("where(type='create') calculate(count)", 'Docker create events'),
    ("where(type='start') calculate(count)", 'Docker start events'),
    ("where(type='stop') calculate(count)", 'Docker stop events'),
    ("where(type='die') calculate(count)", 'Docker die events'),
    ("where(type='attach') calculate(count)", 'Docker attach events'),
    ("where(type='Kill') calculate(count)", 'Docker kill events'),
    ('where(type) groupby(type) calculate(count)', 'Docker events by type'),
    ('where(stats.networks.eth0.rx_bytes!=null) calculate(average:stats.networks.eth0.rx_bytes)', 'Docker rx bytes'),
    ('where(stats.networks.eth0.rx_packets!=null) calculate(average:stats.networks.eth0.rx_packets)', 'Docker rx packets'),
    ('where(stats.networks.eth0.rx_errors!=null) calculate(average:stats.networks.eth0.rx_errors)', 'Docker rx errors'),
    ('where(stats.networks.eth0.rx_dropped!=null) calculate(average:stats.networks.eth0.rx_dropped)', 'Docker rx dropped'),
    ('where(stats.networks.eth0.tx_bytes!=null) calculate(average:stats.networks.eth0.tx_bytes)', 'Docker tx bytes'),
    ('where(stats.networks.eth0.tx_packets!=null) calculate(average:stats.networks.eth0.tx_packets)', 'Docker tx packets'),
    ('where(stats.networks.eth0.tx_errors!=null) calculate(average:stats.networks.eth0.tx_errors)', 'Docker tx errors'),
    ('where(stats.networks.eth0.tx_dropped!=null) calculate(average:stats.networks.eth0.tx_dropped)', 'Docker tx dropped'),
]
for q, d in docker:
    add(q, d, "r7insight-docker")

# ============================================================
# Source 10: Rapid7 docs - Additional examples
# ============================================================
additional = [
    # Select clause examples
    ('select(source_address, destination_address) where(direction="OUTBOUND")', 'Select addresses outbound'),
    ('select(user, action, timestamp) where(action="MEMBER_ADDED_TO_SECURITY_GROUP")', 'Select group add details'),
    # Having examples
    ('where(result!="SUCCESS") groupby(destination_user) calculate(count) having(count>=10)', 'Users with 10+ failures'),
    ('groupby(source_address) calculate(count) having(count>=100)', 'Source addresses with 100+ events'),
    ('groupby(destination_port) calculate(unique:source_address) having(unique:source_address>=5)', 'Ports from 5+ sources'),
    # Sort examples
    ('groupby(source_address) calculate(count) sort(desc)', 'Top source addresses'),
    ('groupby(user) calculate(sum:incoming_bytes) sort(desc) limit(20)', 'Top 20 users by bytes'),
    ('groupby(destination_port) calculate(count) sort(asc)', 'Least used ports'),
    # Timeslice examples
    ('calculate(count) timeslice(1h)', 'Hourly event count'),
    ('where(result!="SUCCESS") calculate(count) timeslice(30m)', '30min failed auth trend'),
    ('where(severity="HIGH") calculate(count) timeslice(1d)', 'Daily high severity trend'),
    # Complex pipeline examples
    ('where(result ISTARTS-WITH "FAILED" AND service NOT IN [kerberos, krbtgt]) groupby(destination_user) calculate(count) having(count>=5) sort(desc) limit(10)', 'Top 10 non-Kerberos failed accounts'),
    ('where(direction=OUTBOUND AND geoip_country_code NOT IN [US, GB, CA, AU, DE]) groupby(geoip_country_name) calculate(sum:outgoing_bytes) sort(desc) limit(25)', 'Top outbound data by country'),
    ('select(user, geoip_country_name, service) where(result=SUCCESS AND geoip_country_name!="United States") groupby(user, geoip_country_name) calculate(count) sort(desc)', 'Foreign successful logins'),
    # Strict equality
    ('where(status==200)', 'Strict equality status'),
    ('where(status!==404)', 'Strict not-equal status'),
    # Triple quotes
    ("where(description='''This is a \"complex\" value''')", 'Triple single quote value'),
    ('where(description="""Another \'complex\' value""")', 'Triple double quote value'),
    # from() clause - log source selection
    ('from(event_type = "asset_auth") where(result = "FAILED_ACCOUNT_LOCKED")', 'From clause with where'),
    ('from(event_type = "vpn") where(result = SUCCESS) groupby(user)', 'From VPN source'),
    ('from(log_type = "firewall") where(direction = "OUTBOUND") groupby(destination_address)', 'From firewall logs'),
    # loose modifier
    ('where(user=admin, loose)', 'Loose matching user'),
    ('where("mimikatz", loose) groupby(hostname, process.cmd_line)', 'Loose keyword search'),
    ('where("password reset", loose) groupby(user)', 'Loose phrase search'),
    ('where(action = "login" AND status = "failed", loose) groupby(source_address)', 'Loose compound condition'),
    # Keyword AND regex search
    ('where("login failed")', 'Keyword search phrase'),
    ('where(/error.*timeout/i)', 'Regex error timeout'),
    ('where(/\\bsudo\\b/)', 'Sudo regex boundary'),
    # ALL() function
    ('where(ALL(source_address, destination_address)="10.0.0.1")', 'ALL both addresses'),
    ('where(ALL(user, source_user)=NOCASE("admin"))', 'ALL user fields nocase'),
    # IP function variations
    ('where(source_address=IP(10.0.0.0/8))', 'Source IP private range'),
    ('where(destination_address=IP(192.168.0.0/16))', 'Dest IP private range'),
    ('where(source_address IN [IP(10.0.0.0/8), IP(172.16.0.0/12), IP(192.168.0.0/16)])', 'Source IP all RFC1918'),
    # Mixed operators
    ('where(status>=400 AND status<500)', 'Client error status range'),
    ('where(incoming_bytes>1000000)', 'Large incoming transfers'),
    ('where(duration>=300)', 'Long duration events'),
    # Nested where
    ('where(where(geoip_country_code NOT IN ["US", "GB"] AND result="SUCCESS"))', 'Nested where foreign success'),
    ('where(where("source_address"!="10.0.0.1" AND "destination_port"="443"))', 'Nested where quoted fields'),
    # Postfix negation
    ('where(source_address NOT IN [IP(10.0.0.0/8), IP(172.16.0.0/12)])', 'Postfix NOT IN IP ranges'),
    ('where(user NOT CONTAINS "service")', 'Postfix NOT CONTAINS'),
    ('where(hostname NOT STARTS-WITH "test-")', 'Postfix NOT STARTS-WITH'),
    # Percentile
    ('calculate(percentile(95):response_time)', 'P95 response time'),
    ('calculate(pctl(99):duration)', 'P99 duration'),
    ('calculate(percentile(50):incoming_bytes)', 'Median incoming bytes'),
    # Standard deviation
    ('groupby(source_address) calculate(standarddeviation:response_time)', 'Stddev response time'),
    ('groupby(user) calculate(sd:incoming_bytes)', 'SD incoming bytes per user'),
    # Complex Windows Security queries
    ('where((data.eventCode=4688 AND eventvwr.msc) OR (data.eventCode=4656 AND eventvwr.msc) OR (data.eventCode=4688 AND wevtutil.exe) OR (data.eventCode=4656 AND wevtutil.exe)) groupby(data.computerName) calculate(count) limit(1000)', 'Event viewer usage'),
    # IIN operator
    ('where(geoip_country_code IIN ["us", "gb", "de"])', 'Case-insensitive country code'),
    # Multi-field OR (comma-separated)
    ('where(source_address,destination_address="10.0.0.1")', 'Multi-field OR match'),
    # Unquoted value after contains
    ('where(process.name CONTAINS powershell)', 'Unquoted contains value'),
    ('where(user ICONTAINS admin)', 'Unquoted icontains value'),
    # List string operators
    ('where(process.name CONTAINS-ANY ["cmd.exe", "powershell.exe", "wscript.exe"])', 'Contains any process'),
    ('where(hostname ICONTAINS-ALL ["dc", "prod"])', 'Contains all hostname parts'),
    ('where(url STARTS-WITH-ANY ["http://malware", "https://evil"])', 'Starts with any URL'),
    # group_by alternative spelling (common in real queries)
    # Note: groupby not group_by in LEQL, but it's accepted
    ('groupby(source_json.Workload, source_json.AccountType)', 'O365 workload by account type'),
]
for q, d in additional:
    add(q, d, "rapid7-additional")

# ============================================================
# Source 11: Common detection patterns
# ============================================================
detections = [
    # Brute force detection
    ('where(result="FAILED_BAD_PASSWORD") groupby(destination_user) calculate(count) having(count>=20) sort(desc)', 'Brute force detection'),
    ('where(result ISTARTS-WITH "failed") groupby(source_address) calculate(count) having(count>=50)', 'Brute force by source IP'),
    # Lateral movement
    ('where(logon_type=3 AND result=SUCCESS) groupby(source_address, destination_address) calculate(count)', 'Network logon mapping'),
    ('where(logon_type=10 AND result=SUCCESS) groupby(source_address, destination_user)', 'RDP sessions'),
    # Data exfiltration
    ('where(direction=OUTBOUND AND outgoing_bytes>100000000) groupby(source_address, destination_address) calculate(sum:outgoing_bytes) sort(desc)', 'Large data transfers'),
    ('where(direction=OUTBOUND AND destination_port NOT IN [80, 443, 53]) groupby(destination_port) calculate(count) sort(desc) limit(50)', 'Non-standard outbound ports'),
    # Account anomalies
    ('where(action="ACCOUNT_CREATED") groupby(source_user, target_user)', 'Account creation tracking'),
    ('where(action="PASSWORD_CHANGED") groupby(source_user, target_user)', 'Password changes'),
    ('where(action="ACCOUNT_DISABLED") groupby(source_user, target_user)', 'Account disablement'),
    ('where(action="ACCOUNT_LOCKED_OUT") groupby(destination_user) calculate(count) sort(desc)', 'Account lockouts'),
    # Malware indicators
    ('where(process.exe_file.hashes.md5="d41d8cd98f00b204e9800998ecf8427e")', 'Known malware MD5'),
    ('where(process.hash_reputation.reputation="MALICIOUS") groupby(process.name, asset)', 'Malicious processes'),
    # DNS anomalies
    ('where(query!="unknown" AND public_suffix NOT IN [com, net, org, edu, gov]) groupby(public_suffix) calculate(count) sort(desc)', 'Unusual DNS TLDs'),
    ('where(query CONTAINS ".onion") groupby(user, query)', 'Tor domain lookups'),
    # VPN monitoring
    ('where(service="sslvpn-session" AND result=SUCCESS) groupby(source_address, geoip_country_name, user)', 'VPN sessions by location'),
    ('where(service="sslvpn-session" AND result ISTARTS-WITH "FAILED") groupby(user) calculate(count) having(count>=5)', 'Failed VPN attempts'),
    # Cloud monitoring
    ('where(source_json.Workload="Exchange" AND action="MailItemsAccessed") groupby(source_user)', 'Exchange mailbox access'),
    ('where(source_json.Workload="SharePoint" AND action CONTAINS "File") groupby(action) calculate(count)', 'SharePoint file operations'),
    ('where(source_json.Workload="AzureActiveDirectory" AND action="UserLoggedIn") groupby(source_user, geoip_country_name)', 'Azure AD logins by country'),
]
for q, d in detections:
    add(q, d, "detection-patterns")

print(json.dumps(queries, indent=2))
print(f"Total unique queries: {len(queries)}", file=sys.stderr)
