package paniic_recover_test

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestPanicVsExits(t *testing.T) {
	fmt.Println("start")
	os.Exit(-1)
}


// 1. will execute defer()
// 2. outprint stack information
func TestPanicVsExits2(t *testing.T) {
	defer func() {
		fmt.Println("Finally")
	}()
	fmt.Println("start")
	panic(errors.New("Something Wrong"))
}

// recover() --< recover from panic
func TestPanicAndRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from", err)
		}
	}()
	fmt.Println("start")
	panic(errors.New("Something Wrong"))
}

