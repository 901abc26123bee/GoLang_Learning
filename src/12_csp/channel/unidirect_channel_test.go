package channel_test

import (
	"fmt"
	"testing"
)

// *	unidirectional channels
// 		除了可以建立可讀（read）可寫（write）的 channel 之外，
// 		還可以建立「只可讀（receive-only）」或「只可寫（send-only）」的 channel：
func TestUniDirectionalChannel(t *testing.T) {
	roc := make(<-chan int) //  receive only channel
	soc := make(chan<- int) //  send only channel

	fmt.Printf("receive only channel type is '%T' \n", roc)
	fmt.Printf("send only channel type is '%T' \n", soc)

	// receive only channel type is '<-chan int'
	// send only channel type is 'chan<- int'
}

// *	透過 unidirectional channels 可以增加型別的安全性（type-safety）
// 		但如果我們希望在某一個 goroutine 中只能從 channel 讀取資料，
// 		但在 main goroutine 中可以對這個 channel 讀和寫資料時，
// 		可以透過 go 提供的語法來將 bi-directional channel 轉換成 unidirectional channel

// <-chan 表示 receive only channel type
func greet2(roc <-chan string) {
	fmt.Println("Hello " + <-roc + "!")

	// receive only channel 不能傳送資料
	// invalid operation: cannot send to receive-only type <-chan string
	// roc <- "foo"
}
func TestChangeToUniDirectional(t *testing.T) {
	fmt.Println("Test start")
	c := make(chan string)

	go greet2(c)

	c <- "Jhon"
	fmt.Println("Test stop")

	// Test start
	// Hello Jhon!
	// Test stop
	// fmt.Printf("Try to comsume data %v\n", 	<-c) // panic: test timed out after 30s

}