// @author hongjun500
// @date 2023/8/4 17:31
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description: goroutine 的 channel

package main

var c = make(chan int, 10)
var a string

func f() {
	a = "你好, golang"
	// 断点执行顺序 2
	c <- 0
}

func Hello() {
	go f()
	// 这里会阻塞，直到 c 中有数据
	// 断点执行顺序 1
	<-c
	// 断点执行顺序 3
	print(a)
}
