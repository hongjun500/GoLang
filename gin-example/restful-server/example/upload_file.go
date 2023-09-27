// @author hongjun500
// @date 2023/9/27 15:46
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package example

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DST 指定一个上传文件的路径地址
const DST = "D:\\goProjects\\upload_gofile\\"

func Upload(context *gin.Context) {
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
}

func Uploads(context *gin.Context) {
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

}
