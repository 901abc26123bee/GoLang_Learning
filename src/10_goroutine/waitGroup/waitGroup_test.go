package waitgroup_test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

/*
	WaitGroup 的用法適合用在需要將單一任務拆成許多次任務，待所有任務完成後才繼續執行的情境。

		💡 這種做法適合用在單純等待任務完成，而不需要從 goroutine 中取得所需資料的情況，
			如果會需要從 goroutine 中返回資料，那麼比較好的做法是使用 channel。

	使用 sync.WaitGroup package 提供的：
		var wg sync.WaitGroup 可以建立 waitgroup，預設 counter 是 0
		wg.Add(delta int) 增加要等待的次數（increment counter），也可以是負值，通常就是要等待完成的 goroutine 數目
		wg.Done() 會把要等待的次數減 1（decrement counter），可以使用 defer wg.Done()
		wg.Wait() 會阻塞在這，直到 counter 歸零，也就是所有 WaitGroup 都呼叫過 done 後才往後執行

*/

var start time.Time

func init() {
	start = time.Now()
}

// 這裡的 wg 需要把 pointer 傳進去 goroutine 中，
// 如果不是傳 pointer 進去而是傳 value 的話，將沒辦法有效把 main goroutine 中的 waitGroup 的 counter 減 1。
func service(wg *sync.WaitGroup, instance int) {
	time.Sleep(time.Duration(instance) * 500 * time.Millisecond)
	fmt.Println("Service called on instance: ", instance, time.Since(start))
	wg.Done() // 4. 減少 counter
}

func TestWaitGroup(t *testing.T) {
	fmt.Println("main() started ", time.Since(start))
	var wg sync.WaitGroup // 1. 建立 waitgroup（empty struct）

	for i := 1; i <=3; i++ {
		wg.Add(1) // 2. 增加 counter
		go service(&wg, i)  // 一共啟動了 3 個 goroutine
	}
	wg.Wait() // 3. blocking 直到 counter 為 0
	fmt.Println("main() stopped ", time.Since(start))

	// main() started  369.956µs
	// Service called on instance:  1 502.524689ms
	// Service called on instance:  2 1.004953792s
	// Service called on instance:  3 1.50483851s
	// main() stopped  1.504949952s
}

func notifying(wg *sync.WaitGroup, s string) {
	fmt.Printf("Staring to notifying %s ... \n", s)
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	fmt.Printf("Finish notifying %s\n", s)
	wg.Done()
}

func notify(services ...string) {
	var wg sync.WaitGroup

	for _, service := range services {
		wg.Add(1) // 添加 counter 的次數
		go notifying(&wg, service)
	}
	wg.Wait() // block 在這，直到 counter 歸零後才繼續往後執行

	fmt.Println("All service notified!")
}

func TestWaitGroup2(t *testing.T) {
	notify("Service-1", "Service-2", "Service-3")

	/*
		Staring to notifying Service-3 ...
		Staring to notifying Service-1 ...
		Finish notifying Service-1
		Staring to notifying Service-2 ...
		Finish notifying Service-2
		Finish notifying Service-3
		All service notified!
	*/
}

// ----------------------------------------------------------------------------
// 如果我們需要使用到 goroutine 中回傳的資料，那個應該要使用 channel 而不是 waitGroup，例如：
func notifying2(res chan string, s string) {
	fmt.Printf("Starting to notifying %s...\n", s)
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	res <- fmt.Sprintf("Finish notifying %s", s)
}

func notify2(services ...string) {
	res := make(chan string)
	var count = 0

	for _, service := range services {
		count++
		go notifying2(res, service)
	}

	for i := 0; i < count; i++ {
		fmt.Println(<- res)
	}
	fmt.Println("All service notified!")
}

func TestWaitFoeReturnValueFromChannel(t *testing.T) {
	notify2("Service-1", "Service-2", "Service-3")

	/*
		Starting to notifying Service-3...
		Starting to notifying Service-2...
		Finish notifying Service-2
		Starting to notifying Service-1...
		Finish notifying Service-1
		Finish notifying Service-3
		All service notified!
	*/
}


