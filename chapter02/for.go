// @author hongjun500
// @date 2023/8/11 17:25
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package main

import "fmt"

func forfunc() {
	// 循环只有for
	// 不需要小括号，但大括号是必须
	// 分号 ; 隔开的三个表达式
	// 初始化语句，条件表达式，后置语句
	sum := 0
	for i := 0; i < 100; i++ {
		sum++
	}
	fmt.Println("sum值", sum)

	// 初始化语句和后置语句不是必须 例如：
	i := 100 // 本质上这里的 i 赋值 100 就是初始化语句
	for i > 0 {
		i-- // 而这里的 i-- 则是后置语句，作用其实就是改变初始语句赋的值
		// 使其条件表达式的值为 false 结束循环，否则将陷入无限循环
	}
	fmt.Println("i值", i)

	fmt.Println("break关键字部分-------------------------------------")

	// break 关键字可以结束其所在循环结构体的执行（这个示例会打印出 4 个 0123456789）
	for i := 0; i < 10; i++ {
		if i == 4 {
			break
			// 如果满足了 i==0 时就会结束所在的一层循环
		}
		for i := 0; i < 10; i++ {
			fmt.Print(i)
		}
		fmt.Print("\n")
	}
	fmt.Println("continue关键字部分-------------------------------------")

	// continue 关键字可以结束条件判断其所在循环的执行（这个示例会打印出 10 个 012356789）
	for i := 0; i < 10; i++ {
		for i := 0; i < 10; i++ {
			if i == 4 {
				continue
				// 如果满足了 i==0 时就会跳出这个满足这个条件的循环，而不是循环体
			}
			fmt.Print(i)
		}
		fmt.Print("\n")
	}
	fmt.Println("0-100的累加和-------------------------------------")

	sum = 0
	for i < 100 {
		sum = sum + i
		i++
	}
	fmt.Println("100 的累加和 sum = ", sum)

	// 下例类似于java中的while
	a := 0
	for a < 10 {
		// if语句不需要小括号，必需大括号，类似for
		if a != 0 && a%2 == 0 {
			fmt.Println("偶数", a)
		}
		a++
	}

	value := pow(2, 3, 100)
	fmt.Println(value)

	/*fmt.Println(Sqrt(2))
	fmt.Println(SqrtFloat64(2))*/

	// 这两个结果不一样
	fmt.Println(1 / 2)
	fmt.Println(1 / 2.0)
}
