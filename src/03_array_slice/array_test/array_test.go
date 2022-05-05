package array_test

import "testing"

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