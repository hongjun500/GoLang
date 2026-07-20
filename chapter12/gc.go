package main

import (
	"fmt"
	"runtime"
)

// Go 用的是并发的三色标记-清除(tri-color mark & sweep)垃圾回收器，主打"低延迟"。
// 两个最该知道的调节旋钮：
//   - GOGC：触发下一次 GC 的堆增长比例，默认 100（即堆相对上次存活量翻倍时触发）。
//           调大 → GC 更少、吞吐更高，但峰值内存更大；调小 → 更省内存但 GC 更频繁。
//           关闭：GOGC=off。运行期可用 debug.SetGCPercent 调整。
//   - GOMEMLIMIT：软内存上限(Go 1.19+)。到达上限会更激进地 GC，适合容器里防 OOM。
//
// 观察 GC：GODEBUG=gctrace=1 go run .  会在每次 GC 时打印一行统计。
func gcDemo() {
	var m runtime.MemStats

	readAndPrint := func(tag string) {
		runtime.ReadMemStats(&m)
		// HeapAlloc：当前堆上仍在使用的字节数；NumGC：已发生的 GC 次数
		fmt.Printf("[%s] HeapAlloc=%d KB, TotalAlloc=%d KB, NumGC=%d\n",
			tag, m.HeapAlloc/1024, m.TotalAlloc/1024, m.NumGC)
	}

	readAndPrint("初始")

	// 故意制造一批垃圾：分配大量临时切片，让它们很快变成不可达
	garbage := make([][]byte, 0, 1000)
	for i := 0; i < 1000; i++ {
		garbage = append(garbage, make([]byte, 10*1024)) // 每个 10KB
	}
	readAndPrint("分配后")

	// 丢弃引用并手动触发一次 GC（生产代码几乎不该手动 GC，这里只为演示）
	garbage = nil
	runtime.GC()
	readAndPrint("GC 后")
	_ = garbage
}
