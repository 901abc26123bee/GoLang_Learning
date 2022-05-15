package struct_test

import (
	"fmt"
	car "go_learning/src/07_object_orientation/car"
	"reflect"
	"testing"
	"time"
)

type Person struct {
	firstName string
	lastName  string
}
// 等同於
type Person2 struct {
	firstName, lastName string
}

type rectangle struct {
	color string
}

func TestStructDefinition(t *testing.T) {
	// anonymous struct
	foo := struct {
		language string
	} {
		language: "golang",
	}

	// 三種宣告 Person struct 的方式：
	// 使用 new syntax：第二種和第三種寫法是一樣的
	var user1 *Person       // nil
	user2 := &Person{}      // {}，user2.firstName 會是 ""
	user3 := new(Person)    // {}，user3.firstName 會是 ""

	fmt.Println(foo) // {golang}
	fmt.Println(user1) // <nil>
	fmt.Println(user2) // &{}
	fmt.Println(user3) // &{}
	fmt.Println(*user3) // &{}
	fmt.Println(&user3) // 0xc000114038
}

func TestStructCreation(t *testing.T) {
	// 方法一：根據資料輸入的順序決定誰是 firstName 和 lastName
  alex := Person{"Alex", "Anderson"}

  // 直接取得 struct 的 pointer
  alex2 := &Person{"Alex", "Anderson"}

	// 方法二（建議）
	alex3 := Person{
			firstName: "Alex",
			lastName:  "Anderson",
	}

  // 方法三：先宣告再賦值
  var alex4 Person
	alex4.firstName = "Alex"
	alex4.lastName = "Anderson"

	// when printing structs, the plus flag (%+v) adds field names
	fmt.Printf("%+v\n", alex)   // {firstName:Alex lastName:Anderson}
	fmt.Println(alex)  // {Alex Anderson}
	fmt.Println(alex2) // &{Alex Anderson}
	fmt.Println(alex3) // {Alex Anderson}
	fmt.Println(alex4) // {Alex Anderson}
}

//----------------------------- struct pointer -----------------------------------
/*
	當 pointer 指稱到的是 struct 時，可以直接使用這個 pointer 來對該 struct 進行設值和取值。
	在 golang 中可以直接使用 pointer 來修改 struct 中的欄位。
	一般來說，若想要透過 struct pointer（&v）來修改該 struct 中的屬性，
	需要先解出其值（*p）後使用 (*p).X = 10，但這樣做太麻煩了，
	因此在 golang 中允許開發者直接使用 p.X 的方式來修改：
*/

func TestStructModByPointer(t *testing.T) {
	p := &Person{
		firstName: "Alex",
		lastName:  "Anderson",
	}
	//  golang 中允許開發者直接使用 p.X 的方式來修改：
	p.firstName = "Bill"
	// 一般來說，需要先解出其值（*p）後使用 (*p).X = 10，但這樣做太麻煩了，
	(*p).lastName = "Xeon"
	fmt.Printf("%+v\n", p) // &{firstName:Bill lastName:Xeon}
	fmt.Printf("%+v\n", *p) // {firstName:Bill lastName:Xeon}


	// 另外，使用 struct pointer 時才可以修改到原本的物件，否則會複製一個新的：
	r1 := rectangle{"Green"}

	// 複製新的，指稱到不同位置 address
	r2 := r1
	r2.color = "Pink"
	fmt.Println(r2) // Pink
	fmt.Println(r1) // Green

	// 指稱到相同位置 address
	r3 := &r1
	r3.color = "Red"
	fmt.Println(r3) // Red
	fmt.Println(r1) // Red
}

//----------------------------- nested struct -----------------------------------

// 在 struct 內關聯另一個 struct（nested struct)
// STEP 1：定義外層 struct
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

// STEP 2：定義內層 struct
type contactInfo struct {
	email   string
	zipCode int
}

func TestNestedStruc(t *testing.T) {
	// STEP 3：建立變數
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
				email:   "jim@gmail.com",
				zipCode: 94000,
		},
	}

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
	}

	// STEP 4：印出變數
	fmt.Printf("%+v\n", jim) // {firstName:Jim lastName:Party contact:{email:jim@gmail.com zipCode:94000}}
	fmt.Println(jim)         // {Jim Party {jim@gmail.com 94000}}

	fmt.Printf("%+v\n", alex) // {firstName:Alex lastName:Anderson contact:{email: zipCode:0}}
	fmt.Println(alex)         // {Alex Anderson { 0}}
}

