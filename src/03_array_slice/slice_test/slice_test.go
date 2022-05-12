package slicetest

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

/*
	從切片的定義我們能推測出，切片在編譯期間的生成的類型只會包含切片中的元素類型，即 int 或者 interface{} 等。
	cmd/compile/internal/types.NewSlice 就是編譯期間用於創建切片類型的函數：

	Extra 字段是一個只包含切片內元素類型的結構，
	也就是說切片內元素的類型都是在編譯期間確定的，編譯器確定了類型之後，會將類型存儲在 Extra 字段中幫助程序在運行時動態獲取。

	切片與數組的關係非常密切，切片引入了一個抽象層，提供了對數組中部分連續片段的引用，
	而作為數組的引用，我們可以在運行區間可以修改它的長度和範圍。
	當切片底層的數組長度不足時就會觸發擴容，切片指向的數組可能會發生變化，
	不過在上層看來切片是沒有變化的，上層只需要與切片打交道不需要關心數組的變化。

	Go 語言中包含三種初始化切片的方式：
	1. 通過下標的方式獲得數組或者切片的一部分；
	2. 使用字面量初始化新的切片；
	3. 使用關鍵字 make 創建切片：

	在 Slice 中會包含
		1. Pointer to Array：這個 pointer 會指向實際上在底層的 array。
		2. Capacity：從 slice 的第一個元素開始算起，它底層 array 的元素數目
		3. Length：該 slice 中的元素數目
*/

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5) // func make([]type, len, capacity)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	// t.Log(s2[0], s2[1], s2[2], s2[3]) // error, since len == 3

	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
}


func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMenory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun","Jul","Aug","Sep","Oct","Nov","Dec"}
	Q2 := year[3 : 6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5 : 8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknown"
	t.Log(Q2)
	t.Log(year)
}

// slice can only be compared to nil
// func TestSliceCompare(t *testing.T) {
// 	a := []int{1, 2, 3, 4}
// 	b := []int{1, 2, 3, 4}
// 	if a == b {
// 		t.Log("equal")
// 	}
// }

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 5, 3, 4}
	d := [...]int{1, 2, 3, 4}
	// f := [...]int{1, 2, 3, 4, 5, 6}

	t.Log(a == b) // false
	t.Log(a == d) // true
	// t.Log(a == f) // invalid operation: cannot compare a == f (mismatched types [4]int and [6]int)
}

func TestArrayandSlice(t *testing.T) {
		// -------------------------------- Slice 的建立 --------------------------------
		// slice 型別：[]T
		// slice := make([]T, len, cap)
	  // 方式一：建立一個帶有資料的 string slice，適合用在知道 slice 裡面的元素有哪些時
		people := []string{"Aaron", "Jim", "Bob", "Ken"}

		// 方式二：透過 make 可以建立「空 slice」，適合用會對 slice 中特定位置元素進行操作時
		people2 := make([]int, 4)  // len=4 cap=4，[0,0,0,0]

		// 方式三：空的 slice，一般會搭配 append 使用
		var people3 []string

		// 方式四：大該知道需要多少元素時使用，搭配 append 使用
		people4 := make([]string, 0, 5) // len=0 cap=5, []
		fmt.Print(people)
		fmt.Print(people2)
		fmt.Print(people3)
		fmt.Print(people4)

		// ---------------------------------- Arrays 的建立 ------------------------------
		// 在 Array 中陣列的元素是固定的：
		// 陣列型別：[n]T
		// 1. 先定義再賦值
		// 2. 使用 [...]T{} 可以根據元素的數目自動建立陣列

		// 先定義再賦值
    var a [2]string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a)  // [Hello World]

  	// 定義且同時賦值
    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes)  // [2 3 5 7 11 13]

		// 沒有使用 ...，建立出來的會是 slice
		arr := []string{"North", "East", "South", "West"}
		fmt.Println(reflect.TypeOf(arr).Kind(), len(arr))  // slice 4

		// 使用 ...，建立出來的會是 array
		arrWithDots := [...]string{"North", "East", "South", "West"}
		fmt.Println(reflect.TypeOf(arrWithDots).Kind(), len(arrWithDots))  // array 4
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

