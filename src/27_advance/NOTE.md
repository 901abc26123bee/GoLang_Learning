## 01 WorkSpace 和GOPATH
### 1. Go 語言源碼的組織方式
1. 在工作區中，一個代碼包的導入路徑實際上就是從 src 子目錄，到該包的實際存儲位置的相對路徑
2. Go 語言源碼的組織方式就是以環境變量 GOPATH、工作區、src 目錄和代碼包為主線的。一般情況下，Go 語言的 source code file 都需要被存放在環境變量 GOPATH 包含的某個工作區（目錄）中的 src 目錄下的某個代碼包（目錄）中。

### 2. 了解源碼安裝後的結果
1. go程式碼文件以及安裝後的結果文件都會放到某個工作區的 src 子目錄下。
2. 那麼在安裝後如果產生了歸檔文件（xxx.a : 編譯後生成的靜態庫文件），就會放進該工作區的 pkg 子目錄；
   如果產生了可執行文件，就可能會放進該工作區的 bin 子目錄。
3. go程式碼文件會以代碼包的形式組織起來，一個代碼包其實就對應一個目錄。安裝某個代碼包而產生的歸檔文件是與這個代碼包同名的。

放置它的相對目錄就是該代碼包的導入路徑的直接父級。比如，一個已存在的代碼包的導入路徑是
`github.com/labstack/echo`, 那麼執行命令 `go install github.com/labstack/echo`
生成的歸檔文件的相對目錄就是 github.com/labstack， 文件名為 echo.a。

歸檔文件的相對目錄與 pkg 目錄之間還有一級目錄，叫做平台相關目錄。平台相關目錄的名稱是由 build（也稱“構建”）的目標操作系統、下劃線和目標計算架構的代號組成的。比如，構建某個代碼包時的目標操作系統是 Linux，目標計算架構是 64 位的，那麼對應的平台相關目錄就是 linux_amd64。
因此，上述代碼包的歸檔文件就會被放置在當前工作區的子目錄 pkg/linux_amd64/github.com/labstack 中。

```
GOPATH
  |--- WorkSpace1
  |--- WorkSpace2 ----- src
                   |      |--- github.com/labstack/echo
                   |      |--- ...
                   |
                   |
                   |--- pkg
                   |      |--- linux_amd64 (accroding to your computer)
                   |                    |--- github.com/labstack
                   |                    |                     |- echo.a
                   |                    |--- ...
                   |--- bin
```

### 3. 理解構建和安裝 Go 程序的過程
構建使用命令go build，安裝使用命令go install。構
建和安裝代碼包的時候都會執行編譯、打包等操作，並且，這些操作生成的任何文件都會先被保存到某個臨時的目錄中。

- build `source code file` (dependency code)，結果文件只會存在於臨時目錄中。這裡的構建的主要意義在於檢查和驗證。

- build `entry file` (with main(){}, entry point of the executable programs)，結果文件會被搬運到 `entry file` 所在的目錄中(under same directory)。

安裝操作會先執行構建，然後還會進行鏈接操作，並且把結果文件搬運到指定目錄。
- install `source code file` (dependency code)，結果文件會被 move to 它所在工作區的 pkg 目錄下的某個子目錄中。
- install `entry file` (with main(){}, entry point of the executable programs)，那麼結果文件會被 move to 它所在工作區的 bin 目錄中，或者環境變量GOBIN指向的目錄中。

### go build
- `-a` : 不但目標代碼包總是會被編譯，它依賴的代碼包也總會被編譯，即使依賴的是標準庫中的代碼包也是如此。
- `-i`: 如果不但要編譯依賴的代碼包，還要安裝它們的歸檔文件(xxx.a:編譯後生成的靜態庫文件)。

怎麼確定哪些代碼包被編譯了
- 運行go build命令時加入標記-x，這樣可以看到go build命令具體都執行了哪些操作。
  另外也可以加入標記-n，這樣可以只查看具體操作而不執行它們。
- 運行go build命令時加入標記-v，這樣可以看到go build命令編譯的代碼包的名稱。它在與-a標記搭配使用時很有用。


命令go get會自動從一些主流公用代碼倉庫（比如 GitHub）下載目標代碼包，並把它們安裝到環境變量GOPATH包含的第 1 工作區的相應目錄中。如果存在環境變量GOBIN，那麼僅包含命令go程式碼文件的代碼包會被安裝到GOBIN指向的那個目錄。

常用的幾個標記有下面幾種。

-u：下載並安裝代碼包，不論工作區中是否已存在它們。
-d：只下載代碼包，不安裝代碼包。
-fix：在下載代碼包後先運行一個用於根據當前 Go 語言版本修正代碼的工具，然後再安裝代碼包。
-t：同時下載測試所需的代碼包。
-insecure：允許通過非安全的網絡協議下載和安裝代碼包。 HTTP 就是這樣的協議。

Go 語言官方提供的go get命令是比較基礎的，其中並沒有提供依賴管理的功能。目前 GitHub 上有很多提供這類功能的第三方工具，比如glide、gb以及官方出品的dep、vgo等等，它們在內部大都會直接使用go get。

## 02 命令源碼文件 `entry file` (with main(){}, entry point of the executable programs)
- 命令源碼文件是程序的運行入口，是每個可獨立運行的程序必須擁有的。
- Module Programming，會將代碼拆分到多個文件，甚至拆分到不同的代碼包中。但無論怎樣，對於一個獨立的程序來說，命令源碼文件永遠只會也只能有一個。如果有與命令源碼文件同包的源碼文件，那麼它們也應該聲明屬於main包。

1. 命令源碼文件怎樣接收參數
    - Go 語言標準庫中有一個代碼包專門用於接收和解析命令參數。這個代碼包的名字叫flag。
      調用flag包的StringVar函數的代碼
      還有一個與flag.StringVar函數類似的函數，叫flag.String。
      這兩個函數的區別是，後者會直接返回一個已經分配好的用於存儲命令參數值的地址。
    - 函數flag.Parse用於真正解析命令參數，並把它們的值賦給相應的變量
      對該函數的調用必須在所有命令參數存儲載體的聲明和設置之後，並且在讀取任何命令參數值之前進行。

