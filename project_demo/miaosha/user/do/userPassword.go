package do

import (
	"gorm.io/gorm"
	"hongjun.com/miaosha/conf"
	"hongjun.com/miaosha/user/model"
)

/**
 * @author hongjun500
 * @date 2023/4/9
 * @tool ThinkPadX1隐士
 * Created with GoLand 2022.2
 * Description: userPassword.go
 */
type UserPassword struct {
	gorm.Model
	// 主键
	// id int `json:"id"`
	// 密文密码
	EncryptPassword string `json:"encrypt_password"`
	// 用户Id
	UserId uint `json:"user_id"`
}

func (pwd *UserPassword) UserRegister(model *model.UserModel) {
	db := conf.GetDb()
	pwd.UserId = model.Id
	pwd.EncryptPassword = model.EncryptPassword
	db.Create(pwd)
}

func (pwd *UserPassword) TableName() string {
	return "user_password"
}
