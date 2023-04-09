package main

import (
	"hongjun.com/miaosha/routers"
	"hongjun.com/miaosha/user/api"
)

func main() {

	// 使用默认中间件
	router := routers.InitRouter()

	router.POST("/userRegister", api.UserRegister)

	// 启动并监听在8080端口
	router.Run(":8080")
}
