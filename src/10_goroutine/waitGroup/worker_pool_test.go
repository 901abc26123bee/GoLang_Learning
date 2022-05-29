package waitgroup_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
	Worker Pool

	worker pool 指的是有許多的 goroutines 同步的進行一個工作。要建立 worker pool，會先建立許多的 worker goroutine，
	這些 goroutine 中會：

		進行相同的 job
		有兩個 channel，一個用來接受任務（task channel），一個用來回傳結果（result channel）
		都等待 task channel 傳來要進行的 tasks
		一但收到 tasks 就可以做事並透過 result channel 回傳結果
*/

// 程式來源：https://medium.com/rungo/anatomy-of-channels-in-go-concurrency-in-go-1ec336086adb
// STEP 3：在 worker goroutines 中會做相同的工作
// tasks is receive only channel
// results is send only channel
func sqrWorker(tasks <-chan int, results chan<- int, instance int) {
	// 一旦收到 tasks channel 傳來資料，就可以動工並回傳結果
	for num := range tasks {
			time.Sleep(500 * time.Millisecond) // 模擬會阻塞的任務
			fmt.Printf("[worker %v] Sending result of task %v \n", instance, num)
			results <- num * num
	}
}


func TestWorkerPool(t *testing.T) {
	fmt.Println("[main] main() started")

	// STEP 1：建立兩個 channel，一個用來傳送 tasks，一個用來接收 results
	tasks := make(chan int, 10)
	results := make(chan int, 10)

	// STEP 2 啟動三個不同的 worker goroutines
	for i := 1; i <= 3; i++ {
			go sqrWorker(tasks, results, i)
	}

	// STEP 4：發送 5 個不同的任務
	for i := 1; i <= 5; i++ {
			tasks <- i // non-blocking（因為 buffered channel 的 capacity 是 10）
	}

	fmt.Println("[main] Wrote 5 tasks")

  // STEP 5：發送完任務後把 channel 關閉（非必要，但可減少 bug）
	close(tasks)

	// STEP 6：等待各個 worker 從 result channel 回傳結果
	for i := 1; i <= 5; i++ {
			result := <-results // blocking（因為 buffer 是空的）
			fmt.Println("[main] Result", i, ":", result)
	}

	fmt.Println("[main] main() stopped")

	// 當所有 worker 都剛好 blocking 的時候，控制權就會交回 main goroutine，這時候就可以立即看到計算好的結果。
	
	/*
		[main] main() started
		[main] Wrote 5 tasks
		[worker 1] Sending result of task 1
		[worker 3] Sending result of task 3
		[main] Result 1 : 1
		[main] Result 2 : 9
		[worker 2] Sending result of task 2
		[main] Result 3 : 4
		[worker 3] Sending result of task 5
		[worker 1] Sending result of task 4
		[main] Result 4 : 25
		[main] Result 5 : 16
		[main] main() stopped
	*/

}

// ------------------------------- WorkerGroup 搭配 WaitGroup ---------------------------------
// 但有些時候，我們希望所有的 tasks 都執行完後才讓 main goroutine 繼續往後做，這時候可以搭配 WaitGroup 使用：

// 這時會等到所有的 worker 都完工下班後，才開始輸出計算好的結果。
// 搭配 WaitGroup 的好處是可以等到所有 worker 都完工後還讓程式繼續，但相對的會需要花更長的時間在等待所有人完工：

func sqrWorker2(wg *sync.WaitGroup, tasks <-chan int, results chan<- int, instance int) {
	defer wg.Done()

	// 一旦收到 tasks channel 傳來資料，就可以動工並回傳結果
	// read from chan tasks
	for num := range tasks {
		time.Sleep(500 * time.Millisecond) // 模擬會阻塞的任務
		fmt.Printf("[worker %v] Sending result of task %v \n", instance, num)
		results <- num * num
	}
}

func TestWorkerPoolWithWaitGroup(t *testing.T) {
	fmt.Println("[main] main() started")

	var wg sync.WaitGroup

	tasks := make(chan int, 10)
	results := make(chan int, 10)

	for i := 1; i <= 3; i++ {
			wg.Add(1)
			go sqrWorker2(&wg, tasks, results, i)
	}

	for i := 1; i <= 5; i++ {
			tasks <- i // non-blocking（因為 buffered channel 的 capacity 是 10）
	}

	fmt.Println("[main] Wrote 5 tasks")

	close(tasks)  // 有用 waitGroup 的話這個 close 不能省略

	// 直到所有的 worker goroutine 把所有 tasks 都做完後才繼續往後
	wg.Wait()

	for i := 1; i <= 5; i++ {
			result := <-results // blocking（因為 buffer 是空的）
			fmt.Println("[main] Result", i, ":", result)
	}

	fmt.Println("[main] main() stopped")

	/*
		[main] main() started
		[main] Wrote 5 tasks
		[worker 3] Sending result of task 3
		[worker 2] Sending result of task 2
		[worker 1] Sending result of task 1
		[worker 3] Sending result of task 4
		[worker 2] Sending result of task 5
		[main] Result 1 : 9
		[main] Result 2 : 4
		[main] Result 3 : 1
		[main] Result 4 : 16
		[main] Result 5 : 25
		[main] main() stopped
	*/
}