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
	// *使用 expression 方式定義
	// sqrt(a^2 + b^2)
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(compute(hypot))    // 5, hypot(3, 4)
	fmt.Println(compute(math.Pow)) // 81,math.Pow(3, 4)

	// *anonymous function
	// 沒有參數
  func() {
		fmt.Println("Hello anonymous") // Hello anonymous
	}()

	// 有參數
	func(i, j int) {
			fmt.Println(i + j) // 3
	}(1, 2)
}


// --------------------------------- closure --------------------------------
// 函式也可以在執行後回傳另一個函式（閉包）

// Example 1:
// fibonacci is a function that returns
// a function that returns an int.
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


// Example 2:
// 在 golang 的函式同樣適用閉包的概念，可以利用閉包把某個變數保留起來：
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func TestClousures(t *testing.T) {
	// pos 中的 sum 和 neg 中的 sum 是不同變數
	// pos == neg == func(x int) int inside adder()
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2 * i))
	}
}


// Example 3:
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func TestIntSeq(t *testing.T) {
	// We call intSeq, assigning the result (a function) to nextInt.
	// This function value captures its own i value,
	// which will be updated each time we call nextInt.
	nextInt := intSeq()

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	newInts := intSeq()
  fmt.Println(newInts()) // 1
}
// ----------------------------- 透過 structure 增加參數的可擴充性 -----------------------------------
// 透過 structure 增加參數的可擴充性


// 如果原本的 function 只需要兩個參數，可以這樣寫：
func add1(x, y int) int {
    return x + y
}

func TestAdd(t *testing.T) {
    fmt.Println(add1(1 ,2))
}


// about parameter :
// 比較好的作法是去定義 struct。如此，未來如果參數需要擴充，只需要改動 struct 和 func 內就好，
// 不用去改動使用這個 function 地方的參數：
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

// ------------------------------ func() return type ----------------------------------

// 沒有回傳值
// 對於沒有回傳值的函式可以不用定義回傳的型別：
func hello() {
	fmt.Println("Hello Go")
}

// 單一回傳值只需要在定義回傳型別的地方給一個型別就好：
func add3(i, j int) int {
    return i + j
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


// *回傳帶有命名的值
// 在 Go 中可以在 func 定義回傳 type 的地方定義要回傳的變數，
// 最後呼叫 return 的時候，該函式會自動去拿這兩個變數。這種做法稱作 naked return，
// 但最好只使用在內容不多的函式中，否則會嚴重影響程式的可讀性：

// 用來說明回傳的內容
func swapWithSpecifiedName(x, y string) (a, b string) {
    a = y
    b = x
    return
}

func TestNakedReturn(t *testing.T) {
    foo, bar := swapWithSpecifiedName("hello", "world")
    fmt.Println(foo, bar) // world hello
}

// ------------------------------- Methods | Function Receiver ---------------------------------
// *Go語言的Receiver是綁定function到特定type成為其method的一個參數。
// *換句話說，一個function加了receiver即成為一個type的method。

// *以Go的function和method的差別在於,是否有receiver。
// *method有reciever，function沒有。

// Receiver參數必須指定一個 型態T 或 指向該型態的指標(pointer) *T。
// Receiver指定的T稱為base type，不可以是interface或pointer，且必須定義在與method同個package中。
// 一旦function定義了receiver成為base type的method後，只有該型態的變數可以.來呼叫method

// *Receiver又分為value reciever及pointer receiver
// 		*Value receiver的型態前不加*，method的receiver為複製值；
// 		*Pointer receiver的型態前加*，method的receiver為指標。

// 因為 Go 本身並不是物件導向程式語言（object-oriented programming language），
// 所以只能用 Type 搭配在函式中使用 receiver 參數來實作出類似物件程式語言的功能：
// 💡 提示：method 就只是帶有 receiver 參數 的函式。
type Person struct {
	name string
	age int
}

func getInfo() string {
	return "I am a function"
}

// function with receiver --> method of type Person
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
	fmt.Println(getInfo()) // I am a function
}

// ------------------------ Methods | Function Receiver : 1. Value Receiver ---------------------------
// 如果單純要呈現某個 instance 的屬性值，這時候可以使用 value receiver：
/*
		1. 透過 type deck []string 來定義一個名為 deck 的型別。要留意的是，
			deck 的本質上仍然是 []string 它可以使用 slice type 的方法，
			也可以把 slice 帶入指定為 deck 型別的函式內使用。

		2. (d deck) 為 deck 添加一個 receiver function

		3. print 是函式名稱

		4. 當我們呼叫 cards.print() 時，
			這個 cards 就會變成這裡指稱到的 d，這個 d 很類似在 JavaScript 中的 this 或 self，
			但在 Go 中慣例上不會使用 this 和 self 來取名，慣例上會使用該 type 的前一兩個字母的縮寫。

		*如果用物件導向的概念來說明，那麼 deck 就類似一個 class，我們在這個 class 中添加了 print() 的方法，
		*同時也可以用 cards := deck {...} 來產生一個名為 cards 的 deck instance。

*/
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

	// 0 Ace of Diamonds
	// 1 Five of Diamonds
	// 2 Six of Spades
}


