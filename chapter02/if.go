// @author hongjun500
// @date 2023/8/11 17:30
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package main

import "math"

// if 语句的表达式
func pow(x, y, z float64) float64 {
	// math.Pow(x, y) 求 x 的 y 次方
	value := math.Pow(x, y)
	if value < z {
		return value
	}
	return z
}

// if 语句中的变量声明可以在连着 if 后面的 if/else 中使用
func powElse(x, y, z float64) float64 {
	// math.Pow(x, y) 求x的y次方
	if value := math.Pow(x, y); value < z {
		return value
	} else if value == 0 {
		return value
	} else if value == 100 {
		return value
	} else {
		value = 9999
	}
	// return value, 报错，if/else 语句已经结束，value 变量已经被销毁
	return z
}
