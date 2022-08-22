package godomain

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	fmt.Println(reverseStringList([]string{"111", "222", "333", "444", "555"}))
}

func TestDomain(t *testing.T) {
	domain1, err1 := ParseDomain("com")
	if err1 != nil {
		fmt.Println("1", err1)
	}
	fmt.Println(domain1)
	domain2, err2 := ParseDomain("example.com")
	if err2 != nil {
		fmt.Println("2", err2)
	}
	fmt.Println(domain2)
	domain3, err3 := ParseDomain("www.example.com")
	if err3 != nil {
		fmt.Println("3", err3)
	}
	fmt.Println(domain3)
	domain4, err4 := ParseDomain("xxx.www.example.com")
	if err4 != nil {
		fmt.Println("4", err4)
	}
	fmt.Println(domain4)
	fmt.Println(domain3.GetL2())
	fmt.Println(domain3.GetL2Only())
	fmt.Println(domain3.GetL3())
	fmt.Println(domain3.GetL3Only())
	fmt.Println(domain3.GetLevel())
}
