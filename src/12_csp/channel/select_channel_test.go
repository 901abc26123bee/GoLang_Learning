package channel_test

import (
	"fmt"
	"testing"
	"time"
)

/*
		select case
		和 switch case 不同，select case 中 case 接收的是 channel（而不是 boolean）。
		當程式執行到 select 的位置時，會阻塞（blocking）在那，直到有任何一個 case 收到 channel 傳來的資料後（除非有用 default）才會 unblock，
		因此通常會有另一個 channel 用來實作 Timeout 機制。

		select {
		case res := <-ch1: // 如果需要取用 channel 中的值
				fmt.Println(res)
		case <-ch1: // 如果不需要取用 channel 中的值
				fmt.Println("receive value")
		}

		⚠️ 和 for value := range channel 不同，select case 中，case 後的 channel 要記得加上 <- 把資料取出來。


		select case 運作的流程是：
	*	如果所有的 case 都沒有接收到 channel 傳來的資料，那麼 select 會一直阻塞（block）在那，直到有任何的 case 收到資料後（unblock）才會繼續執行
	*	如果同一時間有多個 case 收到 channel 傳來的資料（有多個 channel 同時 non-blocking），那個會從所有這些 non-blocking 的 cases 中隨機挑選一個，接著才繼續執行
    *   如果同時對全部 all select case channel 送資料，則會隨機選取到不同的 Channel。
    *   假設沒有送 value 進去 Channel 就會造成 panic, 執行後會發現變成 deadlock，造成 main 主程式爆炸，這時候可以直接用 default 方式解決此
*/

var start time.Time

func init() {
    start = time.Now()
}

func service1(c chan string) {
    time.Sleep(1 * time.Second)
    c <- "Hello from service 1"
}

func service2(c chan string) {
    time.Sleep(3 * time.Second)
    c <- "Hello from service 2"
}

func TestSelectUnBufferedChannel(t *testing.T) {
    fmt.Println("main() started", time.Since(start))

    chan1 := make(chan string)
    chan2 := make(chan string)

    go service1(chan1)
    go service2(chan2)

    fmt.Println("[main] select(blocking)")
    select {
    case res := <-chan1:
        fmt.Println("[main] get response from service 1", res, time.Since(start))
    case res := <-chan2:
        fmt.Println("[main] get response from service 2", res, time.Since(start))
    }

    fmt.Println("main() stopped", time.Since(start))

    // main() started 373.395µs
    // [main] select(blocking)
    // [main] get response from service 1 Hello from service 1 1.003474803s
    // main() stopped 1.003517131s
}

// ------------------------------------------------------------------------------------------------
/*
    隨機選取
    上面的程式中使用的是 unbuffered channel，所以對該 channel 任何的 send 或 receive 都會出現阻塞。
    我們可以使用 buffered channel 來模擬實際上 web service 處理回應的情況：

    由於 buffered channel 的 capacity 是 2，但傳入 channel 的 size 並沒有超過 2（沒有 overflow），
    因此程式會繼續執行而不會發生阻塞（non-blocking）

    當 buffered channel 中的有資料時，直到整個 buffer 都被清空為止前，
    從 buffered channel 讀取資料的動作都是 non-blocking 的，而且在下面的程式碼中又只讀取了一個值出來，
    因此整個 case 的操作都會是 non-blocking 的

    由於 select 中的所有 case 都是 non-blocking 的，因此 select 會從所有的 case 中隨機挑一個加以執行

*/

func TestSelectBufferedChannel(t *testing.T) {
    fmt.Println("main() started", time.Since(start))

    chan1 := make(chan string, 2)
    chan2 := make(chan string, 2)

    // buffered channel：因為 channel 中的資料沒有 overflow (> 2)，所以不會阻塞
    chan1 <- "Value 1"
    chan1 <- "Value 2"

    chan2 <- "Value 1"
    chan2 <- "Value 2"

    // buffered channel 中有資料時，讀取資料會是 non-blocking 的
    // 由於 select 中的 case 都是 non-blocking 的，因此會隨機挑選一個執行
    select {
    case res := <-chan1:
        fmt.Println("[main] get response from service 1", res, time.Since(start))
    case res := <-chan2:
        fmt.Println("[main] get response from service 2", res, time.Since(start))
    }

    fmt.Println("main() stopped", time.Since(start))

    // main() started 424.28µs
    // [main] get response from service 2 Value 1 454.178µs
    // main() stopped 457.036µs
}

