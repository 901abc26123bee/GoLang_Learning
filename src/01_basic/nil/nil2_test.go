package nil

import (
	"fmt"
	"testing"
	"unsafe"
)

// The Sizes of Nil Values With Types of Different Kinds May Be Different
// For the standard Go compiler, the sizes of two values of different types of the same kind whose zero values can be represented as the predeclared nil are always the same. For example, the sizes of all values of all different slice types are the same.
func TestNilSize(t *testing.T) {
	var p *struct{} = nil
	fmt.Println( unsafe.Sizeof( p ) ) // 8

	var s []int = nil
	fmt.Println( unsafe.Sizeof( s ) ) // 24

	var m map[int]bool = nil
	fmt.Println( unsafe.Sizeof( m ) ) // 8

	var c chan string = nil
	fmt.Println( unsafe.Sizeof( c ) ) // 8

	var f func() = nil
	fmt.Println( unsafe.Sizeof( f ) ) // 8

	var i interface{} = nil
	fmt.Println( unsafe.Sizeof( i ) ) // 16
}

func TestCompare(t *testing.T) {
	// 一些类型为不可比较类型的变量。
	var s []int
	var m map[int]int
	var f func()()

	_ = s == nil
	_ = m == nil
	_ = f == nil
	_ = 123 == interface{}(nil)
	_ = true == interface{}(nil)
	_ = "abc" == interface{}(nil)
	fmt.Println(s == nil)
	fmt.Println(m == nil)
	fmt.Println(f == nil)
	fmt.Println(123 == interface{}(nil))
	fmt.Println(true == interface{}(nil))
	fmt.Println("abc" == interface{}(nil))
}

// Two Nil Values of the Same Type May Be Not Comparable
/* 	In Go, map, slice and function types don't support comparison. 
		Comparing two values, including nil values, of an incomparable types is illegal.
		The following comparisons fail to compile.
		
		var _ = ([]int)(nil) == ([]int)(nil)
		var _ = (map[string]int)(nil) == (map[string]int)(nil)
		var _ = (func())(nil) == (func())(nil)
		
	但是，映射類型、切片類型和函數類型的任何值都可以和類型不確定的裸nil標識符比較。
		var _ = ([]int)(nil) == nil
		var _ = (map[string]int)(nil) == nil
		var _ = (func())(nil) == nil

*/