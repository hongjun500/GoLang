package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	restful_server "github.com/hongjun500/GoLang-master/gin-example/restful-server"
)

func main() {

	router := gin.Default() // 默认使用了 Logger 和 Recovery 中间件
	router.GET("/", func(context *gin.Context) {
		time.Sleep(5 * time.Second)
		context.String(http.StatusOK, "Welcome gin Server")
	})
	authGroup(router)
	getThirdPartyApi(router)
	restful_server.ExampleFunc(router)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {
		// 服务连接
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)      // 创建一个接收信号的通道
	signal.Notify(quit, os.Interrupt) // 转发接收到的信号
	<-quit                            // 阻塞在此，直到接收到信号
	log.Println("Shutdown Server ...")
	// 创建一个 5 秒超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

// authGroup 认证组
func authGroup(router *gin.Engine) {
	// 定义一个和账号对应的 map
	var secrets = gin.H{
		"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
		"austin": gin.H{"email": "austin@example.com", "phone": "666"},
		"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
	}

	authorization := router.Group("/user", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
	authorization.GET("/secrets", func(context *gin.Context) {
		user := context.MustGet(gin.AuthUserKey).(string)
		// ok代表对应键是否存在，secret是键对应的值
		if secret, ok := secrets[user]; ok {
			context.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			context.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})
}

// getThirdPartyApi 获取第三方 api
func getThirdPartyApi(router *gin.Engine) {
	router.GET("/someDataFromReader", func(context *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			context.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		context.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
}