//--------------------------- Struct field Tag(meta-data) -------------------------------------

// Struct field Tag(meta-data)
/*
		struct field tag 會在 struct 的 value 後面使用 backtick 來表示，例如 json:"name"：

		在 field tag 中還能帶入其他關鍵字，舉例來說：
		omitempty：指的是該欄位沒值的話，就不要顯示欄位名稱
		-：表示忽略掉該欄位，Marshal 時該欄位不會出現在 JSON 中，Unmarshal 時該欄位也不會被處理
*/
type User struct {
	Name          string    `json:"name"`
	Password      string    `json:"-"`
	PreferredFish []string  `json:"preferredFish,omitempty"`
	CreatedAt     time.Time `json:"createdAt"`
}

// Anonymous fields
// 在 struct 中不一定要替欄位建立名稱，而是可以直接使用 data types，而 Go 會使用這個 data type 當作欄位名稱：
type AnonymousField struct {
	string // 相似於 string string
	bool   // 相似於 bool bool
	int    // 相似於 int int
}

func TestAnonymousField(t *testing.T) {
	anonymousField := AnonymousField{
		"person", true, 22,
	}
	fmt.Printf("%+v\n", anonymousField) // {string:person bool:true int:22}
}

// Function Fields
// struct 中的 field 也可以是 function
type GetDisplayNameType func(string, string) string

type Person3 struct {
    FirstName, LastName string
    GetDisplayName      GetDisplayNameType
}

func TestStructWithFuncField(t *testing.T) {
	p := Person3{
		FirstName: "Aaron",
		LastName:  "Chen",
		GetDisplayName: func(firstName, lastName string) string {
					return firstName + "-" + lastName
			},
	}

	displayName := p.GetDisplayName(p.FirstName, p.LastName)
	fmt.Println(displayName) // Aaron-Chen
}

// Promoted fields / Anonymous / Embedded

// Struct is a sequence of fields each with name and type.
// 定義 Promoted fields 的 struct
// 在 Golang 中 struct 的 fields name 可以省略，沒有 field name 的 name 被稱作 anonymous 或 embedded。
// 在這種情況下，會直接使用 「Type 的名稱」來當作 field name：
// ⚠️ 如果 nested anonymous struct 中的欄位和其 parent struct 的欄位名稱有衝突時，則該欄位不會被 promoted。
// https://medium.com/golangspec/promoted-fields-and-methods-in-go-4e8d7aefb3e3
type Person4 struct {
	name string
	age  int32
}

func (p Person4) IsAdult() bool {
	return p.age >= 18
}

type Employee struct {
	position string
}

func (e Employee) IsManager() bool {
	return e.position == "manager"
}

// 直接使用 「Type 的名稱」來當作 field name：Promoted fields
type Record struct {
	Person4
	Employee
}

func TestStructPromotedFields(t *testing.T) {
	// 明確的定義 embedded 的結構
	record2 := Record{
		Person4 {
			name: "Lisa",
			age: 18,
		},
		Employee {
			position: "manager",
		},
	}
	fmt.Printf("%+v\n", record2) // {Person4:{name:Lisa age:18} Employee:{position:manager}}

	record3 := Record{
		Person4 {
			"Lisa",
			18,
		},
		Employee {
			"manager",
		},
	}
	fmt.Printf("%+v\n", record3) // {Person4:{name:Lisa age:18} Employee:{position:manager}}
}

// 在 Promoted fields 中設值

// 對於 anonymous (embedded) fields 的欄位（field）或方法（method）稱作 prompted，
// 它們就像一般的欄位一樣，但是不能跳過 Type 的名稱直接用 struct literals 的方式來賦值：
// 錯誤用法：不能在未明確定義 promoted fields 名稱的情況下，使用 struct literals 設值
func TestPromotedFieldWithStructLiteralsAndSeting(t *testing.T) {
	// 跳過 Type 的名稱直接用 struct literals 的方式來賦值 --> unknown field xxx in struct literal
	// record := Record{
	// 		name:     "record",
	// 		age:      29,
	// 		position: "software engineer",
	// }

	// fmt.Printf("%+v", record)


	record := Record{}
	record.name = "record"
	record.age = 29
	record.position = "software engineer"

	fmt.Printf("%+v\n", record) // {Person4:{name:record age:29} Employee:{position:software engineer}}
}

