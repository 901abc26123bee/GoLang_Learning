package goroutine_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
		// go func() {
		// 	fmt.Println(i)
		// }()
	}
	time.Sleep(time.Millisecond * 50)
}
/*
	Overview
	Goroutines can be thought of as a lightweight thread
	that has a separate independent execution and which can execute concurrently with other goroutines.
	It is a function or method that is executing concurrently with other goroutines.
	It is entirely managed by the GO runtime.
	Golang is a concurrent language. Each goroutine is an independent execution.
	It is goroutine that helps achieve concurrency in golang

	Start a go routine
	Golang uses a special keyword ‘go’  for starting a goroutine.
	To start one just add go keyword before a function or method call.
	That function or method will now be executed in the goroutine.
	Note that it is not the function or method which determines if it is a goroutine.
	If we call that method or function with a go keyword then that function or method is said to be executing in a goroutine.

*/

// ------------------------------ time.Sleep ----------------------------------
// 使 Goroutine 休眠，讓其他的 Goroutine 在 main 結束前有時間執行完成。
// 缺點：休息指定時間可能會比 Goroutine 需要執行的時間長或短，太長會耗費多餘的時間，太短會使其他 Goroutine 無法完成
func start() {
	fmt.Println("Starting...")
}

func processing() {
	fmt.Println("Processing...")
}

func end() {
	fmt.Println("End")
}

func TestNormalFunction(t *testing.T) {
	start()
	processing()
	end()

	// Starting...
	// Processing...
	// End
}

func TestFunctionWithGo(t *testing.T) {
	start()
	// processing()) will be called as a goroutine which will execute asynchronously.
	// goroutine will be executed concurrently in the background.
	go processing()
	end()

	// Starting...
	// End
	// Processing...
}

// ------------------------------ sync.WaitGroup ----------------------------------
// 產生與想要等待的 Goroutine 同樣多的 WaitGroup Counter
// 將 WaitGroup 傳入 Goroutine 中，在執行完成後叫用 wg.Done() 將 Counter 減一
// wg.Wait() 會等待直到 Counter 減為零為止
// 優點: 避免時間預估的錯誤
// 缺點: 需要手動配置對應的 Counter
func say(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(s , i)
	}
}

func TestSyncWaitGroup(t *testing.T) {
	// 生與想要等待的 Goroutine 同樣多的 WaitGroup Counter
	// 將 WaitGroup 傳入 Goroutine 中，在執行完成後叫用 wg.Done() 將 Counter 減一
	// wg.Wait() 會等待直到 Counter 減為零為止
	wg := new(sync.WaitGroup)
	wg.Add(2)
	
	go say("world", wg)
	go say("hello", wg)

	wg.Wait()

	// world 0
	// hello 0
	// world 1
	// hello 1
	// world 2
	// hello 2
	// world 3
	// hello 3
	// world 4
	// hello 4
}

//----------------------------- channel -----------------------------------
// Channel
// 最後介紹的是使用 Channel 等待, 原為 Goroutine 溝通時使用的，但因其阻塞的特性，使其可以當作等待 Goroutine 的方法。
// 優點: 避免時間預估的錯誤 語法簡潔
// Channel 阻塞的方法為 Go 語言中等待的主要方式。
func say2(s string, c chan string) {
	for i := 0; i < 5; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(s, i)
	}
	c <- "FINISH"
}

func TestWaitByUnbufferedChannel(t *testing.T) {
	ch := make(chan string)

	// 起了兩個 Goroutine(say2("world", ch), say2("hello", ch)) ，
	// 因此需要等待兩個 FINISH 推入 unbuffered Channel 中才能結束 Main Goroutine。
	go say2("world", ch)
	go say2("hello", ch)

	<-ch
	<-ch

	// hello 0
	// world 0
	// hello 1
	// world 1
	// world 2
	// hello 2
	// hello 3
	// world 3
	// hello 4
	// world 4
}

