package main

import (
	"fmt"
	"net/http"
	"time"
)

// http://localhost:8080/
// http://localhost:8080/time
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\":\"%s\"}", t)
		w.Write([]byte(timeStr))
	})
	http.HandleFunc("/test/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"test\":\"%s\"}", t)
		w.Write([]byte(timeStr))
	})
	http.ListenAndServe(":8080", nil)
}