// 在 Promoted fields 中取值
// 不論有沒有使用明確的 promoted fields 名稱，都可以取值：
func TestAccessFieldsInPromotedFields(t *testing.T) {
	record := Record{
		Person4: Person4{
				name: "Bill",
				age:  29,
		},
		Employee: Employee{
				position: "software engineer",
		},
	}

	// 不論有沒有使用明確的 promoted fields 名稱，都可以取值
	fmt.Println("name", record.name)             // name Bill
	fmt.Println("Person.age", record.Person4.age) // Person.age 29

	fmt.Println("position", record.position)                   // position software engineer
	fmt.Println("Employee.position", record.Employee.position) // Employee.position software engineer
}


// Person 有 Name 且可以 Introduce，而 Saiyan 是 Person，因此它也有 Name 且可以 Introduce：

// STEP 1：建立 Person struct 與其 Method
type Person5 struct {
	Name string
}

// Methods | Function Receiver
// 因為 Go 本身並不是物件導向程式語言（object-oriented programming language），
// 所以只能用 Type 搭配在函式中使用 receiver 參數來實作出類似物件程式語言的功能：
func (p *Person5) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

// STEP 2：建立 Saiyan struct，並將 Person embed 在內
// 意思是 Saiyan 是 Person，而不是 Saiyan「有一個」Person
type Saiyan struct {
	*Person5
	Power int
}

func TestAccessFieldsInStruct(t *testing.T) {
	// STEP 3：建立 goku
	goku := Saiyan{
		Person5: &Person5{"Goku"},
		Power:  9001,
	}

	// STEP 4：可以直接使用 goku.Name，也可以使用 goku.Person.Name
	fmt.Println(goku.Name)        // Goku
	fmt.Println(goku.Person5.Name) // Goku

	// STEP 5：方法在使用時也一樣
	goku.Introduce()        // Hi, I'm Goku
	goku.Person5.Introduce() // Hi, I'm Goku
}

// ---------------------------- Interface Fields (Nested interface) ------------------------------------
/*
		Interface Fields (Nested interface)
		struct 中的欄位也可以是 interface，以 Employee 這個 struct 來說，
		其中的 salary 欄位其型別是 Salaried 這個 interface，
		也就是是說 salary 這個欄位的值，一定要有實作出 Salaried 的方法，
		如此 salary 才會符合該 interface 的 type：

		- struct 中的 { salary Salaried } 表示 salary 要符合 Salaried interface type
		- 要符合該 interface type，表示 salary 要實作 Salaried interface 中所定義的 method signatures
		- 在定義 ross 變數時，因為 Salary 這個 struct 已經實作了 Salaried，因此可以放到 salary 這個欄位中
*/

type Salaried interface {
	getSalary() int
}

type Salary struct {
	basic, insurance, allowance int
}

// Salary 實作了 getSalary() 的方法，因此可以算是 Salaried type（polymorphism）
func (s Salary) getSalary() int {
	return s.basic + s. insurance + s. allowance
}

type Employee6 struct {
	firstName, lastName string
	salary Salaried // 只要 salary 實作了 Salaried ==> Salaried interface type
}

func TestStructWithInterfaceType(t *testing.T) {
	ross := Employee6{
		firstName: "Ross",
		lastName: "Geller",
		// 因為 Salary struct 已經實作了 Salaried，因此可以當作 salary 的欄位值
		salary: Salary {
			11000, 500, 400,
		},
	}
	fmt.Println("Ross's salary is", ross.salary.getSalary()) // Ross's salary is 11900
}

/*
		anonymously nested interface
		同樣的，當該 struct 的欄位沒有填寫時（anonymous fields），interface 中所定義的方法也可以被 promoted：

		- 在定義 Employee struct 時使用了 Salaried 作為 anonymous field
		- 在對 Employee 時，因為 Salary struct 有實作 Salaried，因此可以當作 Employee struct 中 Salaried 的值
		- 由於 promoted 這作用，可以直接使用 ross.getSalary() 方法，而不需要使用 ross.Salaried.getSalary()

*/

