package untilanyonereplay

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

func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func FirstResponseWithBufferedChan() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	fmt.Println("Before: ", runtime.NumGoroutine())
	fmt.Println(FirstResponse())
	time.Sleep(time.Second * 1)
	fmt.Println("After: ", runtime.NumGoroutine())

	// Before:  2
	// The result is from 6
	// After:  11 ==> blocked due to unbuffered channel
}

func TestFirstResponseWithBufferedChan(t *testing.T) {
	fmt.Println("Before: ", runtime.NumGoroutine())
	fmt.Println(FirstResponseWithBufferedChan())
	time.Sleep(time.Second * 1)
	fmt.Println("After: ", runtime.NumGoroutine())

	// Before:  2
	// The result is from 3
	// After:  2
}
