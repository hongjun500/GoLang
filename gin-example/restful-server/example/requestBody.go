package example

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
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
		context.JSON(http.StatusOK, gin.H{
			"success": "ok",
		})
		// 因为现在 c.Request.Body 是 EOF，所以这里会报错。
	} else if errB := context.ShouldBind(&objB); errB == nil {
		context.JSON(http.StatusOK, gin.H{
			"success": "ok",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"success": "ok",
			"info":    "other",
		})
	}
}

func RequestBodys(context *gin.Context) {
	// 读取 c.Request.Body 并将结果存入上下文。
	if errA := context.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		context.JSON(http.StatusOK, gin.H{
			"success": "ok",
		})
		// 这时, 复用存储在上下文中的 body。
	} else if errB := context.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		context.JSON(http.StatusOK, gin.H{
			"success": "ok",
		})
		// 可以接受其他格式
	} else if errB2 := context.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
		context.JSON(http.StatusOK, gin.H{
			"success": "ok",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"success": "ok",
			"info":    "other",
		})
	}
}
