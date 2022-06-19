package main

import (

)
// *怎樣在運行命令源碼文件的時候傳入參數，又怎樣查看參數的使用說明
// *Go 語言標準庫中有一個代碼包專門用於接收和解析命令參數。這個代碼包的名字叫flag。

/*
var name string

func init() {
	//* 有一個與flag.StringVar函數類似的函數，叫flag.String。這兩個函數的區別是，後者會直接返回一個已經分配好的用於存儲命令參數值的地址。
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	//* 函數flag.Parse用於真正解析命令參數，並把它們的值賦給相應的變量
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
*/

/*
//* 1. 命令源碼文件怎樣接收參數
$ go run demo2.go -name="Robert"
	Hello, Robert!
$ go run demo2.go --help
	Usage of /var/folders/r4/5hf5nt4d7yq4vzg87tq3d9cc0000gn/T/go-build2819551025/b001/exe/demo2: (臨時生成的可執行文件的完整路徑)
		-name string
					The greeting object. (default "everyone")

如果我們先構建這個命令源碼文件再運行生成的可執行文件:
$ go build demo2.go
$ ./demo2 --help
	Usage of ./demo2:	(生成的可執行文件的 relative 路徑)
		-name string
					The greeting object. (default "everyone")
 */