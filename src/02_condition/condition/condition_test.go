package condition_test

import (
	"fmt"
	"testing"
	"time"
)

func TestIfMutiSpect(t *testing.T) {
	// if v, err := somFunc(); err == nil {
	// 	t.Log("successfully")
	// } else {
	// 	t.Log("fail")
	// }
}

// switch
// 在 Go 的 switch 中不需要使用 break（Go 執行時會自動 break）
// switch 和 if, for 類似，在最前面都可以加上 statement
// 會從上往下開始判斷每一個 case，並在配對到的 case 後終止

// case 後面可以是 express
// 在 case 後面不一定一定要是字串，是可以放 express 的：
func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
			case 0, 2:
				t.Log("Even")
			case 1, 3:
				t.Log("Odd")
			default:
				t.Log("It is not 0-3")
		}
	}
}


// switch 後可以不帶變數（可用來簡化 if ... else）
// switch 後可以不帶變數（可用來簡化 if ... else）
// switch 可以不帶變數，直接變成 if...elseif...else 的簡化版：
func TestSwitchMutiCaseConditioin(t *testing.T) {
	currentHour := time.Now().Hour()
	switch {
	case currentHour < 12:
			fmt.Println("Good morning!")
	case currentHour < 17:
			fmt.Println("Good afternoon.")
	default:
			fmt.Println("Good evening")
	}
}

func TestIterator(t *testing.T) {
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// 如果只需要 index，可以簡寫成：
	// for i := range cardValues { ... }
	for i, value := range cardValues {
		fmt.Println(i, value)
	}
}

func TestIterator2(t *testing.T) {
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for value := range cardValues {
		fmt.Println(value)
	}

	for i := range cardValues {
		fmt.Println(i)
	}
}