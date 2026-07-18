package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/hongjun500/GoLang/chapter09/find"
)

func main() {
	log.Println("goroutine + WaitGroup（正确的等待方式）----------")
	// 输出结果的顺序每次都不一样，因为两个协程是并发执行的
	var wg sync.WaitGroup
	/*  wg.Add(1)
	go func() {
		defer wg.Done() // defer 保证即使 panic 也会 Done，避免死等
		say("world")
	}()
	*/
	wg.Go(func() {
		say("world")
	})
	say("hello")
	wg.Wait() // 阻塞直到计数归零

	log.Println("无缓冲通道：并发求和---------------")
	// 通道是引用类型，通道类型的空值是 nil
	// 通道的声明形式：var 通道名 chan 通道类型
	// 通道的初始化形式：通道名 = make(chan 通道类型, [缓冲大小])
	// 简而言之，零值是 nil；声明 var c chan int，初始化 make(chan int, [size])

	s := []int{1, 7, 6, 0, -7}
	c := make(chan int)
	s1 := s[:len(s)/2]
	s2 := s[len(s)/2:]
	go sum(s1, c)
	go sum(s2, c)
	// 谁先算完谁先被接收，所以 x / y 的先后不确定，但两者之和确定
	// 从通道 c 中接收数据，将其赋值给 x 和 y，因为是并发执行的，所以 x 和 y 的值不一定是 s1 和 s2 的和
	x, y := <-c, <-c // 从通道 c 中接收
	fmt.Println("x =", x, "y =", y, "总和 =", x+y)

	log.Println("有缓冲通道------------")
	c2 := make(chan int, 1)
	c2 <- 1 // 缓冲未满，不阻塞
	fmt.Printf("通道 c2 的长度 len=%d, 容量 cap=%d\n", len(c2), cap(c2))
	closeCh()

	log.Println("select 语句部分------------------")
	sCh := make(chan int)
	quit := make(chan int)
	selectCh(sCh, quit)

	log.Println("time 包的定时时间和 select 语句的 default 分支--------")
	times()

	log.Println("互斥锁 sync.mutex：并发安全计数器-------")
	counter := NewSafeCounter()
	var wg2 sync.WaitGroup
	for range 100 {
		// wg2.Add(1)
		// go func() {
		// 	defer wg2.Done()
		// 	counter.Inc("key")
		// }()
		wg2.Go(func() {
			counter.Inc("key")
		})
	}
	wg2.Wait()
	fmt.Println("100 个 goroutine 并发自增后 key =", counter.Value("key")) // 稳定输出 100

	log.Println("等价二叉树------------------")
	// New(1) 与 New(1) 值集相同 -> true；New(1) 与 New(2) 不同 -> false
	fmt.Println("New(1) 与 New(1) 是否相同:", find.Same(find.New(1), find.New(1)))
	fmt.Println("New(1) 与 New(2) 是否相同:", find.Same(find.New(1), find.New(2)))

	log.Println("通道通信（内存模型 happens-before）------------------")
	Hello()
}

func say(s string) {
	// for i := 0; i < 5; i++ {
	for range 5 { // Go 1.22 之后引入的 for range over int
		// 休眠 100 毫秒
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// sum 函数计算切片 s 的和，并将结果发送到通道 c 中
func sum(s []int, c chan int) {
	sums := 0
	for _, v := range s {
		sums += v
	}
	c <- sums // 将和送入 c
}

func closeCh() {
	// 关闭通道的语义：
	//  - 关闭后不能再发送（会 panic），但仍可接收
	//  - 缓冲区里已有的值会被依次读完，读完后才返回 "零值 + ok=false"
	//  - 通道不像文件，通常不需要显式关闭，只在"需要通知接收方不再有值"时关闭（如结束 range）
	log.Println("通道关闭部分------------------")
	cc := make(chan int, 10)
	cc <- 1
	cc <- 2
	cc <- 3

	// 如果下面这行注释了会导致运行时恐慌，因为 for range 循环会一直读取通道 cc 中的数据，直到通道 cc 被关闭
	close(cc)

	// range 会把缓冲区里剩余的值全部读出，通道关闭且读空后自动结束循环
	for v := range cc {
		fmt.Println("range 读取数据：", v)
	}
	// 由于上面的 cc 已经关闭，所以这里的 x, y, z 都是 int 类型的零值 0
	x, y, z := <-cc, <-cc, <-cc
	fmt.Printf("通道被关闭后，读取的数据是通道类型的零值：x:%d, y:%d, z:%d\n", x, y, z)

	// 关闭且读空后，comma-ok 的 ok 为 false，值为类型零值
	v, ok := <-cc
	fmt.Printf("读空后再接收(comma-ok形式展示通道类型零值): value=%d, ok=%v\n", v, ok)
}

func times() {
	// 注意：time.Tick 返回的 Ticker 无法被回收，只适合"整个程序生命周期都要用"的场景。
	// 短生命周期请用 time.NewTicker + defer ticker.Stop()
	tick := time.Tick(100 * time.Millisecond)
	// time.After() 函数返回一个通道（channel），该通道会在指定时间后发送一个事件（当前时间）
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			println("tick.")
		case <-boom:
			println("BOOM!")
			return
		default:
			// 所有 case 都未就绪时立即执行 default，避免阻塞（这就是"忙等/轮询“）
			println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
