package godomain

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	fmt.Println(reverseStringList([]string{"111", "222", "333", "444", "555"}))
}

func TestDomain(t *testing.T) {
	testDomain := []string{
		"super.admin.gmail.google.com",
		"admin.gmail.google.com",
		"gmail.google.com",
		"google.com",
		"com",
		"..g  oog   le.com.  ",
	}
	for _, s := range testDomain {
		d := Parse(s)
		fmt.Println("_____", s, "_____")
		fmt.Println("str: ", d)
		fmt.Println("tld: ", d.Tld())
		fmt.Println("dom: ", d.Domain())
		fmt.Println("rec: ", d.Record())
	}
	d := Parse("admin.gmail.google.com")
	fmt.Println(d)
	fmt.Println(d.Tld())
	fmt.Println(d.Domain())
	fmt.Println(d.Record())
}
