package transfer

import (
	"bytes"
	"reflect"
	"testing"
	"unsafe"
)

/*
		最常用的字串/數字轉換（string to int and int to string）
		Atoi and Itoa，基本上內部也是使用 ParseInt and FormatInt 來進行轉換。
		這兩種轉換最常用到，Golang 為他特別做了兩個獨立的 func。

		基本上這兩種預設如下（詳細請繼續往後看）

		Atoi = ParseInt(s, 10, 0)      // default: base 10, type: int
		Itoa = FormatInt(int64(i), 10) // default: base 10

		i, err := strconv.Atoi("-18") // result: i = -18
    s := strconv.Itoa(-18)        // result: s = "-18"

		字串轉其他格式（Parse 系列）
		b, err := strconv.ParseBool("true")         // result: true, bool
		f, err := strconv.ParseFloat("3.14", 64)    // result: 3.14, float64
		i, err := strconv.ParseInt("-18", 10, 64)   // result:  -18, int64
		u, err := strconv.ParseUint("18", 10, 64)   // result:   18, int64
*/

// ------------------------------------------------------------------------------------------------
// [徹底弄清Golang中[]byte與string轉換](https://www.gushiciku.cn/pl/pDSl/zh-tw)
/*
	標準轉換
		//* string to []byte
    s1 := "hello"
    b := []byte(s1)

    //* []byte to string
    s2 := string(b)
*/

// 強轉換
func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// 測試強轉換功能
func TestBytes2String(t *testing.T) {
	x := []byte("Hello Gopher!")
	y := Bytes2String(x)
	z := string(x)

	if y != z {
			t.Fail()
	}
}

// 測試強轉換功能
func TestString2Bytes(t *testing.T) {
	x := "Hello Gopher!"
	y := String2Bytes(x)
	z := []byte(x)

	if !bytes.Equal(y, z) {
			t.Fail()
	}
}

// 測試標準轉換string()效能
func Benchmark_NormalBytes2String(b *testing.B) {
	x := []byte("Hello Gopher! Hello Gopher! Hello Gopher!")
	for i := 0; i < b.N; i++ {
			_ = string(x)
	}
}

// 測試強轉換[]byte到string效能
func Benchmark_Byte2String(b *testing.B) {
	x := []byte("Hello Gopher! Hello Gopher! Hello Gopher!")
	for i := 0; i < b.N; i++ {
			_ = Bytes2String(x)
	}
}

// 測試標準轉換[]byte效能
func Benchmark_NormalString2Bytes(b *testing.B) {
	x := "Hello Gopher! Hello Gopher! Hello Gopher!"
	for i := 0; i < b.N; i++ {
			_ = []byte(x)
	}
}

// 測試強轉換string到[]byte效能
func Benchmark_String2Bytes(b *testing.B) {
	x := "Hello Gopher! Hello Gopher! Hello Gopher!"
	for i := 0; i < b.N; i++ {
			_ = String2Bytes(x)
	}
}

// -benchmem 可以提供每次操作分配記憶體的次數，以及每次操作分配的位元組數。
/*
$ go test -bench="." -benchmem
Benchmark_NormalBytes2String-4          40320345                29.02 ns/op           48 B/op          1 allocs/op
Benchmark_Byte2String-4                 1000000000               0.4932 ns/op          0 B/op          0 allocs/op
Benchmark_NormalString2Bytes-4          36297324                30.32 ns/op           48 B/op          1 allocs/op
Benchmark_String2Bytes-4                1000000000               0.4381 ns/op          0 B/op          0 allocs/op
*/
// *強轉換方式的效能會明顯優於標準轉換。

/*
	原理分析
*	[]byte: 在go中，在go標準庫builtin中有如下說明：

	byte is an alias for uint8 and is equivalent to uint8 in all ways. It is
	used, by convention, to distinguish byte values from 8-bit unsigned
	integer values.
	type byte = uint8
	byte是uint8的別名

	type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
	}

*	string
	關於string型別，在go標準庫builtin中有如下說明：

	string is the set of all strings of 8-bit bytes, conventionally but not
	necessarily representing UTF-8-encoded text. A string may be empty, but
	not nil. Values of string type are immutable.
	type string string
	string是8位位元組的集合，通常但不一定代表UTF-8編碼的文字。string可以為空，但是不能為nil。 string的值是不能改變的。

	type stringStruct struct {
    str unsafe.Pointer
    len int
	}

*	綜上，string與[]byte在底層結構上是非常的相近（後者的底層表達僅多了一個cap屬性，因此它們在記憶體佈局上是可對齊的），
	這也就是為何builtin中內建函式copy會有一種特殊情況 copy(dst []byte, src string) int 的原因了。
*/

