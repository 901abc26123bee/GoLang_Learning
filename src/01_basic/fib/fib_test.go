package fib

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	// var a int = 1
	// var b int = 2

	// var (
	// 	a int = 1
	// 	b int = 2
	// )

	// var (
	// 	a int = 1
	// 	b = 2
	// )

	a := 1
	b := 1
	// fmt.Print(a)
	t.Log(a)
	for i := 0; i < 5; i++ {
		// fmt.Print(" ", b)
		t.Log("", a)
		// tmp := a
		// a = b
		// b = tmp + a
		a, b = b, a
	}
	fmt.Println()
}
