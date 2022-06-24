package other_test

import (
	"testing"
)

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func Test_unaddressable(t *testing.T) {
	const num = 123
	//_ = &num // 常量不可尋址。
	//_ = &(123) // 基本類型值的字面量不可尋址。

	var str = "abc"
	_ = str
	//_ = &(str[0]) // 對字符串變量的索引結果值不可尋址。
	//_ = &(str[0:2]) // 對字符串變量的切片結果值不可尋址。
	str2 := str[0]
	_ = &str2 // 但這樣的尋址就是合法的。

	//_ = &(123 + 456) // 算術操作的結果值不可尋址。
	num2 := 456
	_ = num2
	//_ = &(num + num2) // 算術操作的結果值不可尋址。

	//_ = &([3]int{1, 2, 3}[0]) // 對數組字面量的索引結果值不可尋址。
	//_ = &([3]int{1, 2, 3}[0:2]) // 對數組字面量的切片結果值不可尋址。
	_ = &([]int{1, 2, 3}[0]) // 對切片字面量的索引結果值卻是可尋址的。
	//_ = &([]int{1, 2, 3}[0:2]) // 對切片字面量的切片結果值不可尋址。
	//_ = &(map[int]string{1: "a"}[0]) // 對字典字面量的索引結果值不可尋址。

	var map1 = map[int]string{1: "a", 2: "b", 3: "c"}
	_ = map1
	//_ = &(map1[2]) // 對字典變量的索引結果值不可尋址。

	//_ = &(func(x, y int) int {
	//	return x + y
	//}) // 字面量代表的函數不可尋址。
	//_ = &(fmt.Sprintf) // 標識符代表的函數不可尋址。
	//_ = &(fmt.Sprintln("abc")) // 對函數的調用結果值不可尋址。

	dog := Dog{"little pig"}
	_ = dog
	//_ = &(dog.Name) // 標識符代表的函數不可尋址。
	//_ = &(dog.Name()) // 對方法的調用結果值不可尋址。

	//_ = &(Dog{"little pig"}.name) // 結構體字面量的字段不可尋址。

	//_ = &(interface{}(dog)) // 類型轉換錶達式的結果值不可尋址。
	dogI := interface{}(dog)
	_ = dogI
	//_ = &(dogI.(Named)) // 類型斷言表達式的結果值不可尋址。
	// named := dogI.(Named)
	// _ = named
	//_ = &(named.(Dog)) // 類型斷言表達式的結果值不可尋址。

	var chan1 = make(chan int, 1)
	chan1 <- 1
	//_ = &(<-chan1) // 接收表達式的結果值不可尋址。
}