/*
	* 一、CSP
		CSP 是一門形式語言（類似於 ℷ calculus），用於描述並發系統中的互動模式
		Golang，其實只用到了 CSP 的很小一部分，即理論中的 Process/Channel（對應到語言中的 goroutine/channel）：
		這兩個並發原語之間沒有從屬關係， Process 可以訂閱任意個 Channel，Channel 也並不關心是哪個 Process 在利用它進行通信；
		Process 圍繞 Channel 進行讀寫，形成一套有序阻塞和可預測的並發模型。
		This approach ensures that only one goroutine access to the data at a given time.

	*	二、Golang CSP
		與主流語言通過共享內存來進行並發控制方式不同，
		Go 語言採用了 CSP 模式。這是一種用於描述兩個獨立的並發實體通過共享的通訊 Channel（管道）進行通信的並發模型。
		Golang 就是藉用CSP模型的一些概念為之實現並發進行理論支持，其實從實際上出發，go語言並沒有，完全實現了CSP模型的所有理論，僅僅是藉用了 process和channel這兩個概念。 
		process是在go語言上的表現就是 goroutine 是實際並發執行的實體，每個實體之間是通過channel通訊來實現數據共享。
			- Go協程goroutine: 是一種輕量線程，它不是操作系統的線程，而是將一個操作系統線程分段使用，通過調度器實現協作式調度。是一種綠色線程，微線程，它與Coroutine協程也有區別，能夠在發現堵塞後啟動新的微線程。 
			- 通道channel: 類似Unix的Pipe，用於協程之間通訊和同步。協程之間雖然解耦，但是它們和Channel有著耦合。

	* 三、Channel
		Channel 在 gouroutine 間架起了一條管道，在管道里傳輸數據，實現 gouroutine 間的通信；
		由於它是線程安全的，所以用起來非常方便；channel 還提供 “先進先出” 的特性；
		它還能影響 goroutine 的阻塞和喚醒。
		* Do not communicate by sharing memory; instead, share memory by communicating.
			這就是 Go 的並發哲學，它依賴 CSP 模型，基於 channel 實現。

				chan T // 聲明一個雙向通道
				chan<- T // 聲明一個只能用於發送的通道
				<-chan T // 聲明一個只能用於接收的通道COPY

		因為 channel 是一個引用類型，所以在它被初始化之前，它的值是 nil，channel 使用 make 函數進行初始化。
		- 可以向它傳遞一個 int 值，代表 channel 緩衝區的大小（容量），構造出來的是一個緩衝型的 channel；
		- 不傳或傳 0 的，構造的就是一個非緩衝型的 channel。

		* Channel 分為兩種：帶緩衝、不帶緩衝。對不帶緩衝的 channel 進行的操作實際上可以看作 “同步模式”，帶緩衝的則稱為 “異步模式”

		同步模式下，發送方和接收方要同步就緒，只有在兩者都 ready 的情況下，數據才能在兩者間傳輸（後面會看到，實際上就是內存拷貝）。
		否則，任意一方先行進行發送或接收操作，都會被掛起，等待另一方的出現才能被喚醒。
		異步模式下，在緩衝槽可用的情況下（有剩餘容量），發送和接收操作都可以順利進行。
		否則，操作的一方（如寫入）同樣會被掛起，直到出現相反操作（如接收）才會被喚醒。
		小結一下：
			* 同步模式下，必須要使發送方和接收方配對，操作才會成功，否則會被阻塞；
			* 異步模式下，緩衝槽要有剩餘容量，操作才會成功，否則也會被阻塞。

	* 四、Goroutine
		- 用戶空間 避免了內核態和用戶態的切換導致的成本
		- 可以由語言和框架層進行調度
		- 更小的棧空間允許創建大量的實例


*/

/*
		Concurrency and Channels in Go

		Concurrency is not the same with running tasks at the same time.
		Running tasks as parallel is a mix of hardware and software requirements.
		But concurrency is about our code design.

		In computer science, concurrency is the ability of different parts or units of a program, algorithm, or problem to be executed out-of-order or in partial order,
		without affecting the final outcome.

		When a thread tries to read a shared variable, another thread may change the value of the shared variable.
		That may cause inconsistency.

		* Concurrency Solutions in Go
			We may solve concurrency problems using Mutexes, Semaphores, Locks etc.
			Basically, when one thread tries to access the shared resource, it locks the critical section,
			in this way other threads can not access the shared resource until the section is unlocked.

			* Go solves those problems in another way.
			* It uses goroutines instead of threads and uses channels instead of accessing to a shared state.


			* The philosophy behind the Go’s concurrency:
			* Do not communicate by sharing memory; instead, share memory by communicating.

		Go channels are highly recommended to use in Go applications when dealing with concurrency.
		But if your problem can not be solved with channels, you may still use other solutions by sync package.
		This package provides low-level components like mutexes.
*/