2. 怎樣在運行命令源碼文件的時候傳入參數，又怎樣查看參數的使用說明
    如果想查看該命令源碼文件的參數說明，可以這樣做：
    $ go run demo2.go --help
3. 怎樣自定義命令源碼文件的參數使用說明
    - 最簡單的一種方式就是對變量flag.Usage重新賦值。 flag.Usage的類型是func()，即一種無參數聲明且無結果聲明的函數類型。
      flag.Usage變量在聲明時就已經被賦值了，所以我們才能夠在運行命令 `go run xxx.go --help` 時看到正確的結果
      - 我們在調用flag包中的一些函數（比如StringVar、Parse等等）的時候，實際上是在調用flag.CommandLine變量的對應方法。
    - flag.CommandLine相當於默認情況下的命令參數容器。所以，通過對flag.CommandLine重新賦值，可以定制當前命令源碼文件的參數使用說明。這樣做的好處依然是更靈活地定制命令參數容器。

## 03 源碼文件
- 源碼文件是不能被直接運行的源碼文件(xxx.go)，它僅用於存放program實體，這些program實體可以被其他代碼使用（只要遵從 Go 語言規範的話）。
- 源碼文件(xxx.go)聲明的包名可以與其所在目錄的名稱不同，只要這些文件聲明的 package name一致就可以。
- 同目錄下的源碼文件的代碼包聲明語句要一致。如果目錄中有命令源碼文件，那麼其他種類的源碼文件也應該聲明屬於main包。
- 源碼文件(xxx.go) package name 可以與其所在的目錄的名稱不同。while `go build`，生成的結果文件的主名稱與其父目錄的名稱一致。

- 應該讓聲明的包名與其父目錄的名稱一致。
- 通過名稱，Go 語言自然地把程序實體的訪問權限劃分為了包級私有的和公開的。對於包級私有的程序實體，即使你導入了它所在的代碼包也無法引用到它。
- internal包中聲明的公開代碼實體(with capital Letter)僅能被該代碼包的直接父包及其子包中的代碼引用。

## 04 | 程序實體的那些事兒（上）
### variables declarations
- Go 語言中的類型推斷 `var x = "hello"`, 在編譯期自動解釋表達式類型的能力, 表達式類型就是對表達式進行求值後得到結果的類型
- 短變量聲明的用法 `x := "hello"`, Go 語言的類型推斷再加上一點點語法糖。只能在函數體內部使用短變量聲明。編寫if、for或switch語句的時候，我們經常把它安插在初始化子句中，並用來聲明一些臨時的變量。而相比之下，第一種方式更加通用，它可以被用在任何地方。
- Go 語言是靜態類型的，所以一旦在初始化變量時確定了它的類型，之後就不可能再改變。這就避免了在後面維護的一些問題(得代碼重構變得更加容易)。另外，這種類型的確定是在編譯期完成的，因此不會對程序的運行效率產生任何影響。

### variables redeclarations
- 可以對同一個代碼塊中({...}, func, gobal context)的變量進行重聲明。
- 變量重聲明的前提條件如下。
  1. 由於變量的類型在其初始化時就已經確定了，所以對它再次聲明時賦予的類型必須與其原本的類型相同，否則會產生編譯錯誤。
  2. 變量的重聲明只可能發生在某一個代碼塊中。
  3. 變量的重聲明只有在使用短變量聲明時才會發生，否則也無法通過編譯。
  4. 被“聲明並賦值”的變量必須是多個，並且其中至少有一個是新的變量。這時我們才可以說對其中的舊變量進行了重聲明。
  5. 變量重聲明其實算是一個語法糖（或者叫便利措施）。它允許我們在使用短變量聲明時不用理會被賦值的多個變量中是否包含舊變量。
For Example:
```go
func main() {
	var err error
	n, err := io.WriteString(os.Stdout, "Hello, everyone!\n") // 這裡對`err`進行了重聲明。
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("%d byte(s) were written.\n", n)
}
```

## 05 | 程序實體的那些事兒（）


## 07 | array和slice
- `可以把slice看做是對array的一層簡單的封裝`，因為在`每個slice的底層數據結構中，一定會包含一個array`。
- array可以被叫做slice的底層array，而`slice也可以被看作是對array的某個連續片段的引用`。
- Go 語言的slice類型屬於引用類型，同屬引用類型的還有map據類型以及struct類型。
- 在 Go 語言中，如果傳遞的值是引用類型的，那麼就是“傳引用”。如果傳遞的值是值類型的，那麼就是“傳值”。
- 在array和slice之上都可以應用索引表達式，得到的都會是某個元素。我們在它們之上也都可以應用slice表達式，也都會得到一個新的slice。





## 09 | map的操作和約束
- map key type 不能是哪些類型？典型回答是：map的鍵類型不可以是function類型、map類型和slice類型。
  - Go 語言規範規定，在`鍵類型的值之間必須可以施加操作符==和!=`。(哈希碰撞: 不同值的哈希值是可能相同的)
    換句話說，鍵類型的值必須要支持判等操作。由於`函數類型、map類型和slice類型的值並不支持判等操作，所以map的鍵類型不能是這些類型`。
  - 另外，如果`鍵的類型是接口類型的，那麼key的實際類型也不能是上述三種類型`，否則在程序運行過程中會引發 panic（即運行時恐慌）
    ```go
    var badMap2 = map[interface{}]int{
      "1":   1,
      []int{2}: 2, // 這樣的聲明躲過了 Go 語言編譯器的檢查(因為從語法上說，這樣做是可以的),but result in panic。
      3:    3,
    }
    ```
    如果鍵的類型是數組類型，那麼還要確保該類型的元素類型不是函數類型、map類型或slice類型。
    鍵的類型是結構體類型，那麼還要保證其中字段的類型的合法性。比如map[[1][2][3][]string]int。

