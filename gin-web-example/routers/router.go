package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/Go-master/gin-web-example/pkg/setting"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	// 强制日志颜色化
	gin.ForceConsoleColor()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{"success": "ok", "msg": "hello go"})
	})
	return router
}
