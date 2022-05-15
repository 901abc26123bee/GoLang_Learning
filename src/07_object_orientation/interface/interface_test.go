package interface_test

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
)

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}

// --------------------------------- interface -------------------------------------------
/*
	interface 的概念有點像是的藍圖，
	先定義某個方法的名稱（function name）、
	會接收的參數及型別（list of argument type）、
	會回傳的值與型別（list of return types）。
	定義好藍圖之後，並不去管實作的細節，實作的細節會由每個型別自行定義實作（implement）。
*/
// 任何型別，只要符合定義規則的話，就可以被納入 bot interface 中
type bot interface {
	// getGreeting 這個函式需要接收兩個參數（string, int），並回傳 (string, error) 才符合入會資格
	getGreeting(string, int) (string, error)

	// getBotVersion 這個函式需要回傳 float 才符合入會資格
	getBotVersion() float64
}

/*
	Interface 是什麼？
	透過 interface 可以定義一系列的 method signatures 來讓 Type 透過 methods 加以實作，
	*也就是說 interface 可以用來定義 type 有哪些行為（behavior）。

	interface 就像藍圖一樣，在裡面會定義函式的名稱、接收的參數型別以及最終回傳的資料型別，
	而 Type 只需要根據這樣的藍圖加以實作（implement）出這些方法。在 Go 中，
	*Type 不需要明確使用 implement 關鍵字來說明它實作了哪個 interface，
	*只要它符合了該 interface 中所定義的 method signature，就等於自動實作了該 interface。

	舉例來說，定義「狗」的 interface 包含方法「走路」、「吠」，
	只要有一個 Type 它能夠提供「走路」和「吠」的方法，那個這個 Type 就自動實作（implement）了「狗」這個 interface，
	不需要額外使用 implement 關鍵字。

	*另外，任何資料型別，只要實作了該 interface 之後，都可以被視為該 interface 的 type（polymorphism）。
*/

/*
	interface 可以被賦值
		- interface 沒有被賦值前，其 type 和 value 都會是 nil
		- interface 被賦值後，它的型別值會變成實作它的 Type 的型別和值

	*interface 可以被想成是帶有 (value, type) 的元組（tuple），
	*當我們呼叫某個 interface value 的方法時，實際上就是將該 value 去執行與該 type 相同名稱的方法（method）：
		- interface 的變數「動態值（dynamic value / concrete value）」會是實作此 interface 的 Type 的 value
		- interface 的變數「動態型別（dynamic type / concrete type）」會是實作此 interface 的 Type 的型別
		- interface 沒有「靜態值（static value）」
		- interface 的「靜態型別（static type）」，則是該 interface 的本身，例如 type Shape interface{}，
			這個 interface 所建立的變數，其靜態型別即是 Shape

	interface 的 dynamic type 又稱作 concrete type，因為當我們想要存取該 interface 的型別時，它回傳的會是 dynamic value，原本的 static type 會被隱藏。
*/

type Shape interface {
	Area() float64
}

// interface 沒有被賦值前，其 type 和 value 都會是 `nil`
func TestInterfacesWithoutInitialize(t *testing.T) {
	var s Shape
	fmt.Printf("Type of s is %T\n", s) // Type of s is <nil>
	fmt.Println("value of s is", s)    // value of s is <nil>
}

// Rect 實作了 Shape interface
// Rect 所建立的變數同時會符合 Rect（Struct Type）和 Shape（Interface Type）
type Rect struct {
	width  float64
	height float64
}

// 當一個 type(Rect) 實作（implement）了某個 interface 後，
// 該 Type 產生的變數除了會是原本的 Type 外，也同時屬於該 interface type，及 polymorphism
// type Rect 會實作 interface Shape, 但並不需要開發者主動去宣告它, Interface 會隱性的被 implement
func (r Rect) Area() float64 {
	return r.width * r.height
}

// 如果 interface 有賦值的話，
// 則可以看到顯示的 dynamic type 和 dynamic value 會是實作該 interface 的 Type 的 method 和 value：
func TestIntefaceWithInitialize(t *testing.T) {
	// interface 被賦值後，它的型別值會變成實作它的 Type 的型別和值
	// 當把 Rect 作為 Shape interface 的值後，
	// Shape 的 Type（dynamic type）會變成 Rect、Value（dynamic value）會變成 Rect 的值（{3, 5}）
	var s Shape = Rect{3, 5}
	fmt.Printf("(%T, %v) \n", s, s) // (interface_test.Rect, {3 5})
	fmt.Println(s.Area())           // 15 => 可以直接用 Shape interface 來呼叫方法
}

// ---------------------------- 同時符合多個 interfaces 的 Type ------------------------------------
// *同時符合多個 interfaces 的 Type

