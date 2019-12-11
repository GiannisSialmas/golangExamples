# Routine Profiling
```
go tool pprof http://127.0.0.1:8080/debug/pprof/goroutine
```

Result looks like
```
File: main
Type: goroutine
Time: Dec 11, 2019 at 11:45pm (EET)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top5
Showing nodes accounting for 6, 100% of 6 total
Showing top 5 nodes out of 30
      flat  flat%   sum%        cum   cum%
         3 50.00% 50.00%          3 50.00%  runtime.gopark
         1 16.67% 66.67%          1 16.67%  net/http.(*connReader).backgroundRead
         1 16.67% 83.33%          1 16.67%  runtime.notetsleepg
         1 16.67%   100%          1 16.67%  runtime/pprof.writeRuntimeProfile
         0     0%   100%          1 16.67%  internal/poll.(*FD).Accept
```


Goroutines spend most of their time with **runtime.gopark** which is fine, because it’s the goroutine 
scheduler. So this is kinda what we expected, we have no problem here.

# CPU Profiling
```
go tool pprof http://127.0.0.1:8080/debug/pprof/profile
```

# Heap profiling
```
go tool pprof http://127.0.0.1:8080/debug/pprof/heap
```
-alloc_object


???
-inuse_space — amount of memory allocated and not released yet
-inuse_object s— amount of objects allocated and not released yet
-alloc_space — total amount of memory allocated (regardless of released)
-alloc_objects — total amount of objects allocated (regardless of released)
???

# Block profile
```
go tool pprof http://127.0.0.1:8080/debug/pprof/block
```

# Escape analysis
`go build -gcflags=-m -o main main.go` to see which objects escape to the heap 
