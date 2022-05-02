package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// 携带基础中间件Logger,Recovery启动
	/*func Default() *Engine {
		debugPrintWARNINGDefault()
		engine := New()
		engine.Use(Logger(), Recovery())
		return engine
	}*/

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "hongjun500",
		})
	})
	// 默认监听并在 0.0.0.0:8080 上启动服务
	// 可指定端口启动
	r.Run(":8089")
}
