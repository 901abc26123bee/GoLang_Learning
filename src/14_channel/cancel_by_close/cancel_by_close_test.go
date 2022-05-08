package canclebyclose

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
		case <-cancelChan: // get message from channel
			return true
		default:
			return false // blocked
	}
}

func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{} // since only send one message ==> only one groutines would be cancelled
}

func cancel_2(cancelChan chan struct{}) {
	close(cancelChan) // channel broadcasting ==> all groutines would be cancelled
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelChan chan struct{}) {
			for {
				if isCancelled(cancelChan) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Done")
		}(i, cancelChan)
	}
	cancel_2(cancelChan)
	time.Sleep(time.Millisecond * 1)
}