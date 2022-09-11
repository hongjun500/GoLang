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

	// 文章标签路由组
	tagsRouterGroup := router.Group("/api/v1")
	{
		tagsRouterGroup.GET("/tags", v1.GetTags)
		tagsRouterGroup.POST("/tags", v1.AddTag)
		tagsRouterGroup.PUT("/tags/:id", v1.EditTag)
		tagsRouterGroup.DELETE("/tags/:id", v1.DeleteTag)
	}

	// 文章路由组
	articleRouterGroup := router.Group("/api/v1")
	{
		articleRouterGroup.GET("/articles/:id", v1.GetArticle)
		articleRouterGroup.GET("/articles", v1.ListArticles)
		articleRouterGroup.POST("/articles/:id", v1.AddArticle)
		articleRouterGroup.PUT("/articles/:id", v1.EditArticle)
		articleRouterGroup.DELETE("/articles/:id", v1.DeleteArticle)
	}

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{"success": "ok", "msg": "hello go"})
	})

	return router
}