// type Shape interface {
// 	Area() float64
// }

type Object interface {
	Volume() float64
}

type Cube struct {
	side float64
}

func (c Cube) Area() float64 {
	return 6 * (c.side * c.side)
}

func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}

func TestStructImplementMultiInterface(t *testing.T) {
	c := Cube{3}
	// 把 Cube 指派給 Shape 或 Object interface 所建立的變數
	var s Shape = c
	var o Object = c
	fmt.Println(c.Area(), c.Volume())                                // 54 27
	fmt.Println("volume of s of interface type Shape is", s.Area())  // 54
	fmt.Printf("Shape (%T, %v) \n", s, s)                            // Shape (interface_test.Cube, {3})
	fmt.Println("area of o of interface type Object is", o.Volume()) // 27
	fmt.Printf("Object (%T, %v) \n", o, o)                           // Object (interface_test.Cube, {3})
}

// *Type Assertions
// value := i.(Type)   // 得到實作 interface 的 Type 的值
// i ： type interface
// Type：實作該 interface 的 type
func TestTypeAssertions(t *testing.T) {
	var s Shape = Cube{3}

	// 原本不能執行 s.Volume()，但把 s 轉換成 Cube 後得到 c，即可使用 c.Volume()
	c := s.(Cube) // i.(T) 可以把 interface(i) 當中的 T 的值取出來

	fmt.Println("volume of s of interface type Shape is", c.Area())  // 54
	fmt.Printf("Shape (%T, %v) \n", c, c)                            //Shape (main.Cube, {3})
	fmt.Println("area of o of interface type Object is", c.Volume()) // 27
	fmt.Printf("Object (%T, %v) \n", c, c)                           // Object (main.Cube, {3})
}

// 然而，如果「Cube Type 沒有實作 interface 的方法」，
// 或者「雖然 Cube Type 有實作 interface 的方法，但 i 並沒有該 Type 的 concrete value 的話」都會報錯。例如：
func TestTypeAssertionsPanic(t *testing.T) {
	// 由於 s 沒有被實際賦職，因此 c 沒有 dynamic/concrete value（nil）
	// 因此使用 c.Area() 時會出現 panic
	var s Shape
	c := s.(Cube)
	fmt.Println(c.Area()) // panic: interface conversion: interface_test.Shape is nil
}

func TestTypeAssertionsPanic2(t *testing.T) {
	// 為了避免這種情況可以使用：
	// value, ok := i.(Type)
	// ok ：當 i 有 dynamic type 和 dynamic value 時，ok 會是 true，value 則會是 dynamic value；
	// 否則 ok 會是 false，value 會是該 Type 的 zero value。
	var s Shape
	c, ok := s.(Cube)
	fmt.Println(c.Area(), ok) // 0 false
}

// *將變數的 interface 轉換成另一個 interface
// v := i.(I)    // i 是原本的 interface，I 轉變成的 interface
// Type Assertions 除了用來確保某一個 interface 是否有 dynamic value / concrete value 之外，
// 也可以用來將一個變數從原本的 interface 轉成另一個 interface：
type Person interface {
	getFullName() string
}

type Salaried interface {
	getSalary() int
}

type Employee struct {
	firstName string
	lastName  string
	salary    int
}

func (e Employee) getFullName() string {
	return e.firstName + " " + e.lastName
}

func (e Employee) getSalary() int {
	return e.salary
}

func TestChangeInterface(t *testing.T) {
	// johnPerson 原本是 Person interface，只能使用 Person interface 中的 getFullName 方法
	var johnPerson Person = Employee{"John", "Adams", 50000}
	fmt.Printf("full name: %v \n", johnPerson.getFullName()) // full name: John Adams

	// 使用 i.(I)，可以把原本屬於 Person interface 的 johnPerson 轉成 Salaried interface
	johnSalary := johnPerson.(Salaried)
	fmt.Printf("salary : %v \n", johnSalary.getSalary()) // salary : 50000
}

// ------------------------------ Embedding interfaces ----------------------------------
// Embedding interfaces
// 將兩個 interface 組合成一個新的 interface：
type Shape2 interface {
	Area() float64
}

type Object2 interface {
	Volume() float64
}

// `Material` interface 是由 Shape 和 Object 組合而成
type Material interface {
	Shape2
	Object2
}

// ------------------------------ 使用 Interface 建立自己的 function ---------------------------------

// *使用 Interface 建立自己的 function
// 也就是說，任何 Type 底下只要有 String() 這個方法且回傳 string，都會被歸到 Stringer interface。
// *Example 1:
type Stringer interface {
	String() string
}

type Person2 struct {
	Name string
	Age  int
}

