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

/*
	檔案名稱一定要用 _test.go 當結尾
	func 名稱開頭要用 Benchmark
	for 循環內要放置要測試的程式碼
	b.N 是 go 語言內建提供的循環，根據一秒鐘的時間計算
	跟測試不同的是帶入 b *testing.B 參數


	$ $ go test -v -bench=. -run=none -benchmem .
	goos: darwin
	goarch: amd64
	BenchmarkPrintInt2String01-4    10000000               125 ns/op              16 B/op          2 allocs/op
	BenchmarkPrintInt2String02-4    30000000                37.8 ns/op             3 B/op          1 allocs/op
	BenchmarkPrintInt2String03-4    30000000                38.6 ns/op             3 B/op          1 allocs/op
	PASS

	基本的 benchmark 測試也是透過 go test 指令，不同的是要加上 -bench=.，這樣才會跑 benchmark 部分，
	否則預設只有跑測試程式， -4 代表目前的 CPU 核心數，也就是 GOMAXPROCS 的值，
	另外 -run 可以用在跑特定的測試函示，但是假設沒有指定 -run 時，你會看到預設跑測試 + benchmark，
	所以這邊補上 -run=none 的用意是不要跑任何測試，只有跑 benchmark，
	最後看看輸出結果，其中 10000000 代表一秒鐘可以跑 1000 萬次，每一次需要 140 ns，
	如果你想跑兩秒，請加上此參數在命令列 -benchtime=2s，但是個人覺得沒什麼意義。

	1 allocs/op 代表每次執行都需要搭配一個記憶體空間，而一個記憶體空間為 3 Bytes。
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