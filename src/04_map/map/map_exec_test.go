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


type Student struct {
	Name string
}

// var list map[string]Student

func TestMapValueOperator(t *testing.T) {
	list := make(map[string]Student)
	student := Student{Name: "Olive"}
	list["student"] = student
	// map[string]Student 的value是一個Student結構值，
	// *所以當list["student"] = student,是一個值拷貝過程。
	// *而list["student"]則是一個值引用。那麼值引用的特點是只讀。
	// 所以對list["student"].Name = "LDB"的修改是不允許的。
	// list["student"].Name = "Hens"

	/*
			方法1:
			先做一次值拷貝，做出一個tmpStudent副本,然後修改該副本，然後再次發生一次值拷貝複制回去，
			但是這種會在整體過程中發生2次結構體值拷貝，性能很差。
	*/
	tmpStudent := list["student"]
	tmpStudent.Name = "LDB"
	list["student"] = tmpStudent

	fmt.Println(list["student"])
}

// var list2 map[string]*Student

func TestMapValueOperator2(t *testing.T) {
	list2 := make(map[string]*Student)
	student := Student{Name: "Olive"}
	list2["student"] = &student

	/*
			方法2:
			將map的類型的value由Student值，改成Student指針
			* 指針本身是常指針，不能修改，只讀屬性，但是指向的Student是可以隨便修改的，而且這裡並不需要值拷貝。只是一個指針的賦值。
	*/
	list2["student"].Name = "LDB"

	fmt.Println(*list2["student"])
}

type student struct {
	Name string
	Age  int
}

func TestMapValueIteration(t *testing.T) {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	for k, v := range m {
		fmt.Println(k, "-->", v.Name, v.Age)
	}

	// *foreach中，stu是結構體的一個拷貝副本，
	// *所以m[stu.Name]=&stu實際上一致指向同一個指針， 最終該指針的值為遍歷的最後一個struct的值拷貝。
	/*
		map中的3個key均指向數組中最後一個結構體。
		zhou --> wang 22
		li --> wang 22
		wang --> wang 22
	*/
}

func TestMapValueIteration2(t *testing.T) {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "lion", Age: 23},
		{Name: "wang", Age: 22},
	}

	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}

	for k, v := range m {
		fmt.Println(k, "-->", v.Name, v.Age)
	}

	/*
		zhou --> zhou 24
		lion --> lion 23
		wang --> wang 22
	*/
}