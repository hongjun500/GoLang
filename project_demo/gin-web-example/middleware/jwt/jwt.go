package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/Go-master/gin-web-example/pkg/e"
	"github.com/hongjun500/Go-master/gin-web-example/pkg/util"
	"net/http"
	"time"
)

/**
 * @author hongjun500
 * @date 2022/9/11 15:23
 * @tool ThinkPadX1隐士
 * Created with GoLand 2021.2
 * Description:
 */

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := context.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				// token解析失败
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				// token过期
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsgByCode(code),
				"data": data,
			})

			context.Abort()
			return
		}
		context.Next()
	}
}
