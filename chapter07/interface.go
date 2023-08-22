// @author hongjun500
// @date 2023/8/21 15:27
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description: 接口类型

package main

var (
	teacher, student Study
)

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
