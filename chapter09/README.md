# 并发

> `Go` 的并发哲学：**不要通过共享内存来通信，而要通过通信来共享内存**。

## `Goroutine` `Go` 运行时管理的轻量级**线程**

```go
go f(x, y, z)
```

- 初始栈大小为 2KB, 可动态伸缩，一个进程轻松跑几十万个 `goroutine`
- 由 `Go` `runtime` 的 `GMP` **调度器**在用户态调度，切换成本远低于 OS 线程，且不需要用户自己管理线程池
- `main` 函数本身就跑在一个 `goroutine` 中, 当 `main` 函数返回时, 该进程退出, 所有的 `goroutine` 都会被强制结束

### 如何等待 `goroutine` 结束? 用 `sync.WaitGroup` 或者 `channel` 来等待 `goroutine` 结束

```go
var wg sync.WaitGroup
func worker () {
	fmt.Println("worker")
}
for range 10 {
	wg.Go(func(){
		worker()
	})
}
wg.Wait()
```

## `Channel` 类型化的通信管道

- 通道是**引用类型**，**零值**是 `nil`
- 声明：`var Name chan T`
- 初始化：`ch := make(chan T, [size])`, size 表示缓冲区大小, 如果不设置 size, 则为无缓冲通道
  示例：

```go
ch := make(chan int)

ch <- v    // 将 v 发送至通道 ch
v := <-ch  // 从 ch 接收值并赋予 v
```

| 类型                     | 发送何时阻塞                           | 收何时阻塞       |
| ------------------------ | -------------------------------------- | ---------------- |
|                          |                                        |                  |
| 无缓冲 `make(chan T)`    | 直到有接收方就绪（收发必须同步准备好） | 直到有发送方就绪 |
| 有缓冲 `make(chan T, n)` | 缓冲区满时                             | 缓冲区空时       |

> 通道通过 `make` 创建, 默认是无缓冲的, 也就是说只有在对应的接收（`<- chan`） 通道准备好接收时, 才允许进行发送（`chan <-`） 通道的操作。

```go
ch := make(chan int, 2)
cap(ch) // 2  容量
len(ch) // 当前缓冲区里的元素个数
```

### 遍历与关闭

```go
for v := range ch { // 通道关闭且读空后循环自动结束
    fmt.Println(v)
}
```

- 只在 **需要通知接收方** **不会再有值** 时，才关闭通道
- 由发送方关闭通道，绝不由接收方关闭通道（否则发送方可能向已关闭通道发送 → `panic`））

### 关闭通道的三条铁律

- 向已关闭的通道发送 → panic: send on closed channel
- 关闭已关闭的通道 / 关闭 nil 通道 → panic
- 从已关闭通道接收：先把缓冲区剩余值读完，读空后返回 零值, `false`：

```go
v, ok := <-ch // ok == false 表示"通道已关闭且已读空"
```

> ⚠️ 常见误区：并非一关闭就返回零值——缓冲区里的存量数据会先被正常读出。

## 通道的选择器(多路复用) select

```go
select {
case v := <-ch1:
    // ch1 可接收
case ch2 <- x:
    // ch2 可发送
case <-time.After(time.Second):
    // 超时兜底
default:
    // 所有 case 都未就绪时立即执行（非阻塞轮询）
}
```

- `select` 会阻塞到某个分支可以继续执行为止, 当多个分支都准备好的时候会随机选择一个执行;
- `default` 有此分支且当其它所有 `case` 分支都在阻塞时就会执行立刻执行。
- `nil channel` 技巧：对于 `nil` 通道的收发永远阻塞，因此可以把某个 `case` 的通道临时置为 `nil` 来"动态关闭"该分支。

## 互斥锁与原子操作

> 当确实需要共享状态时，用锁保护：

```go
type SafeCounter struct {
    mu sync.Mutex
    v  map[string]int
}
func (c *SafeCounter) Inc(k string) { c.mu.Lock(); defer c.mu.Unlock(); c.v[k]++ }
```

