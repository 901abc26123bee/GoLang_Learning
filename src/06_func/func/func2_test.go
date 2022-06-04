package func_test

import (
	"fmt"
	"math"
	"testing"
)

// compute 這個函式可接收其他函式作為參數
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func TestFuncAsParam(t *testing.T) {
	// *使用 expression 方式定義
	// sqrt(a^2 + b^2)
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(compute(hypot))    // 5, hypot(3, 4)
	fmt.Println(compute(math.Pow)) // 81,math.Pow(3, 4)

	// *anonymous function
	// 沒有參數
  func() {
		fmt.Println("Hello anonymous") // Hello anonymous
	}()

	// 有參數
	func(i, j int) {
			fmt.Println(i + j) // 3
	}(1, 2)
}


// --------------------------------- closure --------------------------------
// 函式也可以在執行後回傳另一個函式（閉包）

// Example 1:
// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	position := 0
	cache := map[int]int{}

	return func() int {
		position++
		if position == 1 {
			cache[position] = 0
			return 0
		} else if position <= 3 {
			cache[position] = 1
		} else {
			cache[position] = cache[position-2] + cache[position-1]
		}
		return cache[position]
	}
}

func TestFuncReturnFunc(t *testing.T) {
	f := fibonacci() // return a anonymous function
	for i := 0; i < 10; i++ {
			fmt.Println(f()) // execute the returned anonymous function by fibonacci
	}
}


// Example 2:
// 在 golang 的函式同樣適用閉包的概念，可以利用閉包把某個變數保留起來：
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func TestClousures(t *testing.T) {
	// pos 中的 sum 和 neg 中的 sum 是不同變數
	// pos == neg == func(x int) int inside adder()
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2 * i))
	}
}


// Example 3:
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func TestIntSeq(t *testing.T) {
	// We call intSeq, assigning the result (a function) to nextInt.
	// This function value captures its own i value,
	// which will be updated each time we call nextInt.
	nextInt := intSeq()

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	newInts := intSeq()
  fmt.Println(newInts()) // 1
}
// ----------------------------- 透過 structure 增加參數的可擴充性 -----------------------------------
// 透過 structure 增加參數的可擴充性


// 如果原本的 function 只需要兩個參數，可以這樣寫：
func add1(x, y int) int {
    return x + y
}

func TestAdd(t *testing.T) {
    fmt.Println(add1(1 ,2))
}


// about parameter :
// 比較好的作法是去定義 struct。如此，未來如果參數需要擴充，只需要改動 struct 和 func 內就好，
// 不用去改動使用這個 function 地方的參數：
// STEP 1：定義參數的 structure
type addOpts struct {
	x int
	y int
	z int    // STEP 4：如果新增一個參數
}

// STEP 2：把參數的型別指定為 structure
func add(opts addOpts) int {
	// STEP 5：接收新的參數 z
	// return opts.x + opts.y
	return opts.x + opts.y + opts.z
}

// STEP 3：使用 add，參數的地方使用 structure
func TestStrut(t *testing.T) {
	// STEP 6：不用改用舊有參數的寫法
	result := add(addOpts{
		x: 10,
		y: 5,
	})

	newResult := add(addOpts{
		x: 10,
		y: 5,
		z: 7,
	})
	fmt.Println(result, newResult)
}

// ------------------------------ func() return type ----------------------------------

// 沒有回傳值
// 對於沒有回傳值的函式可以不用定義回傳的型別：
func hello() {
	fmt.Println("Hello Go")
}

// 單一回傳值只需要在定義回傳型別的地方給一個型別就好：
func add3(i, j int) int {
    return i + j
}

// return multi result
// 多個回傳值，會需要在定義回傳型別的地方給多個型別：
// func nameOfFunction(<arguments>) (<type>)
func swap(x, y string) (string, string) {
    return y, x
}

func TestMultiReturnValue(t *testing.T) {
    a, b := swap("hello", "world")
    fmt.Println(a, b)
}

// 回傳一個 function
// 名稱為 foo 的 function 會回傳一個 function
// 這個回傳的 function 會回傳 int
func foo() func() int {
    return func() int {
        return 100
    }
}

func TestReturnFunc(t *testing.T) {
    bar := foo()            // bar 會是一個 function
    fmt.Printf("%T\n", bar) // func() int
    fmt.Println(bar())      // 100
}


// *回傳帶有命名的值
// 在 Go 中可以在 func 定義回傳 type 的地方定義要回傳的變數，
// 最後呼叫 return 的時候，該函式會自動去拿這兩個變數。這種做法稱作 naked return，
// 但最好只使用在內容不多的函式中，否則會嚴重影響程式的可讀性：

// 用來說明回傳的內容
func swapWithSpecifiedName(x, y string) (a, b string) {
    a = y
    b = x
    return
}

func TestNakedReturn(t *testing.T) {
    foo, bar := swapWithSpecifiedName("hello", "world")
    fmt.Println(foo, bar) // world hello
}

// IIFE (Immediately Invoked Function Expression)
func TestIIFE(t *testing.T) {
	slice := []string{"a", "a"}

	// 使用 IIFE 的寫法
	func(slice []string) {
		slice[0] = "b"
		slice[1] = "b"
	}(slice)

	fmt.Println(slice) // [b b]
}