- key type 長度越短求哈希越快。類型的寬度是指它的單個值需要佔用的字節數。比如，bool、int8和uint8類型的一個值需要佔用的字節數都是1，因此這些類型的寬度就都是1(8bit)。
- 對array類型的值求哈希實際上是`依次求得它的每個元素的哈希值並進行合併`，所以速度就`取決於它的元素類型以及它的長度`。對一個數組來說，我可以任意改變其中的元素值，但在變化前後，它卻代表了兩個不同的key。
- 對struct類型的值求哈希實際上就是對它的所有字段值求哈希並進行合併，所以關鍵在於它的各個字段的類型以及字段的數量。
- 對於接口類型，具體的哈希算法，則由值的實際類型決定。

- 由於map是引用類型，當僅聲明而不初始化一個map的變量的時候，它的值會是nil。
- 在這樣僅聲明map變量上試圖通過key獲取對應的元素值，或者添加key-value，會成功嗎？
  `除了添加key-value，在一個值為nil的map上做任何操作都不會引起錯誤`。
  當試圖在一個值為nil的map中添加key-value的時候，Go 語言的運行時系統就會立即拋出一個 panic。
- 非原子操作需要加鎖， map並發讀寫需要加鎖，map操作不是並發安全的，判斷一個操作是否是原子的可以使用 `go run race `命令做數據的競爭檢測

## 10 | channel的基本操作
### 基本特性。
  - FIFO 的設計
  1. 在同一時刻，`Go 語言的運行時系統只會執行對同一個channel的任意個發送操作中的某一個 or 對同一個channel的任意個接收操作中的某一個`。即使這些操作是並發執行的也是如此。這裡所謂的並發執行，
  - 不同的 goroutine 之中，並有機會在同一個時間段內被執行。另外，對於channel中的同一個元素值來說，`發送操作和接收操作之間也是互斥的`。
  - 進入channel的並不是在接收操作符右邊的那個元素值，而是它的副本。元素值從channel取出時會被移動。這個移動操作實際上包含了兩步，
    1. 第一步是生成正在channel中的這個元素值的副本，並準備給到接收方，
    2. 第二步是刪除在channel中的這個元素值。

  2. `channel處理元素值時都是一氣呵成的，絕不會被打斷(“不可分割”)`。
    - 發送操作要么還沒複製元素值，要么已經復製完畢，絕不會出現只複製了一部分的情況。
    - 接收操作在準備好元素值的副本之後，一定會刪除掉channel中的原值，絕不會出現channel中仍有殘留的情況。這既是為了保證channel中元素值的完整性，也是為了保證channel操作的唯一性。
    - 對於channel中的同一個元素值來說，它只可能是某一個發送操作放入的，同時也只可能被某一個接收操作取出。

  3. 一般情況下，`發送操作包括了“複製元素值”`和`“放置副本到channel內部”`這兩個步驟。在這兩個步驟完全完成之前，發起這個發送操作的那句代碼會一直阻塞在那裡。也就是說，在它之後的代碼不會有執行的機會，直到這句代碼的阻塞解除。更細緻地說，`在channel完成發送操作之後，運行時系統會通知這句代碼所在的 goroutine，以使它去爭取繼續運行代碼的機會。`
    另外，`接收操作通常包含了“複製channel內的元素值”“放置副本到接收方”“刪掉原值”三個步驟`。在所有這些步驟完全完成之前，發起該操作的代碼也會一直阻塞，直到該代碼所在的 goroutine 收到了運行時系統的通知並重新獲得運行機會為止。如此阻塞代碼其實就是`為了實現操作的互斥和元素值的完整`。

### 發送操作和接收操作在什麼時候可能被長時間的阻塞？
1. buffered channel
  - 如果`channel已滿`，那麼`對它的所有發送操作都會被阻塞，直到channel中有元素值被接收走`。
    這時，`channel會優先通知最早因此而等待的、那個發送操作所在的 goroutine，後者會再次執行發送操作`。由於發送操作在這種情況下被阻塞後，`它們所在的 goroutine 會順序地進入channel內部的發送等待隊列，所以通知的順序是公平的`。
  - 如果`channel已空`，那麼`對它的所有接收操作都會被阻塞，直到channel中有新的元素值出現`。
    這時，channel會通知最早等待的那個接收操作所在的 goroutine，並使它再次執行接收操作。因此而等待的、所有接收操作所在的 goroutine，都會按照先後順序被放入channel內部的接收等待隊列。
2. unbuffered channel
  - `無論是發送操作還是接收操作，一開始執行就會被阻塞，直到配對的操作也開始執行，才會繼續傳遞`。
    - `unbuffered channel是在用同步的方式傳遞數據`。也就是說，只有收發雙方對接上了，數據才會被傳遞。
    - `數據是直接從發送方復製到接收方的`，中間並不會用unbuffered channelchannel做中轉。
    - 相比之下，`buffered channel則在用異步的方式傳遞數據`。在大多數情況下，buffered channel會作為收發雙方的中間件。元素值會先從發送方復製到buffered channel，之後再由buffered channel複製給接收方。但是，`當發送操作在執行的時候發現空的channel中，正好有等待的接收操作，那麼它會直接把元素值複製給接收方`。
3. 由於錯誤使用channel而造成的阻塞
  - `對於值為nil的channel`，不論它的具體類型是什麼，對它的`發送操作和接收操作都會永久地處於阻塞狀態`。它們所屬的 goroutine 中的任何代碼，都不再會被執行。注意，由於`channel類型是引用類型，所以它的零值就是nil`。只聲明該類型的變量但沒有用make函數對它進行初始化時，該變量的值就會是nil。


