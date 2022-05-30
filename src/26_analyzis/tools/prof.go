package main

import (
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)



func fillMatrix(m [][]int) {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			m[i][j] = s.Intn(100000)
		}
	}
}

func caculate(m [][]int) {
	for i := 0; i < len(m); i++ {
		tmp := 0
		for j := 0; j < len(m[0]); j++ {
			tmp += m[i][j]
		}
	}
}

func main() {
	// output file
	f, err := os.Create("cpu.prof")
	if err != nil{
		log.Fatal("Could not create CPU profiler: ", err)
	}
	defer f.Close()

	// get system info
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("Could not start CPU profiler: ", err)
	}
	defer pprof.StopCPUProfile()

	row := 10000000
	col := 10000000
	x := make([][]int, row, col)
	fillMatrix(x)
	caculate(x)

	f1, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("Could not create memory profile: ", err)
	}
	// $ go build prof.go --> $ ./prof --> $ go tool pprof mem.prof
	runtime.GC()
	if err := pprof.WriteHeapProfile(f1); err != nil {
		log.Fatal("Could not write heap profile: ", err)
	}
	f1.Close()

	f2, err := os.Create("goroutine.prof")
	if err != nil {
		log.Fatal("Could not create goroutine profile: ", err)
	}
	if gProf := pprof.Lookup("goroutine"); gProf == nil {
		log.Fatal("Could not write goroutine profile: ")
	} else {
		gProf.WriteTo(f2, 0)
	}
	f2.Close()
}

