package godomain

import (
	"fmt"
	"regexp"
	"strings"
)

const domainSeparator = "."
const domainRegex = "^([a-z0-9.-]*)$"

type Domain struct {
	tld   string
	l2    string
	l3    string
	level int
}

func ParseDomain(name string) (Domain, error) {
	m, _ := regexp.MatchString(domainRegex, name)
	if !m {
		return Domain{}, fmt.Errorf("domain invalid by regex %v", domainRegex)
	}
	d := reverseStringList(strings.Split(name, domainSeparator))
	switch len(d) {
	case 1:
		return Domain{}, fmt.Errorf("unsupported domain level tld-only")
	case 2:
		return Domain{tld: d[0], l2: d[1], level: 2}, nil
	case 3:
		return Domain{tld: d[0], l2: d[1], l3: d[2], level: 3}, nil
	default:
		return Domain{}, fmt.Errorf("unsupported domain level greater than 3")
	}
}

func (d Domain) String() string {
	switch d.level {
	case 2:
		return fmt.Sprintf("%v.%v", d.l2, d.tld)
	case 3:
		return fmt.Sprintf("%v.%v.%v", d.l3, d.l2, d.tld)
	}
	return ""
}

func (d Domain) GetTLD() string {
	return d.tld
}

func (d Domain) GetL2() string {
	return fmt.Sprintf("%v.%v", d.l2, d.tld)
}

func (d Domain) GetL2Only() string {
	return d.l2
}

func (d Domain) GetL3() (string, error) {
	if d.l3 == "" {
		return "", fmt.Errorf("domain level 3 is not set")
	}
	return fmt.Sprintf("%v.%v.%v", d.l3, d.l2, d.tld), nil
}

func (d Domain) GetL3Only() (string, error) {
	if d.l3 == "" {
		return "", fmt.Errorf("domain level 3 is not set")
	}
	return d.l3, nil
}

func (d Domain) GetLevel() int {
	return d.level
}

func reverseStringList(strings []string) []string {
	for i := 0; i < len(strings)/2; i++ {
		j := len(strings) - i - 1
		strings[i], strings[j] = strings[j], strings[i]
	}
	return strings
}
