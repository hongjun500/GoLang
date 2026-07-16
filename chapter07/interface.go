// @author hongjun500
// @date 2023/8/21 15:27
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description: 接口类型

package main

import "fmt"

// 编译期实现断言：一行零开销的”契约检查“
var (
	_ Study = Teacher{}       // Teacher 用值接收者实现 Info，值即可满足接口
	_ Study = (*Student)(nil) // *Student 用指针接收者实现 Info，只有指针能满足接口
)

// teacher 存的是 Teacher 值；student 存的是 *Student 指针。
// student 之所以必须是指针：Info() 用了指针接收者，Student 值不在其接口方法集内。
var (
	teacher, student Study
)

// Study 是本章核心接口：任何类型只要实现了 Info() 方法，
// 就"隐式"满足 Study（Go 的鸭子类型，无需显式 implements 声明）
type Study interface {
	Info()
}

// MyInt 演示"可以在本包自定义的命名类型上定义方法"（底层类型是 int）。
// 注意：不能直接给内建的 int/string 或其它包的类型定义方法。
type MyInt int

// IsPositive 是定义在 MyInt 上的方法。
func (m MyInt) IsPositive() bool {
	return m > 0
}

// describe 用类型断言的 comma-ok 写法，安全地探测接口底层的具体类型（失败不 panic）。
func describe(s Study) {
	if t, ok := s.(Teacher); ok {
		fmt.Printf("[断言] 这是一位老师：%s\n", t.Name)
		return
	}
	if st, ok := s.(*Student); ok {
		fmt.Printf("[断言] 这是一位学生：%s\n", st.Name)
		return
	}
	fmt.Println("[断言] 未知类型")
}

// assertType 用 type switch 一次性区分多种具体类型，v 即断言成功后的具体值。
func assertType(s Study) {
	switch v := s.(type) {
	case Teacher:
		fmt.Printf("[type switch] Teacher: %s\n", v.Name)
	case *Student:
		fmt.Printf("[type switch] *Student: %s\n", v.Name)
	default:
		fmt.Printf("[type switch] 未知类型: %T\n", v)
	}
}
