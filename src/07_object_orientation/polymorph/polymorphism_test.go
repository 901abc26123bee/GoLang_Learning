package polymorph_test

import (
	"fmt"
	"math"
	"testing"
)

// Interfaces 試圖解決的問題
// 	1. 問題：共用相同邏輯但帶入不同型別參數的函式
// 		 在程式中的任何 type，只要這個 type 的函式有符合到該 interface 的定義，就可以歸類到該 interface 底下。

// ------------------------- polymorphism with struct/receiver function -------------------------------
type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {

}

func (g *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}

type JavaProgrammer struct {

}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.println(\"Hello World\")"
}

// since Programmer is an interface, writeFirstProgram can only accept pointer of an instance as parameter
// 因為 JavaProgrammer type & GoProgrammer type 的 WriteHelloWorld method 符合 Programmer interface 的規範
// 所以 JavaProgrammer & GoProgrammer 同樣屬於 Programmer interface
func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolygon(t *testing.T) {
	goProg := &GoProgrammer{}
	javaProg := new(JavaProgrammer)
	writeFirstProgram(goProg) // *polymorph_test.GoProgrammer  fmt.Println("Hello World")
	writeFirstProgram(javaProg) // *polymorph_test.JavaProgrammer  System.out.println("Hello World")
}

//------------------------------ polymorphism with -----------------------------------------
// STEP 1：定義 Bot type，它本質上是 interface
type Bot interface {
	// 在程式中的任何 type，只要是名稱為 getGreeting 而且回傳 string 的函式
	// 將自動升級變成 "Bot" 這個 type 的成員
	getGreeting() string
}

// STEP 2：宣告兩個 struct type
type EnglishBot struct{}
type SpanishBot struct{}

// STEP 3：此 receiver function 名稱為 getGreeting 且回傳 string，因此屬於 Bot type
func (EnglishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi There!"
}

// STEP 4：此 receiver function 名稱為 getGreeting 且回傳 string，因此屬於 Bot type
func (SpanishBot) getGreeting() string {
	return "Hola!"
}

// STEP 5：printGreeting 可以傳入 Bot interface
func printGreeting(b Bot) {
	fmt.Println(b.getGreeting())
}

func TestPolygon2(t *testing.T) {
	eb := EnglishBot{}
	sb := SpanishBot{}

	// STEP 6：現在 eb 和 sb 都算是 Bot type
	printGreeting(eb)
	printGreeting(sb)
}

// ---------------------------------- polymorphism with self defined type ------------------------------------------
// STEP 1: 函式名稱為 Abs 且回傳 float64 即屬於 Abser type
type Abser interface {
	Abs() float64
}

// STEP 2：定義一個 MyFloat type 且其 receiver function 符合 Abser interface 的規範
// MyFloat 會屬於 Abser type
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
			return float64(-f)
	}
	return float64(f)
}

// STEP 3：定義一個 Vertex type 且其 receiver function 符合 Abser interface 的規範
// Vertex 會屬於 Abser type
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func TestPolygon3(t *testing.T) {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f                // MyFloat 可以 implement Abser type
	fmt.Println(a.Abs()) // 1.4142135623730951

	a = &v               // *Vertex 可以 implement Abser type
	fmt.Println(a.Abs()) // 5

	// Cannot use 'v' (type Vertex) as type Abser.
	// Vertex 不能 implement Abser type，因為 Abs 這個方法有 pointer receiver
	// a = v
}

