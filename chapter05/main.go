package main

import (
	"fmt"
)

func main() {
	// 数组
	arrFunc()
	println("切片部分------------------")
	// 切片
	spliceFunc()

	println("map 部分------------------")
	maps()

	println("无锁 map 扩展部分-----------")

	syncMapFunc()

	mapSliceFunc()
	var out []string
	for i := 0; i < 10; i++ {
		//out = append(out, strconv.Itoa(i))
		out[i] = fmt.Sprintf("%d", i)
	}
	fmt.Println(out)
}
