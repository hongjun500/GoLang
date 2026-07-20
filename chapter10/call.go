package main

import (
	"fmt"
	"reflect"
)

// Calculator 提供一个可被反射调用的导出方法
type Calculator struct{}

// Add 必须是可导出方法，反射才能拿到并调用（未导出方法不在反射可见的方法集里）
func (c Calculator) Add(a, b int) int {
	return a + b
}

// reflectCall 演示按名称动态调用方法
func reflectCall() {
	v := reflect.ValueOf(Calculator{})
	method := v.MethodByName("Add") // 通过方法名获取 reflect.Value

	// 实参要逐个包成 reflect.Value，再放进切片
	args := []reflect.Value{
		reflect.ValueOf(10),
		reflect.ValueOf(20),
	}
	result := method.Call(args)
	fmt.Println("反射调用 Add(10, 20) =", result[0].Int()) // 30
}