### 發送操作和接收操作在什麼時候會引發 panic？
1. 對於一個`已初始化，但並未關閉的channel`來說，`收發操作一定不會引發 panic`。但是`channel一旦關閉，再對它進行發送操作，就會引發 panic`。
2. 如果`試圖關閉一個已經關閉了的channel`，也會引發 panic。
  - `接收操作`是`可以感知到channel的關閉的，並能夠安全退出`。當把接收表達式`data, ok := <-ch`的結果同時賦給兩個變量時，第二個變量的類型就是一定bool類型。`它的值如果為false就說明channel已經關閉，並且再沒有元素值可取了`。
  - `如果channel關閉時，裡面還有元素值未被取出`，那麼接收表達式的第一個結果(data)，`仍會是channel中的某一個元素值`，而`第二個結果(ok)值一定會是true`。因此，通過接收表達式的第二個結果值，來`判斷channel是否關閉是可能有延時`的。
    - 由於channel的收發操作有上述特性，所以除非有特殊的保障措施，`不要讓接收方關閉channel，而應讓發送方做這件事`。

- chanel有點像socket的同步阻塞模式，只不過channel的發送端和接收端共享一個緩衝，socket則是發送這邊有發送緩衝，接收這邊有接收緩衝，而且socket接收端如果先close的話，發送端再發送數據的也會引發panic（linux上會觸發SIG_PIPE信號，不處理程序就崩潰了）。

## 11 | channel的高級玩法
### unidirectional channels
- 可以使用帶range子句的for語句從通道中獲取數據，也可以通過select語句操縱通道。
1. `<-chan` receive only channel  (only receive data from goroutine)
2. `chan<-` send only channel (only send data to goroutine)
- 最主要的用途就是約束其他代碼的行為, 這種約束一般會出現在接口類型聲明中的某個方法定義上
  ```go
  type Notifier interface {
    SendInt(ch chan<- int)
  }
  ```
  在調用SendInt函數的時候，只需要把一個元素類型匹配的雙向通道傳給它就行了，因為 Go 語言在這種情況下會自動地把雙向通道轉換為函數所需的單向通道。
  ```go
  intChan1 := make(chan int, 3)
  SendInt(intChan1)
  ```

- use unidirectional channels as return type, 意味著得到該通道的程序，只能從通道中接收元素值。這實際上就是對函數調用方的一種約束了。
  ```go
  func getIntChan() <-chan int {
    num := 5
    ch := make(chan int, num)
    for i := 0; i < num; i++ {
      ch <- i
    }
    close(ch)
    return ch
  }
  ```
- 在函數類型中使用了單向通道，那麼就相等於在約束所有實現了這個函數類型的函數。
  ```go
  intChan2 := getIntChan()
  for elem := range intChan2 {
    fmt.Printf("The element in intChan2: %v\n", elem)
  }
  ```
  這裡的for語句也可以被稱為帶有range子句的for語句。
  1. 這樣一條for語句會不斷地嘗試從intChan2種取出元素值，即使intChan2被關閉，它也會在取出所有剩餘的元素值之後再結束執行。
  2. 當intChan2中沒有元素值時，它會被阻塞在有for關鍵字的那一行，直到有新的元素值可取。
  3. 假設intChan2的值為nil，那麼它會被永遠阻塞在有for關鍵字的那一行。

### select
- select語句只能與通道聯用，它一般由若干個分支組成。每次執行這種語句的時候，一般只有一個分支中的代碼會被運行。
  由於select語句是專為通道而設計的，所以每個case表達式中都只能包含操作通道的表達式，比如接收表達式。如果把接收表達式的結果賦給變量的話，還可以把這裡寫成賦值語句或者短變量聲明。
  ```go
  // 準備好幾個通道。
  intChannels := [3]chan int{
    make(chan int, 1),
    make(chan int, 1),
    make(chan int, 1),
  }
  // 隨機選擇一個通道，並向它發送元素值。
  index := rand.Intn(3)
  fmt.Printf("The index: %d\n", index)
  intChannels[index] <- index
  // 哪一個通道中有可取的元素值，哪個對應的分支就會被執行。
  select {
    case <-intChannels[0]:
      fmt.Println("The first candidate case is selected.")
    case <-intChannels[1]:
      fmt.Println("The second candidate case is selected.")
    case elem := <-intChannels[2]:
      fmt.Printf("The third candidate case is selected, the element is %d.\n", elem)
    default:
      fmt.Println("No candidate case is selected!")
  }
  ```
在使用select語句的時候，需要注意:
1. with 默認分支，if all cases are blocked，或者說都沒有滿足求值的條件，那麼默認分支就會被選中並執行。
2. without 默認分支，一旦所有的case表達式都沒有滿足求值條件，那麼select語句就會被阻塞。直到至少有一個case表達式滿足條件為止。一旦發現某個通道關閉了，我們就應該及時地屏蔽掉對應的分支或者採取其他措施
  ```go
  intChan := make(chan int, 1)
  time.AfterFunc(time.Second, func() {
    close(intChan)
  })
  select {
    case _, ok := <-intChan:
      if !ok {
        fmt.Println("The candidate case is closed.")
        break
      }
    fmt.Println("The candidate case is selected.")
  }
  ```
3. select語句只能對其中的每一個case表達式各求值一次。所以，如果想連續或定時地操作其中的通道的話，需要通過在for語句中嵌入select語句的方式實現。但這時要注意，簡單地在select語句的分支中使用break語句，只能結束當前的select語句的執行，而並不會對外層的for語句產生作用。這種錯誤的用法可能會讓這個for語句無休止地運行下去。


