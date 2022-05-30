package benchmark

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	benchmark
	go 可以使用 benchmark 來做一些效能測試，需要調整程式的效能時就會需要用到．
	要效能測試的 function 名稱是要 Benchmark 開頭

	go test -bench
下指令執行 benchmark 跑效能測試，在 go test 後加上 -bench=.，最後一個 . 是代表當前 package．
*/


func TestConcatStringByAdd(t *testing.T) {
	assert := assert.New(t)
	elems := []string{"1", "2", "3", "4", "5", "6"}
	ret := ""
	for _, elem := range elems {
		ret += elem
	}
	assert.Equal("123456", ret) // 0.034s
}

func TestConcatStringByBytesBuffer(t *testing.T) {
	assert := assert.New(t)
	var buf bytes.Buffer
	elems := []string{"1", "2", "3", "4", "5", "6"}
	for _, elem := range elems {
		buf.WriteString(elem)
	}
	assert.Equal("123456", buf.String()) // 0.012s
}

func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5", "6"}
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
	}
	b.StopTimer()
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		var buf bytes.Buffer
		elems := []string{"1", "2", "3", "4", "5", "6"}
		for _, elem := range elems {
			buf.WriteString(elem)
		}
	}
	b.StopTimer()
}

/*
		~/Desktop/go_learning/src/19_unit_test/benchmark$ go test -bench=.
		goos: darwin
		goarch: amd64
		pkg: go_learning/src/19_unit_test/benchmark
		cpu: Intel(R) Core(TM) i5-7267U CPU @ 3.10GHz
		BenchmarkConcatStringByAdd-4             7053465               170.7 ns/op
		BenchmarkConcatStringByBytesBuffer-4    15981808                69.78 ns/op
		PASS
		ok      go_learning/src/19_unit_test/benchmark  3.628s
*/
// for window : go test -bench="."

/*
		~/Desktop/go_learning/src/19_unit_test/benchmark$ go test -bench=. -benchmem
		goos: darwin
		goarch: amd64
		pkg: go_learning/src/19_unit_test/benchmark
		cpu: Intel(R) Core(TM) i5-7267U CPU @ 3.10GHz
		BenchmarkConcatStringByAdd-4             6226160               187.5 ns/op            24 B/op          5 allocs/op
		BenchmarkConcatStringByBytesBuffer-4    16007701                70.99 ns/op           64 B/op          1 allocs/op
		PASS
		ok      go_learning/src/19_unit_test/benchmark  3.496s

*/