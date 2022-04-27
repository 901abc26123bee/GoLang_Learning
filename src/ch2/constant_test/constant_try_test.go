package constan_test

import (
	// "fmt"
	"testing"
)

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestConstantTry(t *testing.T) {
	// fmt.Print(Monday, Tuesday, Wednesday)
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConstantTry1(t *testing.T) {
	a := 7
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}