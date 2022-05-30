package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	row := 100
	col := 100
	x := make([][]int, row)
	for i := range x {
		x[i] = make([]int, col)
	}
	fillMatrix(x)
	caculate(x)
}
