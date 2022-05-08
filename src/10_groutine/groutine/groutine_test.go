package groutine_test

import (
	"fmt"
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
	1. goroutines 是由 Go runtime 所管理的輕量化的 thread
	2. goroutines 會在相同的 address space 中執行，因此要存取共享的記憶體必須要是同步的（synchronized）。
	3.	傳統的 Apache 伺服器來說，當每分鐘需要處理 1000 個請求時，每個請求如果都要 concurrently 的運作，
			將會需要建立 1000 個 threads 或者分派到不同的 process 去做，
				- 如果 OS 的每個 thread 都需要使用 1MB 的 stack size 的話，就會需要 1GB 的記憶體才能撐得住這樣的流量。
				- 但相對於 goroutine 來說，因為 stack size 可以動態增長，因此可以擴充到 1000 個 goroutines，
					每個 goroutine 只需要 2KB（Go 1.4 之後）的 stack size。
*/