/*
------------------------------ capacity and length---------------------------------
	在 Slice 中會包含
		1. Pointer to Array：這個 pointer 會指向實際上在底層的 array。
		2. Capacity：從 slice 的第一個元素開始算起，它底層 array 的元素數目
		3. Length：該 slice 中的元素數目
*/
func TestSliceLenAndCapacity(t *testing.T) {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)  // len=6 cap=6 [2 3 5 7 11 13]

	s = s[:0]
	printSlice(s)  // len=0 cap=6 []

	s = s[:4]
	printSlice(s)  // len=4 cap=6 [2 3 5 7]

	s = s[2:]
	printSlice(s)  // len=2 cap=4 [5 7]
}
// 使用 make 建立 slice 時，可以指定該 slice 的 length 和 capacity：
func TestCreateSliceWithMakers(t *testing.T) {
	// make(T, length, capacity)
	a := make([]int, 5) // len(a)=5, cap(a)=5, [0 0 0 0 0]
	fmt.Print(a)
	// 建立特定 capacity 的 slice
	b := make([]int, 0, 5) // len=0 cap=5, []
	b = b[:cap(b)]         // len=5 cap=5, [0 0 0 0 0]
	b = b[1:]              // len=4 cap=4, [0 0 0 0]
}


// 若 length 的數量不足時，將無法將元素放入 slice 中，這時候可以使用 append 來擴展 slice：
func TestLengthNotEnough(t *testing.T) {
	scores := make([]int, 0, 10)
	fmt.Println(len(scores), cap(scores)) // 0, 10
	//scores[7] = 9033  // 無法填入元素，因為 scores 的 length 不足

	// 這時候可使用 append 來擴展 slice
	scores = append(scores, 5)
	fmt.Println(scores) // [5]
	fmt.Println(len(scores), cap(scores)) // 1, 10

	// 但要達到原本的目的，需要使用切割 slice *****
	scores = scores[0:8]
	fmt.Println(len(scores), cap(scores)) // 8, 10
	scores[7] = 9033
	fmt.Println(scores) // [5 0 0 0 0 0 0 9033]
}

// zero value
// slice 的 zero value 是 nil，也就是當一個 slice 的 length, capacity 都是 0，而且沒有底層 array 時：
func TestSliceWithZeroValue(t *testing.T) {
	var s []int
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)   // len=0 cap=0 []

	if s == nil {
			fmt.Println("nil!")  // nil!
	}
}

// slice of integer, boolean, struct
func TestSliceType(t *testing.T) {
	// integer slice
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	// boolean slice
	r := []bool{ true, false, true, true, false, true}
	fmt.Println(r)

	// struct slice
	s := []struct {
			i int
			b bool
	}{
			{2, true},
			{3, false},
			{5, true},
			{7, true},
			{11, false},
			{13, true},
	}
	fmt.Println(s)
}

// -------------------------- slices of slices -------------------------------
func TestSliceOfSlices(t *testing.T) {
	arr := make([][]int, 3)
	fmt.Println(arr) // [[] [] []]

  // 賦值
  arr[0] = []int{1}
	arr[1] = []int{2}
	arr[2] = []int{3}
	fmt.Println(arr) // [[1] [2] [3]]

	arr2 := [][]int{
		[]int{1},
		[]int{2},
	}
	fmt.Println(arr2) // [[1] [2]]

	board := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}
	fmt.Println(board) // [[- - -] [- - -] [- - -]]

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	fmt.Println(board) // [[X - X] [O - X] [- - O]]

	//X - X
	//O - X
	//- - O
	for i := 0; i < len(board); i++ {
			fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func TestSliceIterator(t *testing.T) {
	// 定義內部元素型別為字串的陣列
	fruits := []string{"apple", "banana", "grape", "orange"}

	// 取得 slice 的 length 和 capacity
	fmt.Println(len(fruits)) // length, 4
	fmt.Println(cap(fruits)) // capacity, 4

	// append 為陣列添加元素（不會改變原陣列）
	fruits = append(fruits, "strawberry")

	// range syntax: fruits[start:end] 可以截取陣列，從 "start" 開始到 "end-1" 的元素
	fmt.Println(fruits[0:2]) //  [apple, banana]
	fmt.Println(fruits[:3])  //  [apple, banana, grape]
	fmt.Println(fruits[2:])  //  [grape, orange, strawberry]

	// 透過 range cards 可以疊代陣列元素
	for _, fruit := range fruits {
			fmt.Println(fruit)
	}
	// 在疊代陣列時會使用到 for i, card := range cards{…}，
	// 其中如果 i (index of elem) 在疊代時沒有被使用到，但有宣告它時，程式會噴錯；因此如果不需要用到這個變數時，需把它命名為 _。
}

// -------------------------- 切割（:）------------------------------
func TestSliceSlice(t *testing.T) {
	// 在 golang 中使用 : 來切割 slice 時，並不會複製原本的 slice 中的資料，而是建立一個新的 slice，
	// 但實際上還是指稱到相同位址的底層 array（並沒有複製新的），因此還是會改到原本的元素：
	scores := []int{1, 2, 3, 4, 5}
	newSlice := scores[2:4]
	fmt.Println(newSlice)  // 3, 4
	newSlice[0] = 999      // 把原本 scores 中 index 值為 3 的元素改成 999
	fmt.Println(scores)    // 1, 2, 999, 4, 5
}

/*
		有時我們只需要使用原本 slice 中的一小部分元素，
		但由於透過 : 的方式重新分割 slice 時，該 slice 仍然會指稱到底層相同的 array，
		這導致雖然我們只用了原本 slice 中的一小部分元素，但其他多餘的部分因為仍然被保留在底層的 array 中，
		進而使得該 array 無法被 garbage collection。因此若只需要使用原本 slice 中一小部分的元素內容時，
		建議可以使用 copy 或 append 建立新的 slice 與其對應的新的 array。

		另外，使用 : 時，不能超過它本身的 capacity，否則會導致 runtime panic：
*/


func TestOutOfCapacity(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}

	// over the slice capacity will cause runtime panic
	overSliceCap := s[:6]
	fmt.Println(overSliceCap)
	// panic: runtime error: slice bounds out of range [:6] with capacity 5 [recovered]
	// 				panic: runtime error: slice bounds out of range [:6] with capacity 5

	// 由於無法直接透過 : 來擴充 slice 的 capacity，
	// 因此若有需要擴充原 slice 的 capacity 時，需要透過 append 或 copy 的方式來擴充原本 slice 的 capacity。
}


