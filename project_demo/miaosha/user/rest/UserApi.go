package rest

import (
	"github.com/gin-gonic/gin"
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
	// var userService := service.UserService{}
	context.ShouldBind(&model.UserModel{})
	// userModel := do.UserDO{}
	// userModel.UserRegister()
}
