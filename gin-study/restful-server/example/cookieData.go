package example

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
 * @author hongjun500
 * @date 2022/5/4 16:08
 * @tool ThinkPadX1隐士
 * Created with GoLand 2021.2
 * Description: cookie相关
 */

func CookieData(context *gin.Context) {
	cookie, err := context.Cookie("gin_cookie")
	if err != nil {
		cookie = "NotSet"
		context.SetCookie("gin_cookie", "test_gin_cookie", 3600, "/", "localhost", false, true)
	}
	fmt.Printf("Cookie value: %s \n", cookie)
}
