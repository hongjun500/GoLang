package service

import (
	"hongjun.com/miaosha/user/model"
)

/**
 * @author hongjun500
 * @date 2022/5/29 18:43
 * @tool ThinkPadX1隐士
 * Created with GoLand 2021.2
 * Description:
 */

type UserServiceStruct struct {
}

type UserService interface {
	// UserRegister 用户注册
	UserRegister(model *model.UserModel)
}
