package channel_test

import (
	"fmt"
	"testing"
)

// Multiple Channels
// 當前的 goroutine 阻塞時，就會切換到其他 goroutine

func square(c chan int) {
	fmt.Println("[square] wait fot testNum")
	num := <-c // ubuffere channel has data --> comsume data
	fmt.Println("[square] sent square to squareChan (blocking)")
	c <- num * num // 3. blocking, switch to other go routine
}

func cube(c chan int) {
	fmt.Println("[cube] wait for testNum (blocking)")
	num := <-c // 2. blocking, switch to other go routine
	fmt.Println("[cube] sent square to cubeChan")
	c <- num * num * num // blocking
}

func TestMultiChannels(t *testing.T) {
	fmt.Println("[Test] TestMultiChannels() started")
	// unbuffered channels
	squareChan := make(chan int)
	cubeChan := make(chan int)

	go square(squareChan)
	go cube(cubeChan)

	testNum := 3


	fmt.Println("[Test] sent testNum to squareChan (blocking)")

	// block `go square(squareChan)` and switch to 	`go cube(cubeChan)`
	squareChan <- testNum // 1. block, switch to other goroutine

	fmt.Println("[Test] resuming")

	fmt.Println("[Test] sent testNum to cubeChan")
	cubeChan <- testNum

	fmt.Println("[Test] resuming")

	fmt.Println("[Test] reading from channels (blocking)")
	squareVal, cubVal := <-squareChan, <-cubeChan

	fmt.Println(squareVal, cubVal)
	fmt.Println("[Test] TestMultiChannels() stopped")

	/*
			[Test] TestMultiChannels() started
			[Test] sent testNum to squareChan (blocking)
			[cube] wait for testNum (blocking)
			[square] wait fot testNum
			[square] sent square to squareChan (blocking)
			[Test] resuming
			[Test] sent testNum to cubeChan
			[Test] resuming
			[Test] reading from channels (blocking)
			[cube] sent square to cubeChan
			9 27
			[Test] TestMultiChannels() stopped
	*/
}