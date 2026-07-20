package main

import "log"

//
//	逃逸分析:   go build -gcflags="-m" .
//	基准测试:   go test -bench=. -benchmem .
//	CPU 剖析:   go test -bench=BenchmarkConcat -cpuprofile=cpu.out .
//	            go tool pprof cpu.out
//	内存剖析:   go test -bench=BenchmarkConcat -memprofile=mem.out .
//
// main 里只演示"看得见"的部分：GC/内存统计、逃逸的两种写法对比。
func main() {
	log.Println("逃逸分析：同一件事，栈上 vs 堆上----------")
	escapeDemo()
	log.Println("GC 与内存统计：runtime.ReadMemStats 观察堆----------")
	gcDemo()
	log.Println("pprof 接入方式（看注释，动手跑）----------")
	pprofDemo()
}
