// @author hongjun500
// @date 2024-04-08-20:53
// @tool ThinkPadX1 隐士
// Created with GoLand 2022.2
// Description: gin_engine.go

package web

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/hongjun500/GoLang-master/common-web/docs"
	"github.com/hongjun500/GoLang-master/common-web/internal/conf"
	"github.com/hongjun500/GoLang-master/common-web/internal/response"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func Bootstrap() {
	engine := gin.Default()

	initSwagger(engine)
	initHealthCheckRouter(engine)
	gin.SetMode(conf.GlobalGinConfigProperties.GinMode)
	srv := &http.Server{
		Addr:        fmt.Sprintf(":%s", conf.GlobalApplicationConfigProperties.Port),
		ReadTimeout: time.Duration(conf.GlobalApplicationConfigProperties.ReadTimeout) * time.Second,
		Handler:     engine,
	}
	g.Go(func() error {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

func initSwagger(engine *gin.Engine) {
	// 设置 Swagger 路由
	engine.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler), /*,
		ginSwagger.InstanceName(conf.GlobalApplicationConfigProperties.ApplicationName*/
		/*ginSwagger.URL("http://localhost:"+
		conf.GlobalApplicationConfigProperties.Port+
		"/swagger/"+"/docs.json"),*/
		/*ginSwagger.PersistAuthorization(true)*/)
	// docs.Title = conf.GlobalAdminServerConfigProperties.ApplicationName
}

// initHealthCheckRouter 初始化健康检查路由
// @Summary 初始化健康检查路由
// @Description 初始化健康检查路由
// @Tags 健康检查
// @Accept json
// @Produce json
// @Success 200 {object} response.CommonReturnType
// @Router /health [get]
func initHealthCheckRouter(engine *gin.Engine) {
	data := map[string]interface{}{}
	data["msg"] = "hello, health check success!"
	data["api_info"] = "common-web"
	engine.GET("health", func(context *gin.Context) {
		response.CreateSuccess(context, data)
	})
}
