// Author hongjun500
// @date 2024-04-05
// @tool ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: common_return_type.go

package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommonReturnType struct {
	Status string `json:"status"` // success or fail
	Data   any    `json:"data"`   // if status is success, return data, else return error message
	/**
	{
		"err_code": 300000,
		"err_msg": "用户名已存在"
	}
	*/
}

// createAny 创建通用返回类型
func createAny(context *gin.Context, status string, data any) {
	context.JSON(http.StatusOK, CommonReturnType{
		Status: status,
		Data:   data,
	})
}

// Created 创建成功返回类型，不带数据
func Created(context *gin.Context) {
	createAny(context, "success", nil)
}

// CreateSuccess 创建失败返回类型
func CreateSuccess(context *gin.Context, data any) {
	createAny(context, "success", data)
}

// CreateFail 创建失败返回类型
func CreateFail(context *gin.Context, data any) {
	createAny(context, "fail", data)
}
