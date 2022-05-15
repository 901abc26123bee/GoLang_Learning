package customertype_test

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(op int) int

func timesSpent(inner IntConv) func(op int) int{
	return func(n int) int {
		start := time.Now()
		ret := inner(n)

		fmt.Println("time spent", time.Since(start).Seconds())
		return ret
	}
}

func slowFn(op int) int {
	time.Sleep(time.Second + 1)
	return op
}

func TestFn2(t *testing.T) {
	tsSF := timesSpent(slowFn)
	t.Log(tsSF(10))
}
