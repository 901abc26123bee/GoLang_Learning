package loop_test

import (
	"testing"
)
// [Go-For Range 性能研究](https://www.cnblogs.com/leeyongbard/p/10394820.html)

// -------------------------------------------- slice loop -------------------------------------------------
const N  = 10000
const text = "123456789abcdefghijklmnopqrstuvwxyz"

func ForSlice(s []string) {
	len := len(s)
	for i := 0; i < len; i++ {
			_, _ = i,s[i]
	}
}

func RangeForSlice(s []string) {
	for i, v := range s {
			_, _ = i, v
	}
}

func RangeForSliceWithoutValueCopy(s []string) {
	for i, _ := range s {
			_, _ = i, s[i]
	}
}

func initSlice() []string{
	s := make([]string,N)

	for i := 0;i < N;i++ {
			s[i] = text
	}
	return s
}

func BenchmarkForSlice(b *testing.B) {
	s := initSlice()

	b.ResetTimer()
	for i := 0;i < b.N;i++ {
			ForSlice(s)
	}
}

func BenchmarkRangeForSlice(b *testing.B) {
	s := initSlice()

	b.ResetTimer()
	for i := 0;i < b.N;i++  {
			RangeForSlice(s)
	}
}

func BenchmarkRangeForSliceWithoutValueCopy(b *testing.B) {
	s := initSlice()

	b.ResetTimer()
	for i := 0;i < b.N;i++  {
		RangeForSliceWithoutValueCopy(s)
	}
}

/*
$ go test -bench=. -benchmem
BenchmarkForSlice-4                               173174              8010 ns/op
BenchmarkRangeForSlice-4                          288541              3803 ns/op
BenchmarkRangeForSliceWithoutValueCopy-4          155236              8017 ns/op
*/

// 常規的 For 循環的性能更高。
// * 主要是因為 for range 是每次對循環元素的拷貝，而 for 循環，它獲取集合內元素是通過 s[i]，這種索引指針引用的方式，要比拷貝性能高得多
// * 既然是元素拷貝的問題，通過 _ 捨棄了元素的複制，然後通過 s[i] 方式獲取迭代的元素, 結果和常規的 for 循環一樣

// -------------------------------------------- map loop -------------------------------------------------
// 對於 map 來說，我們並不能使用 for i=0;i<N;i++ 的方式，大部分我們使用 for range 的方式：

func RangeForMap1(m map[int]string) {
	for k, v := range m{
			_,_ = k,v
	}
}

func RangeForMap2(m map[int]string) {
	for k, _ := range m{
			_,_ = k,m[k]
	}
}

// 初始化 map
func initMap() map[int]string  {
	m := make(map[int]string,N)

	for i := 0;i < N;i++ {
			m[i] = text
	}

	return m
}

func BenchmarkRangeForMap1(b *testing.B) {
	m := initMap()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
			RangeForMap1(m)
	}
}

func BenchmarkRangeForMap2(b *testing.B) {
	m := initMap()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
			RangeForMap2(m)
	}
}

/*
$ go test -bench=. -benchmem
BenchmarkForSlice-4                               181840              7583 ns/op
BenchmarkRangeForSlice-4                          337288              3644 ns/op
BenchmarkRangeForSliceWithoutValueCopy-4          179557              7093 ns/op
BenchmarkRangeForMap1-4                            10000            129494 ns/op
BenchmarkRangeForMap2-4                             6168            234277 ns/op
*/

// 相比較 slice，Map 遍歷的性能更差。(RangeForMap1)
// 使用上面優化遍歷 slice 的方式優化遍歷 map，減少值拷貝，(RangeForMap2)
// 優化後的結果性能明顯下降了，這和我們上面測試 slice 不一樣，這次沒有提升反而下降了

/*
For Range 原理

	遍歷 slice 前先是對要遍歷的 slice 做一個拷貝，然後獲取 slice 的長度作為循環次數，
	循環體中每次循環 會先獲取元素值，遍歷過程中每次迭代都會對 index 和 value 進行賦值，
	如果數據量比較大或 者 value 為 string 時，對 value 的賦值操作可能是多餘的，
	所以在上面我們使用 range 遍歷 slice 的時候，可以 忽略 value，使用 slice[index] 的方式提升性能

	遍歷 map 是通過 key 查找 value 的性能消耗, 高於賦值消耗(如果數據量比較大)，這就是為什麼優化沒有起到作用
*/