- `sync.Mutex` 互斥锁
- `sync.RWMutex`：读写锁，读多写少时读并发更高
- `sync/atomic` 无锁提供原子操作，适合简单的计数器等场景，性能由于 ``Mutex` 锁的竞争而更高
- `sync.Once`：保证某段代码（如初始化/单例模式）**只执行一次**
- `sync.WaitGroup`：等待一组 `goroutine` 完成
- `sync.Cond`：条件变量，配合锁使用，等待/通知某个条件发生

## `Context` 上下文,控制 `goroutine` 生命周期

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()
select {
case <-ctx.Done():
    return ctx.Err() // 超时或被取消
case res := <-work:
    return res
}
```

- 用于跨 `goroutine` **传递取消信息**、**超时**、**截止时间**、**请求域数据**
- HTTP/gRPC 服务器为每个请求携带 `ctx`, 请求结束即取消，防止`goroutine` 泄漏
- `context.Background()`：空上下文，通常用于主函数、初始化和测试

## 底层原理解析

### GMP 调度模型

- `G`（`Goroutine`）：待执行的任务；`M`（`Machine`）：`OS` 线程；`P`（`Processor`）：调度上下文，数量默认 = `GOMAXPROCS`（= `CPU` 核数）。
- `P` 持有一个本地可运行 `G` 队列，`M` 必须绑定 `P` 才能执行 `G`；空闲 `P` 会 `work-stealing`（从别的 `P` 偷 `G`），实现负载均衡。
- 当 `G` 发生阻塞系统调用时，`M` 会与 `P` 解绑，`P` 转交给其他 `M`，避免"一个阻塞拖垮全部"

### `Go` 内存模型

- 单个 `goroutine` 内：代码顺序即执行顺序（对外可见性除外）。
  跨 `goroutine` 的可见性需要同步事件建立 `happens-before 关系，常见的有：
  - 对通道的发送 `happens-before` 对应的接收完成；
  - 通道的关闭 `happens-before` 因关闭而返回的接收；
    \-` Mutex.Unlock happens-before` 后续的 `Lock`。
- 本章 [`channel.go`](channel.go) 就是活教材：`<-c` 返回后，`f` 中对 `a` 的写入保证可见，无需额外加锁

## 避坑指南

1. `goroutine` 泄漏：`goroutine` 阻塞在收/发上且永远无法继续，就永久泄漏（内存只增不减）:

- 典型：接收方提前退出，发送方还阻塞在 `ch <-` 上。
- 解决：`close` 通道、用带缓冲通道、或用 `context` 取消

2. 数据竞争（data race）：多个 goroutine 无同步地读写同一变量，结果未定义

> 一定要用 `race detector` 检测：`go run -race .` / `go test -race ./...`

3. `for range` 循环变量捕获（`Go` 1.22 之前的经典坑）：

```go
// Go 1.22 之前：所有 goroutine 打印同一个 i（循环结束后的值）
for i := 0; i < 3; i++ {
    go func() { fmt.Println(i) }()   // 老版本需 i := i 拷贝一份
}
```

> Go 1.22+ 每次迭代都是新变量，此坑已修复；但老代码/老版本仍需注意

4. 无缓冲通道死锁：主 `goroutine` 里 `ch <- 1` 却没有任何接收方 → `fatal error: all goroutines are asleep - deadlock!`
5. `WaitGroup` 用错：`wg.Add` 必须在 `go` 之前；`Add` 的数量要和 `Done` 对齐；不要拷贝 `WaitGroup`（传指针）
6. `time.Tick` 泄漏：它返回的 `Ticker` 无法回收，只用于整个程序生命周期；短期用 `time.NewTicker` + `defer t.Stop()`
7. `Mutex` 不可拷贝：`sync.Mutex`/`WaitGroup` 一旦使用就不能按值拷贝，结构体嵌了锁就要用指针接收者/传指针。`go vet` 会检查
