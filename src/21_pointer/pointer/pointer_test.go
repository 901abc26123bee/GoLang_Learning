package pointer_test

import (
	"fmt"
	"testing"
)

// Pointers Operation
// & (ampersand) 和 *(Asterisk)
/*
	當我們使用 &variable 時，會回傳該變數 value 的 address，表示給我這個變數值的「記憶體位置」。
	當我們使用 *pointer 時，會回傳該 address 的 value，表示給我這個記憶體位置指稱到的「值」。
	⚠️ 但若 * 是放在 type 前面，那個這麼 * 並不是運算子，而是對於該 type 的描述（type description）。
		因此在 func 中使用的 *type 是指要做事的對象是指稱到該型別（type）的指標（pointer），也就是這個 function 只能被 person 的指標（pointer to a person）給呼叫。
*/
// 會需要「mutate」原本資料的 methods 就需要傳入的是 pointer
// 單純是「顯示」原本資料用的 methods 就不需要傳入 pointer

func TestPointer(t *testing.T) {
	// *T 是一種型別，指的是能夠指向該 T 的值的指標，它的 zero value 是 nil
	// *T means pointer to value of type T
	// * : get value by address(&variable)
	var pointer *int
	fmt.Println(pointer) // nil
	fmt.Println(&pointer) // 0xc00000e048

	// &variable 會產生該 variable 的 pointer
	i := 42
	p := &i         // & 稱作 address of pointer
	fmt.Println(p)  // 0xc0000b4008
	fmt.Println(*p) // 透過 pointer 來讀取到 i ==> 42
}

type person struct {
	firstName string
	lastName string
}
// 當 function receiver 這裡使用了 *type 時
// 這裡拿到的 p 會變成 pointer，指的是存放 p 的記憶體位址
func (p *person) updateNameFromPointer(newFirstName string) {
  // *variable 表示把該指摽對應的值取出
    (*p).firstName = newFirstName
}

// 當沒有使用 *type 時
// 每次傳進來的 p 都會是複製一份新的（by value）
func (p person) updateName(newFirstName string) {
    p.firstName = newFirstName
}

func TestPassByPointer(t *testing.T) {

  jim := person{
		firstName: "Jim",
		lastName:  "Party",
	}

  jim.updateNameFromPointer("Aaron")  // It works as expected
	fmt.Println(jim.firstName) // Aaron
  jim.updateName("Bella")  // It doesn't work as expected
	fmt.Println(jim.firstName) // Aaron
}

// ------------------------------ pass by value vs pass by pointer ----------------------------------
// 為什麼需要使用指標（Pointer）
// Go 是一個 pass by value 的程式語言，也就是每當我們把值放入函式中時，Go 會把這個值完整的複製一份，並放到新的記憶體位址

/*
	在 Go 中雖然都是 pass by value，但要不要使用 Pointer 端看該變數的型別，某些型別會表現得類似 pass by reference，例如 slice，這時候就可以不用使用 Pointer 即可修改原變數的值。
		1. Value type: int, float, boolean, striing, strut --> using pointers to change these things in function
		2. Reference typr: slices, maps, channels, pointers, functions --> don't worry about pointers with these
*/

// 直接傳入參數時，thePerson 會複製一份新的
func updateFirstName(thePerson person) {
	thePerson.firstName = "Aaron"
}

// 透過 *type 傳入 pointer，會參照到原本的 thePerson
func updateFirstNameWithPointer(thePerson *person) {
	(*thePerson).firstName = "Aaron"
}


func TestPointerWithStrut(t *testing.T) {
	jim := person{
		firstName: "Jim",
		lastName: "Anderson",
	}

	fmt.Println(jim)   // {Jim Anderson}
	updateFirstName(jim)
	fmt.Println(jim)   // {Jim Anderson}

	jimPointer := &jim
	updateFirstNameWithPointer(jimPointer)
	fmt.Println(jim)   // {Aaron Anderson}
}

// --------------------------  function receiver with *type --------------------------------------
// 當我們在 function receiver 中使用 *type 後，這個函式將會自動把帶入的參數變成 pointer of the type

// 當 * 放在 type 前面時，這個 * 並不是運算子
// *type 指關於這個 type 的描述，也就是 pointer to the type
// *person 表示要針對「指稱到 person 的指標」做事


// 因為這裡有指稱要帶入的是 *person
func (p *person) updateName2(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("Current person is: %+v\n", p)
}

func TestFuncReceiverPointer(t *testing.T) {
	jim := person{
			firstName: "Jim",
			lastName:  "Party",
	}

	// 原本是這樣
	// jimPointer := &jim
	// jimPointer.updateName2("Aaron")

	// 所以，可以縮寫成，該 function 會自動去取 jim 的指標（記憶體位址）
	jim.updateName2("Aaron")
	jim.print()  // Current person is: {firstName:Aaron lastName:Party}
}


// --------------------------  pointer of pointer type --------------------------------------

func TestPointerAddress(t *testing .T) {
	name := "bill"
	namePointer := &name

	fmt.Println("1", namePointer)  // 0xc00008e1e0
	fmt.Println("2", &namePointer) // 0xc0000ae018
	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	fmt.Println("3", namePointer)  // 0xc00008e1e0
	fmt.Println("4", &namePointer) // 0xc0000ae028
}

// ---------------------------- new(T) ------------------------------------
// new 是 golang 中內建的函式，使用 new(T) 分配這個 Type 所需的記憶體，並回傳一個可以指稱到它的 pointer，概念上和 : = T{} 差不多：

type Vertex struct {
	X, Y float64
}
func TestNew(t *testing.T) {
	v := new(Vertex)   // new 回傳的是指稱到該變數的 pointer
	v.X = 10
	fmt.Println(v)    // &{10, 0}
	fmt.Println(&v)  	// 0xc000114038
	fmt.Println(*v)		// {10 0}
	fmt.Println(*&v)  // &{10 0}
}

