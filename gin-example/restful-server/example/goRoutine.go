package example

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

/**
 * @author hongjun500
 * @date 2022/5/3 15:14
 * @tool ThinkPadX1隐士
 * Created with GoLand 2021.2
 * Description: goroutine的使用
 */

func LongAsync(context *gin.Context) {
	// 创建在goroutine中使用的副本
	contextCp := context.Copy()

	go func() {
		// 睡眠五秒模拟一个长任务
		time.Sleep(5 * time.Second)
		// 使用了goroutine,这里是副本
		// 打印请求路径
		log.Println("Done! in path " + contextCp.Request.URL.Path)
	}()
}

func Sync(context *gin.Context) {
	// 睡眠五秒模拟一个长任务
	time.Sleep(5 * time.Second)
	// 因为没有使用 goroutine，不需要拷贝上下文
	// 打印请求路径
	log.Println("Done! in path " + context.Request.URL.Path)

}
