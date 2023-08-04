package main

import (
	"time"

	"github.com/hongjun500/GoLang-master/chapter09/find"
)

func main() {
	// 输出结果为 hello world hello world hello world hello world，不一定是这个顺序，因为两个协程是并发执行的
	go say("world")
	say("hello")

	// 通道（channel）是用来传递数据的一个数据结构
	s := []int{1, 7, 6, 0, -7}

	// 通道是引用类型，通道类型的空值是 nil
	// 通道的声明形式：var 通道名 chan 通道类型
	// 通道的初始化形式：通道名 = make(chan 通道类型, [缓冲大小])
	c := make(chan int)
	// {1, 7}
	s1 := s[:len(s)/2]
	// {6, 0, -7}
	s2 := s[len(s)/2:]
	go sum(s1, c)
	go sum(s2, c)
	// 从通道 c 中接收数据，将其赋值给 x 和 y，因为是并发执行的，所以 x 和 y 的值不一定是 s1 和 s2 的和
	x, y := <-c, <-c // 从通道 c 中接收
	println("x=", x, "y=", y)

	c2 := make(chan int, 1)
	c2 <- 1
	println("通道 c2 的容量：", cap(c2))

	closeCh()

	println("select 语句部分------------------")
	sCh := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 2; i++ {
			// 从通道 sCh 中接收数据，对应的 case 语句不再阻塞
			println(<-sCh)
		}

		quit <- 0
	}()
	selectCh(sCh, quit)

	println("time 包的定时时间和 select 语句的default分支")
	times()

	println("等价二叉树部分------------------")
	tree1 := find.New(1)
	println("tree1 的值为：", tree1.String())
	tree2 := find.New(1)
	println("tree2 的值为：", tree2.String())
	println("tree1 和 tree2 是否存储了相同的值：", find.Same(tree1, tree2))

	println("通道的通信部分------------------")
	Hello()
}

func say(s string) {
	for i := 0; i < 5; i++ {
		// 休眠 100 毫秒
		time.Sleep(100 * time.Millisecond)
		println(s)
	}
}

func sum(s []int, c chan int) {
	sums := 0
	for _, v := range s {
		sums += v
	}
	c <- sums // 将和送入 c
}

func closeCh() {
	// 关闭通道
	// 关闭通道后，通道不再接收数据，但是仍然可以从通道中读取数据
	// 关闭一个已经关闭的通道会导致运行时恐慌
	// 通道和文件不同，通道通常不需要关闭，只有在必须告诉接收者不再有值需要发送的时候才有必要关闭，例如终止一个 range 循环
	// 关闭通道后，仍然可以从通道中读取数据，但是读取的数据是通道类型的零值
	println("通道部分------------------")
	cc := make(chan int, 10)
	cc <- 1
	cc <- 2
	cc <- 3

	// 如果下面这行注释了会导致运行时恐慌，因为 for range 循环会一直读取通道 cc 中的数据，直到通道 cc 被关闭
	close(cc)
	println("通道 cc 已经关闭")
	for i := range cc {
		println("读取通道 cc 中的数据：", i)
	}
	// 由于上面的 cc 已经关闭，所以这里的 x, y, z 都是 int 类型的零值 0
	x, y, z := <-cc, <-cc, <-cc
	println("通道被关闭后，读取的数据是通道类型的零值：x, y, z 为", x, y, z)

}

func selectCh(ch, quit chan int) {
	x, y := 99, 1
	time.Tick(100 * time.Millisecond)
	for {
		select {
		// 检查通道 ch 是否可以进行发送数据
		// 如果通道没有接收数据，就会阻塞
		case ch <- x:
			x, y = y, x+y
		// 同理，检查通道 ch 是否可以进行接收数据
		// 如果通道没有发送数据，就会阻塞
		case <-quit:
			println("quit")
			return
		}
	}
}

func times() {
	// time.Tick() 函数返回一个通道（channel），该通道会定期（每隔一段时间）发送一个事件（当前时间）
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
			println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
