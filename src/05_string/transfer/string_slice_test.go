package transfer

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

/*
	In logic, a string value denotes a piece of text.
	In memory, a string value stores a sequence of bytes,
	which is the UTF-8 encoding representation of the piece of text denoted by the string value.
	We can learn more facts on strings from the article strings in Go later.
*/

func Test_String_To_Slice(t *testing.T) {
	s := strings.Split("a,b,c", ",")
	fmt.Println(s) // [a b c]

	s2 := strings.Fields(" a \t b \n")
	fmt.Println(s2) // [a b]
}

func Test_String_To_Slice2(t *testing.T) {
	a := "aaahikhl"
	ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a))
	b := *(*[]byte)(unsafe.Pointer(&ssh))
	fmt.Printf("%v", b) // [97 97 97 104 105 107 104 108]
}

func Test_String_By_Index(t *testing.T) {
	s := "aaahikhl"
	fmt.Println(s[2], s[5]) // 97 107
}

/*
	StringHeader 是 string 在go的底層結構。
	type StringHeader struct {
		Data uintptr
		Len  int
	}
	SliceHeader 是 slice 在go的底層結構。
	type SliceHeader struct {
		Data uintptr
		Len  int
		Cap  int
	}

	如果想要在底層轉換二者，只需要把 StringHeader 的地址強轉成 SliceHeader 就行。那麼go有個很強的包叫 unsafe
*/

// 翻轉含有中文、數字、英文字母的字符串
func Test_String_Reverse(t *testing.T) {
	src := "你好abc啊哈哈"
	dst := reverse([]rune(src))
	fmt.Printf("%v\n", string(dst)) // 哈哈啊cba好你
}

func reverse(s []rune) []rune {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
/*
	rune關鍵字，從golang源碼中看出，它是int32的別名（-2^31 ~ 2^31-1），比起byte（-128～127），可表示更多的字符。
	*由於rune可表示的範圍更大，所以能處理一切字符，當然也包括中文字符。
	*在平時計算中文字符，可用rune。 因此將字符串轉為rune的切片，再進行翻轉，完美解決。
*/

// * Use _ in numeric literals for better readability
func Test_Literals(t *testing.T) {
	fmt.Println(0_33_77_22 == 0337722) // true
	fmt.Println(0x_Bad_Face == 0xBadFace) // true
	fmt.Println(6_9) // 69
	fmt.Println(0b1011_0111 + 0xA_B.Fp2i) // (183+687.75i)
}

/*
	* 在Go中，一個rune值表示一個Unicode碼點。
		一般說來，我們可以將一個Unicode碼點看作是一個Unicode字符.
		但是，我們也應該知道，有些Unicode字符由多個Unicode碼點組成。
		每個英文或中文Unicode字符值含有一個Unicode碼點。

*	一個rune字面量由若干包在一對單引號中的字符組成。包在單引號中的字符序列表示一個Unicode碼點值。
	rune字面量形式有幾個變種，其中最常用的一種變種是將一個rune值對應的Unicode字符直接包在一對單引號中。
	A rune literal is expressed as one or more characters enclosed in a pair of quotes.
	The enclosed characters denote one Unicode code point value.
	There are some minor variants of the rune literal form.
	The most popular form of rune literals is just to enclose the characters denoted by rune values between two single quotes.
	'a' // 一個英文字符
	'π'
	'眾' // 一個中文字符

	在日常編程中, rune字面量形式多用做字符串的雙引號字面量形式中的轉義字符

*/

// https://learnku.com/articles/23411/the-difference-between-rune-and-byte-of-go
/*
	* different between [] rune and [] byte in golang
		byte is an alias for uint8 and is equivalent to uint8 in all ways. It is
		used, by convention, to distinguish byte values from 8-bit unsigned
		integer values.
		type byte = uint8

		rune is an alias for int32 and is equivalent to int32 in all ways. It is
		used, by convention, to distinguish character values from integer values.
		type rune = int32

	* byte 表示一個字節，rune 表示四個字節
*/

func Test_byte_rune_slice_compare(t *testing.T) {
	first := "fisrt"
	fmt.Println([]rune(first))	// [102 105 115 114 116]
	fmt.Println([]byte(first))	// [102 105 115 114 116]

	second := "go語言"
	fmt.Println([]rune(second)) // [103 111 35486 35328]
	fmt.Println([]byte(second)) // [103 111 232 170 158 232 168 128]
	// 中文字符串每個佔三個字節
}

func Test_substring_by_slice(t *testing.T) {
	s := "golangcaff"
	fmt.Println(s[:3]) // gol

	s2 := "截取中文"
	fmt.Println(s2[:2]) // ��

	res := []rune(s2)
	fmt.Println(string(res[:2])) // 截取

	// 為什麼 s[:n] 無法直接截取呢，
	// * 猜測 如果直接截取的話，底層會將中文轉化成 []byte， 而不是 []rune
}