// Person 這個 type 有 String 方法，且會回傳 string，因此可被歸在 Stringer interface
func (p Person2) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func TestDefineMethodWithInterface(t *testing.T) {
	a := Person2{"Arthur", 42}
	z := Person2{"Zaphod Beeblebrox", 9001}

	fmt.Printf("%T, %+v \n", a, a) // interface_test.Person2, Arthur (42 years)
	fmt.Printf("%T, %+v \n", z, z) // interface_test.Person2, Zaphod Beeblebrox (9001 years)
	fmt.Println(a)                 // Arthur (42 years)
	fmt.Println(z)                 // Zaphod Beeblebrox (9001 years)
}

// 幫 IPAddr type 客製化自己的 String 方法
type IPAddr [4]byte

// String() string 符合 Stringer  interface
func (ip IPAddr) String() string {
	var ips []string
	for _, ipNumber := range ip {
		ips = append(ips, strconv.Itoa(int(ipNumber)))
	}
	return strings.Join(ips, ".")
}

func TestDefineMethodWithInterface2(t *testing.T) {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
	// loopback: 127.0.0.1
	// googleDNS: 8.8.8.8
}

// *Example 2:
// STEP 1：建立一個 logWriter 的 type
type logWriter struct{}

// STEP 2：根據 Writer Interface 的定義（https://golang.org/pkg/io/#Writer）
// 來撰寫 logWriter 的 Write function
// 如此，它將會被歸類在 Writer Interface 內
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}

func TestDefineMethodWithInterface3(t *testing.T) {
	resp, err := http.Get("https://pjchender.github.io")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// STEP 3：建立 logWriter
	lw := logWriter{}

	// STEP 4：因為 logWriter 已經歸類在 Writer Interface，所以可以帶入 io.Copy 內
	io.Copy(lw, resp.Body)
	// <html>...</html>
	// Just wrote this many bytes:  14860
}

//------------------------ empty interface, Type assertions, Type switchs -------------------------------
// *empty interface
// 沒有定義任何方法的 interface 稱作 empty interface，
// 由於所有的 types 都能夠實作 empty interface，因此它的值會是 any type：
type I interface{}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func TestEmptyInterface(t *testing.T) {
	var i I
	describe(i) // (<nil>, <nil>)

	i = 42
	describe(i) //(42, int)

	i = "hello"
	describe(i) // (hello, string)
}

// *Type assertions
// 前面有提到，interface 可以想成是 (value, type) 的元組，
// 透過 type assertion 則提供了一種方式可以存取該 interface value 底層的 concrete value：
func TestTypeassertion(t *testing.T) {
	// 斷定 interface 的 concrete type 是 T，並將 T 的 value 指派到變數 t
	// t := i.(T)    // 如果型別不正確會直接 panic

	// 如果要檢測某 interface 是否包含某一個 type，則需要接收兩個回傳值－ underlying value 和 assertion 是否成功
	// t, ok := i.(T)
	// 如果 i 有 T，則 t 會得到 underlying value，而 ok 會是 true
	// 如果 i 沒有 T，則 t 會得到 type T 的 zero value，且 ok 會是 false
}

// *Type switches
// type switch 很適合做為某一個會接收多種型別方法的寫法，
// 在該方法中透過 type switch 的方式來根據不同的型別回傳不同的內容或執行不同的行為：
func foo(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("i stored as type string: ", i)
	case int:
		fmt.Println("i stored as type int: ", i)
	default:
		fmt.Println("i stored something else: ", i)
	}
}

// 如此 foo 這個 methods 就可以接多種不同的參數
func TestTypeSwitch(t *testing.T) {
	foo("this is string")
	foo(30)
	foo(true)
	// i stored as type string: this is string
	// i stored as type int: 30
	// i stored something else: true
}

// *CAUTION : 不同於 Type Assertions 中 i.(T) 中的 T 指的是型別，Type switch 中 i.(type) 的 type 是固定的關鍵字，
// *只能用 type ，不能用其他字，且只能在 switch 中使用
// type switch 和一般的 switch 語法相同，
// 只是 switch 判斷的內容是使用 type assertion（i.(type)）、在 case 的地方則是判斷某一 interface value 的型別：
// 透過 i.(type) 可以取得該 interface 的 dynamic type：
func do(i interface{}) {
	switch v := i.(type) {
	case int:
			fmt.Printf("Twice %v is %v \n", v, v*2)
	case string:
			fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
			fmt.Printf("I don't know about type %T!\n", v)
	}
}

func TestTypeSwitch2(t *testing.T) {
	var i interface{}
	i = "Hello"
	fmt.Printf("(%v, %T)\n", i, i) // (Hello, string)

	do(21)      // Twice 21 is 42
	do("hello") // "hello" is 5 bytes lon
	do(true)    // I don't know about type bool!
}