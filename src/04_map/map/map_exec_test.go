package map_test

import (
	"fmt"
	"testing"
)

func TestMapWithFuncValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op + op }
	m[3] = func(op int) int { return op + op + op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[3] = true
	n := 3
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	mySet[1] = true
	mySet[5] = true
	t.Log(len(mySet))

	n = 5
	delete(mySet,5)
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}

/*
	在 map 中因為 key 本身是有資料型別的，因此只能使用 [] 來設值和取值，不能使用 .：

		1. 使用 len() 可以取得 map 鍵值的數量
		2. 使用 delete(map, key) 可以移除 map 中的鍵值
		3. 使用 value, isExist := m[key] 可以用來取值，同時判斷該 key 是否存在
*/
func TestMapOperations(t *testing.T) {
	colors := make(map[int]string)

	// 先增 map 中的 key-value
	colors[10] = "#ffffff"
	colors[1] = "#4bf745"

	// 使用 len 可以取得 map 鍵值的數量
	fmt.Println(len(colors)) // 2

	// 移除 map 中的資料，delete(m, key)
	delete(colors, 10)

	// 取值
	fmt.Println(colors[10])

	// 檢查某 Map 中是否存在該 key
	// value, isExist = m[key]
	// 該 key 存在的話 isExist 會是 true，否則會是 false
	// 該 key 存在的話 value 會回傳該 map 的值，否則回傳該其 zero value
	value, isExist := colors[1] // #4bf745 true
	value, isExist = colors[10] // "", false

	fmt.Println(value, isExist)
}

// 因為 map 是 reference type，所以不用使用 pointer
func printMap(c map[string]string) {
	for color, hex := range c {
			fmt.Printf("%v: %v\n", color, hex)
	}
}

func TestMapIterator(t *testing.T) {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// 疊代 map
	printMap(colors)
}