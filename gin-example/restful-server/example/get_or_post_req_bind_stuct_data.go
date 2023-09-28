package example

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hongjun500/GoLang-master/gin-example/restful-server/common"
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
		person.Name = "getMethod"
		person.Address = "router.GET"
		person.Birthday = time.Now()
		common.Create(context, person)
		return
	} else {
		common.CreateFail(context)
	}
}
