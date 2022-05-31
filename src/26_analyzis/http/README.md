[Go 程式崩了？煎魚教你用 PProf 工具來救火！](https://www.gushiciku.cn/pl/gd3C/zh-tw)
[Profiling A Guide to Profiling and Code Optimization](https://medium.com/happyfresh-fleet-tracker/danny-profiling-1c60a19d30de)
## produce Profile with HTTP
1. import _ "net/http/pprof"

2. visit  `http://<host>:<port>/debug/pprof`

3. `go tool pprof _http://<host>:<port>/debug/pprof/profile?seconds=10`   (default採樣 30s)
  ```sh
  # 等待 10 秒，執行完畢後會產生一個檔案 Saved profile in XXXX,
  $ go tool pprof http://localhost:8081/debug/pprof/profile
  Fetching profile over HTTP from http://localhost:8081/debug/pprof/profile
  Saved profile in /Users/wengzhenyuan/pprof/pprof.samples.cpu.001.pb.gz
  Type: cpu
  Time: May 31, 2022 at 9:46pm (CST)
  Duration: 30s, Total samples = 0
  No samples were found with the default sample value type.
  Try "sample_index" command to analyze different sample values.
  Entering interactive mode (type "help" for commands, "o" for options)
  (pprof) %

  # 使用 pprof UI介面打開檔案, need top install Graphviz beforehand
  $ go tool pprof -http=:8080 /Users/wengzhenyuan/pprof/pprof.samples.cpu.001.pb.gz
  Serving web UI on http://localhost:8080
  $ go tool pprof -http=:6001 /Users/wengzhenyuan/pprof/pprof.samples.cpu.001.pb.gz
  Serving web UI on http://localhost:6001
  ```
  
  /debug/pprof/profile：訪問這個鏈接會自動進行 CPU profiling，持續 30s，並生成一個文件供下載
  /debug/pprof/block：Goroutine阻塞事件的記錄。默認每發生一次阻塞事件時取樣一次。
  /debug/pprof/goroutines：活躍Goroutine的信息的記錄。僅在獲取時取樣一次。
  /debug/pprof/heap： 堆內存分配情況的記錄。默認每分配512K字節時取樣一次。
  /debug/pprof/mutex: 查看爭用互斥鎖的持有者。
  /debug/pprof/threadcreate: 系統線程創建情況的記錄。僅在獲取時取樣一次。

  allocs：檢視過去所有記憶體分配的樣本，訪問路徑為 $HOST/debug/pprof/allocs 。
  block：檢視導致阻塞同步的堆疊跟蹤，訪問路徑為 $HOST/debug/pprof/block 。
  cmdline：當前程式的命令列的完整呼叫路徑。
  goroutine：檢視當前所有執行的 goroutines 堆疊跟蹤，訪問路徑為 $HOST/debug/pprof/goroutine 。
  heap：檢視活動物件的記憶體分配情況， 訪問路徑為 $HOST/debug/pprof/heap 。
  mutex：檢視導致互斥鎖的競爭持有者的堆疊跟蹤，訪問路徑為 $HOST/debug/pprof/mutex 。
  profile：預設進行 30s 的 CPU Profiling，得到一個分析用的 profile 檔案，訪問路徑為 $HOST/debug/pprof/profile 。
  threadcreate：檢視建立新OS執行緒的堆疊跟蹤，訪問路徑為 $HOST/debug/pprof/threadcreate 。
  如果你在對應的訪問路徑上新增 ?debug=1 的話，就可以直接在瀏覽器訪問

4. go-torch
  ```sh
  # download go_torch
  $ go get github.com/uber/go-torch
  $ cd $GOPATH/src/github.com/uber/go-torch
  $ git clone https://github.com/brendangregg/FlameGraph.git
  # 火焰圖分析
  $ go-torch -seconds 10 http://<host>:<port>/debug/pprof/profile
  ```


run the program `go run fb_server.go`