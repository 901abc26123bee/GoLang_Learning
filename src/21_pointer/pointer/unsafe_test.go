package pointer_test

import (
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

// Go 的指針是不支持指針運算和轉換

// unsafe.Pointer是特别定義的一種指針類型（譯註：類似C語言中的void*類型的指針），它可以包含任意類型變量的地址
// 它表示任意類型且可尋址的指針值，可以在不同的指針類型之間進行轉換
// 第一是 unsafe.Pointer 可以讓你的變量在不同的指針類型轉來轉去，也就是表示為任意可尋址的指針類型。
// 第二是 uintptr 常用於與 unsafe.Pointer 打配合，用於做指針運算


func TestChangePointerType(t *testing.T) {
	num := 5
	numPointer := &num

	// flnum := (*float32)(numPointer) // cannot convert numPointer (variable of type *int) to *float32
	flnum := (*float32)(unsafe.Pointer(numPointer))
	flnum2 := *(*float32)(unsafe.Pointer(numPointer))
	flnum3 := unsafe.Pointer(&numPointer)
	fmt.Printf("Value : %v, Type : %T",flnum, reflect.TypeOf(flnum)) // Value : 0xc00001a310, Type : *reflect.rtype
	fmt.Printf("Value : %v, Type : %T",flnum2, reflect.TypeOf(flnum2)) // Value : 7e-45, Type : *reflect.rtype
	// 		--> fail to get the value which (*float32)(unsafe.Pointer(numPointer)) point at
	fmt.Printf("Value : %v, Type : %T",flnum3, reflect.TypeOf(flnum3)) // Value : 0xc00000e048, Type : *reflect.rtype
}

type Customer struct {
	Name string
	Age int
}

func TestUnsafe(t *testing.T) {
	i := 10
	f := *(*float64)(unsafe.Pointer(&i))
	fmt.Println(unsafe.Pointer(&i))	// 0xc0000ac1d8
	fmt.Println(f) // 5e-323
	fmt.Printf("Value : %v, Type : %T",f, reflect.TypeOf(f)) // Value : 5e-323, Type : *reflect.rtype
}

type Num struct{
	i string
	j int64
}

// 結構體的一些基本概念：
// 1. 結構體的成員變量在內存存儲上是一段連續的內存
// 2. 結構體的初始地址就是第一個成員變量的內存地址
// 3.基於結構體的成員地址去計算偏移量。就能夠得出其他成員變量的內存地址
func TestPointerOffset(t *testing.T) {
	n := Num{i: "Lala", j: 5}
	nPointer := unsafe.Pointer(&n)

	niPointer := (*string)(unsafe.Pointer(nPointer))
	*niPointer = "Xeon"

	// uintptr：uintptr 是 Go 的內置類型。返回無符號整數，可存儲一個完整的地址。後續常用於指針運算
	// unsafe.Offsetof：返回成員變量 x 在結構體當中的偏移量。更具體的講，就是返回結構體初始位置到 x 之間的字節數。
	// 需要注意的是入參 ArbitraryType 表示任意類型，並非定義的 int。它實際作用是一個佔位符
	// uintptr 類型是不能存儲在臨時變量中的。因為從 GC 的角度來看，uintptr 類型的臨時變量只是一個無符號整數，並不知道它是一個指針地址
	njPointer := (*int64)(unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(n.j)))
	*njPointer = 9

	fmt.Printf("n.i: %s, n.j: %d", n.i, n.j) // n.i: Xeon, n.j: 9
}

// this case is sutible for unsafe
type MyInt int

func TestConvert(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := *(*[]MyInt)(unsafe.Pointer(&a))
	fmt.Println(b) // [1 2 3 4 5]
}

func TestAtomic(t *testing.T) {
	var shareBufPtr unsafe.Pointer
	writeDataFn := func() {
		data := []int{}
		for i := 0; i < 100; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(&shareBufPtr, unsafe.Pointer(&data))
	}
	readDataFn := func() {
		data := atomic.LoadPointer(&shareBufPtr)
		fmt.Println(data, *(*[]int)(data))
	}
	var wg sync.WaitGroup
	writeDataFn()
	for i := 0; i <10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				writeDataFn()
				time.Sleep(time.Millisecond * 100)
			}
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				readDataFn()
				time.Sleep(time.Millisecond * 100)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

/*

0xc000126078 [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99]
0xc000126288 [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99]
0xc00000c030 [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99]
... total 100
*/