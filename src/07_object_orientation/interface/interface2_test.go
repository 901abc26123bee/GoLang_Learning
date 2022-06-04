package interface_test

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh-tw/02.6.md
/*
	什麼是 interface
*	簡單的說，interface 是一組 method 簽名的組合，我們透過 interface 來定義物件的一組行為。
	interface 型別定義了一組方法，如果某個物件實現了某個介面的所有方法，則此物件就實現了此介面。
*/
// * Go 裡面的物件導向，沒有任何的私有、公有關鍵字，透過大小寫來實現（大寫開頭的為公有，小寫開頭的為私有），方法也同樣適用這個原則。

type Human1 struct {
	name string
	age int
	phone string
}

type Student1 struct {
	Human1 //匿名欄位 Human
	school string
	loan float32
}

type Employee1 struct {
	Human1 //匿名欄位 Human
	company string
	money float32
}

// Human 物件實現 Sayhi 方法
func (h *Human1) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human 物件實現 Sing 方法
func (h *Human1) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

// Human 物件實現 Guzzle 方法
func (h *Human1) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// Employee 過載 Human 的 Sayhi 方法
func (e *Employee1) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //此句可以分成多行
}

// Student 實現 BorrowMoney 方法
func (s *Student1) BorrowMoney(amount float32) {
	s.loan += amount // (again and again and...)
}

// Employee 實現 SpendSalary 方法
func (e *Employee1) SpendSalary(amount float32) {
	e.money -= amount // More vodka please!!! Get me through the day!
}

// 定義 interface

type Men1 interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap1 interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent1 interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}
// -------------------------------------------------------------------------------------------------------------
// interface 值
// * duck-typing
// interface 可以被任意的物件實現。我們看到上面的 Men interface 被 Human、Student 和 Employee 實現。
// 同理，一個物件可以實現任意多個 interface，例如上面的 Student 實現了 Men 和 YoungChap 兩個 interface。
// * 任意的型別都實現了空 interface（我們這樣定義：interface{}），也就是包含 0 個 method 的 interface。
type Human2 struct {
	name string
	age int
	phone string
}

type Student2 struct {
	Human2 //匿名欄位
	school string
	loan float32
}

type Employee2 struct {
	Human2 //匿名欄位
	company string
	money float32
}

// Human 實現 SayHi 方法
func (h Human2) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human 實現 Sing 方法
func (h Human2) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

// Employee 過載 Human 的 SayHi 方法
func (e Employee2) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
	}

// Interface Men 被 Human,Student 和 Employee 實現
// 因為這三個型別都實現了這兩個方法
type Men2 interface {
	SayHi()
	Sing(lyrics string)
}

