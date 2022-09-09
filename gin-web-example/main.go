package main

import (
	"fmt"
	"github.com/hongjun500/Go-master/gin-web-example/pkg/setting"
	"github.com/hongjun500/Go-master/gin-web-example/routers"
	"net/http"
)

func main() {

	// 使用默认的中间件
	//gin.SetMode(gin.ReleaseMode)
	router := routers.InitRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
