package main

import (
	"fmt"
)

func main() {
	// 数组
	arrExample()

	// 切片
	sliceExample()
	// map
	mapExample()

	// sync.Map
	syncMapExample()

	// 切片映射
	mapSliceExample()

	// 向切片追加元素：nil 切片不能按下标赋值，应使用 append
	var out1 []string
	for i := 0; i < 10; i++ {
		// out[i] = fmt.Sprintf("%d", i) // 编译通过，但是运行报错，因为 out1 初始 为 nil，不能直接通过索引赋值
		out1 = append(out1, fmt.Sprintf("%d", i))
	}
	fmt.Println("out1:", out1)

	// 若必须按下标赋值，需先分配长度
	out2 := make([]string, 10)
	for i := 0; i < 10; i++ {
		out2[i] = fmt.Sprintf("%d", i)
	}
	fmt.Println("out2:", out2)
}
