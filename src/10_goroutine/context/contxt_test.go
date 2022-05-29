package context_test

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

/*
	*	context.Context 的使用方法和設計原理 —
	*		多個 Goroutine 同時訂閱 ctx.Done() 管道中的消息，一旦接收到取消信號就立刻停止當前正在執行的工作。

	* context package 最重要的就是處理多個 goroutine 的情況，特別是用來送出取消或結束的 signal。

		我們可以在一個 goroutine 中建立 Context 物件後，傳入另一個 goroutine；
		另一個 goroutine 即可以透過 Done() 來從該 context 中取得 signal，
		一旦這個 Done channel 關閉之後，這個 goroutine 即會關閉並 return。

		Context 也可以是受時間控制，它也可以在特定時間後關閉該 signal channel，
		我們可以定義一個 deadline 或 timeout 的時間，時間到了之後，Context 物件就會關閉該 signal channel。

		更好的是，一旦父層的 Context 關閉其 Done channel 之後，子層的 Done channel 則會自動關閉。

*/

/*
	重要概念
	* 不要把 Context 保存在 struct 中，而是直接當作第一個參數傳入 function 或 goroutine 中，通常會命名為 ctx
	* server 在處理傳進來的請求時應該要建立一個 Context，而使用該 server 的方法則應該要接收 Context 作為參數
	* 雖然函式可以允許傳入 nil Context，但千萬不要這麼做，如果你不確定要用哪個 Context，可以使用 context.TODO
	* 只在 request-scoped data 這種要交換處理資料或 API 的範疇下使用 context Values，不要傳入 optional parameters 到函式中。
	* 相同的 Context 可以傳入多個不同的 goroutine 中使用，在多個 goroutines 中同時使用 Context 是安全的（safe）

		func DoSomething(ctx context.Context, arg Arg) error {
     		... use ctx ...
		}

		type Context interface {
			Deadline() (deadline time.Time, ok bool)
			Done() <-chan struct{}
			Err() error
			Value(key interface{}) interface{}
		}
		context 包中提供的 context.Background、context.TODO、context.WithDeadline 和 context.WithValue 函數會返回實現該接口的私有結構體
*/

// ------------------------------------------------------------------------------------------------
/*
*	context.Background()
		context.Background() 會回傳一個不是 nil 的 empty Context，
		這個 Context 絕不會被取消（canceled）、不會有值、也不會有 deadline。
		這通常會用在 main function、初始化（initialization）或測試中使用，
		可以作為處理請求時最高層的 Context（top-level Context）。

*	context.TODO()
		context.TODO() 會回傳一個不是 nil 的 empty Context。
		它通常會使用在還不清楚要使用哪個 Context 時，或還無法取得 Context 的情況下使用。

*	context.WithCancel()
		context.WithCancel() 函式會回傳 Context 物件和 CancelFunction。
		這個 Context 的 Done channel 會在 cancel function 被呼叫到時關閉，或是父層的 Done channel 關閉時亦會關閉。

			func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

			* 1. 重複呼叫 cancel() 不會有任何效果
			2. context 建議當成函式或 goroutine 的參數傳入，並且命名為 ctx，並不建議把它保存在 struct 中
			3. Context 可以有父子層的關係，也就是一個 Context 可以產生另一個 Context，
					但一旦父層 Context 取消／關閉時，所有根據這個 Context 所產生的 Context 也會一併關閉


*	context.WithDeadline()

			func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)

		在 context.WithDeadline() 中可以指定一個時間（time.Time）當作 deadline，一旦時間到時，就會自動觸發 cancel
		context.WithDeadline() 同樣會回傳 cancel，因此也可以主動呼叫 cancel
		如果父層的 context 被 cancel 的話，子層的 context 也會一併被 cancel

*	context.WithTimeout()
		超過一定的時間後就會停止該 function

		func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		context.WithTimeout() 的用法和 context.WithDeadline() 幾乎相同，差別只在於 WithTimeout() 帶入的參數是時間區間（time.Duration）
		實際上，WithTimeout() 的底層仍然是呼叫 WithDeadline()，只是它會幫忙做掉 time.Add() 的動作
*/
func square(ctx context.Context, c chan int) {
	i := 0
	for {
		select {
		case <-ctx.Done(): // STEP 2：監聽 context Done
			return // kill goroutine
		case c <- i * i:
			i++
		}
	}
}

