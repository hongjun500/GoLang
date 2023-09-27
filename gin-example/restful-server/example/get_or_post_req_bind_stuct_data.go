package example

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/**
 * @author hongjun500
 * @date 2022/5/4 15:41
 * @tool ThinkPadX1隐士
 * Created with GoLand 2021.2
 * Description: get或者post请求绑定结构体数据
 */

type Persons struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func GetOrPostParamDataBindStruct(context *gin.Context) {
	var person Persons
	if context.ShouldBind(&person) == nil {
		context.JSON(http.StatusOK, gin.H{
			"status:": "ok",
			"person":  &person,
		})
	} else {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
		})
	}
}
