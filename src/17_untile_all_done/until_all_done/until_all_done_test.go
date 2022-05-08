package untilalldone

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("The result is from %d", id)
}

func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	finnalRet := ""
	for j := 0; j < numOfRunner; j++ {
		finnalRet += <-ch + "\n"
	}
	return finnalRet
}

func TestAllResponse(t *testing.T) {
	fmt.Println("Before: ", runtime.NumGoroutine())
	fmt.Println(AllResponse())
	time.Sleep(time.Second * 1)
	fmt.Println("After: ", runtime.NumGoroutine())

	// Before:  2
	// The result is from 5
	// The result is from 7
	// The result is from 6
	// The result is from 2
	// The result is from 4
	// The result is from 0
	// The result is from 1
	// The result is from 3
	// The result is from 8
	// The result is from 9
	// After:  2  ==> block until all completed due to finnalRet for loop (keep consuming data from channel)
}
