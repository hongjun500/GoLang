package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/Go-master/gin-web-example/models"
	"github.com/hongjun500/Go-master/gin-web-example/pkg/e"
	"github.com/hongjun500/Go-master/gin-web-example/pkg/logging"
	"github.com/hongjun500/Go-master/gin-web-example/pkg/util"
	"net/http"
)

/**
 * @author hongjun500
 * @date 2022/9/11 15:32
 * @tool ThinkPadX1隐士
 * Created with GoLand 2021.2
 * Description: 鉴权路由
 */

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(context *gin.Context) {
	username := context.Query("username")
	password := context.Query("password")
	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	if ok {
		hasExist := models.CheckAuth(username, password)
		if hasExist {
			// 生成token
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			//log.Println(err.Key, err.Message)
			logging.Info(err.Key, err.Message)
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsgByCode(code),
		"data": data,
	})
}
