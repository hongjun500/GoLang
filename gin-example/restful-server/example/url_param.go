package example

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/GoLang-master/gin-example/restful-server/common"
)

type Person struct {
	ID   string `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}

// UrIParam 用占位符取值
func UrIParam(context *gin.Context) {
	var person Person
	if err := context.ShouldBindUri(&person); err != nil {
		common.CreateFail(context, err.Error())
		return
	} else {
		common.Create(context, person)
	}
}
