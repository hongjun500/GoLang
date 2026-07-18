// @author hongjun500
// @date 2023/8/8 14:58
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:  select 多路复用（斐波那契生成器）

package main

import "fmt"

func selectCh(ch, quit chan int) {
	go func() {
		for i := 0; i < 10; i++ {
			// 从通道 sCh 中接收数据，对应的 case 语句不再阻塞
			fmt.Println(<-ch)
		}

		quit <- 0 // 通知生产端退出
	}()

	// 用 select 同时处理“发送下一个斐波那契数”和“接收退出通知”两件事
	x, y := 0, 1
	for {
		select {
		// 检查通道 ch 是否可以进行发送数据
		// 如果通道没有接收数据，就会阻塞
		case ch <- x:
			x, y = y, x+y
		// 同理，检查通道 ch 是否可以进行接收数据
		// 如果通道没有发送数据，就会阻塞
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
