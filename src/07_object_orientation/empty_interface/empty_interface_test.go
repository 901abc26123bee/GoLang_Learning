package emptyinterface_test

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
		return
	}
	if s, ok := p.(string); ok {
		fmt.Println("String", s)
		return
	}
	fmt.Println("Unknown Type")
}

func DoSomething2(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknown Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}

func TestEmptyInterfaceAssertion2(t *testing.T) {
	DoSomething2(10)
	DoSomething2("10")
}