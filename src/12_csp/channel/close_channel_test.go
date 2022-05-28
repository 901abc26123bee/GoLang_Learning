package channel_test

import (
	"fmt"
	"testing"
)

//------------------------------ 讀取 Channels 中的資料 ----------------------------------
// 讀取 Channels 中的資料
// for loop 搭配 close：需要手動 break 迴圈
// 使用 for{} 來顯示內容，但需要手動 break loop：

func squares(c chan int) {
	// 把 0 ~ 9 寫入 channel 後便把 channel 關閉
	for i := 0; i <= 9; i++ {
		c <- i
	}

	close(c)
}

func TestChannelWithForLoop(t *testing.T) {
	fmt.Println("main() started")
	c := make(chan int)

	// 發動 squares goroutine
	go squares(c)

	// 監聽 channel 的值：週期性的 block/unblock main goroutine 直到 squares goroutine close
	for {
		val, ok := <-c
		if ok == false {
			fmt.Println(val, ok, "<-- loop broke case channel closed")
			break // exit loop
		} else {
			fmt.Println(val, ok)
		}
	}

	fmt.Println("main() close")
}

/*
	for range 搭配 close：會自動關閉迴圈
	* 透過 for i := range c 可以重複取出 channel 的值，直到該 channel 被關閉。
	* - 如同上一段的程式碼，單純透過 for 需要自行判斷 channel 是否已經 close，如果是的話需要自行使用 break 把 for loop 終止。
	* - 在 Go 中則提供了 for range loop，只要該 channel 被關閉後，loop 則會自動終止。
	* - 需要特別留意，如果是在 main goroutine 中使用 for val := range channel {} 的寫法時，最後 channel 沒有被 close 的話程式會 deadlock。
	* 	但如果是在其他的 goroutine 中使用，即使沒有 close 也不會 deadlock，但為了不必要的 bug 產生，一般都還是將其關閉。
				for val := range channel {
					當 channel 關閉時會自動 break loop
				}
	* 需要特別留意，使用 for val := range channel {} 的寫法時，如果最後 channel 沒有被 close 的話程式會 deadlock。
	*/
func TestChannelWithRanges(t *testing.T) {
	fmt.Println("main() started")
	c := make(chan int)

	// 發動 squares goroutine
	go squares(c)

	// 使用 for range 的寫法，一但 channel close，loop 會自動 break
	for val := range c {
			fmt.Println(val)
	}

	fmt.Println("main() close")
}
