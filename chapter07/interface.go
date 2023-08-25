// @author hongjun500
// @date 2023/8/21 15:27
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description: 接口类型

package main

var (
	teacher, student Study
)

type MyInt int

var i int
var j MyInt

// 变量i和j具有不同的静态类型，尽管它们具有相同的基础类型，但如果不进行转换，它们就无法相互分配。

type Study interface {
	Info()
}

// 对接口实现类型断言
func assertType(i Study) {
	switch i.(type) {
	case Teacher:
		println("Teacher")
	case *Student:
		println("Student")
	default:
		println("default")
	}
}
