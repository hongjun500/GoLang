package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @author hongjun500
 * @date 2022/5/4 16:18
 * @tool ThinkPadX1隐士
 * Created with GoLand 2021.2
 * Description:通用返回信息
 */

type ReturnType struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

func Create(status string, data any, context *gin.Context) {
	context.JSON(http.StatusOK, ReturnType{
		Status: status,
		Data:   data,
	})
}

func CreateSuccess(data any, context *gin.Context) {
	Create("success", data, context)
}
func CreateFail(errMsg any, context *gin.Context) {
	Create("fail", errMsg, context)
}
