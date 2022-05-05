package client_test

import (
	"go_learning/src/09_package/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSeries(5))
	t.Log(series.Square(8))
}
