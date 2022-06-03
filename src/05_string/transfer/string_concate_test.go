package stringbyte_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)
/*
	*	 + 連接適用於短小的、常量字符串（明確的，非變量），因為編譯器會給我們優化。
	*	 Join是比較統一的拼接，不太靈活
	*	 fmt和buffer基本上不推薦
	*	 builder從性能和靈活性上，都是上佳的選擇。
*/
// [Go语言字符串高效拼接（二](https://www.flysnow.org/2018/11/05/golang-concat-strings-performance-analysis.html)
// ----------------------------------------------------------------
// 字符串拼接 shorter one

// +號拼接, String Concat
// golang 裡面的字符串都是不可變的，每次運算都會產生一個新的字符串，所以會產生很多臨時的無用的字符串，會給 gc 帶來額外的負擔
func StringPlusOld() string{
	var s string
	s+="Name"+":"+"Kenny"+"\n"
	s+="Age"+":"+"55"+"\n"
	s+="Email"+":"+"kenny@company.com"
	return s
}

// go test -bench=. -benchmem
func BenchmarkStringPlusOld(b *testing.B) {
	for i:=0;i<b.N;i++{
		StringPlusOld()
	}
}

// fmt 拼接
// Sprint系列函數會把傳入的數據生成並返回一個字符串。
func StringFmtOld() string{
	return fmt.Sprint("Name"+":"+"Kenny"+"\n","Age"+":"+"55"+"\n","Email"+":"+"kenny@company.com")
}

func BenchmarkStringFmtOld(b *testing.B) {
	for i:=0;i<b.N;i++{
		StringFmtOld()
	}
}

// Join 拼接
// strings.Join函數進行拼接，接受一個字符串數組，轉換為一個拼接好的字符串
// * join會先根據字符串數組的內容，計算出一個拼接之後的長度，然後申請對應大小的內存，一個一個字符串填入，
// * 在已有一個數組的情況下，這種效率會很高，但是本來沒有，去構造這個數據的代價也不小
func StringJoinOld() string{
	s:=[]string{"Name"+":"+"Kenny"+"\n","Age"+":"+"55"+"\n","Email"+":"+"kenny@company.com"}
	return strings.Join(s,"")
}


func BenchmarkStringJoinOld(b *testing.B) {
	for i:=0;i<b.N;i++{
		StringJoinOld()
	}
}

// buffer 拼接
// 使用的是bytes.Buffer進行的字符串拼接，它是非常靈活的一個結構體，
// 不止可以拼接字符串，還是可以byte,rune等，並且實現了io.Writer接口，寫入也非常方便。
// 可以當成可變字符使用，對內存的增長也有優化，如果能預估字符串的長度，還可以用 buffer.Grow() 接口來設置 capacity
func StringBufferOld() string {
	var b bytes.Buffer
	b.WriteString("Name")
	b.WriteString(":")
	b.WriteString("Kenny")
	b.WriteString("\n")
	b.WriteString("Age")
	b.WriteString(":")
	b.WriteString("55")
	b.WriteString("\n")
	b.WriteString("Email")
	b.WriteString(":")
	b.WriteString("kenny@company.com")
	return b.String()
}

func BenchmarkStringBufferOld(b *testing.B) {
	for i:=0;i<b.N;i++{
		StringBufferOld()
	}
}

// builder 拼接
// 為了改進buffer拼接的性能，從go 1.10 版本開始，增加了一個builder類型，用於提升字符串拼接的性能。它的使用和buffer幾乎一樣。
/*
	* Builder要比Buffer性能好很多，這個問題原因主要還是[]byte和string之間的轉換，Builder恰恰解決了這個問題。
		func (b *Builder) String() string {
			return *(*string)(unsafe.Pointer(&b.buf))
		}
*/
func StringBuilderOld() string {
	var b strings.Builder
	b.WriteString("Name")
	b.WriteString(":")
	b.WriteString("Kenny")
	b.WriteString("\n")
	b.WriteString("Age")
	b.WriteString(":")
	b.WriteString("55")
	b.WriteString("\n")
	b.WriteString("Email")
	b.WriteString(":")
	b.WriteString("kenny@company.com")
	return b.String()
}

func BenchmarkStringBuilderOld(b *testing.B) {
	for i:=0;i<b.N;i++{
		StringBuilderOld()
	}
}

// Bytes Append
func TestBytesAppend(t *testing.T) {
	var b []byte
	s := "test-string"
	b = append(b, s...)
	str := string(b)
	fmt.Print(str)
}

// String Copy
func TestStringCopy(t *testing.T) {
	ts := "test-string"
	n := 5
	tsl := len(ts) * n
	bs := make([]byte, tsl)
	bl := 0

	for bl < tsl {
			bl += copy(bs[bl:], ts)
	}

	str := string(bs)
	fmt.Print(str)
}

