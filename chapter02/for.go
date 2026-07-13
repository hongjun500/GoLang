// @author hongjun500
// @date 2023/8/11 17:25
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package main

import (
	"fmt"
	"log"
)

func forExample() {
	// 循环只有for
	// 不需要小括号，但大括号是必须
	// 分号 ; 隔开的三个表达式
	// 初始化语句，条件表达式，后置语句
	sum := 0
	for i := 0; i < 100; i++ {
		sum++
	}
	fmt.Println("计算 0 到 100 的和，结果sum值为：", sum)

	// 初始化语句和后置语句不是必须 例如：
	i := 100 // 本质上这里的 i 赋值 100 就是初始化语句
	for i > 0 {
		i-- // 而这里的 i-- 则是后置语句，作用其实就是改变初始语句赋的值
		// 使其条件表达式的值为 false 结束循环，否则将陷入无限循环
	}
	fmt.Println("i值", i)

	log.Println("break 关键字部分-------------------------------------")

	// break 关键字可以结束其所在循环结构体的执行（这个示例会打印出 4 个 0123456789）
	for i := 0; i < 10; i++ {
		if i == 4 {
			break
			// 如果满足了 i==0 时就会结束所在的一层循环
		}
		for i := 0; i < 10; i++ {
			fmt.Print(i)
		}
		println("")
	}
	log.Println("continue 关键字部分-------------------------------------")

	// continue 关键字可以结束条件判断其所在循环的执行（这个示例会打印出 10 个 012356789）
	for i := 0; i < 10; i++ {
		for i := 0; i < 10; i++ {
			if i == 4 {
				continue
				// 如果满足了 i==0 时就会结束所在的一层循环
			}
			fmt.Print(i)
		}
		println("")
	}

	log.Println("类似于 java 中的 while 循环部分-------------------------------------")
	whilesum := 0
	for i < 100 {
		whilesum = whilesum + i
		i++
	}
	fmt.Println("whilesum = ", whilesum)

	a := 0
	for a < 10 {
		if a != 0 && a%2 == 0 {
			fmt.Println("偶数", a)
		}
		a++
	}

	log.Println("range 关键字部分-------------------------------------")
	// range 关键字可以遍历数组、切片、字符串、map、通道
	str := "hello world"
	for k, v := range str {
		fmt.Printf("对于字符串 %s，下标为%d, 的值为%v \n", str, k, string(v))
	}

	log.Println("range 关键字部分只遍历下标即索引-------------------------------------")
	// 也可以只遍历下标
	for k := range str {
		fmt.Printf("对于字符串 %s，下标为%d \n", str, k)
	}

	log.Println("range 关键字部分对于切片的遍历-------------------------------------")
	var arr []string
	arr = append(arr, "见到", "你,", "很高兴！")
	// for 的遍历使用关键字 range
	// 1.只获取下标
	for key := range arr {
		fmt.Println(key)
	}
	// 2.只获取值
	for _, v := range arr {
		fmt.Println(v)
	}

	rangeStringDemo()
}

func rangeStringDemo() {
	log.Println("range 关键字部分对于字符串的遍历-------------------------------------")
	str := "你好, world"
	for k, v := range str {
		fmt.Printf("对于字符串 %s，下标为%d, 的值为%v \n", str, k, string(v))
	}
	log.Println("对比使用 for 循环遍历字符串的方式")
	for i := 0; i < len(str); i++ {
		fmt.Printf("对于字符串 %s，下标为%d, 的值为%v \n", str, i, string(str[i]))
	}
	fmt.Printf("字符串 【%s】的字节数为%d \n", str, len(str))

	runes := []rune(str)   // 把字符串按字符拆成 rune 切片
	fmt.Println(len(runes))  // 这才是"字符个数"（9），而不是字节数（13）
	fmt.Println(string(runes[0])) // 按"第几个字符"取值，比如取第0个字符"你"
}
