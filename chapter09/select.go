// @author hongjun500
// @date 2023/8/8 14:58
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package main

import "time"

func selectCh(ch, quit chan int) {
	go func() {
		for i := 0; i < 2; i++ {
			// 从通道 sCh 中接收数据，对应的 case 语句不再阻塞
			println(<-ch)
		}

		quit <- 0
	}()

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
