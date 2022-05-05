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
	s := strconv.Itoa(10)
	t.Log("str" + s)
	if i, err := strconv.Atoi("15"); err == nil {
		t.Log(10 + i)
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


func TestStringFunction(t *testing.T) {
	// 根據「空白」，將字串拆成 slices：
	fields := strings.Fields("I am learning Go")
  fmt.Println(fields)  // [I am learning Go]
	fmt.Println(fields[2])
}


func TestStringFunction2(t *testing.T) {
	type fruits []string
	favorites := fruits{"appele", "banana", "Orange", "Grupe"}
	fmt.Println(favorites)
	fmt.Println(strings.Join(favorites, "-"))
}

// TrimLeft(s, cutset string) 會從 s 的左邊開始，依序判斷每一個字元是否被包含在 cutset 中，cutset 中可能不只包含一個字元，只要有在 cutset 中則會被 Trim 掉；
// TrimPrefix(s, prefix string) 一樣會從 s 的左邊開始，但是會把 prefix 當成一個完整的字串，而不是當成包含許多獨立字元的 cutset
func TestStringFunction3(t *testing.T) {
	auth := "123111123/1234"

	fmt.Println(strings.TrimLeft(auth, "123"))   // /1234
	fmt.Println(strings.TrimPrefix(auth, "123")) // 111123/1234

	auth = "123123/1234"

	fmt.Println(strings.TrimLeft(auth, "123"))   // /1234
	fmt.Println(strings.TrimPrefix(auth, "123")) // 123/1234
}

// String Interpolation, injection syntax
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
