package select_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func otherTask() string {
	fmt.Println("working on something else...")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
	return "other"
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

// time exceeded control
func TestSelect(t *testing.T) {
	select {
	case ret := <- AsyncService():
		t.Log(ret)
	case <- time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}

// muti select
func TestSelect2(t *testing.T) {
	select {
		case ret := <- AsyncService():
			t.Log(ret)
		case ret := <- AsyncService():
			t.Log(ret)
		default:
			t.Error("No returned")
	}
}