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