/*
* 對於[]byte與string而言，兩者之間最大的區別就是string的值不能改變。

		對於[]byte來說，以下操作是可行的：
		b := []byte("Hello Gopher!")
				b [1] = 'T'

		string，修改操作是被禁止的：
		s := "Hello Gopher!"
				s[1] = 'T'

		而string能支援這樣的操作：
		s := "Hello Gopher!"
		s = "Tello Gopher!"

	string在底層都是結構體 stringStruct{str: str_point, len: str_len}
* string結構體的str指標指向的是一個字元常量的地址， 這個地址裡面的內容是不可以被改變的，因為它是隻讀的，但是這個指標可以指向不同的地址。
	以下操作的含義是不同的：

	s := "S1" // *分配儲存"S1"的記憶體空間，s結構體裡的str指標指向這塊記憶體
	s = "S2"  // *分配儲存"S2"的記憶體空間，s結構體裡的str指標轉為指向這塊記憶體

	b := []byte{1} // 分配儲存'1'陣列的記憶體空間，b結構體的array指標指向這個陣列。
	b = []byte{2}  // 將array的內容改為'2'

* 因為string的指標指向的內容是不可以更改的，所以每更改一次字串，就得重新分配一次記憶體，之前分配的空間還需要gc回收，這是導致string相較於[]byte操作低效的根本原因。
*/

/*
* 強轉換的實現細節
	萬能的unsafe.Pointer指標
	* 在go中，任何型別的指標*T都可以轉換為unsafe.Pointer型別的指標，它可以儲存任何變數的地址。
	* 同時，unsafe.Pointer型別的指標也可以轉換回普通指標，而且可以不必和之前的型別*T相同。
	* 另外，unsafe.Pointer型別還可以轉換為uintptr型別，該型別儲存了指標所指向地址的數值，從而可以使我們對地址進行數值計算。

	而string和slice在reflect包中，對應的結構體是reflect.StringHeader和reflect.SliceHeader，它們是string和slice的執行時表達。
	
	type StringHeader struct {
    Data uintptr
    Len  int
	}

	type SliceHeader struct {
			Data uintptr
			Len  int
			Cap  int
	}

	從string和slice的執行時表達可以看出，除了SilceHeader多了一個int型別的Cap欄位，Date和Len欄位是一致的。
	所以，它們的記憶體佈局是可對齊的，這說明我們就可以直接通過unsafe.Pointer進行轉換。
*/

/*
		Q1. 為啥強轉換效能會比標準轉換好？
		對於標準轉換，無論是從[]byte轉string還是string轉[]byte都會涉及底層陣列的拷貝。
		而強轉換是直接替換指標的指向，從而使得string和[]byte指向同一個底層陣列。這樣，當然後者的效能會更好。

		Q2. 為啥在上述測試中，當x的資料較大時，標準轉換方式會有一次分配記憶體的操作，從而導致其效能更差，而強轉換方式卻不受影響?
		標準轉換時，當資料長度大於32個位元組時，需要通過mallocgc申請新的記憶體，
		之後再進行資料拷貝工作。而強轉換隻是更改指標指向。所以，當轉換資料較大時，兩者效能差距會愈加明顯。


		Q3. 既然強轉換方式效能這麼好，為啥go語言提供給我們使用的是標準轉換方式？
		Go是一門型別安全的語言，而安全的代價就是效能的妥協。
		但是，效能的對比是相對的，這點效能的妥協對於現在的機器而言微乎其微。另外強轉換的方式，會給我們的程式帶來極大的安全隱患。
		a := "hello"
		b := String2Bytes(a)
		b[0] = 'H'
		a是string型別，前面我們講到它的值是不可修改的。
		* 通過強轉換將a的底層陣列賦給b，而b是一個[]byte型別，它的值是可以修改的，
		* 所以這時對底層陣列的值進行修改，將會造成嚴重的錯誤（通過defer+recover也不能捕獲）。

		unexpected fault address 0x10b6139
		fatal error: fault
		[signal SIGBUS: bus error code=0x2 addr=0x10b6139 pc=0x1088f2c]

*/