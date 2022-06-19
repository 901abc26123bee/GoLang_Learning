package interface_test

import (
	"fmt"
	"testing"
)
// *(1) interface的賦值問題
/*
	在golang中對多態的特點體現從語法上並不是很明顯。 我們知道發生多態的幾個要素：
		1、有interface接口，並且有接口定義的方法。
		2、有子類去重寫interface的接口。
	*	3、有父類指針指向子類的具體對象
	滿足上述3個條件，就可以產生多態效果，就是，父類指針可以調用子類的具體方法
*/
type People interface {
	Speak(string) string
}

type Student struct {}

func (stu *Student) Speak(think string) (talk string) {
	if think == "love" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func TestInterfaceAssignment(t *testing.T) {
	var peo People = &Student{}
	think := "love"
	// Student{}已經重寫了父類People{}中的Speak(string) string方法，那麼只需要用父類指針指向子類對象即可。
	// （People为interface类型，就是指针类型）
	fmt.Println(peo.Speak(think)) // You are a good boy
}

// ----------------------------------------------------------------------------

//* (2) interface的內部構造(非空接口iface情況)
// People擁有一個Show方法的，屬於非空接口，People的內部定義應該是一個iface結構體
type people interface {
	Show()
}

type student struct {}



func (stu *student) Show() () {}

func live() people {
	var stu *student
	return stu
}

func TestNotEmptyInterface2(t *testing.T) {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
	// BBBBBBB
}
/*
	stu是一個指向nil的空指針，但是最後return stu 會觸發匿名變量 People = stu值拷貝動作，
	所以最後live()放回給上層的是一個People insterface{}類型，也就是一個iface struct{}類型。
*	stu為nil，只是iface中的data 為nil而已。但是iface struct{}本身並不為nil.
*/


// ----------------------------------------------------------------------------

// *(3) interface內部構造(空接口eface情況)
// Foo()的形參x interface{}是一個空接口類型eface struct{}。
// 所以 x 結構體本身不為nil，而是data指針指向的p為nil。
func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
func TestEmptyInterface4(t *testing.T) {
	var p *int = nil
	Foo(p) // non-empty interface
}

// ----------------------------------------------------------------------------

// *(4) inteface{}與*interface{}
type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}

type point *interface {
}

func TestInterface(t *testing.T) {
	s := S{}
	p := &s
	f(s)
	f(p)
	// g(s) // cannot use s (type S) as type *interface {} in argument to g:
	// g(p) // cannot use p (type *S) as type *interface {} in argument to g:
	var ps point
	g(ps) //* 函數 func g(x *interface{}) 只能接受 *interface{}

	/*
	Golang是強類型語言，
	* interface是所有golang類型的父類
	* 函數中 func f(x interface{}) 的 interface{} 可以支持傳入golang的任何類型，包括指針，
	* 但是函數 func g(x *interface{}) 只能接受 *interface{}
	*/
}