// ------------------------------------------------------------------------------------------------
/*
    default
    default case 本身是非阻塞的（non-blocking），同時它也會使得 select statement 總是變成 non-blocking，
    也就是說，不論是 buffered 或 unbuffered channel 都會變成非阻塞的。

    當有任何資料可以從 channel 中取出時，select 就會執行該 case，但若沒有，就會直接進到 default case。
    簡單的說，當 channel 本身就有值時，就不會走到 default，但如果 channel 執行的當下沒有值，還需要等其他 goroutine 設值到 channel 的話，就會直接走到 default。

    如果沒有資料送進 channel，也就是註解掉 ch <- 1 的話，程式會出現 panic（fatal error: all goroutines are asleep - deadlock!），
    這是因為它認為應該要從 channel 取到值，但卻沒有得到任何東西，雖然加上 default 後可以解決，
    但會使得 select case 不會被阻塞住，導致還沒收到 channel 的訊息前，main goroutine 就執行完畢了：
*/
func TestSelectDefault(t *testing.T) {
    ch := make(chan int, 1)

    // 沒有送 value 進去 Channel 就會造成 panic --> 直接用 default to avoid
    // 主程式 main 就不會因為讀不到 channel value 造成整個程式 deadlock
    select {
    case <-ch:
        fmt.Println("random 01")
    case <-ch:
        fmt.Println("random 02")
    default:
        fmt.Println("exit")
    }
}

// ------------------------------------------------------------------------------------------------
/*
    Timeout 超時機制：使用 time.After
    單純使用 default case 並不是非常有用，有時我們希望的是有 timeout 的機制，
    也就是超過一定時間後，沒有收到任何回應時，才做預設的行為，這時候我們可以使用 time.After 來完成：

    time.After() 和 time.Tick() 都是會回傳 time.Time 型別的 receive channel（<- channel）
*/
func TestSelectWithTimeOut(t *testing.T) {
    ch := make(chan int)

    select {
    case <-ch:
        fmt.Println("receive value from channel")

    // 超過一秒沒有收到主要 channel 的 value，就會收到 time.After 送來的訊息
    // time.After 是回傳 chan time.Time。
    case <-time.After(1 * time.Second):
        fmt.Println("timeout after 1 second")
    }

    // timeout after 1 secon
}

func TestSelectWithTimeOut2(t *testing.T) {
    fmt.Println("main() started", time.Since(start))

    chan1 := make(chan string, 1)
    chan2 := make(chan string, 1)

    go service1(chan1)
    go service2(chan2)

    select {
    case res := <-chan1:
        fmt.Println("get response from service 1", res, time.Since(start))
    case res := <-chan2:
        fmt.Println("get response from service 2", res, time.Since(start))
    case <-time.After(2 * time.Second):
        fmt.Println("No response received", time.Since(start))
    }
    fmt.Println("main() stopped", time.Since(start))

    // main() started 332.171µs
    // get response from service 1 Hello from service 1 1.004916005s
    // main() stopped 1.005034108s
}

func TestSelectWithTimeOut3(t *testing.T) {
    timeout := make(chan bool, 1)
    go func() {
        time.Sleep(2 * time.Second)
        timeout <- true
    }()
    ch := make(chan int)
    select {
    case <- ch:
    case <- timeout:
        fmt.Println("timeout 01")
    }
    // timeout 01
}

// ------------------------------------- empty select -----------------------------------------
/*
    empty select
    如同 for{} 迴圈可以不帶任何條件一樣，select {} 也可以不搭配 case 使用（稱作，empty select）。

    從前面的例子中可以看到，因為 select statement 會一直阻塞（blocking），直到其中一個 case unblocks 時，才會繼續往後執行，
    * 但因為 empty select 中並沒有任何的 case statement，因此 main goroutine 將會永遠阻塞在那，如果沒有其他 goroutine 可以持續運行的話，最終導致 deadlock。
*/
func service() {
    fmt.Println("Hello from service ")
}

func TestEmptySelect(t *testing.T) {
    fmt.Println("main() started")

    go service()

    // 這個 select 會永遠 block 在這
    select {}

    fmt.Println("main() stopped")

    // main() started
    // Hello from service
    // fatal error: all goroutines are asleep - deadlock!
}

// * 如果在 main goroutine 使用 empty select 後，main goroutine 將會完全阻塞，
// * 需要靠其他的 goroutine 持續運作才不至於進入 deadlock：
func service3() {
    for {
        fmt.Println("Hello from service ")
        time.Sleep(time.Millisecond * 400)
    }
}

func TestEmptySelect2(t *testing.T) {
    fmt.Println("main() started")

    go service3()

    // 這個 select 會永遠 block 在這
    select {}

    fmt.Println("main() stopped")

    // main() started
    // Hello from service
    // Hello from service
    // ...
}

func init() {
    start = time.Now()
}

func service4() {
    for {
        fmt.Println("Hello from service1 ", time.Since((start)))
        time.Sleep(time.Millisecond * 500)
    }
}

func service5() {
    for {
        fmt.Println("Hello from service2 ", time.Since((start)))
        time.Sleep(time.Millisecond * 700)
    }
}

func TestEmptySelect3(t *testing.T) {
    fmt.Println("main() started")

    go service4()
    go service5()

    // 這個 select 會永遠 block 在這，service1 和 service2 輪流輸出訊息
    select {}

    fmt.Println("main() stopped")  // 這行不會被執行到

    // main() started
    // Hello from service2  457.154µs
    // Hello from service1  492.037µs
    // Hello from service1  502.024131ms
    // Hello from service2  704.660326ms
    // Hello from service1  1.006027462s
    // ...
}


