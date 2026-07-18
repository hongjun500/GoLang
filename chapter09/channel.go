// @author hongjun500
// @date 2023/8/4 17:31
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description: 通道通信演示 Go 内存模型的 happens-before 规则
//
// 内存模型保证：对通道的一次发送，happens-before 对应的接收「完成」。
// 因此 <-c 返回后，f 中对 a 的赋值一定对 Hello 可见（无需额外加锁）。
package main

import "fmt"

var c = make(chan int, 10)
var a string

func f() {
	a = "你好, golang" // 断点执行顺序 1
	c <- 0           // 断点执行顺序 2, 发送 happens-before 接收完成
}

func Hello() {
	go f()
	// 这里会阻塞，直到 c 中有数据
	<-c // 接收完成后，顺序 1 的写入对当前 goroutine 保证可见

	fmt.Println(a) // 断点执行顺序 3
}
