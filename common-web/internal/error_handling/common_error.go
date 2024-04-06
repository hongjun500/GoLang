// @author hongjun500
// @date 2024-04-05-22:18
// @tool ThinkPadX1 隐士
// Created with GoLand 2022.2
// Description: common_error.go

package error_handling

import "strings"

// 通用错误码，自定义业务错误码
const (
	UnknownError             = 100000 // 未知错误
	CustomError              = 100001 // 自定义错误
	ParameterValidationError = 100002 // 参数校验错误
	DatabaseError            = 100003 // 数据库错误
)

// BusinessErrorConstant 错误信息常量
var BusinessErrorConstant = map[int]string{
	UnknownError:             "未知错误",
	CustomError:              "自定义错误",
	ParameterValidationError: "参数校验错误",
	DatabaseError:            "数据库错误",
}

// CommonError 通用错误结构体
type CommonError struct {
	Err     error  `json:"err"`      // 原始错误信息
	ErrCode int    `json:"err_code"` // 错误码
	ErrMsg  string `json:"err_msg"`  // 错误信息, 经过业务处理后的错误信息
}

func (e *CommonError) Error() string {
	if e.ErrCode == 0 { // 如果错误为初始化值，则直接返回错误信息
		return e.ErrMsg
	}
	if msg, ok := BusinessErrorConstant[e.ErrCode]; ok && strings.EqualFold(msg, e.ErrMsg) {
		return msg
	}
	return e.ErrMsg

}

// NewCommonError 创建通用错误 example
//
// error_handling.NewCommonError(error_handling error_handling) len(args) == 1
//
// error_handling.NewCommonError(error_handling.UnknownError int) len(args) == 1
//
// error_handling.NewCommonError(error_handling error_handling, error_handling.UnknownError int) len(args) == 2
//
// error_handling.NewCommonError(error_handling.UnknownError int, errMsg string) len(args) == 2
//
// error_handling.NewCommonError(error_handling error_handling, error_handling.UnknownError int, errMsg string) len(args) == 3
func NewCommonError(args ...any) *CommonError {
	err := &CommonError{}

	switch len(args) {
	case 1:
		arg := args[0]
		switch v := arg.(type) {
		case error:
			err.Err = v
			err.ErrCode = UnknownError
			err.ErrMsg = BusinessErrorConstant[UnknownError]
		case int:
			err.ErrCode = v
			err.ErrMsg = BusinessErrorConstant[v]
		default:
			panic("common_error.go: if there is only one argument, it should be error_handling or int")
		}
	case 2:
		switch arg0 := args[0].(type) {
		case error:
			err.Err = arg0
			if arg1, ok := args[1].(int); ok {
				err.ErrCode = arg1
				err.ErrMsg = BusinessErrorConstant[arg1]
			} else {
				panic("common_error.go: if the first argument is error_handling, the second argument should be int")
			}
		case int:
			err.ErrCode = arg0
			if arg1, ok := args[1].(string); ok {
				err.ErrMsg = arg1
			} else {
				panic("common_error.go: if the first argument is int, the second argument should be string")
			}
		default:
			panic("common_error.go: if there are 2 arguments, the first argument should be error_handling or int")
		}
	case 3:
		if arg0, ok := args[0].(error); ok {
			err.Err = arg0
			if arg1, ok := args[1].(int); ok {
				err.ErrCode = arg1
				if arg2, ok := args[2].(string); ok {
					err.ErrMsg = arg2
				} else {
					panic("common_error.go: if there are 3 arguments, the third argument should be string")
				}
			} else {
				panic("common_error.go: if there are 3 arguments, the second argument should be int")
			}
		} else {
			panic("common_error.go: if there are 3 arguments, the first argument should be error_handling")
		}
	default:
		panic("common_error.go: unsupported number of arguments")
	}
	return err
}
