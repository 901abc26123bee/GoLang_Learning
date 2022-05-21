package paniic_recover_test

import (
	"fmt"
	"testing"
)

// ---------------------------- defer ------------------------------------
// *	defer 這個 statement 可以用來在函式最終要回傳前被執行，類似 clean-up 的動作。
//    有三個最主要的原則：

// * 1. deferred function’s arguments are evaluated when the defer statement is evaluated.
// *	defer 中的值在執行時就寫入
// 		這個 defer 中所使用到的參數是在執行到的時候就已經被帶入：
func TestDeferValue(t *testing.T) {
	language := "Go"
	defer fmt.Print(language + "\n") // 2,  language = "Go"

	language = "Java"
	fmt.Print("Hello ") // 1

	// 輸出：Hello Go
}


// * 2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.
// *	defer 是後進先出（last-in-first-out）
// 		當有多個 defer 時，採用的是 後進先出（last-in-first-out） 原則：
func TestDeferExecutedOrder(t *testing.T) {
	language := "Go"
	defer fmt.Print(" to " + language + "\n") // 3,  language = "Go"

	language = "Java"
	defer fmt.Print("from " + language) // 2,  language = "Java"
	fmt.Print("Hello ") // 1

	// 輸出：Hello from Java to Go
}


// * 3. Deferred functions may read and assign to the returning function’s named return values.
// *	defer 可以用在函式回傳具名的變數
// 當函式回傳的值是具名的變數時，defer 可以去修改最終回傳出去的值：

// result 2
func c() (i int) {
	defer func() { i++ }()
	return 1
}

func TestDeferReturnFunction(t *testing.T) {
	fmt.Println(c()) // 2
}


// -------------------------- panic & recover ------------------------
/*
	Panic is a built-in function that stops the ordinary flow of control and begins panicking.
	When the function F calls panic, execution of F stops, any deferred functions in F are executed normally,
	and then F returns to its caller.
	To the caller, F then behaves like a call to panic.
	The process continues up the stack until all functions in the current goroutine have returned,
	at which point the program crashes. Panics can be initiated by invoking panic directly.
	They can also be caused by runtime errors, such as out-of-bounds array accesses.

	Recover is a built-in function that regains control of a panicking goroutine.
	Recover is only useful inside deferred functions.
	During normal execution, a call to recover will return nil and have no other effect.
	If the current goroutine is panicking,
	a call to recover will capture the value given to panic and resume normal execution.
*/
// Panic
// *	panic 會停止原本的 control flow，並進入 panicking。
// 		當函式 F 呼叫 panic 時，F 會停止執行，並且執行 F 中的 deferred function，最終 F 會回到呼叫它的函式（caller）。
// 		回到 Caller 之後，F 也會進入 panic，直到當前的 goroutine 都 returned，並導致程式 crash。
//		F() is panic --> F() defer()... --> F caller() is panic --> F caller() defer()... --> crash

// Recover
// *	recover 則是可以讓 panicking 的 goroutine 重新取得控制權，它只有在 deferred function 中執行是有用的。
// 		在一般函式運行的過程中，呼叫 recover 會回傳 nil 並且沒有任何效果，
// 		一旦當前的 goroutine panicking 時，recover 會攔截 panic 中給的 value，並讓函式回到正常的執行。

func f() {
	defer func() {
	// 可以取得 panic 的回傳值
			r := recover()
			if r != nil {
					fmt.Println("Recovered in f", r) // 8
			}
	}()

	fmt.Println("Calling g.") // 1
	g(0) // 2
	fmt.Println("Returned normally from g.") // 9
}

func g(i int) {
	if i > 3 {
			panicNumber := fmt.Sprintf("%v", i)
			fmt.Println("Panicking!", panicNumber) // 5
			// log.Fatalln(panicNumber)  // 如果使用 fatal 程式會直接終止，defer 也不會執行到
			panic(panicNumber) // 7
	}
	defer fmt.Println("Defer in g", i) // 6
	fmt.Println("Printing in g", i) // 3
	g(i + 1) // 4
}

func TestDeferPanicRecover(t *testing.T) {
	f()
	fmt.Println("Returned normally from f.")
	/*
	Calling g.
	Printing in g 0
	Printing in g 1
	Printing in g 2
	Printing in g 3
	Panicking! 4
	Defer in g 3
	Defer in g 2
	Defer in g 1
	Defer in g 0
	Recovered in f 4
	Returned normally from f.
	*/
}
