# 性能诊断与调优

> 一句话原则：**先测量，再优化**。凭直觉猜性能瓶颈，十有八九猜错。

## 逃逸分析 (Escape Analysis)

> 编译器在**编译期**决定变量分配在栈上还是堆上。

- 栈：随函数返回自动回收，零 GC 压力，快
- 堆：需要 GC 追踪回收，有开销
- 规则：编译器只要**无法证明**变量生命周期不超过当前函数，就让它逃逸到堆

```bash
go build -gcflags="-m" .
# ./escape.go:NN: moved to heap: u
# ./escape.go:NN: &User{...} escapes to heap
```

常见的"隐性逃逸"来源：
- 返回局部变量的地址
- 把具体值赋给 `interface{}` / `any`（装箱）
- 闭包捕获的变量
- 大小在编译期不确定的分配（如 `make` 长度是变量）

## GC 与内存

> `Go` 用并发的**三色标记-清除**垃圾回收器，主打低延迟。

两个最该知道的调节旋钮：

| 旋钮 | 作用 | 备注 |
| ---- | ---- | ---- |
| `GOGC` | 触发下次 GC 的堆增长比例，默认 `100`（堆翻倍时触发） | 调大→GC 更少吞吐更高但更吃内存；`GOGC=off` 关闭；运行期用 `debug.SetGCPercent` |
| `GOMEMLIMIT` | 软内存上限（Go 1.19+） | 逼近上限时更激进 GC，容器里防 OOM 利器 |

```bash
GODEBUG=gctrace=1 go run .   # 每次 GC 打印一行统计
```

代码里用 `runtime.ReadMemStats` 可以观察 `HeapAlloc` / `NumGC` 等指标（见 `gc.go`）。

## pprof 剖析

两种接入方式：

**方式一 · 基准测试出 profile（最常用）**

```bash
go test -bench=. -benchmem .                          # 先看哪段慢/分配多
go test -bench=BenchmarkConcatPlus -cpuprofile=cpu.out .
go tool pprof cpu.out            # 交互式：top / list 函数名 / web
go tool pprof -http=:8081 cpu.out  # 浏览器看火焰图
```

**方式二 · 长跑服务在线剖析**

```go
import _ "net/http/pprof"
go func() { log.Println(http.ListenAndServe("localhost:6060", nil)) }()
```

```bash
# CPU：采 30 秒（期间要有真实负载）
go tool pprof -http=:6061 http://localhost:6060/debug/pprof/profile?seconds=30

# 堆内存（当前 in-use）
go tool pprof -http=:6061 http://localhost:6060/debug/pprof/heap

# 分配量（allocs，看谁分配多，不只看当前占用）
go tool pprof -http=:6061 http://localhost:6060/debug/pprof/allocs

# goroutine
go tool pprof -http=:6061 http://localhost:6060/debug/pprof/goroutine

# 阻塞（需 SetBlockProfileRate）
go tool pprof -http=:6061 http://localhost:6060/debug/pprof/block

# 锁竞争（需 SetMutexProfileFraction）
go tool pprof -http=:6061 http://localhost:6060/debug/pprof/mutex
```

读 pprof 两个关键指标：
- `flat`：函数**自身**消耗（不含子函数）→ 定位自身热点
- `cum`：**累计**消耗（含子函数）→ 定位总耗时大户

## 基准测试与 PGO

- 写法见 `bench_test.go`：`BenchmarkConcatPlus`（`+=` 拼接）vs `BenchmarkConcatBuilder`（`strings.Builder`），跑一下能直观看到 `ns/op` 和 `allocs/op` 差一到两个数量级
- Go 1.20+ 引入 **PGO（Profile-Guided Optimization，剖析引导优化）**：把生产环境采到的 `cpu.out` 命名为 `default.pgo` 放到 `main` 包目录，`go build` 会自动读取，用真实热点指导内联等优化，通常能白捡几个百分点的性能

```bash
# 采集生产/基准 profile → 命名为 default.pgo → 正常 build 即自动启用
go test -bench=. -cpuprofile=default.pgo .
go build .   # 自动检测到 default.pgo 并启用 PGO
```

## 底层原理解析

- **为什么逃逸这么重要**：栈分配就是移动一下栈指针，几乎零成本；堆分配要走内存分配器 + 后续 GC。减少逃逸 = 减少 GC 压力，这是 Go 性能优化的第一性原理
- **三色标记**：白（待回收）、灰（已发现待扫描）、黑（已扫描存活）。GC 与用户代码**并发**执行，靠写屏障(write barrier)保证并发正确性，从而把 STW（Stop-The-World）压到亚毫秒级
- **`GOGC` 的本质**是"用内存换 CPU"：堆允许长得越大，GC 频率越低，单位时间花在 GC 上的 CPU 越少，但峰值内存越高
- **`strings.Builder` 为何快**：它内部持有一个 `[]byte` 并按扩容策略增长，`String()` 通过 `unsafe` 零拷贝转换；而 `+=` 每次都分配新字符串并整体复制旧内容
- **`b.Loop()`（Go 1.24+）**：比手写 `for i := 0; i < b.N; i++` 更准，能防止被测代码被编译器优化消除，且自动处理计时

## 避坑指南

- **过早优化**：没有 profile 数据就动手优化 = 浪费时间 + 增加复杂度。先 `-bench`/pprof 定位，再动手
- **benchmark 被优化掉**：结果没被使用时，编译器可能整段删掉，测出"0 ns/op"。用 `b.Loop()`，或把结果赋给包级变量 `sink`
- **误读 flat/cum**：想找"哪个函数自己最慢"看 `flat`；想找"哪条调用路径最贵"看 `cum`。二者混淆会把优化方向带偏
- **手动 `runtime.GC()` / `debug.FreeOSMemory()` 上生产**：几乎总是坏主意，会打乱 GC 的自适应节奏。只在测试/演示里用
- **内存"泄漏"在 Go 里多是"忘了释放引用"**：常见于 goroutine 泄漏、切片截取后仍持有底层大数组（`small := big[:10]` 会让整个 `big` 无法回收）、全局 map 只增不删
- **`time.Now()` 手动计时代替 benchmark**：单次计时受调度/GC 抖动影响极大，不可靠。基准测试应交给 `testing.B`
- **PGO 用了过期的 profile**：profile 要能代表真实负载；用错负载采的 profile 反而可能帮倒忙