type Salaried2 interface {
	getSalary2() int
}

type Salary2 struct {
	basic, insurance, allowance int
}

// Salary2 實作了 getSalary2() 的方法，因此可以算是 Salaried2 type（polymorphism）
func (s Salary2) getSalary2() int {
	return s.basic + s.insurance + s.allowance
}

type Employee7 struct {
	firstName, lastName string
	Salaried2
}

func TestStructWithInterfaceTypePromoted(t *testing.T) {
	ross := Employee7{
			firstName: "Ross",
			lastName:  "Geller",
			// 因為 Salary 實作了 Salaried，因此可以作為 Salaried 的欄位值
			Salaried2: Salary2{
					1000, 50, 50,
			},
	}

	// 由於 method 會被 promoted，因此可以直接呼叫 ross.getSalary() 的方法
	// 而不需要使用 ross.Salaried.getSalary()
	fmt.Println("Ross's salary is", ross.getSalary2()) // Ross's salary is 1100
}

// ----------------------------------------------------------------
// 匯出的欄位（Exported fields）
// 如同 package 中的變數一樣，struct 中的欄位只有在 `欄位名稱` 以大寫命名時才會 export 出去，其他 package 中才能取用得到：

// 如果想要使用 car package 中的 Car Type 時，主要沒有想要對 unexported field 做事，則不會報錯，
// 在沒有 exported 出來的 fields 則會取得 zero value：


// 錯誤發生！price 並沒有 export 出來被使用
// unknown field 'Price' in struct literal of type Car (but does have price)
func TestExportStruct(t *testing.T) {
	c := car.Car{
		Name:  "Toyota",
		// price: 1000,
	}

	// 如果想要使用 car package 中的 Car Type 時，主要沒有想要對 unexported field 做事，
	// 則不會報錯，在沒有 exported 出來的 fields 則會取得 zero value：
	fmt.Println(c) // {Toyota 0 0}

	totota := car.Toyota{
		Name:  "Toyota",
		Price: 1000,
		Amount: 20,
	}
	fmt.Println(totota) // {Toyota 1000 20}
}

// ------------------------------ struct 的比較（Struct comparison） ----------------------------------

// struct 的比較（Struct comparison）
// 當兩個 struct 的 type 和 field value 都相同時，兩個 struct 可以被視為相同：
func TestStructComparison(t *testing.T) {
	p := Person{
		firstName: "Aaron",
		lastName:  "Chen",
	}

	a := Person{
		firstName: "Aaron",
		lastName:  "Chen",
	}

	fmt.Println(p == a)  // true
}

// 但若在 struct 中有 field 的 type 是無法比較的話（例如，map），那麼這兩個 struct 將無法進行比較：
// 會跳出錯誤訊息：
// invalid operation: p == a (struct containing map[string]int cannot be compared)
type Person7 struct {
	FirstName, LastName string
	leaves              map[string]int
}

// ---------------------------- 辨認 Struct Type 的名稱 ------------------------------------
// 使用 reflect.TypeOf 和 reflect.ValueOf().Kind() 可以用來判斷該 struct 的 struct type 名稱，以及變數的實際 type：

func TestCheckStructTypeByReflections(t *testing.T) {
	u := User{
		Name:     "Sammy the Shark",
		Password: "fisharegreat",
	}

	fmt.Println(u) // {Sammy the Shark fisharegreat [] 0001-01-01 00:00:00 +0000 UTC}
	fmt.Println(reflect.TypeOf(u))         // struct_test.User
	fmt.Println(reflect.ValueOf(u).Kind()) // struct

	up := &User{
			Name:      "Sammy the Shark",
			Password:  "fisharegreat",
			CreatedAt: time.Now(),
	}

	fmt.Println(*up) // {Sammy the Shark fisharegreat [] 2022-05-15 00:16:17.882858 +0800 CST m=+0.000851678}
	fmt.Println(reflect.TypeOf(up))         // *struct_test.User
	fmt.Println(reflect.ValueOf(up).Kind()) // ptr
}