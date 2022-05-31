package string_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A, B, C, D, E, F, G, H, I, J"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConvert(t *testing.T) {
	// strconv.Itoa : int -> string
	s := strconv.Itoa(10)
	t.Log("str" + s) // str10
	// strconv.Atoi : string -> int
	if i, err := strconv.Atoi("15"); err == nil {
		t.Log(10 + i) // 25
	}
}

func TestCreateString(t *testing.T) {
	str := `
    Hello
    World
    of
    Go
    `
	t.Log(str)

	var s strings.Builder
	s.WriteString("Hello")
	s.WriteString(" ")
	s.WriteString("World")
	fmt.Println(s.String())
}

// Fields: 根據「空白」，將字串拆成 slices：
func TestStringFunctionField(t *testing.T) {
	// 根據「空白」，將字串拆成 slices：
	fields := strings.Fields("I am learning Go")
  fmt.Println(fields)  // [I am learning Go]
	fmt.Println(fields[2]) // learning
}

// Join
/*
	a 是 slice string
	sep 指的是 separator，陣列連接時要用什麼字串連接
	func Join(a []string, sep string) string
*/
func TestStringFunctionJoin(t *testing.T) {
	type fruits []string
	favorites := fruits{"appele", "banana", "Orange", "Grupe"}
	fmt.Println(favorites) // [appele banana Orange Grupe]
	fmt.Println(strings.Join(favorites, "-")) // appele-banana-Orange-Grupe
}

type fruits []string

func (f fruits) toString() string {
	return strings.Join(f, ",")
}

func TestStringFunctionJoin2(t *testing.T) {
	favoriteFruits := fruits{"Apple", "Banana", "Orange", "Guava"}
  fmt.Println(favoriteFruits.toString()) // Apple,Banana,Orange,Guava
}

/*
	Split
	將字串轉成 slice string
	s 是字串
	sep 是 separator
	func Split(s, sep string) []string
*/

// ------------------------------------------------------------------------------------------------
// TrimLeft(s, cutset string) 會從 s 的左邊開始，依序判斷每一個字元是否被包含在 cutset 中，cutset 中可能不只包含一個字元，只要有在 cutset 中則會被 Trim 掉；
// TrimPrefix(s, prefix string) 一樣會從 s 的左邊開始，但是會把 prefix 當成一個完整的字串，而不是當成包含許多獨立字元的 cutset
func TestStringFunctionTrimLeftAndTrimPrefix(t *testing.T) {
	auth := "123111123/1234"

	fmt.Println(strings.TrimLeft(auth, "123"))   // /1234
	fmt.Println(strings.TrimPrefix(auth, "123")) // 111123/1234

	auth = "123123/1234"

	fmt.Println(strings.TrimLeft(auth, "123"))   // /1234
	fmt.Println(strings.TrimPrefix(auth, "123")) // 123/1234
}

// 字串插補（String Interpolation, injection syntax）
func TestStringFunction4(t *testing.T) {
	type person struct {
    firstName string
    lastName  string
	}
	var i int
	var f float64
	var b bool
	var s string
	var aaron person
	fmt.Printf("%v %v %v %q\n", i, f, b, s)     // 0 0 false ""
	fmt.Printf("%+v\n", aaron)                // {firstName: lastName:}

}
