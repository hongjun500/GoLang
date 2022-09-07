package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// DST 指定一个上传文件的路径地址
const DST = "D:\\goProjects\\upload_gofile\\"

func main() {
	router := gin.Default()
	// 可以给文件限制大小
	router.MaxMultipartMemory = 8 << 20 // 8M
	router.POST("/upload", func(context *gin.Context) {
		// 单个文件
		file, err := context.FormFile("file")
		if err != nil {
			log.Fatal("获取文件失败!")
		} else {
			log.Println(file.Filename)
			// 上传文件到指定路径
			context.SaveUploadedFile(file, DST+file.Filename)
			context.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
		}
	})

	router.POST("/uploads", func(context *gin.Context) {
		form, err := context.MultipartForm()
		if err != nil {
			log.Fatal("获取文件失败!")
		} else {
			files := form.File["upload[]"]
			for _, file := range files {
				log.Println(file.Filename)

				// 上传文件到指定路径
				context.SaveUploadedFile(file, DST+file.Filename)
			}
			context.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
		}

	})

	router.Run(":8086")
}