select 規則:
1. 對於每一個case表達式，都至少會包含一個代表發送操作的發送表達式或者一個代表接收操作的接收表達式，同時也可能會包含其他的表達式。
2. select語句包含的候選分支中的`case表達式都會在該語句執行開始時先被求值`，並且`求值的順序是依從代碼編寫的順序從上到下`的。結合上一條規則，在select語句開始執行時，排在最上邊的候選分支中`最左邊的表達式會最先被求值，然後是它右邊的表達式`。僅當最上邊的候選分支中的所有表達式都被求值完畢後，從上邊數第二個候選分支中的表達式才會被求值，順序同樣是從左到右，然後是第三個候選分支、第四個候選分支，以此類推。
3. 對於每一個case表達式，如果其中的`發送表達式或者接收表達式在被求值時，相應的操作正處於阻塞狀態，那麼對該case表達式的求值就是不成功的`。在這種情況下,這個case表達式所在的候選分支是不滿足選擇條件的。
4. `僅當select語句中的所有case表達式都被求值完畢後，它才會開始選擇候選分支`。這時候，它只會挑選滿足選擇條件的候選分支執行。如果所有的候選分支都不滿足選擇條件，那麼默認分支就會被執行。如果這時沒有默認分支，那麼select語句就會立即進入阻塞狀態，直到至少有一個候選分支滿足選擇條件為止。一旦有一個候選分支滿足選擇條件，select語句（或者說它所在的 goroutine）就會被喚醒，這個候選分支就會被執行。
5. 如果select語句發現同時有多個候選分支滿足選擇條件，那麼它就會用一種`偽隨機的算法在這些分支中選擇一個並執行`。即使select語句是在被喚醒時發現的這種情況，也會這樣做。

一條select語句中只能夠有一個默認分支。並且，默認分支只在無候選分支可選時才會被執行，這與它的編寫位置無關。

select語句的每次執行，包括case表達式求值和分支選擇，都是獨立的。不過，至於它的執行是否是並發安全的，就要看其中的case表達式以及分支中，是否包含並發不安全的代碼了。
我把與以上規則相關的示例放在 demo25.go 文件中了。你一定要去試運行一下，然後嘗試用上面的規則去解釋它的輸出內容。

- 如果在select語句中發現某個通道已關閉，那麼應該怎樣屏蔽掉它所在的分支？
  1. 把nil賦給代表了這個通道的變量就可以了。如此一來，對於這個通道（那個變量）的發送操作和接收操作就會永遠被阻塞。
  2. 可以把這個channel重新賦值成為一個長度為0的非緩衝通道，這樣這個case就一直被阻塞了
  ```go
  for {
    select {
    case _, ok := <-ch1:
      if !ok {
        ch1 = nil // 1
        // ch1 = make(chan int) // 2
      }
      case ..... :
    ////
    default:
    //// 
    }
  }
  ```
- 在select語句與for語句聯用時，怎樣直接退出外層的for語句？
  1. 可以用 break和標籤配合使用，直接break出指定的循環體，
  2. 或者goto語句直接跳轉到指定標籤執行
  3. return
  ```go
  // 1. break配合標籤：
  ch1 := make(chan int, 1)
  time.AfterFunc(time.Second, func() { close(ch1) })
  loop:
    for {
      select {
      case _, ok := <-ch1:
        if !ok {
          break loop
        }
      fmt.Println("ch1")
      }
    }
  fmt.Println("END")

  // 2. goto配合標籤：
  ch1 := make(chan int, 1)
  time.AfterFunc(time.Second, func() { close(ch1) })
  for {
    select {
    case _, ok := <-ch1:
      if !ok {
        goto loop
      }
    fmt.Println("ch1")
    }
  }
  loop:
    fmt.Println("END")
  ```

## 13 | 結構體及其方法的使用
- 結構體類型也可以不包含任何字段，可以為這些類型關聯上一些方法，把方法看做是函數的特殊版本。
- (functional programming)函數則是獨立的程序實體。可以聲明有名字的函數，也可以聲明沒名字的函數，還可以把它們當做普通的值傳來傳去。我們能把具有相同簽名的函數抽象成獨立的函數類型，以作為一組輸入、輸出（或者說一類邏輯組件）的代表。
- 方法卻不同，它需要有名字，不能被當作值來看待，最重要的是，它`必須隸屬於某一個類型。方法所屬的類型會通過其聲明中的接收者（receiver）聲明體現出來`。
- 接收者聲明就是在關鍵字func和方法名稱之間的那個圓括號包裹起來的內容，其中必須包含確切的名稱和類型字面量。這個`接收者的類型其實就是當前方法所屬的那個類型`，而接收者的名稱，則用於在當前方法中引用它所屬的類型的當前值。
- 方法隸屬的類型其實並`不局限於結構體類型，但必須是某個自定義的數據類型，並且不能是任何接口類型`。


### Promoted / Anonymous / Embedded fields
```go
type Animal struct {
  scientificName string // 學名。
  AnimalCategory    // 動物基本分類。
}
```
字段聲明AnimalCategory代表了Animal類型的一個嵌入字段。 Go 語言規範規定，如果一個字段的聲明中只有字段的類型名而沒有字段的名稱，那麼它就是一個嵌入字段，也可以被稱為匿名字段。我們可以通過此類型變量的名稱後跟“.”，再後跟嵌入字段類型的方式引用到該字段。也就是說，嵌入字段的類型既是類型也是名稱。

- 只要名稱相同，無論這兩個方法的簽名是否一致，被嵌入類型的方法都會“屏蔽”掉嵌入字段的同名方法。
- 因為嵌入字段的字段和方法都可以“嫁接”到被嵌入類型上，所以即使在兩個同名的成員一個是字段，另一個是方法的情況下，這種“屏蔽”現象依然會存在。
- 不過，即使被屏蔽了，我們仍然可以通過鍊式的選擇表達式，選擇到嵌入字段的字段或方法

- `Go 語言中根本沒有繼承的概念，它所做的是通過嵌入字段的方式實現了類型之間的組合`。(Why is there no type inheritance?)。簡單來說，面向對象編程中的繼承，其實是通過犧牲一定的代碼簡潔性來換取可擴展性，而且這種可擴展性是通過侵入的方式來實現的。類型之間的組合採用的是非聲明的方式，我們不需要顯式地聲明某個類型實現了某個接口，或者一個類型繼承了另一個類型。

