package main

import "errors"

// ErrDivByZero 是被测函数返回的哨兵错误，测试里会用 errors.Is 断言它。
var ErrDivByZero = errors.New("除数不能为 0")

// Divide 返回 a/b；b 为 0 时返回错误。既有正常路径又有错误路径，适合演示表驱动测试。
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return a / b, nil
}

// Fib 计算第 n 个斐波那契数（迭代版），用于演示基准测试。
func Fib(n int) int {
	if n < 2 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
