// @author hongjun500
// @date 2023/9/27 15:46
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package example

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hongjun500/GoLang-master/gin-example/restful-server/common"
)

// DST 指定一个上传文件的路径地址
const DST = "D:\\goProjects\\upload_gofile\\"

func Upload(context *gin.Context) {
	// 单个文件
	file, err := context.FormFile("file")
	if err != nil {
		common.CreateFail(context, "获取文件失败!")
	} else {
		log.Println(file.Filename)
		// 上传文件到指定路径
		context.SaveUploadedFile(file, DST+file.Filename)
		common.Create(context, "uploaded", file.Filename)
	}
}

func Uploads(context *gin.Context) {
	form, err := context.MultipartForm()
	if err != nil {
		common.CreateFail(context, "获取文件失败!")
	} else {
		files := form.File["upload[]"]
		for _, file := range files {
			log.Println(file.Filename)

			// 上传文件到指定路径并以时间戳重命名
			context.SaveUploadedFile(file, DST+fmt.Sprintf("%d", time.Now().Unix())+file.Filename)
			// context.SaveUploadedFile(file, DST+file.Filename)
		}
		common.Create(context, fmt.Sprintf("uploaded %d files", len(files)))
	}

}
