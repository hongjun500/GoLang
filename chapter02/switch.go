// @author hongjun500
// @date 2023/8/14 16:36
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package main

import (
	"fmt"
	"log"
)

func switchExample() {
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

	switchMutilDemo()
	switchTypeDemo()
}

func switchMutilDemo() {
	log.Println("switch 表达式匹配多个值的情况----")
	mutil := []string{"Go", "go", "Golang", "GO"}
	for _, v := range mutil {
		switch v {
		case "Go", "go", "Golang", "GO":
			fmt.Printf("[%s] in gopls.\n", v)
		default:
			fmt.Printf("%s is not in gopls.\n", v)
		}
	}
}

func switchTypeDemo() {
	log.Println("switch 类型开关--------")
	var x any
	x = 42
	switch v := x.(type) {
	case int:
		fmt.Printf("x is an int: %d\n", v)
	case string:
		fmt.Printf("x is a string: %s\n", v)
	default:
		fmt.Printf("x is of unknown type\n")
	}
	var y any
	y = []string{"Go", "Java", "Python"}
	switch v := y.(type) {
	case []string:
		fmt.Printf("%v is a slice of string\n", y)
	default:
		fmt.Printf("%v is of unknown type\n", v)
	}
}