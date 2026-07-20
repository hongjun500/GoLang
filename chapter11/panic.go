package main

import (
	"errors"
	"fmt"
)

// panic / recover 不是 Go 的常规错误处理手段。
// 原则：可预期的错误(文件不存在、参数非法、网络超时)一律用 error 返回；
//      panic 只留给"程序进入了不该发生的状态"这类不可恢复的编程错误。
//
// recover 只有在 defer 函数里调用才生效，典型用途是在"边界"处兜底，
// 防止单个请求的 panic 把整个进程带崩。
func panicRecover() {
	// 把可能 panic 的逻辑包在一个带 recover 的函数里，将 panic 转成 error 往外抛
	if _, err := safeDivide(10, 0); err != nil {
		fmt.Println("safeDivide 兜底成功，把 panic 转成了 error:", err)
	}
	if v, err := safeDivide(10, 2); err == nil {
		fmt.Println("safeDivide(10, 2) 正常执行，无 panic，结果 =", v)
	}
	fmt.Println("即使发生过 panic，主流程依然在继续运行")
}

// safeDivide 演示"在边界处用 recover 把 panic 收敛为 error"的惯用法。
// 命名返回值 err 是关键：defer 里的闭包可以修改它，从而把恢复到的 panic 写回返回值。
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			// r 是 recover 到的任意值(interface{})，这里统一转成 error
			err = fmt.Errorf("从 panic 中恢复: %v", r)
		}
	}()

	if b == 0 {
		// 这里故意 panic，模拟"底层某处炸了"
		panic(errors.New("除数为 0"))
	}
	result = a / b
	return result, nil
}
