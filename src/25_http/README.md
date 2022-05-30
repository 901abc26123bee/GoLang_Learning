# 45 HTTP server
## Default Router
```go
    func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request) {
    	handler := sh.srv.Handler
    	if handler == nil {
    		handler = DefaultServeMux //使用缺省的Router
    	}
    	if req.RequestURI == "*" && req.Method == "OPTIONS" {
    		handler = globalOptionsHandler{}
    	}
    	handler.ServeHTTP(rw, req)
    }
```

## 路由规则
* URL 分為兩種，末尾是/:表示一個子樹，後面可以跟其他子路徑；
  末尾不是/，表示一個葉子，固定的路徑
* 以/結尾的URL可以匹配它的任何子路徑，比如/images 會匹配 /images/cute-cat.jpg
* 它採用最長匹配原則，如果有多個匹配，一定採用匹配路徑最長的那個進行處理 * 如果沒有找到任何匹配項，會返回 404 錯誤


[HttpRouter](https://github.com/julienschmidt/httprouter)

$ go run http_hello.go
$ go run http_router.go