// STEP 1：根據型別 string，定義 color 型別
type color string

// STEP 2：
// (c color)，定義 color 的 function receiver
// describe(description string) string，describe 這個 function 接受一個字串的參數 description，並會回傳 string
func (c color) describe(description string) string {
  // 這裡的 c 就類似 this
  return string(c) + " " + description
}

func TestDescribe(t *testing.T) {
	// 根據型別 color 建立變數 c
	c := color("Red")

	fmt.Println(c.describe("is an awesome color")) // Red is an awesome color
}
// ------------------  Methods | Function Receiver : 2. Pointer Receiver -----------------
// 也可以把某一個 method 定義給某個 Type 的 Pointer，
// 如果是想要修改某一個 instance 中屬性的資料，
// 這時候的 receiver 需要使用 pointer receiver 才能修改到該 instance，否則無法修改到該 instance 的資料。
// 例如，下面程式中的 ScalePointer 這個 method 就是定義給 *Vertex 這個 pointer：

// 💡 補充：同樣的，如果 receiver 接收的是 value receiver 而非 pointer receiver 時，
// 		使用 pointer receiver 去執行某方法也會成功：v.Abs() 等同於 (&v).Abs()。


// *把某一個 method 定義給某個 Type 的 Pointer
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale 這個 methods 會修改到的是 v 的複製，而無法直接修改到 v
func (v Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// ScalePointer 這個 methods 可以修改 v 中的屬性與值
func (v *Vertex) ScalePointer(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func TestVertex(t *testing.T) {
	v := Vertex{3, 4}

	v.Scale(10)
	fmt.Println(v.Abs()) // 5

  // 這裡雖然 v 應該要是 *Vertex，但我們使用的是 Vertex 邏輯上要發生錯誤
  // 但因為 ScalePointer 這個方法本身有 pointer receiver
  // 因此 Go 會自動將 v.ScalePointer(10) 視為 (&v).ScalePointer(10)
  v.ScalePointer(10)   // 等同於 (&v).ScalePointer(10)
  fmt.Println(v.Abs()) // 50，等同於（&v).Abs()
	fmt.Println((&v).Abs()) // 50
}

// *同樣的功能一樣可以改用 function 的方式來寫：
// 使用 *Type 當作 function 的參數，也就是 *Vertex
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleWithPointer(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func TestVertex2(t *testing.T) {
	v := Vertex{3, 4}
	Scale(v, 10)
	fmt.Println(v) // {3 4}
	fmt.Println(Abs(v)) // 5

	// 留意帶進去的變數需要是 Pointer，也就是 &v
	ScaleWithPointer(&v, 10) // 和 receiver 不同，「不能」簡化為 ScaleWithPointer(v, 10)
	fmt.Println(v)      // {30 40}
	fmt.Println(Abs(v)) // 50
}


// ------------------------------ Golang value receiver與pointer receiver差別 ----------------------------------
/*
		此外實作interface的方法時，value type的值無法分派到pointer receiver實作的interface變數；
		反之pointer type的值可以分派到value receiver實作的interface變數。

		這是因為pointer type的method sets同時包含了pointer receiver及value receiver的methods；
		而value type的method sets只有value receiver的methods。節錄Go規格文件的Method sets：

			The method set of any other type T consists of all methods declared with receiver type T.
			The method set of the corresponding pointer type *T is the set of all methods declared with receiver *T or T
			(that is, it also contains the method set of T).

		*任意型態T 的method sets包含了receiver T 的全部methods。
		*Pointer type *T 的method sets包含了receiver *T 及 T 的全部methods

*/
type Worker interface {
    Work()
}

type Employee struct {
    Id   int
    Name string
    Age  int
}

type Employee2 struct {
	Id   int
	Name string
	Age  int
}

// method of pointer receiver
func (e *Employee) Work() {
    fmt.Println(e.Name + " works")
}

// method of value receiver
func (e Employee2) Work() {
	fmt.Println(e.Name + " works")
}

func TestValueAndPointerReceiver(t *testing.T) {
		// 因為Worker的實作為pointer receiver而非value receiver，因此只能接受pointer type的Employee值，因為才有包含pointer receiver的method。
    // var worker Worker = Employee{1, "John", 33} // compile error
		var worker Worker = &Employee{1, "John", 33} // assign pointer of Employee literal to worker
    worker.Work() // John works
}

func TestValueAndPointerReceiver2(t *testing.T) {
	// value receiver實作的interface變數則可同時接收value type或pointer type的值，
	// 因為pointer type的method sets同時包含了pointer receiver及value receiver的methods。
	var worker1 Worker = &Employee2{1, "John", 33} // assign pointer of Employee literal to worker1
	worker1.Work() // John works

	var worker2 Worker = Employee2{2, "Mary", 28} // assign value of Employee literal to worker2
	worker2.Work() // Mary works
}



// IIFE (Immediately Invoked Function Expression)
func TestIIFE(t *testing.T) {
	slice := []string{"a", "a"}

	// 使用 IIFE 的寫法
	func(slice []string) {
		slice[0] = "b"
		slice[1] = "b"
	}(slice)

	fmt.Println(slice) // [b b]
}