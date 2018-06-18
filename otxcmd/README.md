otxcmd
---

CLI interface to AlienVault Open Threat Exchange

```
usage: otxcmd [<flags>] <report> <target>

Flags:
  --help                Show context-sensitive help (also try --help-long and --help-man).
  --apikey="~/.otxcmd"  Path to API key
  --section="general"   Section to query, default general

Args:
  <report>  What to report on [ipv4, domain, hostname]
  <target>  What to report on


otxcmd ipv4 4.2.2.1

Sections:
  general
  geo
  reputation
  url_list
  passive_dns
  malware
  nids_list
  http_scans
ASN:	AS3356 Level 3 Parent, LLC
Base Indicator:
City:
Country Name:	United States
Pulse Info:
  Count: 0
  Pulses:
Type:	IPv4
Type Title:	IPv4
Validation:
Whois:	http://whois.domaintools.com/4.2.2.1


otxcmd --section geo ipv4 4.2.2.1

Area Code:	0
ASN:	AS3356 Level 3 Parent, LLC
Charset:	0
City:
City Data:	true
Continent Code:
Country Code:	US
Country Code(3):	USA
Country Name:	United States
DMA Code:	0
Flag Title:	United States
Flag Url:	/static/img/flags/us.png
Latitude:	37.750999
Longitude:	-97.821999
Postal Code:
Region:
```