- 同時，類型組合也是非侵入式的，它不會破壞類型的封裝或加重類型之間的耦合。我們要做的只是把類型當做字段嵌入進來，然後坐享其成地使用嵌入字段所擁有的一切。如果嵌入字段有哪裡不合需求，我們還可以用“包裝”或“屏蔽”的方式去調整和優化。另外，類型間的組合也是靈活的，我們總是可以通過嵌入字段的方式把一個類型的屬性和能力“嫁接”給另一個類型。這時候，被嵌入類型也就自然而然地實現了嵌入字段所實現的接口。再者，組合要比繼承更加簡潔和清晰，Go 語言可以輕而易舉地通過嵌入多個字段來實現功能強大的類型，卻不會有多重繼承那樣複雜的層次結構和可觀的管理成本。接口類型之間也可以組合。在 Go 語言中，接口類型之間的組合甚至更加常見，我們常常以此來擴展接口定義的行為或者標記接口的特徵。

### Method Receiver
- 方法的接收者類型(Methods,Function Receiver)`必須是某個自定義的數據類型，而且不能是接口類型或接口的指針類型`。所謂的值方法，就是接收者類型是非指針的自定義數據類型的方法。

- Methods,Function Receiver
  1. return type:
    - value receiver: copy of value。在該方法內對該副本的修改一般都不會體現在原值上，除非這個類型本身是某個引用類型（比如slice或map）的別名類型。
    - pointer receiver: copy of pointer。我們在這樣的方法內對該副本指向的值進行修改，卻一定會體現在原值上。
  2. 一個`自定義數據類型的方法集合`中僅會包含它的所有`值方法`，而該類型的`指針類型的方法集合`卻囊括了前者的所有方法，`包括所有值方法(value receiver methods)和所有指針方法(pointer receiver methods)`。
    - value receiver: 只能調用到它的值方法。
    - pointer receiver: Go 語言會適時地為我們進行自動地轉譯，使得我們在這樣的值上也能調用到它的指針方法。
      比如，在Cat類型的變量cat之上，之所以我們可以通過`cat.SetName("monster")`修改貓的名字，是因為 Go 語言把它自動轉譯為了`(&cat).SetName("monster")`，即：`先取cat的指針值，然後在該指針值上調用SetName方法`。
  3. 一個類型的方法集合中有哪些方法與它能實現哪些接口類型是息息相關的。如果一個基本類型和它的指針類型的方法集合是不同的，那麼它們具體實現的接口類型的數量就也會有差異，除非這兩個數量都是零。比如，一個指針類型實現了某某接口類型，但它的基本類型卻不一定能夠作為該接口的實現類型


- 方法的定義感覺本質上也是一種語法糖形式，其本質就是一個函數，聲明中的方法接收者就是函數的第一個入參，在調用時go會把施調變量作為函數的第一個入參的實參傳入，比如
  ```go
  func (t MyType) MyMethod(in int) (out int)
  // 可以看作是
  func MyMethod(t Mytype, in int) (out int)

  myType.MyMethod(123)
  // 就可以理解成是調用
  MyMethod(myType, 123)
  // 如果 myType 是 *MyType 指針類型，
  // 則在調用是會自動進行指針解引用，實際就是這麼調用的 MyMethod(*myType, 123)，這麼一理解，值方法和指針方法的區別也就顯而易見了。
  ```

## 14 | 接口類型的合理運用
- 接口類型聲明中的這些方法所代表的就是該接口的方法集合。一個接口的方法集合就是它的全部特徵。對於任何數據類型，只要它的方法集合中完全包含了一個接口的全部特徵（即全部的方法），那麼它就一定是這個接口的實現類型。(Duck typing)
- 判定一個數據類型的某一個方法實現,有兩個充分必要條件，
  1. 一個是“兩個方法的簽名需要完全一致”
  2. 另一個是“兩個方法的名稱要一模一樣”
- 怎樣才能讓一個接口變量的值真正為nil呢？只聲明它但不做初始化，or 直接把字面量nil賦給它。

- Go 語言團隊鼓勵我們聲明體量較小的接口，並建議我們通過這種接口間的組合來擴展程序、增加程序的靈活性。這是因為相比於包含很多方法的大接口而言，小接口可以更加專注地表達某一種能力或某一類特徵，同時也更容易被組合在一起。
- Go 語言標準庫代碼包io中的ReadWriteCloser接口和ReadWriter接口就是這樣的例子，它們都是由若干個小接口組合而成的。以io.ReadWriteCloser接口為例，它是由io.Reader、io.Writer和io.Closer這三個接口組成的。
  這三個接口都只包含了一個方法，是典型的小接口。它們中的每一個都只代表了一種能力，分別是讀出、寫入和關閉。們編寫這幾個小接口的實現類型通常會很容易。並且，一旦我們同時實現了它們，就等於實現了它們的組合接口io.ReadWriteCloser。即使我們只實現了io.Reader和io.Writer，那麼也等同於實現了io.ReadWriter接口，因為後者就是前兩個接口組成的。可以看到，這幾個io包中的接口共同組成了一個接口矩陣。它們既相互關聯又獨立存在。

## 15 | 關於指針的有限操作
### Container Value Literals
Like struct values, container values can also be represented with composite literals, T{...}, where T denotes container type (except the zero values of slice and map types). Here are some examples:
```go
// An array value containing four bool values.
[4]bool{false, true, true, false}

// A slice value which contains three words.
[]string{"break", "continue", "fallthrough"}

// A map value containing some key-value pairs.
map[string]int{"C": 1972, "Python": 1991, "Go": 2009}
```

