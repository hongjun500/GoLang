package example

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

/**
 * @author hongjun500
 * @date 2022/5/3 15:24
 * @tool ThinkPadX1隐士
 * Created with GoLand 2021.2
 * Description: 记录日志
 */

func LogInfoSave(context *gin.Context) {
	// 禁用控制台颜色
	gin.DisableConsoleColor()

	log.Println("------------------------------日志日志")
	log.Println("------------------------------日志日志")
	log.Println("------------------------------日志日志")
	log.Println("------------------------------日志日志")
	log.Println("------------------------------日志日志")
	// 记录到文件
	f, _ := os.Create("ginLog.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// 同时将日志写入到文件和控制台
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	context.JSON(http.StatusOK, gin.H{
		"success": "ok",
	})
}
