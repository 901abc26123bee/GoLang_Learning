package main

import (
	"errors"
	"fmt"
)

type operate func(x, y int) int

// 方案1。
func calculate(x int, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}

// 方案2。
type calculateFunc func(x int, y int) (int, error)

func genCalculator(op operate) calculateFunc {
	return func(x int, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

// func main() {
// 	// 方案1。
// 	x, y := 12, 23
// 	op := func(x, y int) int {
// 		return x + y
// 	}
// 	result, err := calculate(x, y, op)
// 	fmt.Printf("The result: %d (error: %v)\n",
// 		result, err)
// 	result, err = calculate(x, y, nil)
// 	fmt.Printf("The result: %d (error: %v)\n",
// 		result, err)

// 	// 方案2。
// 	// 把其他的函數作為結果返回
// 	x, y = 56, 78
// 	add := genCalculator(op)
// 	result, err = add(x, y)
// 	fmt.Printf("The result: %d (error: %v)\n",
// 		result, err)
// }

/*
The result: 35 (error: <nil>)
The result: 0 (error: invalid operation)
The result: 134 (error: <nil>)
*/

// ----------------------------------------------------------------------------

func modifyArray(a [3]string) [3]string {
	a[1] = "x"
	return a
}
// 數組是值類型，所以每一次復制都會拷貝它，以及它的所有元素值
func main() {
	array1 := [3]string{"a", "b", "c"}
	fmt.Printf("The array: %v\n", array1)
	array2 := modifyArray(array1)
	fmt.Printf("The modified array: %v\n", array2)
	fmt.Printf("The original array: %v\n", array1)
}
/*
The array: [a b c]
The modified array: [a x c]
The original array: [a b c]
*/
