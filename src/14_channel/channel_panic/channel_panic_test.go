package channelpanic_test

import "testing"

func Test_channel_blocking_case(t *testing.T) {
	// 示例1。
	ch1 := make(chan int, 1)
	ch1 <- 1
	// ch1 <- 2 // channel 已滿，因此這裡會造成阻塞。

	// 示例2。
	ch2 := make(chan int, 1)
	// elem, ok := <-ch2 // channel 已空，因此這裡會造成阻塞。
	// _, _ = elem, ok
	ch2 <- 1

	// 示例3。
	var ch3 chan int
	//ch3 <- 1 // channel 的值為nil，因此這裡會造成永久的阻塞！
	//<-ch3 // channel 的值為nil，因此這裡會造成永久的阻塞！
	_ = ch3
}