// ------------------------------ slice operation ----------------------------------

// 使用 append 可以擴充 slice 底層 array 的 capacity。
// 在底層 array 的 capacity 已經滿的情況下，append 它會創造一個更大的陣列，並且複製原本的值到新陣列中：
// func append(s []T, x ...T) []T
func TestSliceAppend(t *testing.T) {
	var s []int
  printSlice(s) // len=0 cap=0 []

  // 一次添加多個元素
	s = append(s, 2, 3)
	printSlice(s) //len=2 cap=2 [2 3]

  // 一次添加一個元素
	s = append(s, 4)
	printSlice(s) // len=3 cap=4 [2 3 4]
	// cap 的數量再擴展時會是 1, 2, 4, 8, 16, .... 。
}

// append 也可以把另一個 slice append 進到原本的 slice 中。
// 使用 append(a, b...) 就可以把 b 這個 slice append 到 a 這個 slice 中：
func TestAppendSlice(t *testing.T) {
	s := []int{1, 2}
	b := []int{3, 4, 5}

	s = append(s, b...) // 等同於 append(s, b[0], b[1], b[2])
	fmt.Println(s) // [1 2 3 4 5]
}

// func copy(dst, src []T) int
// 回傳的值是複製了多少元素（the number of elements copied）
//  1. cloneScores 的元素數量如果「多於」被複製進去的元素時，會用 zero value 去補。
// 			例如，當 cloneScores 的長度是 4，但只複製 3 個元素進去時，最後位置多出來的元素會補 zero value。
//  2. cloneScores 的元素數量如果「少於」被複製進去的元素時，超過的元素不會被複製進去。
//			例如，當 cloneScores 的長度是 1，但卻複製了 3 個元素進去時，只會有 1 個元素被複製進去。
func TestSliceCopy(t *testing.T) {
	scores := []int{1, 2, 3, 4, 5}

	// STEP 1：建立空 slice 且長度為 4
	cloneScores := make([]int, 4)

	// STEP 2：使用 copy 將前 scores 中的前三個元素複製到 cloneScores 中
	copy(cloneScores, scores[:len(scores)-2])
	fmt.Println(cloneScores) // [1,2,3, 0]
}

// 使用 copy 可以用來擴展擴展 slice 的 capacity。
func TestExtendSliceByCopy(t *testing.T) {
	// 在沒有 copy 時，寫法會像這樣：
	s := []int{1, 2}
	copys := make([]int, len(s), (cap(s)+1)*2) // +1 in case cap(s) == 0
	for i := range s {
		copys[i] = s[i]
	}
	printSlice(s) // len=2 cap=2 [1 2]
	s = copys
	printSlice(s) // len=2 cap=6 [1 2]

	// 使用 copy 的話可以簡化成這樣：
	k := []int{1, 2}
	// s is another slice
	copyk := make([]int, len(k), (cap(k)+1)*2)
	copy(copyk, k)
	printSlice(k) // len=2 cap=2 [1 2]
	k = copyk
	printSlice(k) // len=2 cap=6 [1 2]
}

