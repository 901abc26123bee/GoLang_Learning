package slicetest

import (
	"fmt"
	"sort"
	"testing"
)

// 反向排序
func TestSortReverse(t *testing.T) {
	numbers := []int{1, 5, 3, 6, 2}
	sort.Ints(numbers)
	fmt.Println(numbers) // [1 2 3 5 6]，ascending

	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Println(numbers) // [6 5 3 2 1]，descending

}

// 客製化
// 只要有實作 sort.Interface 的 slice of structs 都可以使用 sort.Sort() 進行排序：
type Programmer struct {
	Name string
	Age  int
}

/* 建立一個符合 sort.interface 的  type */
type byAge []Programmer

func (p byAge) Len() int {
	return len(p)
}

func (p byAge) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p byAge) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func TestCustomizeSorting(t *testing.T) {
	programmers := []Programmer{
			{Name: "Aaron", Age: 30},
			{Name: "Bruce", Age: 20},
			{Name: "Candy", Age: 50},
			{Name: "Derek", Age: 1000},
	}

	sort.Sort(byAge(programmers))

	fmt.Println(programmers) // [{Bruce 20} {Aaron 30} {Candy 50} {Derek 1000}]
}