func TestDuckTyping(t *testing.T) {
	mike := Student2{Human2{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student2{Human2{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee2{Human2{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee2{Human2{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	// 定義 Men 型別的變數 i
	var i Men2

	// i 能儲存 Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	// i 也能儲存 Employee
	i = tom
	fmt.Println("This is tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	// 定義了 slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men2, 3)
	// 這三個都是不同型別的元素，但是他們實現了 interface 同一個介面
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x{
		value.SayHi()
	}
}
//* 透過上面的程式碼，你會發現 interface 就是一組抽象方法的集合，
//  它必須由其他非 interface 型別實現，而不能自我實現，
//* Go 透過 interface 實現了 duck-typing：
//* 即「當看到一隻鳥走起來像鴨子、游泳起來像鴨子、叫起來也像鴨子，那麼這隻鳥就可以被稱為鴨子」。


// ------------------------------------------------------------------------------------------------
// interface 函式參數
/*
	空 interface
	空 interface(interface{}) 不包含任何的 method，
	正因為如此，所有的型別都實現了空 interface。空 interface 對於描述起不到任何的作用（因為它不包含任何的 method），
*	但是空 interface 在我們需要儲存任意型別的數值的時候相當有用，因為它可以儲存任意型別的數值。
	它有點類似於 C 語言的 void* 型別。

	//* 定義 a 為空介面
	var a interface{}
	var i int = 5
	s := "Hello world"
	//* a 可以儲存任意型別的數值
	a = i
	a = s

*	一個函式把 interface{} 作為參數，那麼他可以接受任意型別的值作為參數，如果一個函式回傳 interface{}，那麼也就可以回傳任意型別的值。
*/

/*
	舉個例子：fmt.Println 是我們常用的一個函式，它可以接受任意型別的資料。
	開啟 fmt 的原始碼檔案，會看到這樣一個定義:

	type Stringer interface {
		String() string
	}

	也就是說，任何實現了 String 方法的型別都能作為參數被 fmt.Println 呼叫
*	即如果需要某個型別能被 fmt 套件以特殊的格式輸出，你就必須實現 Stringer 這個介面。如果沒有實現這個介面，fmt 將以預設的方式輸出。
* 實現了 error 介面的物件（即實現了 Error() string 的物件），使用 fmt 輸出時，會呼叫 Error() 方法，因此不必再定義 String() 方法了。
*/

type Human3 struct {
	name string
	age int
	phone string
}

// 透過這個方法 Human 實現了 fmt.Stringer
func (h Human3) String() string {
	return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years -  ✆ " +h.phone+"❱"
}

func TestEmptyInterface3(t *testing.T) {
	Bob := Human3{"Bob", 39, "000-7777-XXX"}
	fmt.Println("This Human is : ", Bob) // This Human is :  ❰Bob - 39 years -  ✆ 000-7777-XXX❱
}

// ------------------------------------------------------------------------------------------------
// interface 變數儲存的型別

/*
* Comma-ok 斷言
	Go 語言裡面有一個語法，可以直接判斷是否是該型別的變數：
		value, ok = element.(T)，
	這裡 value 就是變數的值，ok 是一個 bool 型別，element 是 interface 變數，T 是斷言的型別。
	如果 element 裡面確實儲存了 T 型別的數值，那麼 ok 回傳 true，否則回傳 false。

* switch 測試

*/
type Element interface {}
type List []Element

type Person4 struct {
	name string
	age int
}

// 定義了 String 方法，實現了 fmt.Stringer
func (p Person4) String() string {
	return "(name: " + p.name + " - age: "+strconv.Itoa(p.age)+ " years)"
}

func TestTypeOfInterface(t *testing.T) {
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person4{"Siman", 20}

	// * Comma-ok 斷言
	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		} else {
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}

	// * switch 測試
	for index, element := range list{
		switch value := element.(type) {
			case int:
				fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
			case string:
				fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
			case Person:
				fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
			default:
				fmt.Printf("list[%d] is of a different type", index)
		}
	}
}

// ------------------------------------------------------------------------------------------------
// 嵌入 interface
// container/heap 裡面有這樣的一個定義
type Interface interface {
	sort.Interface // 嵌入欄位 sort.Interface
	Push(x interface{}) // a Push method to push elements into the heap
	Pop() interface{} // a Pop elements that pops elements from the heap
}
// 我們看到 sort.Interface 其實就是嵌入欄位，把 sort.Interface 的所有 method 給隱式的包含進來了。也就是下面三個方法：
/*
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less returns whether the element with index i should sort
	// before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

另一個例子就是 io 套件下面的 io.ReadWriter ，它包含了 io 套件下面的 Reader 和 Writer 兩個 interface：

io.ReadWriter
type ReadWriter interface {
	Reader
	Writer
}
*/

// ------------------------------------------------------------------------------------------------
// reflection
/*
	reflect 流程可分成三個步驟
		1. 將資料本身透過 reflect 轉換為物件結構（reflect.Type 或者 reflect.Value，根據不同的情況呼叫不同的函式
			t := reflect.TypeOf(i)    // 得到型別的 Meta 資料，透過 t 我們能取得型別定義裡面的所有元素
			v := reflect.ValueOf(i)   // 得到實際的值，透過 v 我們取得儲存在裡面的值，還可以去改變值
		2. 將物件轉化為對應的值以及類型
		3. 可修改 reflect 欄位
*/

type MyInt struct {
	number int `format:"number"`
	note string `format:"note"`
}

type MyInterface interface {
	Compute(int) string
	Double()
	discount(int) string	// need to be capital --> Discount(int) string
}

func (myint *MyInt) Compute(number int) string {
	res := myint.number * 2
	return strconv.Itoa(res) //  convert an int value to string
}

func (myint *MyInt) Double() {
	myint.number = myint.number * 2
}

func (myint *MyInt) discount(amount int) string {
	res := myint.number - 1
	return strconv.Itoa(res) //  convert an int value to string
}

func TestReflection(t *testing.T) {

	var myint MyInt
	myint = MyInt{number: 5, note: "lala"}
	// 使用 reflect 一般分成三步，
	// 1. 要去反射是一個型別的值（這些值都實現了空 interface），
	// 		首先需要把它轉化成 reflect 物件（reflect.Type 或者 reflect.Value，根據不同的情況呼叫不同的函式）。

	ty := reflect.TypeOf(myint)    // 得到型別的 Meta 資料，透過 t 我們能取得型別定義裡面的所有元素
	va := reflect.ValueOf(myint)   // 得到實際的值，透過 v 我們取得儲存在裡面的值，還可以去改變值
	fmt.Println(ty, va) // interface_test.MyInt {5 lala}

	// ------------------------------------------------------------------------------------------------
	// 2. 轉化為 reflect 物件之後我們就可以進行一些操作了，也就是將 reflect 物件轉化成相應的值，
	// tag := ty.Elem().Field(0).Tag  // 取得定義在 struct 裡面的標籤
	// name := va.Elem().Field(0).String()  // 取得儲存在第一個欄位裡面的值
	// fmt.Println(name, tag) // panic

	tag := reflect.TypeOf(myint).Field(0).Tag  // 取得定義在 struct 裡面的標籤
	name := reflect.ValueOf(myint) .Field(0).String()  // 取得儲存在第一個欄位裡面的值
	fmt.Println(name, tag) // <int Value> format:"number"

	//* 如果 reflect.TypeOf() 接受的是個指標，因為指標實際上只是個位址值，
	//* 必須要透過 Type 的 Elem 方法取得指標的目標 Type，才能取得型態的相關成員：
	tag2 := reflect.TypeOf(&myint).Elem().Field(0).Tag  // 取得定義在 struct 裡面的標籤
	name2 := reflect.ValueOf(&myint) .Elem().Field(0).String()  // 取得儲存在第一個欄位裡面的值
	fmt.Println(name2, tag2) // <int Value> format:"number"

	// get all property
	for i, n := 0, ty.NumField(); i < n; i++ {
		f := ty.Field(i)
		fmt.Println(f.Name, f.Type)
		// number int
		// note string
	}

	// get all methods
	for i, n := 0, ty.NumMethod(); i < n; i++ {
		f := ty.Method(i)
		fmt.Println(f.Name, f.Type)
		// nothing
	}

	var myinterface MyInterface = &MyInt{8, "hi"} // since we are using pointer receiver
	ity := reflect.TypeOf(myinterface)
	// get all methods
	for i, n := 0, ity.NumMethod(); i < n; i++ {
		f := ity.Method(i)
		fmt.Println(f.Name, f.Type)
		// Compute func(*interface_test.MyInt, int) string
		// Double func(*interface_test.My
	}

	// ------------------------------------------------------------------------------------------------
	// 取得反射值能回傳相應的型別和數值：
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type()) // type: float64
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64) // kind is float64: true
	fmt.Println("value:", v.Float()) // value: 3.4

	// var x1 float64 = 3.4
	// v1 := reflect.ValueOf(x1)
	// v1.SetFloat(7.1)
	// fmt.Println("value v1:", v1) // error

	// 如果要修改相應的值，必須這樣寫： using Elem() with pointer
	var x2 float64 = 3.4
	p := reflect.ValueOf(&x2)
	v2 := p.Elem()
	v2.SetFloat(7.1)
	fmt.Println("value v2:", v2) // value v2: 7.1
}

/*
* 在前面提到的 TypeOf 回傳實例，是基於 Type Interface 定義進行實作，關於 Type Interface 主要的方法如下：
* 因此，透過 TypeOf 獲得中繼資料後，就可使用上面這些方法，
* 但是，在如果不是用 struct 呼叫 NumFiled ，則會 panics。
https://golang.org/src/reflect/type.go

type Type interface {
	Align() int
	FieldAlign() int
	Method(int) Method
	MethodByName(string) (Method, bool)
	NumMethod() int
	Name() string
	PkgPath() string
	Size() uintptr
	String() string
	Kind() Kind
	Implements(u Type) bool
	AssignableTo(u Type) bool
	ConvertibleTo(u Type) bool
	Comparable() bool
	Bits() int
	ChanDir() ChanDir
	IsVariadic() bool
	Elem() Type
	Field(i int) StructField
	FieldByIndex(index []int) StructField
	FieldByName(name string) (StructField, bool)
	FieldByNameFunc(match func(string) bool) (StructField, bool)
	In(i int) Type
	Key() Type
	Len() int
	NumField() int
	NumIn() int
	NumOut() int
	Out(i int) Type
	common() *rtype
	uncommon() *uncommonType
}
*/