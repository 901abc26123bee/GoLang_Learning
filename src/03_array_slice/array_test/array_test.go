package array_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4}
	arr4 := [...][]int{{1, 2, 3, 4}, {9, 8, 7, 7}}
	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1[2], arr2)
	t.Log(arr4)
}

func TestArrayTraversal(t *testing.T) {
	arr3 := [...]int{6, 5, 4, 3}
	for i := 0; i <len(arr3); i++ {
		t.Log(arr3[i])
	}
	for idx, e := range arr3 {
		t.Log(idx, e)
	}
	for _, e := range arr3 {
		t.Log(e)
	}
}

func TestArraySections(t *testing.T) {
	arr1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t.Log(arr1[1:3])
	t.Log(arr1[1:len(arr1)])
	t.Log(arr1[3:])
	t.Log(arr1[:3])
	t.Log(arr1[:])
}

// Arrays 的建立
// 在 Array 中陣列的元素是固定的：
// 1. 陣列型別：[n]T
// 2. 使用 [...]T{} 可以根據元素的數目自動建立陣列：
func TestArrayCreatiojn(t *testing.T) {
	// 先定義再賦值
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a)  // [Hello World]

	// 定義且同時賦值
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)  // [2 3 5 7 11 13]

	// 沒有使用 ...，建立出來的會是 slice
	arr := []string{"North", "East", "South", "West"}
	fmt.Println(reflect.TypeOf(arr).Kind(), len(arr))  // slice 4

	// 使用 ...，建立出來的會是 array
	arrWithDots := [...]string{"North", "East", "South", "West"}
	fmt.Println(reflect.TypeOf(arrWithDots).Kind(), len(arrWithDots))  // array 4
}