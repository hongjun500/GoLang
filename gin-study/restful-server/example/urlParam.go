package example

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

type Person2 struct {
	// ID   string `uri:"id" binding:"required"`
	// 指定以uuid格式的数据，否则报错
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func StartPage(context *gin.Context) {
	var person Person
	if context.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	context.JSON(http.StatusOK, gin.H{
		"success": "ok",
	})
}

// 用占位符取值
func UrIParam(context *gin.Context) {
	var person Person2
	if err := context.ShouldBindUri(&person); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"errMsg": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"uuid": person.ID,
			"name": person.Name,
		})
	}
}
