package main

import (
	"errors"
	"fmt"
)

// ValidationError 携带上下文，用 errors.As 提取
// 当错误需要携带结构化上下文（字段、错误码等）时，就自定义一个错误类型。
// 只要实现 Error() string 就是一个合法的 error。
type ValidationError struct {
	Field string
	Value any
	Msg   string
}

// Error 让 *ValidationError 满足 error 接口。
// 惯例用指针接收者：错误类型通常较"重"，且方便和 errors.As 的目标 **T 配合。
func (e *ValidationError) Error() string {
	return fmt.Sprintf("字段 %q(值=%v) 校验失败: %s", e.Field, e.Value, e.Msg)
}

// 如果想让自定义错误也能挂在错误链上，可以额外实现 Unwrap() error 返回内部错误。

func validateAge(age int) error {
	if age < 0 || age > 150 {
		return &ValidationError{Field: "age", Value: age, Msg: "必须在 0~150 之间"}
	}
	return nil
}

func customError() {
	// 上层再包一层，模拟真实调用链
	err := fmt.Errorf("创建用户失败: %w", validateAge(200))
	fmt.Println("错误信息:", err)

	// errors.As：沿着错误链找"某个具体类型"的错误，找到就把它取出来赋值给 target。
	// 这是从错误里"拿回结构化信息"的标准姿势——比字符串解析可靠得多。
	var ve *ValidationError
	if errors.As(err, &ve) {
		// 现在可以访问 ve 的具体字段，做精细化处理
		fmt.Printf("errors.As 提取到 ValidationError：Field=%s Value=%v\n", ve.Field, ve.Value)
	}

	// 小结：
	//   errors.Is  —— 判断"是不是某个错误值"（哨兵错误）
	//   errors.As  —— 判断"是不是某个错误类型"，并取出该类型实例（自定义错误）
}
