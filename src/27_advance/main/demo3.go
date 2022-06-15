package main

import (
	"flag"
	"fmt"
	"os"
)
//  *怎樣自定義命令源碼文件的參數使用說明
var name string

func init_old() {
	// 1.
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	// 2.
	// flag.ExitOnError的含義是，告訴命令參數容器，當命令後跟--help或者參數設置的不正確的時候，在打印命令參數使用說明後以狀態碼2結束當前程序。
	// *狀態碼2代表用戶錯誤地使用了命令，而flag.PanicOnError與之的區別是在最後拋出“運行時恐慌（panic）”
	// flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError) // 產生另一種輸出效果
	
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

	// 3.
	// 不用全局的flag.CommandLine變量，轉而自己創建一個私有的命令參數容器。我們在函數外再添加一個變量聲明：
var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)
func init() {
	
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	cmdLine.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main_old() {
	// *在調用flag包中的一些函數（比如StringVar、Parse等等）的時候，實際上是在調用flag.CommandLine變量的對應方法。
	// *flag.CommandLine相當於默認情況下的命令參數容器。所以，通過對flag.CommandLine重新賦值，我們可以更深層次地定制當前命令源碼文件的參數使用說明
	// flag.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	// 	flag.PrintDefaults()
	// }
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}

func main() {
	// *os.Args[1:]指的就是我們給定的那些命令參數。這樣做就完全脫離了flag.CommandLine。
	cmdLine.Parse(os.Args[1:])
	fmt.Printf("Hello, %s!\n", name)
}
/*
$ go run demo3.go --help
	Usage of question:
	-name string
			The greeting object. (default "everyone")
*/