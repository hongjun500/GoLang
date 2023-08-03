package main

import "time"

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

	println("select 语句")
	sCh := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 2; i++ {
			println(<-sCh)
		}
		quit <- 0
	}()
	selectCh(sCh, quit)
}

func selectCh(ch, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			println("quit")
			return
		}
	}
}
