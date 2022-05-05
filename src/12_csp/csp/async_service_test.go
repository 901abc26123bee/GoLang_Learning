package csp_test

import (
	"fmt"
	"testing"
	"time"
)

// csp : Communicating Sequential Process
func service() string{
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else...")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

// chan == channel
func AsyncService() chan string {
	retCh := make(chan string)
	// run func in other groutine and return channel
	go func() {
		ret := service()
		fmt.Println("returned result")
		// palce res in to channel
		// since not using buffered channel
		//   --> block the fellowing code until retCh is comsumed by client
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retch := AsyncService()
	otherTask()
	fmt.Println(<-retch) // get result from channel
	time.Sleep(time.Second + 1)
}

// chan == channel
func AsyncBufferService() chan string {
	// buffered channel : make(cham type, capacity)
	retCh := make(chan string, 1)
	// run func in other groutine and return channel
	go func() {
		ret := service()
		fmt.Println("returned result")
		// palce res in to channel
		// since using buffered channel --> won't block fellowing code
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}

func TestAsyncBufferService(t *testing.T) {
	retch := AsyncBufferService ()
	otherTask()
	fmt.Println(<-retch) // get result from channel
	time.Sleep(time.Second + 1)
}