/*
	1. 從一個 goroutine 切換到另一個 goroutine 的時機點是「當正在執行的 goroutine 阻塞時，就會交給其他 goroutine 做事」
	2. unbuffered channel 指的是 buffer size 為 0 的 channel
	3. 對於 unbuffered channel 來說，不論是從 channel 讀資料（需等到被寫入），或把資料寫入 channel 中時（需等到被讀出），都會阻塞該 goroutine
	4. 對於 buffered channel 來說：
			- 從 channel 讀值時若是 empty buffer 時才會阻塞，否則都是 non-blocking
			- 把資料寫入 channel 中時，寫入 channel 中的 value 數目（n + 1）需要超過 buffer size（n），也就是溢出（overflow）時才會使得該 goroutine 被阻塞；而且一旦 buffer channel 中的值開始被讀取，就會被全部讀完
*/

/*
	goroutines vs threads
	1. goroutines 是由 Go runtime 所管理的輕量化的 thread
	2. goroutines 會在相同的 address space 中執行，因此要存取共享的記憶體必須要是同步的（synchronized）。
	3. 傳統的 Apache 伺服器來說，當每分鐘需要處理 1000 個請求時，每個請求如果都要 concurrently 的運作，
		 將會需要建立 1000 個 threads 或者分派到不同的 process 去做，
			- 如果 OS 的每個 thread 都需要使用 1MB 的 stack size 的話，就會需要 1GB 的記憶體才能撐得住這樣的流量。
			- 但相對於 goroutine 來說，因為 stack size 可以動態增長，因此可以擴充到 1000 個 goroutines，
				每個 goroutine 只需要 2KB（Go 1.4 之後）的 stack size。
*/

/*
threads vs goroutines https://gist.github.com/thatisuday/1a357a725113e1c1cdf174a537287afd#file-threasvsgoroutines-md

		OS thread																			 			 					goroutine

由 OS kernel 管理，相依於硬體	goroutines 						 						 是由 go runtime 管理，不依賴於硬體

OS threads 一般有固定 1-2 MB 的 stack size			 			 					goroutines 的 stack size 約 8KB（自從 Go 1.4 開始為 2KB）

在編譯的時候就決定了 stack 的大小，並且不能增長					 						 由於是在 run-time 管理 stack size，透過分配和釋放 heap storage 可以增長到 1GB

不同 thread 之間沒有簡易的溝通媒介，並且溝通時易有延遲	  	 					goroutine 使用 channels 來和其它的 goroutine 溝通，且低延遲

thread 有 identity，透過 TID 可以辨別 process 中的不同 thread			goroutine 沒有 identity

Thread 有需要 setup 和 teardown cost，												  goroutine 是在 go 的 runtime 中建立和摧毀，和 OS threads 相比非常容易，
需要向 OS 請求資源並在完成時還回去																 因為 go runtime 已經為 goroutines 建立了 thread pools，
																															因此 OS 並不會留意到 coroutines

threads 需要先被 scheduled，
在不同 thread 間切換時的消耗很高，因為 scheduler 需要儲存和還原

*/

// [Go 的並發：Goroutine 與 Channel 介紹](https://peterhpchen.github.io/2020/03/08/goroutine-and-channel.html)
// ---------------------------- Goroutines ----------------------------
// 每個 Go 程式預設都會建立一個 goroutine，這被稱作是 main goroutine，也就是函式 main 中執行的內容
// 所有的 goroutines 都是沒有名稱的（anonymous），因為 goroutine 並沒有 identity
func printHello() {
	time.Sleep(1 * time.Millisecond)
	fmt.Println("Hello World")
}

func TestBlockGoroutine(t *testing.T) {
	fmt.Println("main execution started")

	// call function
	go printHello()

	// block here
	time.Sleep(1 * time.Millisecond)
	fmt.Println("main execution stopped")
}

// anonymous goroutine
func TestAnanymous(t *testing.T) {
	fmt.Println("main() started")

	c := make(chan string)

	// anonymous goroutine
	go func(c chan string) {
			fmt.Println("Hello " + <-c + "!")
	}(c)

	c <- "John"
	fmt.Println("main() ended")
}

