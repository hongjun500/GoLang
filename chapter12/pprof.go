package main

import "fmt"

// pprof 是 Go 内置的性能剖析工具，能采集 CPU、堆内存、goroutine、阻塞、锁竞争等画像。
// 有两种接入方式，按场景选一种。
func pprofDemo() {
	fmt.Println("pprof 用法见下方注释；本函数不真正采样，避免污染演示输出")

	// ============ 方式一：基准测试直接出 profile（最常用） ============
	// 见 bench_test.go 里的 BenchmarkConcat：
	//	go test -bench=. -benchmem .                              // 先看哪段慢/分配多
	//	go test . -bench=BenchmarkConcat -cpuprofile=cpu.out      // 采 CPU
	//	go test . -bench=BenchmarkConcat -memprofile=mem.out      // 采内存
	// 注意（Windows PowerShell）：不要写 -cpuprofile=cpu.out .
	// PowerShell 会把 cpu.out 拆成 cpu + .out，再加末尾的 . 就变成「两个包」，
	// 报错 cannot use -cpuprofile flag with multiple packages。
	// 包路径 . 放前面，或写成 -cpuprofile cpu.out（空格形式）即可。
	//	go tool pprof cpu.out        // 进入交互式：top / list 函数名 / web(需 graphviz)
	//	go tool pprof -http=:8080 cpu.out  // 直接开浏览器看火焰图

	// ============ 方式二：长跑服务在线剖析 ============
	// 只需匿名导入，pprof 会把 handler 注册到默认的 http.DefaultServeMux：
	//
	//	import _ "net/http/pprof"
	//	go func() { log.Println(http.ListenAndServe("localhost:6060", nil)) }()
	//
	// 然后访问 / 采集：
	//	go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30  // 30s CPU
	//	go tool pprof http://localhost:6060/debug/pprof/heap                // 堆
	//	curl http://localhost:6060/debug/pprof/goroutine?debug=2           // 所有 goroutine 栈

	// ============ 读 pprof 的关键指标 ============
	//	flat  ：函数自身消耗（不含它调用的子函数）
	//	cum   ：累计消耗（含子函数）—— 定位"总耗时大户"看 cum，定位"自身热点"看 flat
}
