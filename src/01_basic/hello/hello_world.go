package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("Hello, hello!", os.Args[1])
	}

	os.Exit(0)
	// os.Exit(0) // exit the program
}
