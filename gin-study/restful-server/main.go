package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginForm 表单结构体
type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {

	router := gin.Default()
	// Content-Type = application/x-www-form-urlencoded 形式
	/*router.POST("/login", func(context *gin.Context) {
		var loginForm LoginForm
		// 显示绑定
		if context.ShouldBindWith(&loginForm, binding.Form) == nil {
		// 自动绑定会调用显示绑定
		// if context.ShouldBind(&loginForm) == nil {
			if loginForm.User == "user" && loginForm.Password == "password"{
				context.JSON(http.StatusOK, gin.H{"success":"ok"})
			}else {
				context.JSON(http.StatusUnauthorized, gin.H{"success":"fail"})
			}
		}
	})*/

	// Content-Type = multipart/form-data
	/*router.POST("/login-form", func(context *gin.Context) {
		msg := context.PostForm("message")
		// 绑定key,并且如果没有对应的参数则给出一个默认值anonymous
		nick := context.DefaultPostForm("nick", "anonymous")
		context.JSON(http.StatusOK, gin.H{
			"success": "ok",
			"msg":     msg,
			"nick" : nick,
		})
	})*/

	// JSON和PureJSON
	// 使用 unicode 替换特殊 HTML 字符,
	/*
		{"html":"\u003cb\u003eHello, world!\u003c/b\u003e"}
	*/
	router.GET("/json", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 返回的数据不会替换 <
	/*
		{"html":"<b>Hello, world!</b>"}
	*/
	router.GET("/purejson", func(context *gin.Context) {
		context.PureJSON(http.StatusOK, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 当 Content-Type: application/x-www-form-urlencoded 时另外一种获取请求参数的方式
	// 请求的url中带有id和page, message和name在body中
	router.POST("/post", func(context *gin.Context) {
		id := context.Query("id")
		page := context.DefaultQuery("page", "0")
		name := context.PostForm("name")
		message := context.PostForm("message")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		context.JSON(http.StatusOK, gin.H{
			"id":      id,
			"page":    page,
			"message": message,
			"name":    name,
		})
	})

	router.Run(":8086")

}
