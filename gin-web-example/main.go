package main

import (
	"context"
	"fmt"
	"github.com/hongjun500/Go-master/gin-web-example/pkg/setting"
	"github.com/hongjun500/Go-master/gin-web-example/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
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

	go func() {
		// 服务连接
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
