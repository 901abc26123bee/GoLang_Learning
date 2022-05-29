package mutex_test

import (
	"fmt"
	"sync"
	"testing"
)

/*
	mutex
	在 goroutines 中，由於有獨立的 stack，因此並不會在彼此之間共享資料（也就是在 scope 中的變數）；
	然而在 heap 中的資料是會在不同 goroutine 之間共享的（也就是 global 的變數），
	在這種情況下，許多的 goroutine 會試著操作相同記憶體位址的資料，導致未預期的結果發生。 --> race condition
*/


var i int // i == 0

func worker(wg *sync.WaitGroup) {
    i = i + 1
    wg.Done()
}

func TestRaceCondition(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
			wg.Add(1)
			go worker(&wg)
	}

	wg.Wait()
	fmt.Println(i) // value i should be 1000 but it did not
}

/*
	為了要避免多個 goroutine 同時取用到一個 heap 中的變數，
	第一原則是應該要盡可能避免在多個 goroutine 中使用共享的資源（變數）。

	如果無法避免會需要操作共用的變數的話，則可以使用 Mutex（mutual exclusion），
	也就是說在一個時間內只有一個 goroutine（thread）可以對該變數進行操作，
	在對該變數進行操作前，會先把它「上鎖」，操作完後再進行「解鎖」的動作，
	當一個變數被上鎖的時候，其他的 goroutine 都不能對該變數進行讀取和寫入。

	*	mutex 是 map 型別的方法，被放在 sync package 中，
	*	使用 mutex.Lock() 可以上鎖，使用 mutex.Unlock() 可以解鎖：
*/

func worker2(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock() // 上鎖
	i = i + 1
	m.Unlock() // 解鎖
	wg.Done()
}

func TestMutex(t *testing.T) {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker2(&wg, &m) // 把 mutex 的記憶體位址傳入
	}

	wg.Wait()
	fmt.Println(i) // 在使用 mutex 對 heap 中的變數進行上鎖和解鎖後，即可以確保最終的值是 1000
}

/*
		*	⚠️ mutex 和 waitgroup 一樣，都是把「記憶體位址」傳入 goroutine 中使用。

		如同前面所說，第一原則應該是要避免 race condition 的方法，
		也就是不要在 goroutine 中對共用的變數進行操作，
		\在 go 的 CLI 中可以透過下面的指令檢測程式中是否有 race condition 的情況：

		*	$ go run -race program.go
*/

// ------------------------------ Concurrency Pattern ----------------------------------
// fib 會回傳 read-only channel
func fib(length int) <-chan int {
	c := make(chan int, length)

	// run generation concurrently
	go func() {
			for i, j := 0, 1; i < length; i, j = i+j, i {
					c <- i
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