func TestWithCancel(t *testing.T) {
	c := make(chan int)
	// STEP 1：建立可以被 cancel 的 context
	ctx, cancel := context.WithCancel(context.Background())

	go square(ctx, c)
	for i := 0; i < 5; i++ {
		fmt.Println("Next square is", <-c)
	}

	// STEP 3：當所有訊息都從 channel 取出後，使用 cancel 把 square 這個 goroutine 關閉
	cancel()
	time.Sleep(3 * time.Second)
	fmt.Println("Number of active goroutines", runtime.NumGoroutine())

	/*
		Next square is 0
		Next square is 1
		Next square is 4
		Next square is 9
		Next square is 16
		Number of active goroutines 2
	*/
}

func startProcessA(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "Exit")
			return
		case <-time.After(1 * time.Second):
			fmt.Println(name, "keep doing something")
		}
	}
}
func TestWithCancel2(t *testing.T) {
	// 使用 context.WithCancel 取得 ctx 和 cancel
	ctx, cancel := context.WithCancel(context.Background())
	go startProcessA(ctx, "Process A") // 執行 goroutine 並把 context 傳入
	time.Sleep(5 * time.Second)
	fmt.Println("client release connection, need to notify Process A and exit")
	cancel() // 呼叫 cancel 方法
	fmt.Println("Process finish")
	/*
		Process A keep doing something
		Process A keep doing something
		Process A keep doing something
		Process A keep doing something
		client release connection, need to notify Process A and exit
		Process finish
		Process A Exit
	*/
}

// ---------------------------------------------------------------------------------------------

var startTime = time.Now()

func worker(ctx context.Context, durationSecs int) {
	select {
	// STEP 3：deadline 時間到時或主動呼叫 cancel 時，都會進入 ctx.Done()
	case <-ctx.Done():
		fmt.Printf("%0.2fs - worker(%ds) killed!\n", time.Since(startTime).Seconds(), durationSecs)
		return // kills goroutine

	// 模擬做事所需花費的時間
	case <-time.After(time.Duration(durationSecs) * time.Second):
		fmt.Printf("%0.2fs - worker(%ds) completed the job.\n", time.Since(startTime).Seconds(), durationSecs)
	}
}

func TestWithDeadline(t *testing.T) {
	// STEP 1：建立 deadline
	deadline := time.Now().Add(3 * time.Second)

	// STEP 2：將 deadline 傳入並取得 cancel
	ctx, cancel := context.WithDeadline(context.Background(), deadline)

	// STEP 4：如果 main 比其他 goroutine 提早結束時，呼叫 cancel 讓其他 goroutine 結束
	defer cancel()

	go worker(ctx, 2)
	go worker(ctx, 3)
	go worker(ctx, 4)
	go worker(ctx, 6)
	fmt.Println("Number of active goroutines", runtime.NumGoroutine())

	time.Sleep(5 * time.Second)

	fmt.Println("Number of active goroutines", runtime.NumGoroutine())

	/*
		Number of active goroutines 6
		2.00s - worker(2s) completed the job.
		3.00s - worker(3s) completed the job.
		3.00s - worker(6s) killed!
		3.00s - worker(4s) killed!
		Number of active goroutines 2
	*/
}

// ----------------------------------------------------------------------------
func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

// handle 函數沒有進入超時的 select 分支，
// 但是 main 函數的 select 卻會等待 context.Context 超時並打印出 main context deadline exceeded。
func TestWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 因為過期時間大於處理時間，所以我們有足夠的時間處理該請求
	go handle(ctx, 1000*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}

	// process request with 1s
	// main context deadline exceeded
	// --- PASS: TestWithTimeout (2.00s)
}

func TestWithTimeout2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 整個程序都會因為上下文的過期而被中止
	go handle(ctx, 3000*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}

	// handle context deadline exceeded
	// main context deadline exceeded
	// --- PASS: TestWithTimeout2 (2.00s)
}

