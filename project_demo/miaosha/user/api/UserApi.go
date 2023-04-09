package api

import (
	"github.com/gin-gonic/gin"
	"hongjun.com/miaosha/common"
	"hongjun.com/miaosha/user/do"
	"hongjun.com/miaosha/user/model"
)

/**
 * @author hongjun500
 * @date 2022/6/4 13:34
 * @tool ThinkPadX1隐士
 * Created with GoLand 2021.2
 * Description:用户接口
 */

func UserRegister(context *gin.Context) {
	userModel := new(model.UserModel)
	err := context.ShouldBind(userModel)
	if err != nil {
		return
	}
	var user do.UserDO
	user.UserRegister(userModel)
	common.CreateSuccess(true, context)
}
