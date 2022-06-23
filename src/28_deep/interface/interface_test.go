package deep_interface_test

import (
	"fmt"
	"testing"
)

type error interface {
	Error() string
}

type RPCError struct {
	Code int64
	Message string
}

func (e *RPCError) Error() string {
	return fmt.Sprintf("%s, code=%d", e.Message, e.Code)
}

func NewRPCError(code int64, message string) error {
	return &RPCError{ // typecheck3
		Code: code,
		Message: message,
	}
}

func AsErr(err error) error {
	return err
}

func Test_Interface_Runtime_Check(t *testing.T) {
	var rpcErr error = NewRPCError(400, "unknown err") // typecheck1
	err := AsErr(rpcErr) // typecheck2
	println(err) // (0x1148840,0xc000168748)
}
/*
	使用上述 RPCError 結構體時並不關心它實現了哪些接口，Go 語言只會在傳遞參數、返回參數以及變量賦值時才會對某個類型是否實現接口進行檢查，
* Go 語言在編譯期間對代碼進行類型檢查，上述代碼總共觸發了三次類型檢查：
	1. 將 *RPCError 類型的變量賦值給 error 類型的變量 rpcErr；
	2. 將 *RPCError 類型的變量 rpcErr 傳遞給簽名中參數類型為 error 的 AsErr 函數；
	3.n將 *RPCError 類型的變量從函數簽名的返回值類型為 error 的 NewRPCError 函數中返回；
* 從類型檢查的過程來看，編譯器僅在需要時才檢查類型，類型實現接口時只需要實現接口中的全部方法，
	不需要像 Java 等編程語言中一樣顯式聲明。
*/

// ------------------------------------------------------------------------------------------------
/*
*	需要注意的是，與 C 語言中的 void * 不同，interface{} 類型不是任意類型。
* 如果我們將類型轉換成了 interface{} 類型，變量在運行期間的類型也會發生變化，獲取變量類型時會得到 interface{}。
*/
func Print(v interface{}) {
	println(v)
}

/*
* 上述函數不接受任意類型的參數，只接受 interface{} 類型的值，
* 在調用 Print 函數時會對參數 v 進行類型轉換，將原來的 Test 類型轉換成 interface{} 類型，本節會在後面介紹類型轉換的實現原理。
*/
func Test_empty_interface_receiver(t *testing.T) {
	type Test struct {}
	v := Test{}
	Print(v)
}