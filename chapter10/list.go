// @author hongjun500
// @date 2023/8/24 10:59
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description: 反射访问未导出（私有）字段——能读，但不能 Interface()、不能 Set()

package main

import (
	"container/list"
	"fmt"
	"reflect"
)

// reflectUnexported 演示反射访问未导出字段。
//
// container/list.List 的 len 字段是未导出的：
//
//	type List struct {
//		root Element
//		len  int // 未导出
//	}
//
// 三层关键结论（务必记牢）：
//  1. 能"读"未导出字段的值：用 Int()/String() 等类型化取值方法；
//  2. 不能对它调用 Interface()：会 panic（禁止把私有值"逃逸"成 any）；
//  3. 不能 Set()：未导出字段的 Value 带只读标记 flagRO，改它必 panic。
func reflectUnexported() {

	l := list.New()
	l.PushFront("foo")
	l.PushFront("bar")

	// Elem() 解引用 *list.List，FieldByName 按名字取字段
	fv := reflect.ValueOf(l).Elem().FieldByName("len")
	// 1 读：OK
	fmt.Printf("反射读未导出字段 len = %d\n", fv.Int())
	// 2，3 是否可设置：一定是 false
	fmt.Println("未导出字段 CanSet:", fv.CanSet()) // false

	// 下面两行都会 panic，演示"能读不能改"：
	// _ = fv.Interface() // panic: cannot return value obtained from unexported field
	// fv.SetInt(100)     // panic: using value obtained from unexported field
}
