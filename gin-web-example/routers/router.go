package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/Go-master/gin-web-example/pkg/setting"
	v1 "github.com/hongjun500/Go-master/gin-web-example/routers/api/v1"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	// 强制日志颜色化
	gin.ForceConsoleColor()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	// 路由组
	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{"success": "ok", "msg": "hello go"})
	})

	return router
}
