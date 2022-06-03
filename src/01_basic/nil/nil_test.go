package nil

import (
	"fmt"
	"net/http"
	"testing"
)

// 【golang】nil的理解 https://www.gushiciku.cn/pl/poAb/zh-tw
/*
bool -> false
numbers -> 0
string    -> ""
pointers -> nil
slices -> nil
maps -> nil
channels -> nil
functions -> nil
interfaces -> nil
*/
/*
	這個 nil 到底是什麼呢？Go的文件中說到，_nil是預定義的識別符號，代表指標、通道、函式、介面、對映或切片的零值_，
	也就是預定義好的一個變數：

		type Type int
		var nil Type

	nil 並不是Go的關鍵字之一，你甚至可以自己去改變 nil 的值：

			var nil = errors.New("hi")

	這樣是完全可以編譯得過的，但是最好不要這樣子去做。

*/

type tree struct {
	v int
	l *tree
	r *tree
}
// first solution
func (t *tree) Sum() int {
	sum := t.v
	if t.l != nil {
		sum += t.l.Sum()
	}
	if t.r != nil {
		sum += t.r.Sum()
	}
	return sum
}

// 上面的程式碼有兩個問題，一個是程式碼重複 if t.l != nil {...}, 另一個是當 t 是 nil 的時候會panic：


// * 指標接收器的例子：
type person struct {}
// func sayHi(p *person) {
// 		fmt.Println("hi")
// }
func (p *person) sayHi() {
		fmt.Println("hi")
}

func TestTypeReceiver(t *testing.T) {
	var p *person
	p.sayHi() // hi
}


// * 對於指標物件的方法來說，就算指標的值為 nil 也是可以呼叫的，
// 基於此，我們可以對剛剛計算二叉樹和的例子進行一下改造：
func (t *tree) SumNew() int {
	if t == nil {
		return 0
	}
	return t.v + t.l.Sum() + t.r.Sum()
}

// 於 nil 指標，只需要在方法前面判斷一下就ok了，無需重複判斷。
// 換成列印二叉樹的值或者查詢二叉樹的某個值都是一樣的：
func(t *tree) String() string {
	if t == nil {
			return ""
	}
	return fmt.Sprint(t.l, t.v, t.r)
}

// nil receivers are useful: Find
func (t *tree) Find(v int) bool {
	if t == nil {
		return false
	}
	return t.v == v || t.l.Find(v) || t.r.Find(v)
}

// ----------------------------------------------------------------------------
// slices
// 一個為 nil 的slice，除了不能索引外，其他的操作都是可以的，
// 當你需要填充值的時候可以使用 append 函式，slice會自動進行擴充。
/*
---- nil slices ----
	var s []slice
	len(s)  // 0
	cap(s)  // 0
	for range s  // iterates zero times
	s[i]  // panic: index out of range
*/
// ----------------------------------------------------------------------------
// map
// 對於 nil 的map，我們可以簡單把它看成是一個只讀的map，不能進行寫操作，否則就會panic。
/*
---- nil map ----
	var m map[t]
	len(m)  // 0
	for range m // iterates zero times
	v, ok := m[i] // zero(u), false
	m[i] = x // panic: assignment to entry in nil map
*/

// nil 的map有什麼用
// 對於 NewGet 來說，我們需要傳入一個型別為map的引數，並且這個函式只是對這個引數進行讀取，我們可以傳入一個非空的值：
func NewGet(url string, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
			return nil, err
	}
	for k, v := range headers {
			req.Header.Set(k, v)
	}
	return req, nil
}

func TestReadOnlyMap(t *testing.T) {
	// 對於 NewGet 來說，我們需要傳入一個型別為map的引數，並且這個函式只是對這個引數進行讀取，我們可以傳入一個非空的值：
	NewGet("http://google.com", map[string]string{  "USER_AGENT": "golang/gopher",},)
	// 或者這樣傳：
	NewGet("http://google.com", map[string]string{})
	// 但是前面也說了，map的零值是 nil ，所以當 header 為空的時候，我們也可以直接傳入一個 nil ：
	NewGet("http://google.com", nil)
	// 是不是簡潔很多？所以，把 nil map作為一個只讀的空的map進行讀取吧。
}


// ----------------------------------------------------------------------------
// How to Gracefully Close Channels https://go101.org/article/channel-closing.html
// channel
/*
	nil channels
		var c chan
		t<- c      // blocks forever
		c <- x    // blocks forever
		close(c)  // panic: close of nil channel
*/

// 如果在外部呼叫中關閉了a或者b，那麼就會不斷地從a或者b中讀出0，這和我們想要的不一樣，我們想關閉a和b後就停止彙總了
func mergeOld(out chan<- int, a, b <-chan int) {
	for {
			select {
			case v := <-a:
					out <- v
			case v := <- b:
					out <- v
			}
	}
}

// 在知道channel關閉後，將channel的值設為nil，這樣子就相當於將這個select case子句停用了，因為 nil 的channel是永遠阻塞的。
func merge(out chan<- int, a, b <-chan int) {
	for a != nil || b != nil {
			select {
			case v, ok := <-a:
					if !ok {
							a = nil
							fmt.Println("a is nil")
							continue
					}
					out <- v
			case v, ok := <-b:
					if !ok {
							b = nil
							fmt.Println("b is nil")
							continue
					}
					out <- v
			}
	}
	fmt.Println("close out")
	close(out)
}

// ----------------------------------------------------------------------------
/*
interface
	interface並不是一個指標，它的底層實現由兩部分組成，一個是型別，一個值，也就是類似於：(Type, Value)。
	只有當型別和值都是 nil 的時候，才等於 nil 。
*/

/*
	func do() error {   // error(*doError, nil)
			var err *doError
			return err  // nil of type *doError
	}
	func main() {
			err := do()
			fmt.Println(err == nil) // false
	}

	-------------------------------------------------------------------------
	func do() *doError {  // nil of type *doError
		return nil
	}
	func wrapDo() error { // error (*doError, nil)
		return do()       // nil of type *doError
	}
	func main() {
    err := wrapDo()   // error  (*doError, nil)
    fmt.Println(err == nil) // false
	}
*/