// range
// 如果只需要 index：for i := range pow
// 如果只需要 value：for _, value := range pow
func TestSkiceRange(t *testing.T) {
	pow := []int {1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
			fmt.Printf("2**%d = %d\n", i , v)
	}
}

// ----------------------------------------------------------------
/*
	把 slice 當成 function 的參數時
		實際上 slice 並不會保存任何資料，它只是描述底層的 array，
		當我們修改 slice 中的元素時，實際上是修改底層 array 的元素。
		不同的 slice 只要共享相同底層的 array 時，就會看到相同對應的變化。
*/

func changeSliceItem(words []string) {
	words[0] = "Hi"
}
func TestPassSliceToFunction(t *testing.T) {
	// Example 1:
	names := []string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	a := names[0:2]
	b := names[1:3]

	fmt.Println(a, b) // [John Paul] [Paul George]

	b[0] = "XXX"       // a 和 b 這兩個 slice 參照到的是底層相同的 array
	fmt.Println(a, b)  // [John XXX] [XXX George]
	fmt.Println(names) // [John XXX George Ringo]

	// Example 2:
	words := []string{"Hello", "every", "one"}
    fmt.Println(words) // [Hello every one]
    changeSliceItem(words)
    fmt.Println(words) // [Hi every one]
}

// Byte Slice
// Byte Slice 其實本質上就是字串的另一種表示方式，它只是將字串轉換成電腦比較好理解的方式。
// 例如字串 "Hi there!" 轉換成 byte slice 後，就是將這些單字轉換成十進位 的 ASCII 碼：


// 移除特定 index 的元素
func TestRemoveByIndex(t *testing.T) {
	scores := []int{1, 2, 3, 4, 5}
	scores = removeAtIndex(scores, 2)
	fmt.Println(scores)  // [1, 2, 5, 4]
}

func removeAtIndex(source []int, index int) []int {
	// STEP 1：取得最後一個元素的 index
	lastIndex := len(source) - 1

	// STEP 2：把要移除的元素換到最後一個位置
	source[index], source[lastIndex] = source[lastIndex], source[index]

	// STEP 3：除了最後一個位置的元素其他回傳出去
	return source[:lastIndex]
}

//------------------------------------ others ------------------------------------------

// Implicit memory aliasing in for loop
// The warning means, in short, that you are taking the address of a loop variable.
func TestImplicitMemAlas(t *testing.T) {
	versions := []int{7, 8, 9, 10}
	for i, v := range versions {
    // 不要使用 res := createWorkerFor(&v)
		/*
			This happens because in for statements the iteration variable(s) is reused.
			At each iteration,
			the value of the next element in the range expression is assigned to the iteration variable;
			v doesn't change, only its value changes.
			Hence, the expression &v is referring to the same location in memory.
		*/
		// res := createWorkerFor(&versions[i])  // 使用 versions[i]
		fmt.Println("v: ", v)
		fmt.Println("&v: ", v)
		fmt.Println("versions[i]: ", versions[i])
		fmt.Println("&versions[i]: ", &versions[i])
		// v:  7
		// &v:  7
		// versions[i]:  7
		// &versions[i]:  0xc000024100
		// v:  8
		// &v:  8
		// versions[i]:  8
		// &versions[i]:  0xc000024108
	}
}

// preallocating (prealloc) 的問題
// 雖然 append 會自動幫我們擴展 slice 的 capacity，
// 但如果可能的話，在建立 slice 時預先定義好 slice 的 length 和 capacity 將可以避免不必要的記憶體分派（memory allocations）。
// 💡 每次 array 的 capacity 改變時，就會進行一次 memory allocation。

func TestPreAllocate(t *testing.T) {
	// 方法一：使用 append ，搭配 length: 0，並設定 capacity 定好：
	collection := []string{"aa", "bb", "cc", "dd", "ee"}
	to := make([]string, 0, len(collection))
	for _, s := range collection {
			to = append(to, s)
	}
	fmt.Println(to) // [aa bb cc dd ee]


	// 方法二（maybe better）：不使用 append，直接設定 len 和 cap 並透過 index 的方式把值給進去：
	to2 := make([]string, len(collection))
	for i, s := range collection {
			to2[i] = s
	}
	fmt.Println(to2) // [aa bb cc dd ee]
}
