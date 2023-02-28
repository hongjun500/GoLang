package main

import (
	"github.com/gin-gonic/gin"
	"hongjun.com/miaosha/user/rest"
)

func main() {

	// 使用默认中间件
	router := gin.Default()

	router.POST("/userRegister", rest.UserRegister)

	// 启动并监听在8080端口
	router.Run(":8080")
}