// ----------------------------- Go 語言使用 Select 四大用法 -----------------------------------
// [Go 語言使用 Select 四大用法](https://blog.wu-boy.com/2019/11/four-tips-with-select-in-golang/)
// 另外透過 empty select 導致 main goroutine 阻塞的這種方式，可以在 server 啟動兩個不同的 service：

// 檢查 channel 是否已滿
// 判斷是否超過 channel 的 buffer size
func TestSelectTrick(t *testing.T) {
    // STEP 1：建立一個只能裝 buffer size 為 1 資料
    ch := make(chan int, 2)
    ch <- 1

    select {
    case ch <- 2:
        fmt.Println("channel value is", <-ch)
        fmt.Println("channel value is", <-ch)
    default:
        // ch 中的內容超過 1 時，但若把 channel buffer size 的容量改成 2，就不會走到 default
        fmt.Println("channel blocking")
    }
    // ch := make(chan int, 1)
    // channel blocking

    // ch := make(chan int, 2)
    // channel value is 1
    // channel value is 2
}

// 使用 for + select 讀取多個 channel 的 value
func TestSelectTrick2(t *testing.T) {
    tick := time.Tick(100 * time.Millisecond)
    boom := time.After(500 * time.Millisecond)

    for {
        select {
        case <-tick:
            fmt.Println("tick.")
        case <-boom:
            fmt.Println("BOOM!")
            return // 如果沒有 return 的話程式將不會結束，一直卡在 for loop 中
        default:
            fmt.Println("    .")
            time.Sleep(50 * time.Millisecond)
        }
    }
    /*
        .
        .
    tick.
        .
        .
    tick.
        .
        .
    tick.
        .
        .
    tick.
        .
        .
    BOOM!
    */
}

// 如果你有多個 channel 需要讀取，而讀取是不間斷的，就必須使用 for + select 機制來實現
// 要結束 for 或 select 都需要透過 break 來結束
func TestSelectTrick3(t *testing.T) {
    ch1 := make(chan string)
    ch2 := make(chan int, 1)

    defer func() {
        fmt.Println("------ In defer ------")
        close(ch1)
        close(ch2)
    }()

    i := 0

    //STEP 1：建立一個 Go Routine
    go func() {
        fmt.Println("In Go Routine")

        //  STEP 2：透過 for loop 來不停來監控不同 channels 傳回來的資料
    LOOP:
        for {
            //  STEP 3：透過 sleep 讓它每 500 毫秒檢查有無 channel 傳訊息進來
            time.Sleep(500 * time.Millisecond)
            i++
            fmt.Printf("In Go Routine, i: %v, time: %v \n", i, time.Now().Unix())

            // STEP 4：透過 select 判斷不同 channel 傳入的資料
            select {
            case m := <-ch1: // STEP 6：當收到 channel 1 傳入的資料時，就 break
                fmt.Printf("In Go Routine, get message from channel 1: %v \n", m)
                break LOOP // 要在 select 區間直接結束掉 for 迴圈，只能使用 break variable 來結束

            case m := <-ch2: // STEP 7：當收到 channel 2 傳入的資料時，單純輸出訊息
                fmt.Printf("In Go Routine, get message from channel 2: %v\n", m)

            default: // STEP 5：檢查時，如果沒有 channel 丟資料進 channel 則走 default
                fmt.Println("In Go Routine to DEFAULT")
            }
        }
    }()

    ch2 <- 666 // STEP 8：在 sleep 前將訊息丟入 channel2

    fmt.Println("Start Sleep")

    // STEP 9：雖然這裡 sleep，但 go routine 中的 for 迴圈仍然不斷在檢查有無收到訊息
    time.Sleep(4 * time.Second)

    fmt.Println("After Sleep: send value to channel")

    // STEP 10：四秒後把 "stop" 傳進 channel 1，for 迴圈收到訊息後 break
    ch1 <- "stop"

    fmt.Println("------ End ------")

    /*
        Start Sleep
        In Go Routine
        In Go Routine, i: 1, time: 1653704739
        In Go Routine, get message from channel 2: 666
        In Go Routine, i: 2, time: 1653704739
        In Go Routine to DEFAULT
        In Go Routine, i: 3, time: 1653704740
        In Go Routine to DEFAULT
        In Go Routine, i: 4, time: 1653704740
        In Go Routine to DEFAULT
        In Go Routine, i: 5, time: 1653704741
        In Go Routine to DEFAULT
        In Go Routine, i: 6, time: 1653704741
        In Go Routine to DEFAULT
        In Go Routine, i: 7, time: 1653704742
        In Go Routine to DEFAULT
        After Sleep: send value to channel
        In Go Routine, i: 8, time: 1653704742
        In Go Routine, get message from channel 1: stop
        ------ End ------
        ------ In defer ------
    */
}