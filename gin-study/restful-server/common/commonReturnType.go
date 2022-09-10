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
	Status string
	Data   interface{}
}

func Create(resultData interface{}, context *gin.Context) {
	context.JSON(http.StatusOK, ReturnType{
		Status: "success",
		Data:   resultData,
	})
}
