package err_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

var LessThenTwoErrors = errors.New("n should be not less than 2")
var LargerThanHundredErrors = errors.New("n should be not larger than 100")
func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThenTwoErrors
	}
	if n > 100 {
		return nil, LargerThanHundredErrors
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i - 2] + fibList[i - 1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(10); err != nil {
		if err == LessThenTwoErrors {
			fmt.Println("It is less")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}

// nestedError
func GetFibonacci_1(str string) {
	var (
		i int
		err error
		list []int
	)
	if i, err = strconv.Atoi(str); err == nil {
		if list, err = GetFibonacci((i)); err == nil {
			fmt.Println(list)
		} else {
			fmt.Println("Error", err)
		}
	} else {
		fmt.Println("Error",err)
	}
}

// more clearly --> using err to avoid nestedError
func GetFibonacci_2(str string) {
	var (
		i int
		err error
		list []int
	)
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
		return
	}
	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(list)
}