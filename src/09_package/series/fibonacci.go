package series

import "fmt"

// 1. 在 Go 語言中，並沒有區分 public、private 或 protected，而是根據變數名稱的第一個字母大小寫來判斷能否被外部引用。
// 2. 在同一個 package 中變數、函式、常數和 type，都隸屬於同一個 package scope，因此雖然可能在不同支檔案內，但只要隸屬於同一個 package，都可以使用（visible）
// 3. 如果需要讓 package 內的變數或函式等能夠在 package 外部被使用，則該變數的第一個字母要大寫才能讓外部引用（Exported names），否則的話會無法使用
func GetFibonacciSeries(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2] + ret[i-1])
	}
	return ret
}

func Square(n int) int {
	return n * n
}

func init() {
	fmt.Println("init - 1")
}

func init() {
	fmt.Println("init - 2")
}