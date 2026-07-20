package main

import (
	"fmt"
	"reflect"
)

func reflectSettable() {
	var x float64 = 3.4
	// 反例：直接 ValueOf(x) 传的是副本，改它影响不到原变量，所以不可设置
	v := reflect.ValueOf(x)
	fmt.Println("CanSet(值传递):", v.CanSet()) // false
	// v.SetFloat(7.1) // panic: using unaddressable value
	// 正解：传 &x（指针），再 Elem() 解引用，拿到"指向 x 本身"的可寻址 Value
	elem := reflect.ValueOf(&x).Elem()
	fmt.Println("CanSet(指针+Elem):", elem.CanSet()) // true
	elem.SetFloat(7.1)
	fmt.Println("反射修改后 x =", x) // 7.1
}
