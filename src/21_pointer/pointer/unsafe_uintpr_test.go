package pointer_test

import (
	"fmt"
	"testing"
	"unsafe"
)
/*
uintptr和unsafe.Pointer的區別？
	1. unsafe.Pointer只是單純的通用指針類型，用於轉換不同類型指針，它不可以參與指針運算；
	2. 而uintptr是用於指針運算的，GC 不把 uintptr 當指針，也就是說 uintptr 無法持有對象， uintptr 類型的目標會被回收；
	3. unsafe.Pointer 可以和 普通指針 進行相互轉換；
	4. unsafe.Pointer 可以和 uintptr 進行相互轉換。
*/

type W struct {
	b int32
	c int64
}

func TestUnSafePointer(t *testing.T) {
	var w *W = new(W)
	fmt.Println(w.b,w.c) // 0 0
	b := unsafe.Pointer(uintptr(unsafe.Pointer(w)) + unsafe.Offsetof(w.b)) // b is address of w.b
	*((*int)(b)) = 10
	fmt.Println(w.b,w.c) // 10 0
}

/*
	1. uintptr(unsafe.Pointer(w)) 獲取了 w 的指針起始值
	2. unsafe.Offsetof(w.b) 獲取 b 變量的偏移量
	3. 兩個相加就得到了 b 的地址值，
		 將通用指針 Pointer 轉換成具體指針 ((*int)(b))，通過 * 符號取值，然後賦值。
		 *((*int)(b)) 相當於把 (*int)(b) 轉換成 int 了，最後對變量重新賦值成 10，這樣指針運算就完成了。
*/