### Go 語言中的哪些值是不可尋址的（addressable）
  1. 不可變的:
    - `常量的值總是會被存儲到一個確切的內存區域中，並且這種值肯定是不可變的`。`基本類型值的字面量也是一樣`，其實它們本就可以被視為常量，`只不過沒有任何標識符可以代表它們罷了`。
    - 由於 Go 語言中的`字符串值也是不可變的`，所以對於一個字符串類型的變量來說，`基於它的索引或slice的結果值也都是不可尋址的`，因為即使拿到了這種值的內存地址也改變不了什麼。

  2. 臨時結果:
    - 算術操作的結果值屬於一種臨時結果。在我們把這種結果值賦給任何變量或常量之前，即使能拿到它的內存地址也是沒有任何意義的。
    - 可以把各種`對值字面量施加的表達式的求值結果`都看做是臨時結果。
      Go 語言中的表達式有很多種，其中常用的包括以下幾種:
      用於獲得某個元素的`索引表達式`。
      用於獲得某個`slice（片段）的slice表達式`。
      用於訪問某個字段的選擇表達式。
      用於調用某個`函數或方法的調用`表達式。
      用於轉換值的類型的 `類型轉換`錶達式。
      用於判斷值的類型的`類型斷言`表達式。
      向channel發送元素值或從channel那裡接收元素值的接收表達式。
    - 需要特別注意的例外是，`對slice字面量的索引結果值是可尋址的`。因為不論怎樣，每個slice值都會持有一個底層數組，而這個底層數組中的每個元素值都是有一個確切的內存地址的。
      那麼對slice字面量的slice結果值為什麼卻是不可尋址的？這是`因為slice表達式總會返回一個新的slice值，而這個新的slice值在被賦給變量之前屬於臨時結果`。
      - 變量的值本身就不是“臨時的”。對比而言，值字面量在還沒有與任何變量（或者說任何標識符）綁定之前是沒有落腳點的，我們無法以任何方式引用到它們。這樣的值就是“臨時的”。
    - `如果把臨時結果賦給一個變量，那麼它就是可尋址的了`。如此一來，取得的指針指向的就是這個變量持有的那個值了。
  3. 不安全的:
    - `“不安全的”操作很可能會破壞程序的一致性，引發不可預知的錯誤`，從而嚴重影響程序的功能和穩定性。
    - `函數在 Go 語言中是一等公民`，所以`可以把代表函數或方法的字面量或標識符賦給某個變量、傳給某個函數或者從某個函數傳出`。但是，這樣的函數和方法都是不可尋址的。`對函數或方法的調用結果值也是不可尋址的`，這是因為它們都屬於臨時結果。
      1. 函數就是代碼，是`不可變的`。
      2. 拿到指向一段代碼的指針是`不安全的`。
    - 通過對map類型的變量施加索引表達式，得到的結果值不屬於臨時結果，可是，這樣的值卻是不可尋址的。原因是，`map中的每個key-value pair的存儲位置都可能會變化，而且這種變化外界是無法感知的`。
      - map中總會有若干個哈希桶用於均勻地儲存key-value pair。`當滿足一定條件時，map可能會改變哈希桶的數量，並適時地把其中的key-value pair搬運到對應的新的哈希桶中`。
      - 在這種情況下，獲取map中任何元素值的指針都是無意義的，也是不安全的。because不知道什麼時候那個元素值會被搬運到何處，也不知道原先的那個內存地址上還會被存放什麼別的東西。

### 不可尋址的值在使用上有哪些限制？
- 無法使用取址操作符&獲取它們的指針
```go
func New(name string) Dog {
  return Dog{name}
}

func main() {
  New("little pig").SetName("monster")
}
```
- 調用表達式 dog.SetName("monster") 會被自動地轉譯為 (&dog).SetName("monster")，即：先取dog的指針值，再在該指針值上調用SetName方法。
  - 由於New函數的調用結果值屬於臨時結果，是不可尋址的，所以無法對它進行取址操作。因此，編譯器報告兩個錯誤
    1. because 不能取得New("little pig")的地址
    1. result in 不能在New("little pig")的結果值上調用指針方法

例外case:
1. Go 語言中的++和--並不屬於操作符，而分別是自增語句和自減語句的重要組成部分。
  Go 語言規範中的語法定義是，`只要在++或--的左邊添加一個表達式，就可以組成一個自增語句或自減語句`，但是，它還明確了一個很重要的限制，那就是這個`表達式的結果值必須是可尋址的`。這就使得針對值字面量的表達式幾乎都無法被用在這裡。
  - 不過這有一個例外，雖然對`map字面量(literal)`和`map變量索引表達式 map[key]`的結果值都是不可尋址的，但是這樣的表達式卻`可以被用在自增語句和自減語句中`。

2. 在賦值語句中，`賦值操作符左邊的表達式的結果值必須可尋址的`，但是對`map的索引結果值map[key]`也是可以的。

3. 在帶有range子句的for語句中，在`range關鍵字左邊的表達式的結果值也都必須是可尋址的`，不過對`map的索引結果值map[key]`同樣可以被用在這裡。以上這三條規則我們合併起來記憶就可以了。

### 怎樣通過unsafe.Pointer操縱可尋址的值？

```go
dog := Dog{"little pig"}
dogP := &dog
dogPtr := uintptr(unsafe.Pointer(dogP))
```

1. 先聲明了一個Dog類型的變量dog，然後用取址操作符&，取出了它的指針值，並把它賦給了變量dogP。
2. 使用了兩個類型轉換，先把dogP轉換成了一個unsafe.Pointer類型的值，然後緊接著又把後者轉換成了一個uintptr的值，並把它賦給了變量dogPtr。這背後隱藏著一些轉換規則，如下：
  - 一個指針值（比如*Dog類型的值）可以被轉換為一個unsafe.Pointer類型的值，反之亦然。
  - 一個uintptr類型的值也可以被轉換為一個unsafe.Pointer類型的值，反之亦然。
  - `一個指針值無法被直接轉換成一個uintptr類型的值`，反過來也是如此。

