package example

import (
	"github.com/gin-gonic/gin"
	"restful-server/common"
)

/**
 * @author hongjun500
 * @date 2022/5/4 16:15
 * @tool ThinkPadX1隐士
 * Created with GoLand 2021.2
 * Description: 直接通过上下文获取路由参数
 */

func GetRouteParam(context *gin.Context) {
	name := context.Param("name")
	mapData := make(map[string]interface{})
	mapData["name"] = "hongjun500"
	mapData["age"] = 22
	mapData["property"] = []string{
		"11", "222",
	}
	if name != "" {
		/*context.JSON(http.StatusOK, gin.H{
			"name":name,
		})*/
		common.Create(mapData, context)
	}

}