// ----------------------------------- Multi Goroutine share data -----------------------------------------
// 1. Mutex
// * 在多個 goroutine 裡對同一個變數total做加法運算，在賦值時無法確保其為安全的而導致運算錯誤，此問題稱為 Race Condition。
/*
		互斥鎖使用在資料結構(struct)中，用以確保結構中變數讀寫時的安全，它提供兩個方法：
		Lock
		Unlock
		在 Lock 及 Unlock 中間，會使其他的 Goroutine 等待，確保此區塊中的變數安全。
*/
// 在執行緒間使用同樣的變數時，最重要的是確保變數在當前的正確性，在沒有控制的情況下極有可能發生問題，下面有個例子：
func TestRaceCondition(t *testing.T) {
	total := 0
    for i := 0; i < 1000; i++ {
        go func() {
            total++
        }()
    }
    time.Sleep(time.Second)
    fmt.Println(total) // 995
}

// * 互斥鎖(sync.Mutex)
// 		在這種狀況下，可以使用互斥鎖(sync.Mutex)來保證變數的安全：
type SafeNumber struct {
	v int
	mux sync.Mutex
}
func TestShareDataWithMutex(t *testing.T) {
	total := SafeNumber{v: 0}
	for i := 0; i < 1000; i++ {
		go func() {
			total.mux.Lock()
			total.v++
			total.mux.Unlock()
		}()
	}
	time.Sleep(time.Second)
	// total.mux.Lock()
	fmt.Println(total.v) // 1000
	// total.mux.Unlock()
}

// 2. channel
/*
		goroutine1 拉出 total 後，Channel 中沒有資料了
		因為 Channel 中沒有資料，因此造成 goroutine2 等待
		goroutine1 計算完成後，將 total 推入 Channel
		goroutine2 等到 Channel 中有資料，拉出後結束等待，繼續做運算

	* 因為 Channel 推入及拉出時等待的特性，被拉出來做計算的值會保證是安全的。
		因為此範例一定要拉出 Channel 資料才能做運算，所以使用非立即阻塞的 Buffered Channel
		上述的三個例子在 main goroutine 中都使用 time.Sleep 避免程式提前結束。
*/
func TestShareDataWithChannel(t *testing.T) {
	total := 0
	ch := make(chan int, 1)
	ch <- total
	for i := 0; i < 1000; i++ {
		go func() {
			ch <- <- ch + 1
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(<-ch) // 1000
}

// ----------------------------------- Buffered Channel -----------------------------------------
// Unbuffered Channel
// Unbuffered Channel，此種 Channel 只要
// * 推入一個資料會造成推入方的等待
// * 拉出時沒有資料會造成拉出方的等待
// 使用 Unbuffered Channel 的壞處是：
// 	* 如果 推入方的執行一次的時間 較 拉取方 短，會造成 推入方 被迫等待 拉取方 才能在做下一次的處理，這樣的等待是不必要並且需要被避免的。
func TestUnBufferedChannelWaitForComsumed(t *testing.T) {
	ch := make(chan string)

	go func() { // calculate goroutine
			fmt.Println("calculate goroutine starts calculating")
			time.Sleep(time.Second) // Heavy calculation
			fmt.Println("calculate goroutine ends calculating")

			ch <- "FINISH" // goroutine 執行會在此被迫等待

			fmt.Println("calculate goroutine finished")
	}()

	time.Sleep(2 * time.Second) // 使 main 比 goroutine 慢
	fmt.Println(<-ch)
	time.Sleep(time.Second)
	fmt.Println("main goroutine finished")

	/*
	calculate goroutine starts calculating
	calculate goroutine ends calculating
	FINISH
	calculate goroutine finished
	main goroutine finished
	*/
}

func TestUnBufferedChannelWaitForPushingDataIn(t *testing.T) {
	ch := make(chan string)

	go func() { // calculate goroutine
			fmt.Println("calculate goroutine starts calculating")
			time.Sleep(time.Second) // Heavy calculation
			fmt.Println("calculate goroutine ends calculating")

			ch <- "FINISH" // goroutine 執行會在此被迫等待

			fmt.Println("calculate goroutine finished")
	}()

	fmt.Println(<-ch) // main 因拉取的時候 calculate 還沒將資料推入 Channel 中，因此 main 會被迫等待
	time.Sleep(time.Second)
	fmt.Println("main goroutine finished")

	/*
	calculate goroutine starts calculating
	calculate goroutine ends calculating
	calculate goroutine finished
	FINISH
	main goroutine finished
	*/
}

//-------------------------------------- Unbuffered Channel -----------------------------------------
// Buffered Channel 的宣告會在第二個參數中定義 buffer 的長度，
// * 它只會在 Buffered 中資料填滿以後才會阻塞造成等待

func TestUnBufferedChannelBlocked(t *testing.T) {
	ch := make(chan int)
	ch <- 1 // 等到天荒地老
	fmt.Println(<-ch)
	// fatal error: all goroutines are asleep - deadlock!
	/*
		使用 Unbuffered Channel：

		只有一條 Goroutine：main
		推入 1 後因為還沒有其他 Goroutine 拉取 Channel 中的資料，所以進入阻塞狀態
		因為 main 已經在推入資料時阻塞，所以拉取的程式永遠不會被執行，造成死結
	*/
}

func TestBufferedChannel(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println(<-ch) // 1
	// 推入 1 後 Channel 內的資料數為1並沒有超過 Buffer 的長度1，所以不會被阻塞
	// 因為沒有阻塞，所以下一行拉取的程式碼可以被執行，並且完成執行
}

// ------------------------------------- Loop 中的 Channel --------------------------------------
// 在迴圈中的 Channel 可以藉由第二個回傳值 ok 確認 Channel 是否被關閉，如果被關閉的話代表此 Channel 已經不再使用，可以結束巡覽。
func TestLoopChannel(t *testing.T) {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			time.Sleep(500 * time.Millisecond)
		}
		close(c)
	}()
	// infinite foor loop
	for {
		v, ok := <-c
		if !ok { // 判斷 Channel 是否關閉
				break
		}
		fmt.Println(v)
	}
}

