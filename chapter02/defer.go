package main

import "fmt"

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	// 执行步骤 6 将 un 函数压入栈中，栈2
	// 执行步骤 7 执行函数 trace("a")，输出 entering: a 并将返回值 "a" 传递给 un 函数
	defer un(trace("a"))
	// 执行步骤 8
	fmt.Println("in a")
}

func deferExample() {
	// 执行步骤 2 将 un 函数压入栈中，栈1
	// 执行步骤 3 执行函数 trace("b")，输出 entering: b 并将返回值 "b" 传递给 un 函数
	defer un(trace("b"))
	// 执行步骤 4
	fmt.Println("in b")
	// 执行步骤 5
	a()
}

