package ioc

import (
	"fmt"
	"testing"
)

// [GO编程模式：委托和反转控制](https://coolshell.cn/articles/21214.html)

type Widget struct {
	X, Y int
}

type Label struct {
	Widget        // Embedding (delegation)
	Text   string // Aggregation
}

type Button struct {
	Label // Embedding (delegation)
}

type ListBox struct {
	Widget          // Embedding (delegation)
	Texts  []string // Aggregation
	Index  int      // Aggregation
}

// ----------------------------------------------------------------------------
// 對於 Lable 來說，只有 Painter ，沒有Clicker
// 對於 Button 和 ListBox來說，Painter 和Clicker都有。
type Painter interface {
	Paint()
}

type Clicker interface {
	Click()
}

// ----------------------------------------------------------------------------
func (label Label) Paint() {
  fmt.Printf("%p:Label.Paint(%q)\n", &label, label.Text)
}

//因為這個接口可以通過 Label 的嵌入帶到新的結構體，
//所以，可以在 Button 中可以重載這個接口方法以
func (button Button) Paint() { // Override Label.Paint()
    fmt.Printf("Button.Paint(%s)\n", button.Text)
}
func (button Button) Click() {
    fmt.Printf("Button.Click(%s)\n", button.Text)
}
// 重點提示一下，Button.Paint() 接口可以通過 Label 的嵌入帶到新的結構體，
// 如果 Button.Paint() 不實現的話，會調用 Label.Paint() ，
// 所以，在 Button 中聲明 Paint() 方法，相當於Override。


func (listBox ListBox) Paint() {
    fmt.Printf("ListBox.Paint(%q)\n", listBox.Texts)
}
func (listBox ListBox) Click() {
    fmt.Printf("ListBox.Click(%q)\n", listBox.Texts)
}

// ----------------------------------------------------------------------------
func TestOverridee(t *testing.T) {
	// 如果在 Label 結構體裡出現了重名，就需要解決重名，
	// 例如，如果 成員 X 重名，用 label.X表明 是自己的X ，用  label.Wedget.X 表示嵌入過來的。
	label := Label{Widget{10, 10}, "State:"}
	label.X = 11
	label.Y = 12
	fmt.Printf("%+v\n", label) // {Widget:{X:11 Y:12} Text:State:}

	button1 := Button{Label{Widget{10, 70}, "OK"}}
	button2 := Button{Label{Widget{10, 70}, "Cancel"}}

	listBox := ListBox{Widget{10, 40}, []string{"AL", "AK", "AZ", "AR"}, 0}

	for _, painter := range []Painter{label, listBox, button1, button2} {
		painter.Paint()
	}
	/*
		0xc0000900e0:Label.Paint("State:")
		ListBox.Paint(["AL" "AK" "AZ" "AR"])
		Button.Paint(OK)
		Button.Paint(Cancel)
	*/

	for _, widget := range []interface{}{label, listBox, button1, button2} {
		widget.(Painter).Paint()
		if clicker, ok := widget.(Clicker); ok {
			clicker.Click()
		}
		fmt.Println() // print a empty line
	}
	/*
	0xc000026160:Label.Paint("State:")

	ListBox.Paint(["AL" "AK" "AZ" "AR"])
	ListBox.Click(["AL" "AK" "AZ" "AR"])

	Button.Paint(OK)
	Button.Click(OK)

	Button.Paint(Cancel)
	Button.Click(Cancel)
	*/
}