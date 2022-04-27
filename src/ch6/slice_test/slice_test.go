package slice_test

import (
	"fmt"
	"testing"
)

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5) // func make([]type, len, capacity)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	// t.Log(s2[0], s2[1], s2[2], s2[3]) // error, since len == 3

	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
}


func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMenory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun","Jul","Aug","Sep","Oct","Nov","Dec"}
	Q2 := year[3 : 6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5 : 8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknown"
	t.Log(Q2)
	t.Log(year)
}

// slice can only be compared to nil
// func TestSliceCompare(t *testing.T) {
// 	a := []int{1, 2, 3, 4}
// 	b := []int{1, 2, 3, 4}
// 	if a == b {
// 		t.Log("equal")
// 	}
// }

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 5, 3, 4}
	d := [...]int{1, 2, 3, 4}
	// f := [...]int{1, 2, 3, 4, 5, 6}

	t.Log(a == b)
	t.Log(a == d)
	// t.Log(a == f)
}

func TestArrayandSlice(t *testing.T) {
		// Slice 的建立 ----------------------------------------------------------------
		// slice 型別：[]T
		// slice := make([]T, len, cap)
	  // 方式一：建立一個帶有資料的 string slice，適合用在知道 slice 裡面的元素有哪些時
		people := []string{"Aaron", "Jim", "Bob", "Ken"}

		// 方式二：透過 make 可以建立「空 slice」，適合用會對 slice 中特定位置元素進行操作時
		people2 := make([]int, 4)  // len=4 cap=4，[0,0,0,0]

		// 方式三：空的 slice，一般會搭配 append 使用
		var people3 []string

		// 方式四：大該知道需要多少元素時使用，搭配 append 使用
		people4 := make([]string, 0, 5) // len=0 cap=5, []
		fmt.Print(people)
		fmt.Print(people2)
		fmt.Print(people3)
		fmt.Print(people4)



		// Arrays 的建立 ----------------------------------------------------------------
		// 在 Array 中陣列的元素是固定的：

		// 陣列型別：[n]T
		// 陣列的元素只能有 10 個，且都是數值
		// 先定義再賦值
    var a [2]string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a)  // [Hello World]

  	// 定義且同時賦值
    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes)  // [2 3 5 7 11 13]
}