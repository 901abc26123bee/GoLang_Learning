package cocurrency_test

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"testing"
	"time"
)

// ------------------------------ Concurrency Pattern ----------------------------------
// Generator
// fib 會回傳 read-only channel
func fib(length int) <-chan int {
	c := make(chan int, length)

	// run generation concurrently
	go func() {
			for i, j := 0, 1; i < length; i, j = i+j, i {
					c <- i
					time.Sleep(time.Second)
			}

			close(c)
	}()

	// return channel
	return c
}

func TestConcurrency(t *testing.T) {
	for fn := range fib(10) {
			fmt.Println("Current fibonacci number is ", fn)
	}
	/*
	Current fibonacci number is  0
	Current fibonacci number is  1
	Current fibonacci number is  1
	Current fibonacci number is  2
	Current fibonacci number is  3
	Current fibonacci number is  5
	Current fibonacci number is  8
	*/
}

// https://github.com/kevchn/go-concurrency-patterns
// goroutine is launched inside the called function (more idiomatic)
// multiple instances of the generator may be called
func generator(msg string) <-chan string { // returns receive-only channel
	ch := make(chan string)
	go func() { // anonymous goroutine
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Second)
		}
	}()
	return ch
}

func TestGenerator2(t *testing.T) {
	ch := generator("Hello")
	for i := 0; i < 5; i++ {
		fmt.Println(<- ch)
	}
	/*
	Hello 0
	Hello 1
	Hello 2
	Hello 3
	Hello 4
	*/
}


// ----------------------------------------------------------------------------
// WaitGroup 搭配 Channel
// ⚠️ 邏輯上，可以單獨使用 channel 就好，不需要使用 WaitGroup。

func worker(wg *sync.WaitGroup, c chan<- int, i int) {
	defer wg.Done()
	fmt.Println("[worker] start i:", i)
	time.Sleep(time.Second * 1)
	c <- i
	fmt.Println("[worker] finish i:", i)
}
func TestCocurrency3(t *testing.T) {
	numOfFacilities := 6
	var wg sync.WaitGroup

	c := make(chan int, numOfFacilities)

	for i := 0; i < numOfFacilities; i++ {
			fmt.Println("[main] add i: ", i)
			wg.Add(1)
			go worker(&wg, c, i)
	}

	wg.Wait()

	var numbers []int
	for i := 0; i < numOfFacilities; i++ {
			numbers = append(numbers, <-c)
	}
	fmt.Println("[main] ---all finish---", numbers)

	defer close(c)

	/*
	[main] add i:  0
	[main] add i:  1
	[main] add i:  2
	[main] add i:  3
	[main] add i:  4
	[main] add i:  5
	[worker] start i: 5
	[worker] start i: 1
	[worker] start i: 2
	[worker] start i: 3
	[worker] start i: 4
	[worker] start i: 0
	[worker] finish i: 0
	[worker] finish i: 5
	[worker] finish i: 2
	[worker] finish i: 3
	[worker] finish i: 4
	[worker] finish i: 1
	[main] ---all finish--- [0 4 5 1 2 3]
	*/
}


// ----------------------------------------------------------------------------
func controller(c chan string, wg *sync.WaitGroup) {
	fmt.Println("controller() start and block")
	wg.Wait()
	fmt.Println("controller() unblock and close channel")
	close(c)
	fmt.Println("controller() end")
}

func printString(s string, c chan string, wg *sync.WaitGroup) {
	fmt.Println(s)
	wg.Done()
	c <- "Done printing: " + s
}

func TestCocurrency4(t *testing.T) {
	fmt.Println("main() start")
	c := make(chan string)
	var wg sync.WaitGroup
	for i := 1; i <= 4; i++ {
			go printString("Hello ~ "+strconv.Itoa(i), c, &wg)
			wg.Add(1)
	}

	go controller(c, &wg)

	for message := range c {
			fmt.Println(message)
	}

	fmt.Println("main() end")
}

// -------------------------------------------------------------------------------------------
// A Tour of Go: https://tour.golang.org/concurrency/2
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
			sum += v
	}

// STEP 3：把加總後的值丟回 channel
	c <- sum // send sum to c
}

func TestCocurrency5(t *testing.T) {
	s := []int{7, 2, 8, -9, 4, 0}

  // STEP 1：建立一個 channel，該 channel 會傳出 int
    c := make(chan int)

  // STEP 2：使用 goroutine，並把 channel 當成參數傳入
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)

  // STEP 4：從 channel 取得計算好的結果
    x, y := <-c, <-c

  // ⚠️ 寫在這裡的內容會在 channel 傳回結果後才會被執行...

    fmt.Println(x, y, x+y)
}


func TestCocurrency6(t *testing.T) {
	links := []string{
			"http://google.com",
			"http://facebook.com",
			"http://stackoverflow.com",
			"http://golang.com",
			"http://amazon.com",
	}

	for _, link := range links {
			checkLink(link)  // 這裡會被阻塞
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
			fmt.Println(link, "might be down!")
			return
	}
	fmt.Println(link, "is up!")
}