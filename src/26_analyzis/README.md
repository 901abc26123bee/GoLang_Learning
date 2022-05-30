## 準備工作
* 安裝 graphviz
    * brew install graphviz
* 將 $GOPATH/bin 加入到 $PATH
    * Mac OS: 在 .bash_profile 中修改路徑
* 安裝 go-torch
    * go get -u github.com/uber/go-torch
    * 下載並複制 flamegraph.pl ([https://github.com/brendangregg/FlameGraph](https://github.com/brendangregg/FlameGraph))
    至 $GOPATH/bin 路徑下
    * 將 $GOPATH/bin 加入 $PATH