對於`指針值和uintptr類型值之間的轉換，必須使用unsafe.Pointer類型的值作為中轉`。那麼，把指針值轉換成uintptr類型的值有什麼意義嗎？
```go
// dogP的name字段值的起始存儲地址了 = 結構體值在內存中的起始存儲地址 + 結構體中某個字段的值偏移量
namePtr := dogPtr + unsafe.Offsetof(dogP.name)
// 指向dogP的name字段值的指針值 = (*string)(uintptr類型值和指針值之間的轉換)
nameP := (*string)(unsafe.Pointer(namePtr))
```
- 這裡需要與unsafe.Offsetof函數搭配使用。
  - unsafe.Offsetof函數用於獲取兩個值在內存中的起始存儲地址之間的偏移量，以字節為單位。

- 直接用取址表達式&(dogP.name)不就能拿到這個指針值了嗎？
  - 如果我們根本就不知道這個結構體類型是什麼，也拿不到dogP這個變量，那麼還能去訪問它的name字段嗎？
  - 答案是，只要有namePtr就可以。它就是一個`無符號整數`，但同時也是一個指向了程序內部數據的內存地址。它可以直接修改埋藏得很深的內部數據。


## 16 | go語句及其執行規則（上）
- `go函數真正被執行的時間`總會與`其所屬的go語句被執行的時間不同(被P調度執行)`。
  - 當程序執行到一條go語句的時候，Go 語言的運行時系統，`會先試圖從某個存放空閒的 G 的隊列中獲取一個 G（也就是 goroutine）`，已存在的 goroutine 總是會被優先復用。
  - 它`只有在找不到空閒 G 的情況下才會去創建一個新的 G`。
  - 在拿到了一個空閒的 G 之後，Go 語言運行時系統會用這個 G 去包裝當前的那個go函數（或者說該函數中的那些代碼），然後再把這個 G 追加到某個存放可運行的 G 的隊列中。這類隊列中的 G 總是會按照`先入先出的順序`，很快地由運行時系統內部的調度器安排運行。
- 因此，`go函數的執行時間`總是會明顯`滯後於它所屬的go語句的執行時間(被P調度執行)`。這裡所說的“明顯滯後”是對於計算機的 CPU 時鐘和 Go 程序來說的。在大多數時候都不會有明顯的感覺。

- `只要go語句本身執行完畢，Go 程序完全不會等待go函數的執行，它會立刻去執行後面的code。這就是所謂的異步並發地執行`。
  ```go
  package main
  import "fmt"

  func main() {
    // for語句會以很快的速度執行完畢。當它執行完畢時，那 10 個包裝了go函數的 goroutine 往往還沒有獲得運行的機會。
    for i := 0; i < 10; i++ {
      go func() {
        fmt.Println(i) // 不會有任何內容被打印出來
      }()
      // 後面的code先輩立即執行
    }
    // 一旦主 goroutine 中的代碼（也就是main函數中的那些代碼）執行完畢，當前的 Go 程序就會結束運行。
  }
  ```
嚴謹地講，Go 語言並不會去保證這些 goroutine 會以怎樣的順序運行, `默認情況下的執行順序是不可預知的`。由於主 goroutine 會與其他 goroutine 一起接受調度，又因為調度器很可能會在 goroutine 中的代碼只執行了一部分的時候暫停，以期所有的 goroutine 有更公平的運行機會。

## 17 | go語句及其執行規則（下）
### 怎樣才能讓主 goroutine 等待其他 goroutine？
```go
func main() {
	num := 10
	sign := make(chan struct{}, num)

	for i := 0; i < num; i++ {
		go func() {
			fmt.Println(i)
			sign <- struct{}{}
		}()
	}

	// 辦法1。
	//time.Sleep(time.Millisecond * 500)

	// 辦法2。
	for j := 0; j < num; j++ {
		<-sign
	}
}
```
- struct{}類型值的表示法只有一個，即：struct{}{}。並且，它佔用的內存空間是0字節。確切地說，這個值在整個 Go 程序中永遠都只會存在一份。雖然可以無數次地使用這個值字面量，但是用到的卻都是同一個值。
- sync.WaitGroup

### 怎樣讓我們啟用的多個 goroutine 按照既定的順序運行？
- 在go語句被執行時，`傳給go函數的參數i會先被求值，如此就得到了當次迭代的序號`。之後，無論go函數會在什麼時候執行，這個參數值都不會變。
- `讓count變量成為一個信號`，它的值總是下一個可以調用打印函數的`go函數的序號`。
- 這個序號其實就是啟用 goroutine 時，那個當次迭代的序號。也正因為如此，go函數實際的執行順序才會與go語句的執行順序完全一致。此外，這裡的`trigger函數實現了一種自旋（spinning）。除非發現條件已滿足，否則它會不斷地進行檢查`。
- trigger函數會不斷地獲取一個名叫count的變量的值，並判斷該值是否與參數i的值相同。如果相同，那麼就立即調用fn代表的函數，然後把count變量的值加1，最後顯式地退出當前的循環。否則，我們就先讓當前的 goroutine“睡眠”一個納秒再進入下一個迭代。
- 由於當所有啟用的 goroutine 都運行完畢之後，count的值一定會是10，所以就把10作為了第一個參數值。又`由於並不想打印這個10，所以把一個什麼都不做的函數作為了第二個參數值`。
```go
func main() {
	var count uint32
	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
        // 由於trigger函數會被多個 goroutine 並發地調用，所以它用到的非本地變量count，就被多個用戶級線程共用了。因此，對它的操作就產生了競態條件（race condition），破壞了程序的並發安全性。 --> sync/atomic包中聲明了很多用於原子操作的函數 to solve this problem
        // 原子操作函數對被操作的數值的類型有約束，count以及相關的變量和參數的類型 由int變為了uint32
        // 所以這麼做是因為int位數是根據系統決定的，而原子級操作要求速度盡可能的快，所以明確了整數的位數才能最大地提高性能。
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}
	for i := uint32(0); i < 10; i++ {
    // 在go語句被執行時，傳給go函數的參數i會先被求值，如此就得到了當次迭代的序號。之後，無論go函數會在什麼時候執行，這個參數值都不會變。
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}
	trigger(10, func() {}) // 因為我依然想讓主 goroutine 最後一個運行完畢
}
```