package gototest

import (
	"fmt"
	"testing"
)

// 標籤名稱 (label) 是區分大小寫的的。
func myFunc() {
	i := 0
	Here:   //這行的第一個詞，以冒號結束作為標籤
		println(i)
		i++
		goto Here   //跳轉到 Here 去
}

func test() {
	LABEL1:
	for i := 0; i <= 5; i++ {
			for j := 0; j <= 5; j++ {
					if j == 4 {
							continue LABEL1
					}
					fmt.Printf("i is: %d, and j is: %d\n", i, j)
			}
	}
}

func test2() {
	i:=0
	HERE:
			print(i)
			i++
			if i==5 {
					return
			}
			goto HERE
}


func TestGoto(t *testing.T) {
//	myFunc() // inifinite koop
	test()
	test2() // 01234
}