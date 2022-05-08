package func_test

import (
	"fmt"
	"math"
	"testing"
)

// compute 這個函式可接收其他函式作為參數
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func TestFuncAsParam(t *testing.T) {
	// 使用 expression 方式定義
	// sqrt(a^2 + b^2)
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(compute(hypot))    // 5, hypot(3, 4)
	fmt.Println(compute(math.Pow)) // 81,math.Pow(3, 4)
}


// 函式也可以在執行後回傳另一個函式（閉包）
func fibonacci() func() int {
	position := 0
	cache := map[int]int{}

	return func() int {
		position++
		if position == 1 {
			cache[position] = 0
			return 0
		} else if position <= 3 {
			cache[position] = 1
		} else {
			cache[position] = cache[position-2] + cache[position-1]
		}
		return cache[position]
	}
}

func TestFuncReturnFunc(t *testing.T) {
	f := fibonacci() // return a anonymous function
	for i := 0; i < 10; i++ {
			fmt.Println(f()) // execute the returned anonymous function by fibonacci
	}
}

// clousures
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func TestClousures(t *testing.T) {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2 * i))
	}
}

// about parameter :
// 比較好的作法是去定義 struct。如此，未來如果參數需要擴充，只需要改動 struct 和 func 內就好，不用去改動使用這個 function 地方的參數：
// STEP 1：定義參數的 structure
type addOpts struct {
	x int
	y int
	z int    // STEP 4：如果新增一個參數
}

// STEP 2：把參數的型別指定為 structure
func add(opts addOpts) int {
	// STEP 5：接收新的參數 z
	// return opts.x + opts.y
	return opts.x + opts.y + opts.z
}

// STEP 3：使用 add，參數的地方使用 structure
func TestStrut(t *testing.T) {
	// STEP 6：不用改用舊有參數的寫法
	result := add(addOpts{
		x: 10,
		y: 5,
	})

	newResult := add(addOpts{
		x: 10,
		y: 5,
		z: 7,
	})
	fmt.Println(result, newResult)
}

// return multi result
// 多個回傳值，會需要在定義回傳型別的地方給多個型別：

// func nameOfFunction(<arguments>) (<type>)
func swap(x, y string) (string, string) {
    return y, x
}

func TestMultiReturnValue(t *testing.T) {
    a, b := swap("hello", "world")
    fmt.Println(a, b)
}

// 回傳一個 function
// 名稱為 foo 的 function 會回傳一個 function
// 這個回傳的 function 會回傳 int
func foo() func() int {
    return func() int {
        return 100
    }
}

func TestReturnFunc(t *testing.T) {
    bar := foo()            // bar 會是一個 function
    fmt.Printf("%T\n", bar) // func() int
    fmt.Println(bar())      // 100
}


// 回傳帶有命名的值
// 在 Go 中可以在 func 定義回傳 type 的地方定義要回傳的變數，最後呼叫 return 的時候，該函式會自動去拿這兩個變數。這種做法稱作 naked return，但最好只使用在內容不多的函式中，否則會嚴重影響程式的可讀性：

// 用來說明回傳的內容
func swapWithSpecifiedName(x, y string) (a, b string) {
    a = y
    b = x
    return
}

func main() {
    foo, bar := swapWithSpecifiedName("hello", "world")
    fmt.Println(foo, bar)
}


// 因為 Go 本身並不是物件導向程式語言（object-oriented programming language），所以只能用 Type 搭配在函式中使用 receiver 參數來實作出類似物件程式語言的功能：
// 💡 提示：method 就只是帶有 receiver 參數 的函式。
type Person struct {
	name string
	age int
}

func (p Person) getInfo() string {
	return p.name
}

// 如果該函式不需要使用到 receiver 本身，可以簡寫成
func (Person) getInfo2() string {
	return "Xeon"
}

func TestMethodReceiver(t *testing.T) {
	p := Person{name: "Aaron", age: 32}
	fmt.Println(p.getInfo())  // Aaron
	fmt.Println(p.getInfo2())  // Xeon
}

// --------------------------------- Value Receiver -----------------------------------
// 如果單純要呈現某個 instance 的屬性值，這時候可以使用 value receiver：
// 建立一個新的型別稱作 'deck'，它會是帶有許多字串的 slice
// deck 會擁有 slice of string 所帶有的行為（概念類似繼承）
type deck []string

// 建立一個 deck 的 receiver
// 任何型別是 deck type 的變數，都將可以使用 "print" 這個方法
func (d deck) print() {
    for i, card := range d {
        fmt.Println(i, card)
    }
}

func newCard() string {
  return "Five of Diamonds"
}

func TestDeck(t *testing.T) {
	// 使用 deck type 定義變數
	cards := deck{
		"Ace of Diamonds",
		newCard(),
	}

	// 為陣列添加元素（append 本身不會改變原陣列）
	cards = append(cards, "Six of Spades")

	// 因為我們在 deck.go 中為 "deck" 這個型別添加了 print 的 receiver
	// 因此可以直接針對型別為 deck 的變數使用 print() 這個方法
	cards.print()
}



// --------------------------------- Pointer Receiver ---------------------------------