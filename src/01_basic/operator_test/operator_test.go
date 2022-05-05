package operator_test

import (
	"fmt"
	"testing"
)

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 5, 3, 4}
	d := [...]int{1, 2, 3, 4}
	// f := [...]int{1, 2, 3, 4, 5, 6}

	t.Log(a == b)
	t.Log(a == d)
	// t.Log(a == f)
}

//  AND(&), OR(|), XOR(^) and NOT(~)
// bit clear (AND NOT)
// 解釋：
// ^ --> XOR - 異或(exclusive or) ：相同為0，不同為1。也可用「不進位加法」來理解。
// 如果運算符右側數值的第 i 位為 1，那麼計算結果中的第 i 位為 0；
// 如果運算符右側數值的第 i 位為 0，那麼計算結果中的第 i 位為運算符左側數值的第 i 位的值。
func TestBitClear(t *testing.T)() {
	x := 11
	y := (1 << 0) | (1 << 3) //保证 z 中的第 0 位和第 3 位为 0
	z := x &^ y // x 和 y取反後的值進行與操作
	fmt.Printf("x = %b\n", x)
	fmt.Println("\t\t&^")
	fmt.Printf("y = %b\n", y)
	fmt.Println("————————")
	fmt.Printf("z = %04b\n", z)
}

func TestBitClear1(t *testing.T) {
	a := 7
	a = a &^ Readable
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}