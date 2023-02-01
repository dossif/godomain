package godomain

import (
	"fmt"
	"strings"
)

const domainSplitSeparator = "."

// Domain represents level-1, level-2 and level-3... parts of Domain as tld, domain and record
type Domain struct {
	tld    string
	domain string
	record string
}

// Parse parses domainName string to Domain obj
func Parse(domainName string) *Domain {
	d := reverseStringList(strings.Split(trimAll(domainName), domainSplitSeparator))
	switch len(d) {
	case 1:
		return &Domain{tld: d[0], domain: "", record: ""}
	case 2:
		return &Domain{tld: d[0], domain: d[1]}
	default:
		return &Domain{tld: d[0], domain: d[1], record: strings.Join(reverseStringList(d[2:]), domainSplitSeparator)}
	}
}

// String returns record + domain + tld parts.
// Ex: admin.gmail.google.com >>> admin.gmail.google.com
func (d Domain) String() string {
	return trimDot(fmt.Sprintf("%v.%v.%v", d.record, d.domain, d.tld))
}

// Domain returns domain + tld parts.
// Ex: admin.gmail.google.com >>> google.com
func (d Domain) Domain() string {
	return trimDot(fmt.Sprintf("%v.%v", d.domain, d.tld))
}

// Tld returns tld part only.
// Ex: admin.gmail.google.com >>> com
func (d Domain) Tld() string {
	return trimDot(fmt.Sprintf("%v", d.tld))
}

// Record returns record part only.
// Ex: admin.gmail.google.com >>> admin.gmail
func (d Domain) Record() string {
	return trimDot(fmt.Sprintf("%v", d.record))
}

func reverseStringList(strings []string) []string {
	for i := 0; i < len(strings)/2; i++ {
		j := len(strings) - i - 1
		strings[i], strings[j] = strings[j], strings[i]
	}
	return strings
}

func trimAll(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, " ", "")
	s = trimDot(s)
	return s
}

func trimDot(s string) string {
	s = strings.TrimLeft(s, domainSplitSeparator)
	s = strings.TrimRight(s, domainSplitSeparator)
	return s
}
