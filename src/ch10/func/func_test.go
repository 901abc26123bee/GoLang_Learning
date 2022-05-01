package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMutiValue() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	_, b := returnMutiValue()
	t.Log(b)
}

func timesSpent(inner func (op int) int) func(op int) int{
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
	a, _ := returnMutiValue()
	t.Log(a)
	tsSF := timesSpent(slowFn)
	t.Log(tsSF(10))
}


// ...
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

// defer
func TestVarParams(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4, 5))
	t.Log(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

// defer
func Clear() {
	fmt.Println("Clear Resources")
}
func TestDefer(t *testing.T) {
	defer Clear() // execute lastly
	fmt.Println("Start")
	panic("err") // code behind panic() won't be executed, defer still execute
	fmt.Println("End")
}