/*
$ go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: go_learning/src/05_string/transfer
cpu: Intel(R) Core(TM) i5-7267U CPU @ 3.10GHz

BenchmarkStringPlus-4           11657326                95.31 ns/op           72 B/op          2 allocs/op
BenchmarkStringFmt-4             9536484               126.7 ns/op            48 B/op          1 allocs/op 		// 操作內存分配只有1次，分配48字節
BenchmarkStringJoin-4           20595765                55.66 ns/op           48 B/op          1 allocs/op
BenchmarkStringBuffer-4          9295358               119.3 ns/op           112 B/op          2 allocs/op
BenchmarkStringBuilder-4         8405898               137.5 ns/op           120 B/op          4 allocs/op
PASS

*/

// ------------------------------------------------------------------------------------------------
func StringPlus(p []string) string{
	var s string
	l:=len(p)
	for i:=0;i<l;i++{
		s+=p[i]
	}
	return s
}

func StringFmt(p []interface{}) string{
	return fmt.Sprint(p...)
}

func StringJoin(p []string) string{
	return strings.Join(p,"")
}

func StringBuffer(p []string) string {
	var b bytes.Buffer
	l:=len(p)
	for i:=0;i<l;i++{
		b.WriteString(p[i])
	}
	return b.String()
}

func StringBuilder(p []string) string {
	var b strings.Builder
	l:=len(p)
	for i:=0;i<l;i++{
		b.WriteString(p[i])
	}
	return b.String()
}

const BLOG  = "0123456789abcdefghijklmnopqrstuvwxyz"

func initStrings(N int) []string{
	s:=make([]string,N)
	for i:=0;i<N;i++{
		s[i]=BLOG
	}
	return s;
}

func initStringi(N int) []interface{}{
	s:=make([]interface{},N)
	for i:=0;i<N;i++{
		s[i]=BLOG
	}
	return s;
}

// 10000个字符串
func BenchmarkStringPlus10000(b *testing.B) {
	p:= initStrings(10000)
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		StringPlus(p)
	}
}

func BenchmarkStringFmt10000(b *testing.B) {
	p:= initStringi(10000)
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		StringFmt(p)
	}
}

func BenchmarkStringJoin10000(b *testing.B) {
	p:= initStrings(10000)
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		StringJoin(p)
	}
}


func BenchmarkStringBuffer10000(b *testing.B) {
	p:= initStrings(10000)
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		StringBuffer(p)
	}
}

func BenchmarkStringBuilder10000(b *testing.B) {
	p:= initStrings(10000)
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		StringBuilder(p)
	}
}

const (
	sss = "hello world!"
	cnt = 10000
)

var expected = strings.Repeat(sss, cnt)

// need to make array first before using join
func BenchmarkStringJoinWithoutArray(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
			var str string
			for i := 0; i < 10000; i++ {
					str = strings.Join([]string{str, sss}, "")
			}
			result = str
	}
	b.StopTimer()
	if result != expected {
			b.Errorf("unexpected result; got=%s, want=%s", string(result), expected)
	}
}


/*
$ go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: go_learning/src/05_string/transfer
cpu: Intel(R) Core(TM) i5-7267U CPU @ 3.10GHz

test func 	 	 						excuted times per seconds
BenchmarkStringPlusOld-4                 9504038               105.5 ns/op            72 B/op          2 allocs/op
BenchmarkStringFmtOld-4                  9360343               116.0 ns/op            48 B/op          1 allocs/op
BenchmarkStringJoinOld-4                19819828                55.01 ns/op           48 B/op          1 allocs/op
BenchmarkStringBufferOld-4               9703952               151.5 ns/op           112 B/op          2 allocs/op
BenchmarkStringBuilderOld-4              7456269               134.4 ns/op           120 B/op          4 allocs/op

BenchmarkStringPlus10000-4                     5         214465855 ns/op        1838244611 B/op    10053 allocs/op
BenchmarkStringFmt10000-4                   2212            494620 ns/op         2340800 B/op         29 allocs/op
BenchmarkStringJoin10000-4                  8244            133471 ns/op          360448 B/op          1 allocs/op
BenchmarkStringBuffer10000-4                5529            218346 ns/op         1192433 B/op         14 allocs/op
BenchmarkStringBuilder10000-4               3718            335846 ns/op         1979796 B/op         26 allocs/op

BenchmarkStringJoinWithoutArray-4             13          83842733 ns/op        632845640 B/op     10011 allocs/op
PASS
*/

// 表現好的還是Join和Builder。這兩個方法的使用側重點有些不一樣，
// * 如果有現成的數組、切片那麼可以直接使用Join,
// * 但是如果沒有，並且追求靈活性拼接，還是選擇Builder。
// * Join還是定位於有現成切片、數組的（畢竟拼接成數組也要時間），並且使用固定方式進行分解的，比如逗號、空格等，局限比較大。