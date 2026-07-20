package main

import (
	"errors"
	"fmt"
)

// error 就是标准库里这样一个再普通不过的接口：
//
//	type error interface {
//	    Error() string
//	}
//
// 任何实现了 Error() string 的类型都是一个 error。这也是"错误只是值(errors are values)"
// 这句 Go 名言的由来：错误不是语言层面的异常机制，而是能被当普通值传递、比较、组合的东西。
func basicError() {
	// errors.New 返回的是 *errorString，携带一段只读的描述文本
	err := errors.New("打开文件失败")
	fmt.Println("errors.New:", err)

	// fmt.Errorf 可以像 Printf 那样格式化，拼出带上下文的错误信息
	name := "config.yaml"
	err2 := fmt.Errorf("读取配置 %s 时出错", name)
	fmt.Println("fmt.Errorf:", err2)

	// Go 的惯用返回签名是 (T, error)：error 放最后一个返回值，nil 表示成功
	if v, err := parsePositive(10); err == nil {
		fmt.Println("parsePositive(10) =", v)
	}
	// 失败路径：先判 err，再用值。这是 Go 最基础也最重要的编码节奏
	if _, err := parsePositive(-3); err != nil {
		fmt.Println("parsePositive(-3) 出错:", err)
	}
}

// parsePositive 演示最常见的错误返回范式：成功返回 (值, nil)，失败返回 (零值, err)。
func parsePositive(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("期望正数，得到 %d", n)
	}
	return n, nil
}
