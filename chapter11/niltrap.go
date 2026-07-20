package main

import "fmt"

// nil 接口陷阱：Go 里最经典、最坑的错误处理翻车点。
//
// 接口值在底层是一个二元组 (类型, 值)。只有当"类型"和"值"都为 nil 时，
// 接口才 == nil。如果把一个"值为 nil 但类型确定"的具体指针赋给 error 接口，
// 接口的类型部分被填上了 (*MyError)，于是 err != nil 恒成立——即使里面的指针是 nil。

type MyError struct{ Code int }

func (e *MyError) Error() string { return fmt.Sprintf("错误码 %d", e.Code) }

// badReturn 是反例：用具体类型 *MyError 作为返回类型，成功时 return 了 nil 指针。
func badReturn(fail bool) *MyError {
	var e *MyError // e 是 (*MyError)(nil)
	if fail {
		e = &MyError{Code: 500}
	}
	return e // 成功时返回的是"类型为 *MyError、值为 nil"的指针
}

// goodReturn 是正确写法：返回类型就用 error 接口，成功时直接 return nil。
func goodReturn(fail bool) error {
	if fail {
		return &MyError{Code: 500}
	}
	return nil // 干净的 nil 接口：类型和值都为 nil
}

func nilInterfaceTrap() {
	// 反例：把 badReturn 的结果装进 error 接口
	var err error = badReturn(false) // 成功场景，本以为是 nil
	// 陷阱：err 的动态类型是 *MyError（非 nil），所以下面这句竟然为 true！
	fmt.Printf("badReturn 成功但 err != nil? %v (动态类型=%T)\n", err != nil, err)

	// 正确：返回类型直接用 error
	err2 := goodReturn(false)
	fmt.Printf("goodReturn 成功且 err == nil? %v\n", err2 == nil)

	// 结论：
	//   1. 函数返回错误时，返回类型永远写 error，不要写具体的 *XxxError
	//   2. 成功路径直接 `return nil`，不要 `return someNilTypedPointer`
}
