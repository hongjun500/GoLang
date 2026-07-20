package main

import (
	"fmt"
	"reflect"
)

// MyInt 自定义类型，用来演示 Type（具体类型）与 Kind（底层种类）的区别
type MyInt int

func reflectBasic() {
	var x float64 = 3.4

	// TypeOf 返回 reflect.Type：接口里存的"动态类型"
	t := reflect.TypeOf(x)
	// ValueOf 返回 reflect.Value：接口里存的"动态值"
	v := reflect.ValueOf(x)

	fmt.Println("Type :", t)        // float64
	fmt.Println("Kind :", v.Kind()) // float64
	fmt.Println("Value:", v.Float())

	// Type ≠ Kind：Type 是"你声明的具体类型"，Kind 是"编译器眼里的底层种类"
	var mi MyInt = 7
	fmt.Println("Type =", reflect.TypeOf(mi))        // main.MyInt
	fmt.Println("Kind =", reflect.TypeOf(mi).Kind()) // int
	// 结论：判断"是不是整数一族"要看 Kind，判断"是不是某个具体类型"才看 Type
}

// reflectBackToInterface 演示反射第二定律：从反射对象回到接口值。
// Interface() 把 reflect.Value 重新装箱成 any，再用类型断言取回原值。
func reflectBackToInterface() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	// v.Interface() 的静态类型是 any，动态类型才是 float64，所以必须断言
	back := v.Interface().(float64)
	fmt.Printf("Interface() 装箱回原值: %v (类型 %T)\n", back, back)
}