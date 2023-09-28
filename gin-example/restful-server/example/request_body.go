package example

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/hongjun500/GoLang-master/gin-example/restful-server/common"
)

/**
 * @author hongjun500
 * @date 2022/5/4 13:11
 * @tool ThinkPadX1隐士
 * Created with GoLand 2021.2
 * Description: 将request body中的数据绑定到不同的结构体中
	尽量使用context.ShouldBind()
*/

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

var (
	objA = formA{}
	objB = formB{}
)

func RequestBody(context *gin.Context) {

	if errA := context.ShouldBind(&objA); errA == nil {
		common.Create(context, objA)
		return
	} else if errB := context.ShouldBind(&objB); errB == nil {
		common.Create(context, objB)
		return
	} else {
		common.Create(context, "other")
	}
}

func RequestBodys(context *gin.Context) {
	// 读取 c.Request.Body 并将结果存入上下文。
	if errA := context.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		common.Create(context, objA)
		return
		// 这时, 复用存储在上下文中的 body。
	} else if errB := context.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		common.Create(context, objB)
		return
		// 可以接受其他格式
	} else if errB2 := context.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
		common.Create(context, objB)
		return
	} else {
		common.Create(context, "other")
	}
}
