package method

import (
	"fmt"
	"math"
	"testing"
)


// *"A method is a function with an implicit first argument, called a receiver."
// * Go 裡面的物件導向，沒有任何的私有、公有關鍵字，透過大小寫來實現（大寫開頭的為公有，小寫開頭的為私有），方法也同樣適用這個原則。
/*
	* 在使用 method 的時候重要注意幾點

			雖然 method 的名字一模一樣，但是如果接收者不一樣，那麼 method 就不一樣
			method 裡面可以訪問接收者的欄位
			呼叫 method 透過 . 訪問，就像 struct 裡面訪問欄位一樣
			指標作為 Receiver 會對實體物件的內容發生操作，而普通型別作為 Receiver 僅僅是以副本作為操作物件，並不對原實體物件發生操作。

  那是不是 method 只能作用在 struct 上面呢？當然不是囉，
* method可以定義在任何你自訂的型別、內建型別、struct 等各種型別上面。
* struct 只是自訂型別裡面一種比較特殊的型別而已，還有其他自訂型別宣告，可以透過如下這樣的宣告來實現。

*/

// 宣告自訂型別
// type typeName typeLiteral

type ages int

type money float32

type months map[string]int


//------------------------------------------------------------------------------------------------
type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width*r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}


func TestMethod(t *testing.T) {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
}

// ----------------------------------------------------------------
//* 如果一個 method 的 receiver 是 *T，你可以在一個 T 型別的變數 V 上面呼叫這個 method，而不需要 &V 去呼叫這個 method。
//* 如果一個 method 的 receiver 是 T，你可以在一個 *T 型別的變數 P 上面呼叫這個 method，而不需要 *P 去呼叫這個 method。

// 在宣告 const 採用分組方式宣告時，第一個常數可以用於預設值，假設他的值為 0 ，則在同一分組，其他常數預設會用前一個常數值。
// * Go 的 iota 關鍵字，它預設值會是 0，每次當 const 分組陸續呼叫宣告時，就會加 1，直到遇到下一個 const 宣告 iota 時，才會重置為 0。
// 如果常數宣告忽略值時，就會預設採用前一個值。
type Color byte

const(
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Box struct {
	width, height, depth float64
	color Color
}

type BoxList []Box // a slice of boxes

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) {
	//* 如果一個 method 的 receiver 是 T，
	//* 你可以在一個 *T 型別的變數 P 上面呼叫這個 method，而不需要 *P 去呼叫這個 method。
	// (*b).color = c // same as
	b.color = c
}

func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PaintItBlack() {
	for i := range bl {
		//* 如果一個 method 的 receiver 是 *T，
		//* 你可以在一個 T 型別的變數 V 上面呼叫這個 method，而不需要 &V 去呼叫這個 method。
		// (&bl[i]).SetColor(BLACK) // same as
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string {"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func TestMethod2(t *testing.T) {
	fmt.Println(WHITE, BLACK, BLUE, RED, YELLOW) // 0 1 2 3 4

	boxes := BoxList {
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	fmt.Println("The color of the last one is",boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestColor().String())

	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())
	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())
	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor())

	/*
		We have 6 boxes in our set
		The volume of the first one is 64 cm³
		The color of the last one is YELLOW
		The biggest one is YELLOW
		
		Let's paint them all black
		The color of the second one is BLACK
		Obviously, now, the biggest one is BLACK
		Obviously, now, the biggest one is BLACK
	*/
}


// ------------------------------------------------------------------------------------------------
// method 繼承
// 前面一章我們學習了欄位的繼承，那麼你也會發現 Go 的一個神奇之處，method 也是可以繼承的。
// * 如果匿名欄位實現了一個 method，那麼包含這個匿名欄位的 struct 也能呼叫該 method。
type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human // 匿名欄位
	school string
}

type EmployeeA struct {
	Human // 匿名欄位
	company string
}

// 在 human 上面定義了一個 method
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func TestMethodExtendtion(t *testing.T) {
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := EmployeeA{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	mark.SayHi() // Hi, I am Mark you can call me on 222-222-YYYY
	sam.SayHi() // Hi, I am Sam you can call me on 111-888-XXXX
}

// ------------------------------------------------------------------------------------------------
// method 重寫
type HumanB struct {
	name string
	age int
	phone string
}

type StudentB struct {
	HumanB // 匿名欄位
	school string
}

type EmployeeB struct {
	HumanB // 匿名欄位
	company string
}

// Human 定義 method
func (h *HumanB) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Employee 的 method 重寫 Human 的 method
func (e *EmployeeB) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

func TestMethodOverride(t *testing.T) {
	mark := StudentB{HumanB{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := EmployeeB{HumanB{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	mark.SayHi() // Hi, I am Mark you can call me on 222-222-YYYY
	sam.SayHi() // Hi, I am Sam, I work at Golang Inc. Call me on 111-888-XXXX
}