package func_test

import (
	"fmt"
	"testing"
)

// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh-tw/02.3.md
// 函式作為值、型別
// 在 Go 中函式也是一種變數，我們可以透過 type 來定義它，它的型別就是所有擁有相同的參數，相同的回傳值的一種型別
// type typeName func(input1 inputType1 , input2 inputType2 [, ...]) (result1 resultType1 [, ...])

type testInt func(int) bool

func isOdd(integer int) bool {
	if integer % 2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer % 2 == 1 {
		return false
	}
	return true
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func TestFuncAsParams(t *testing.T){
	slice := []int {1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice) // slice =  [1 2 3 4 5 7]
	odd := filter(slice, isOdd)    // 函式當做值來傳遞了
	fmt.Println("Odd elements of slice are: ", odd) // Odd elements of slice are:  [1 3 5 7]
	even := filter(slice, isEven)  // 函式當做值來傳遞了
	fmt.Println("Even elements of slice are: ", even) // Even elements of slice are:  [2 4]
}