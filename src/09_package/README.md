## GOROOT、GOPATH、Go-Modules
[Golang — GOROOT、GOPATH、Go-Modules-三者的關係介紹](https://medium.com/%E4%BC%81%E9%B5%9D%E4%B9%9F%E6%87%82%E7%A8%8B%E5%BC%8F%E8%A8%AD%E8%A8%88/golang-goroot-gopath-go-modules-%E4%B8%89%E8%80%85%E7%9A%84%E9%97%9C%E4%BF%82%E4%BB%8B%E7%B4%B9-d17481d7a655)
GOROOT介紹
### GOROOT
- 在安裝完Golang語言的時候，所謂的安裝路徑其實就是你的GOROOT路徑，也就是說GOROOT存放的Golang語言內建的程式庫的所在位置

  首先我們先看目前GOROOT路徑，透過以下指令：
  ```sh
  go env
  ```
  這樣可以看到當前有關Golang語言的全部相關環境變數設定

- 通常如果你是初次安裝Golang語言並且沒做什麼環境變數設定的話，GOROOT設定路徑就是你當初安裝Golang語言的路徑，而GOPATH通常預設會是使用者目錄下的go資料夾。

- 當執行Golang程式碼，當需要存取套件時，會先去GOROOT路徑下的src資料夾找同等於我們在程式碼中import的路徑下去找有沒有gin這個資料夾，而這資料夾裡面就是包含了所有有關於該套件的程式庫。
- 如果在GOROOT路徑下沒有找到，則會往GOPATH路徑下的src資料夾找同等於我們在程式碼中import的路徑下去找有沒有gin這個資料夾。
- 所以只要GOROOT跟GOPATH路徑下都沒有找到該套件的話，就無法執行該程式碼。

### GOPATH
- 官方的程式庫所在位置就是在GOROOT裡面，而GOPATH就是專門存放第三方套件以供我們程式碼的需要。
- 那通常開發Golang的話，通常會在重新設定GOPATH的位置
   `set GOPATH=<PATH>` : <PATH>就是你想要放置Golang語言專案的地方，但是需要做成以下的分層：
  ```sh
    D:\CodingProject\GolangProject ---> bin
                                   ---> pkg
                                   ---> src
  ```
  也就是<PATH>，實際上是D:\CodingProject\GolangProject，並不是D:\CodingProject\GolangProject\src，依照Golang語言的慣例(強制)，GOPATH是指src路徑的上一層，如果你去看GOROOT的設計，也是這樣的。
  而我們要在GOPATH路徑下主動新增src資料夾，所謂src就是代表source code的意思，也就是放我們開發Golang程式碼的相關專案的原始碼。

- `go run` : 而編譯檔跟執行檔事實上是存在一個暫存資料夾裡面，當運行完此程式就會自動刪除。

- `go build` : 將程式碼編譯為執行檔
  1.可以透過以下指令指定build出來的可執行檔位置放置在哪：  `go build -o bin/main.exe src/webDemo/main.go`
  2. -o 後面第一個參數代表產生出來執行檔路徑要放哪及名稱要取什麼，第二個參數代表要編譯哪一個go程式碼的路徑。
  3. 從這邊就可以知道為什麼我們在GOPATH下要新增一個bin資料夾，依照官方慣例，GOPATH下會有一個bin資料夾專門就是放專案build出來的可執行檔。
  4. 但是go build有個缺點就是每次編譯程式碼，比較沒有效率，當專案架構越來越大，build的速度也就越來越慢。

- `go install` : 因應go build的缺點，因此有go install指令，go install可以做兩件事情：
  - 將套件編譯成.a file
  - 如果是main套件則編譯成可執行檔
  而有了第一項功能，在下次編譯的時候，就不會將套件的程式碼重新編譯，而是直接用編譯好的.a file。
  而.a file就位於GOPATH/pkg裡面。這就是為什麼golangGO慣例要在GOPATH下新增此三個資料夾，都是有用處的。
  - 特別注意：
    go install 如果要在非GOPATH路徑下使用的話，要先設定GOBIN環境變數，否則會出現錯誤
    通常GOBIN環境變數就是設定GOPATH/bin。

GOPATH的缺點
講完了GOROOT、GOPATH，不知道大家有沒有發現GOPATH的一個很大的缺點，那就是你相關的第三方套件只要不是官方程式庫，都需要放置在GOPATH/src的路徑下才可以使用。

## Modules and Packages
- 為了解決不被GOPATH的問題，因此官方在1.11開始推出了Go Modules的功能。Go Modules解決方式很像是Java看到Maven的做法，將第三方程式庫儲存在本地的空間，並且給程式去引用。

- 而採用Go Modules，下載下來的第三方套件都就位在GOPATH/pkg/mod資料夾裡面。`go mod init <module name>`

- `go get -u github.com/xxx`會將需要的套件安裝在GOPATH/pkg/mod資料夾裡面。而且會發現出現一個go.sum的檔案，這個檔案基本上用來記錄套件版本的關係，確保是正確的，是不太需要理會的。

- 只要有開啟go modules功能，go get 就不會像以前一樣在GOPATH/src下放置套件檔案，而是會放在GOPATH/pkg/mod裡面，並且go.mod會寫好引入。

[[Golang] Modules and Packages](https://pjchender.dev/golang/modules-and-packages/)
```sh
# 在 hello 資料夾中
$ go mod init example.com/user/hello        # 宣告 module_path，通常會和該 repository 的 url 位置一致
$ go install .                      # 編譯該 module 並將執行檔放到 GOBIN，因此在 GOBIN 資料夾中會出現 hello 的執行檔

$ go mod tidy             # 移除沒用到的套件
$ go clean -modcache            # 移除所有下載第三方套件的內容
```

💡 安裝到 GOBIN 資料夾的檔案名稱，會是在 go.mod 檔案中第一行定義 module path 中路徑的最後一個。因此若 module_path 是 example.com/user/hello 則在 GOBIN 中的檔名會是 hello；若 module_path 是 example/user 則在 GOBIN 資料夾中的檔名會是 user。

- 若我們在 go module 中有使用其他的遠端（第三方）套件，當執行 go install、go build 或 go run 時，go 會自動下載該 remote module，並記錄在 go.mod 檔案中。
- 這些遠端套件會自動下載到 $GOPATH/pkg/mod 的資料夾中。當有不同的 module 之間需要使用相同版本的第三方套件時，會共用這些下載的內容，因此這些內容會是「唯讀」。
- 若想要刪除這些第三方套件的內容，可以輸入 go clean -modcache。

## 設定 GOPATH
```sh
vim ~/.zshrc
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```


## import

### 一些特殊的 import

1. 點操作

```go
  import(
      . "fmt"
  )
```

這個點操作的含義就是這個套件匯入之後在你呼叫這個套件的函式時，你可以省略字首的套件名，也就是前面你呼叫的 `fmt.Println("hello world")` 可以省略的寫成 `Println("hello world")`

2. 別名操作

別名操作顧名思義我們可以把套件命名成另一個我們用起來容易記憶的名字
```go
  import(
      f "fmt"
  )
```
別名操作的話呼叫套件函式時字首變成了我們的字首，即 `f.Println("hello world")`

3. `_`操作

```go
	import (
	    "database/sql"
	    _ "github.com/ziutek/mymysql/godrv"
	)
```
`_`操作其實是引入該套件，而不直接使用套件裡面的函式，而是呼叫了該套件裡面的 init 函式。