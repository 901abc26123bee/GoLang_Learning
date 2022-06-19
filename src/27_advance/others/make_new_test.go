package make_new_test

import (
	"fmt"
	"sync"
	"testing"
)

func TestMakeAndNew(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println(s1) // [1 2 3 4 5]

	// ----------------------------------------------------------------

	list := new([]int)
	// list = append(list, 1) // first argument to append must be slice; have *[]int
	fmt.Println(list)	// &[]


	// 切片指針的解引用。
	list2 := make([]int, 0) // list類型為切片
	list2 = append(list2, 3)
	list2 = append(list2, 5)
	fmt.Println(list2) // [3 5]

	list3 := append(*list, 1) // list類型為指針
	list3 = append(list3, 3)
	fmt.Println(list3) // [0 3]


	list4 := new([3]int)
	fmt.Println(list4)	// &[0 0 0]
	/*
	new和make的區別：​ 二者都是內存的分配（堆上），
	但是make只用於slice、map以及channel的初始化（非零值）；
	而new用於類型的內存分配，並且內存置為零。
	所以在我們編寫程序的時候，就可以根據自己的需要很好的選擇了。​ make返回的還是這三個引用類型本身；而new返回的是指向類型的指針。
	*/
}


func TestNoAllocateMem(t *testing.T) {
	var i *int
	*i=10
	fmt.Println(*i)
	// 對於引用類型的變量，我們不光要聲明它，還要為它分配內容空間
	// 對於值類型的聲明不需要，是因為已經默認幫我們分配好了。
	/*
	panic: runtime error: invalid memory address or nil pointer dereference [recovered]
					panic: runtime error: invalid memory address or nil pointer dereference
	[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x10ef0d6]
	*/
}

func TestNew(t *testing.T) {
	var i *int
	i = new(int)
	fmt.Println(i, *i) // 0xc00011c190 0
	*i = 10
	fmt.Println(i, *i) // 0xc00011c190 10

	/*
	The new built-in function allocates memory. The first argument is a type,
	not a value, and the value returned is a pointer to a newly
	allocated zero value of that type.
	func new(Type) *Type

	它只接受一個參數，這個參數是一個類型，
	分配好內存後，返回一個指向該類型內存地址的指針。同時請注意它同時把分配的內存置為零，也就是類型的零值。
	*/
}

type user struct {
	lock sync.Mutex
	name string
	age int
}
//*  new，它返回的永遠是類型的指針，指向分配類型的內存地址。
func TestNew2(t *testing.T) {
	 u := new(user)

	 u.lock.Lock() // lock字段不用初始化，直接可以拿來用，不會有無效內存引用異常，因為它已經被零值了
	 u.name = "Dany"
	 u.lock.Unlock()

	 fmt.Println(u)
}

/*
	* make也是用於內存分配的，但它只用於 chan map slice 的內存創建，
	* 這三種類型就是引用類型,返回的類型就是這三個類型本身，所以就沒有必要返回他們的指針了。
		func make(t Type, size ...IntegerType) Type
	注意，因為這三種類型是引用類型，所以必須得初始化，但是不是置為零值，這個和new是不一樣的。
	在使用slice、map以及channel的時候，還是要使用make進行初始化，然後才才可以對他們進行操作。

*/