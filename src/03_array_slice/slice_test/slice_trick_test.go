package slicetest

import (
	"fmt"
	"testing"
)


func log(slice []string) {
	fmt.Printf("%v ({len: %d, cap: %d}) \n", slice, len(slice), cap(slice))
	fmt.Printf("addr of first element %p. add of slice %p\n", &slice[0], &slice)
}

func TestModSliceByPassToFunction(t *testing.T) {
	slice := make([]string, 2, 3)
	log(slice)

	func(slice []string) {
		// 先 append 的話 slice 因為 metadata 改變了，已經不同，
		// 因此使用 slice[0], slice[1] 對於原本的 slice 進行改值將不會有作用
		slice = append(slice, "a")
		log(slice)
		slice[0] = "b"
		slice[1] = "b"

		// 後來才 append 的話，即可先修改到原本的 slice
		// slice = append(slice, "a", "a")
		log(slice)
	}(slice)

	log(slice)
	/*
		[ ] ({len: 2, cap: 3})
		addr of first element 0xc000066480. add of slice 0xc00000c090
		[  a] ({len: 3, cap: 3})
		addr of first element 0xc000066480. add of slice 0xc00000c0c0
		[b b a] ({len: 3, cap: 3})
		addr of first element 0xc000066480. add of slice 0xc00000c0f0
		[b b] ({len: 2, cap: 3})
		addr of first element 0xc000066480. add of slice 0xc00000c120
	*/
}

func TestModSliceByCreateNewOne(t *testing.T) {
	slice := make([]string, 1, 3)

	func(slice []string) {
		// 這個 slice 和原本的 slice 有不同的 metadata
		// 因此無法使用 slice[0] 和 slice[1] 去改到外層的 slice
		slice = slice[1:3]
		slice[0] = "b"
		slice[1] = "b"

		log(slice)
	}(slice)
	log(slice)

	/*
		[b b] ({len: 2, cap: 2})
		addr of first element 0xc000066490. add of slice 0xc00000c090
		[] ({len: 1, cap: 3})
		addr of first element 0xc000066480. add of slice 0xc00000c0c0
	*/
}