[Profiling A Guide to Profiling and Code Optimization](https://medium.com/happyfresh-fleet-tracker/danny-profiling-1c60a19d30de)

$ brew install graphviz
$ go tool pprof -http :6060 cpu.prof

```bash
$ go build prof.go
$ ls
prof		prof.go
$ ./prof
$ ls
cpu.prof		gorutine.prof		mem.prof	 prof 	prof.go
```

## cpu.prof
```bash
$ go-torch cpu.prof
$ go tool pprof cpu.prof
Type: cpu
Time: May 30, 2022 at 11:52pm (CST)
Duration: 219.84ms, Total samples = 120ms (54.59%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 120ms, 100% of 120ms total
Showing top 10 nodes out of 12
      flat  flat%   sum%        cum   cum%
      60ms 50.00% 50.00%       60ms 50.00%  runtime.scanobject
      30ms 25.00% 75.00%       30ms 25.00%  main.caculate (inline)
      20ms 16.67% 91.67%       20ms 16.67%  runtime.heapBitsSetType
      10ms  8.33%   100%       10ms  8.33%  main.fillMatrix
         0     0%   100%       60ms 50.00%  main.main
         0     0%   100%       60ms 50.00%  runtime.gcBgMarkWorker
         0     0%   100%       60ms 50.00%  runtime.gcBgMarkWorker.func2
         0     0%   100%       60ms 50.00%  runtime.gcDrain
         0     0%   100%       60ms 50.00%  runtime.main
         0     0%   100%       20ms 16.67%  runtime.makeslice
(pprof) list fillMatrix
Total: 120ms
ROUTINE ======================== main.fillMatrix in /Users/wengzhenyuan/Desktop/go_learning/src/26_analyzis/tools/prof.go
      10ms       10ms (flat, cum)  8.33% of Total
         .          .     11:
         .          .     12:
         .          .     13:func fillMatrix(m [][]int) {
         .          .     14:   s := rand.New(rand.NewSource(time.Now().UnixNano()))
         .          .     15:
      10ms       10ms     16:   for i := 0; i < len(m); i++ {
         .          .     17:           for j := 0; j < len(m[0]); j++ {
         .          .     18:                   m[i][j] = s.Intn(100000)
         .          .     19:           }
         .          .     20:   }
         .          .     21:}
(pprof)
```

## mem.prof

```bash
$ go tool pprof mem.prof
Type: inuse_space
Time: May 30, 2022 at 11:52pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 231.61MB, 99.57% of 232.61MB total
Dropped 9 nodes (cum <= 1.16MB)
      flat  flat%   sum%        cum   cum%
  228.88MB 98.40% 98.40%   230.60MB 99.14%  main.main
    1.72MB  0.74% 99.14%     1.72MB  0.74%  runtime/pprof.StartCPUProfile
       1MB  0.43% 99.57%     1.50MB  0.65%  runtime.allocm
         0     0% 99.57%   230.60MB 99.14%  runtime.main
         0     0% 99.57%     1.50MB  0.65%  runtime.newm
         0     0% 99.57%     1.50MB  0.65%  runtime.resetspinning
         0     0% 99.57%     1.50MB  0.65%  runtime.schedule
         0     0% 99.57%     1.50MB  0.65%  runtime.startm
         0     0% 99.57%     1.50MB  0.65%  runtime.wakep

			with runtime.GC()
		$ go tool pprof mem.prof
Type: inuse_space
Time: May 31, 2022 at 12:06am (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 2721.04kB, 100% of 2721.04kB total
Showing top 10 nodes out of 20
      flat  flat%   sum%        cum   cum%
 1184.27kB 43.52% 43.52%  1184.27kB 43.52%  runtime/pprof.StartCPUProfile
  512.56kB 18.84% 62.36%   512.56kB 18.84%  runtime.allocm
  512.20kB 18.82% 81.18%   512.20kB 18.82%  runtime.malg
     512kB 18.82%   100%      512kB 18.82%  os.newFile
         0     0%   100%  1184.27kB 43.52%  main.main
         0     0%   100%      512kB 18.82%  os.NewFile
         0     0%   100%      512kB 18.82%  os.init
         0     0%   100%      512kB 18.82%  runtime.doInit
         0     0%   100%  1696.28kB 62.34%  runtime.main
         0     0%   100%   512.56kB 18.84%  runtime.mstart
(pprof)
```