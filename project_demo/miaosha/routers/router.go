package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"hongjun.com/miaosha/user/api"
)

/**
 * @author hongjun500
 * @date 2023/4/9
 * @tool ThinkPadX1隐士
 * Created with GoLand 2022.2
 * Description: routers.routers.go
 */

func InitRouter() *gin.Engine {
	// 强制日志颜色化
	gin.ForceConsoleColor()
	engine := gin.Default()
	// 跨域请求设置
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	// 请求头
	corsConfig.AddAllowHeaders("Authorization")
	v1 := engine.Group("/v1")
	v1.POST("/userRegister", api.UserRegister)

	return gin.Default()
}
