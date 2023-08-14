// @author hongjun500
// @date 2023/8/14 18:12
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package main

import "fmt"

func gotoFunc() {
	i := 0
HERE:
	fmt.Println(i) // i = 0 打印出0
	i++            // 此时 i = 1
	if i < 10 {    // i = 1会进入此判断结构体
		// 跳到HERE	标签处，执行打印，程序会接着执行i++,又进入条件判断以此类推直到i=10的时候程序运行结束
		goto HERE
	}
}
