package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func GetFibonacciSeries(n int) []int {
	ret := make([]int, 2, n)
	ret[0] = 1
	ret[1] = 1
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i - 2] + ret[i - 1])
	}
	return ret
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}

func CreateFBS(w http.ResponseWriter, r *http.Request) {
	var fbs []int
	for i := 0; i < 10000000; i++ {
		fbs = GetFibonacciSeries(50)
	}
	w.Write([]byte(fmt.Sprintf("%v", fbs)))
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/fb", CreateFBS)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
// http://localhost:8081/debug/pprof/
// $ go tool pprof http://localhost:8081/debug/pprof/profile
// $ go-torch http://localhost:8081/debug/pprof/profile
// Writing svg to torch.svg