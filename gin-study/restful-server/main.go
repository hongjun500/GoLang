package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io"
	"log"
	"net/http"
	"os"
	"restful-server/example"
)

// LoginForm 表单结构体
type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {

	router := gin.Default()
	// 强制日志颜色化
	gin.ForceConsoleColor()
	// Content-Type = application/x-www-form-urlencoded 形式
	router.POST("/login", func(context *gin.Context) {
		var loginForm LoginForm
		// 显示绑定
		if context.ShouldBindWith(&loginForm, binding.Form) == nil {
			// 自动绑定会调用显示绑定
			// if context.ShouldBind(&loginForm) == nil {
			if loginForm.User == "user" && loginForm.Password == "password" {
				context.JSON(http.StatusOK, gin.H{"success": "ok"})
			} else {
				context.JSON(http.StatusUnauthorized, gin.H{"success": "fail"})
			}
		}
	})

	// Content-Type = multipart/form-data
	router.POST("/login-form", func(context *gin.Context) {
		msg := context.PostForm("message")
		// 绑定key,并且如果没有对应的参数则给出一个默认值anonymous
		nick := context.DefaultPostForm("nick", "anonymous")
		context.JSON(http.StatusOK, gin.H{
			"success": "ok",
			"msg":     msg,
			"nick":    nick,
		})
	})

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
	// 使用SecureJSON防止json劫持
	router.POST("/secureJson", func(context *gin.Context) {
		// 将输出：while(1);["one","two","three"]
		names := []string{"one", "two", "three"}
		context.SecureJSON(http.StatusOK, names)
	})

	// 只绑定 url 查询字符串 ShouldBindQuery 函数只绑定 url 查询参数而忽略 post 数据
	router.Any("/urlParam", example.StartPage)

	// goroutine的使用
	router.GET("/long_async", example.LongAsync)
	router.GET("/sync", example.Sync)

	// 日志
	// router.GET("/logInfo", example.LogInfoSave)
	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	// 记录到文件。
	f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// 同时将日志写入文件和控制台
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	/*router.POST("/foo", func(c *gin.Context) {
		c.JSON(http.StatusOK, "foo")
	})

	router.GET("/ ", func(c *gin.Context) {
		c.JSON(http.StatusOK, "bar")
	})

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})*/

	router.POST("/requestBody", example.RequestBody)
	router.POST("/requestBodys", example.RequestBodys)

	// log.Fatal(autotls.Run(router, "example1.com", "example2.com"))

	// 使用ShouldBindUri获取占位符数据绑定结构体
	router.GET("/:name/:id", example.UrIParam)

	// 如果是 `GET` 请求，只使用 `Form` 绑定引擎（`query`）。
	router.GET("/getOrPostBingParamStruct", example.GetOrPostParamDataBindStruct)
	// 如果是 `POST` 请求，首先检查 `content-type` 是否为 `JSON` 或 `XML`，然后再使用 `Form`（`form-data`）。
	router.POST("/getOrPostBingParamStruct", example.GetOrPostParamDataBindStruct)

	// cookie

	router.GET("/cookie", example.CookieData)

	router.GET("/user/:name", example.GetRouteParam)
	// 启动并监听在8086端口上
	router.Run(":8086")
}
