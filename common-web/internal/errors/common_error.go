// @author hongjun500
// @date 2024-04-05-22:18
// @tool ThinkPadX1 隐士
// Created with GoLand 2022.2
// Description: common_error.go

package errors

// 通用错误码，自定义业务错误码
const (
	UnknownError             = 100000 // 未知错误
	CustomError              = 100001 // 自定义错误
	ParameterValidationError = 100002 // 参数校验错误
	DatabaseError            = 100003 // 数据库错误
)

var ErrorMsgConstant = map[int]string{
	UnknownError:             "未知错误",
	CustomError:              "自定义错误",
	ParameterValidationError: "参数校验错误",
	DatabaseError:            "数据库错误",
}

// CommonError 通用错误结构体
type CommonError struct {
	ErrCode int    `json:"err_code"` // 错误码
	ErrMsg  string `json:"err_msg"`  // 错误信息, 经过业务处理后的错误信息
	Err     error  `json:"error"`    // 原始错误信息
}

func (e CommonError) Error() string {
	if e.ErrCode == 0 { // 如果错误为初始化值，则直接返回错误信息
		return e.ErrMsg
	}
	return ErrorMsgConstant[e.ErrCode]
}
