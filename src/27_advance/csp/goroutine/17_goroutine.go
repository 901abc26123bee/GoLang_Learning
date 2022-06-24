package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// ----------------------------------主 goroutine 的運行若過早結束------------------------------------------
// func main() {
// 	// for語句會以很快的速度執行完畢。當它執行完畢時，那 10 個包裝了go函數的 goroutine 往往還沒有獲得運行的機會。
// 	for i := 0; i < 10; i++ {
// 		go func() {
// 			fmt.Println(i) // 不會有任何內容被打印出來
// 		}()
// 		// 後面的code先輩立即執行
// 	}
// 	// 一旦主 goroutine 中的代碼（也就是main函數中的那些代碼）執行完畢，當前的 Go 程序就會結束運行。
// }
/*
no output
*/

// ----------------------------------默認情況下的執行順序是不可預知的------------------------------------------
// func main() {
// 	num := 10
// 	sign := make(chan struct{}, num)

// 	for i := 0; i < num; i++ {
// 		go func() {
// 			fmt.Println(i)
// 			sign <- struct{}{}
// 		}()
// 	}

// 	// 辦法1。
// 	//time.Sleep(time.Millisecond * 500)

// 	// 辦法2。
// 	for j := 0; j < num; j++ {
// 		<-sign
// 	}
// }
/*
10
3
10
10
10
10
8
10
8
10
*/

// ------------------------------------多個 goroutine 按照既定的順序運行----------------------------------------
func main() {
	var count uint32
	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
        // 由於trigger函數會被多個 goroutine 並發地調用，所以它用到的非本地變量count，就被多個用戶級線程共用了。因此，對它的操作就產生了競態條件（race condition），破壞了程序的並發安全性。 --> sync/atomic包中聲明了很多用於原子操作的函數 to solve this problem
        // 原子操作函數對被操作的數值的類型有約束，count以及相關的變量和參數的類型 由int變為了uint32
				// 所以這麼做是因為int位數是根據系統決定的，而原子級操作要求速度盡可能的快，所以明確了整數的位數才能最大地提高性能。
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}
	for i := uint32(0); i < 10; i++ {
    // 在go語句被執行時，傳給go函數的參數i會先被求值，如此就得到了當次迭代的序號。之後，無論go函數會在什麼時候執行，這個參數值都不會變。
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}
	trigger(10, func() {}) // 因為我依然想讓主 goroutine 最後一個運行完畢
}
/*
0
1
2
3
4
5
6
7
8
9
*/

