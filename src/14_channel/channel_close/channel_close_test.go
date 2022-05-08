package channelclose_test

import (
	"fmt"
	"sync"
	"testing"
)

/*
	1. Pass data to closed channel ==> cause panic
	2. v, ok <- ch ==> ok is a boolean: true ==> normally accept data from channel, false -> channel is closed
	2. All channel's receivers return from blocked waiting state if 'ok' return false.
			This is often used in broadcasting ==> send exiting signal to all subscribers

	當 channel 已經被關閉時，ok 會是 false，value 則會是 zero value
	⚠️ 只有 sender 可以使用 close，receiver 使用的話會發生 panic。
*/


func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			if data, ok := <-ch; ok { // ok is boolean: true -> normally accept data from channel, false -> channel is closed
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChannels(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}

