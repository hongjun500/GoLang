// @author hongjun500
// @date 2023/8/14 16:36
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package main

import "fmt"

func switchFunc() {
	// switch 语句 从上向下执行(可以没有小括号，大括号必须)
	// 同 if 一样可以有短变量声明
	// 每个 case 结构体后有默认的 break,
	// 并且只会运行表达式对应选中的 case，此时 switch 运行到此结束
	// 如果选中了 case 并且对应 case 包含一个 fallthrough 关键字语句，那么紧随其后的 case 也会执行

	// 可以没有 default 语句
	// 这个 switch 只会打印出 Go
	switch lang := "Go"; lang {
	case "Go":
		fmt.Println("Go.")
	case "Java":
		fmt.Println("Java.")
		fallthrough
	case "Python":
		fmt.Println("Python.")

		/*default:
		fmt.Printf("%s.\n", lang)*/
	}

	switch lang2 := "Go"; lang2 {
	case "Go":
		fmt.Print("Go And ")
		// 在这里停止，执行下一个 case 且只执行下一个
		fallthrough
		// 这里会报错，由于有 fallthrough 的存在，程序将会直接去执行下一个 case 结构体或者是 default 中的代码
		// fmt.Println("log")
	/*case "Java":
		fmt.Println(" Java.")
	case "Python":
		fmt.Println("Python.")
	*/
	default:
		fmt.Printf("%s.\n", lang2)
	}
}
