package main

import (
	"log"
)

func main() {

	log.Println("error 的本质：一个只有 Error() string 的接口----------")
	basicError()
	log.Println("哨兵错误(sentinel error)：预定义的可比较错误值----------")
	sentinelError()
	log.Println("错误包装 %w + errors.Is / As / Unwrap----------")
	wrapError()
	log.Println("自定义错误类型：携带上下文，用 errors.As 提取----------")
	customError()
	log.Println("nil 接口陷阱：返回具体类型的 nil 指针 != nil error----------")
	nilInterfaceTrap()
	log.Println("panic / recover：边界兜底，而非常规控制流----------")
	panicRecover()
}