// * 為了避免將資料推入已關閉的 Channel 中造成 Panic，Channel 的關閉應該由推入的 Goroutine 處理。
// 如果對 Closed Channel 推入資料的話會造成 Panic：
func TestPushDataIntoClosedChannel(t *testing.T) {
	c := make(chan int)
	close(c)
	c <- 0 // panic: send on closed channel
}

// ------------------------------------- range 中的 Channel --------------------------------------
// * range 是可以巡覽 Channel 的，終止條件為 Channel 的狀態為已關閉的(Closed)：
func TestChannelWithRange(t *testing.T) {
	c := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			time.Sleep(500 * time.Millisecond)
		}
		close(c) // 關閉 Channel
	}()
	for i := range c { // 在 close 後跳出迴圈
		fmt.Println(i)
	}
}

// ----------------------------------------- 使用 select 避免等待 -------------------------------------------
// 使用 select 避免等待
// 在 Channel 推入/拉取時，會有一段等待的時間而造成 Goroutine 無法回應，
// 如果此 Goroutine 是負責處理畫面的，使用者就會看到畫面 lag 的情況，這是我們不想見的情況。
/*
		case <-ch:: 會等到沒有阻塞情況時(ch 內有資料)才會執行
		default:: 在所有的 case 都阻塞的情況下執行
		因為有 default 可以設置，當 Channel 阻塞時也可以藉由 default 輸出資訊讓使用者知道。
*/
func TestChannelWithSelect(t *testing.T) {
	ch := make(chan string)

	go func() {
			fmt.Println("calculate goroutine starts calculating")
			time.Sleep(time.Second) // Heavy calculation
			fmt.Println("calculate goroutine ends calculating")

			ch <- "FINISH"
			time.Sleep(time.Second)
			fmt.Println("calculate goroutine finished")
	}()

	for {
			select {
			case <-ch: // Channel 中有資料執行此區域
					fmt.Println("main goroutine finished")
					return
			default: // Channel 阻塞的話執行此區域
					fmt.Println("WAITING...")
					time.Sleep(500 * time.Millisecond)
			}
	}

	/*
		WAITING... # main goroutine 在阻塞時可以回應
		calculate goroutine starts calculating
		WAITING... # main goroutine 在阻塞時可以回應
		WAITING... # main goroutine 在阻塞時可以回應
		calculate goroutine ends calculating
		main goroutine finished # main goroutine 解除阻塞並結束程式
	*/
}