package TestChannel

import (
	"fmt"
	"testing"
)

/*
	1. 從一個 goroutine 切換到另一個 goroutine 的時機點是「當正在執行的 goroutine 阻塞時，就會交給其他 goroutine 做事」
	2. unbuffered channel 指的是 buffer size 為 0 的 channel
	3. 對於 unbuffered channel 來說，
			不論是從 channel 讀資料（需等到被寫入），或把資料寫入 channel 中時（需等到被讀出），都會阻塞該 goroutine
	4. 對於 buffered channel 來說：
			- 從 channel 讀值時若是 empty buffer 時才會阻塞，否則都是 non-blocking
			- 把資料寫入 channel 中時，寫入 channel 中的 value 數目（n + 1）需要超過 buffer size（n），也就是溢出（overflow）時才會使得該 goroutine 被阻塞；
				而且一旦 buffer channel 中的值開始被讀取，就會被全部讀完
*/

/*
	// go routine
	go f(x, y, z)

	// channels
	// 和 maps, slices, channels 一樣需要在使用前被建立，這裡表示定義的 chan 會回傳 int
	ch := make(chan int)

	ch <- v        // Send v to channel ch.
	v, ok := <-ch  // Receive from ch, and assign value to v. ok => channel status
*/

func TestChannel(t *testing.T) {
	// channel 的 zero value 是 nil
	var zeroC chan int
	fmt.Println(zeroC)

	// 一般建立 channel 的方式
	c := make(chan int)                 // unbuffered channel
	fmt.Printf("type of c is %T\n", c)  // type of c is chan int
	fmt.Printf("value of c is %v\n", c) // value of c is 0xc000062060
}

// ----------------------------- unbufffer channel -----------------------------------
/*
	1. 所有的 unbuffered channel 操作預設都是 blocking 的
	2. 當有資料要寫入 channel 時，goroutine 會阻塞住，直到有其他的 goroutine 從該 channel 把值讀出來
	3. 當有資料要讀取 channel 中的值時，goroutine 也會阻塞，直到其他 goroutine 把值寫入 channel 中
	也就是說，當我們是的把資料寫入 channel 或從 channel 中取出資料時，該 goroutine 都會阻塞住，並且將控制權交給其他可以運行的 goroutines
*/

// 只是寫入資料但沒有 channel 讀取 -> deadlock
func TestUnbufferDeadLockWrite(t *testing.T) {
	fmt.Println("main() started")

	// 只是寫入資料但沒有 channel 讀取
	c := make(chan string)
	c <- "John"

	fmt.Println("main() stopped")
}

// 有 goroutine 要讀取 channel 但沒 goroutine 寫入資料到 channel -> deadlock
func greet(c chan string) {
  // 要讀取 channel 但沒人寫入資料
    fmt.Println("Hello " + <-c + "!")
}

func TestUnbufferDeadLockRead(t *testing.T) {
    fmt.Println("main() started")

    c := make(chan string)
    // c <- "John"
    greet(c)

    fmt.Println("main() stopped")
}

// ----------------------------- bufffer channel ----------------------------------
/*
	1. buffered channel 寫值時，需要在 overflow 時才會 block goroutine
	2. unbuffered channel 指的是 buffered size 為 0 的 channel。
	3. unbuffered channel 不論是「從 channel 讀值」（需等到值被其他 goroutine 寫入），或「把值寫入 channel」（需等到值被其他 goroutine 讀出）都會阻塞當下的 goroutine。
	4. 當 buffer size 不是 0 的話，就屬於 buffered channel
		- 「從 channel 讀值」時，只有在 buffered 是空的時才會 blocking
		- 「把值寫入 channel」時，該 goroutine 並不會被阻塞，除非該 buffer 已經填滿（full）且溢出（overflow）。
			當 buffer 已經填滿（full）時，再把新的一筆資料傳入 channel 時會造成溢出（overflow），此時 goroutine 才會被阻塞。
			讀值的動作一旦開始，就會一直到 buffer 變成 empty 為止才會結束。也就是說，讀取 channel 的那個 goroutine 需到等到 buffer 完全清空後才會阻塞。

			寫值：直到該 channel 寫入到 n+1 個值以前，它都不會阻塞當前的 goroutine。
			讀值：從該 channel 讀值時，若 buffer 是 empty 才會阻塞當前的 goroutine。
*/


// 透過 squares goroutine 讀值
func print(c chan int) {
	for i := 0; i <= 3; i++ {
			fmt.Println(<-c)
	}
}

// 由於寫入 channel 的值並沒有超出 buffered channel 的 size，因此 main goroutine 並不會被阻塞，使得 print goroutine 不會有機會取得控制權而被執行
func TestBufferChannelNonBlocking(t *testing.T) {
	fmt.Println("main() started")

  // 建立 buffered size 為 3 的 channel
	c := make(chan int, 3)

	go print(c)

  // 寫入 3 個值
	c <- 1
	c <- 2
	c <- 3

	fmt.Println("main() close")

	// main() started
	// main() close
}

// 在 TestBufferChannelOverflowing goroutine 中，使用了 c <- 4 後，因為超過 buffered channel 的 size，也就是溢出（overflow），因此在這裡會阻塞
// TestBufferChannelOverflowing goroutine 阻塞後，print goroutine 便有機會執行，一旦 print goroutine 開始讀取 channel 的值後，它就會把該 buffer 中的所有值都讀全部讀完
func TestBufferChannelOverflowing(t *testing.T) {
	fmt.Println("main() started")

	c := make(chan int, 3)
	go print(c)

	c <- 1
	c <- 2
	c <- 3
	c <- 4 // 因為超過 buffered size，這裡會 block

	fmt.Println("main() close")

	// Output
	// main() started
	// 1
	// 2
	// 3
	// 4
	// main() close
}