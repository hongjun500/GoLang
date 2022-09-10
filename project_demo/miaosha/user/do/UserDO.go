package do

import (
	"crypto/md5"
	"encoding/hex"
	"gorm.io/gorm"
	"hongjun.com/miaosha/conf"
	"hongjun.com/miaosha/user/model"
)

/**
 * @author hongjun500
 * @date 2022/5/29 18:33
 * @tool ThinkPadX1隐士
 * Created with GoLand 2021.2
 * Description: 用户实体
 */

type UserDO struct {
	gorm.Model
	// 主键
	// Id int
	// 用户名
	Name string `json:"name"`
	// 1男性，2女性
	Gender int8 `json:"gender"`
	// 年龄
	Age int `json:"age"`
	// 手机号
	TelPhone string `json:"tel_phone"`
	// 第三方注册信息 //byphone,bywechat,byalipay
	RegisterMode string `json:"register_mode"`
	// 第三方
	ThirdPartyId string `json:"third_party_id"`
}

type UserPassword struct {
	gorm.Model
	// 主键
	// id int `json:"id"`
	// 密文密码
	EncryptPassword string `json:"encrypt_password"`
	// 用户Id
	UserId uint `json:"user_id"`
}

func (user *UserDO) UserRegister(model *model.UserModel) {
	user.Name = model.Name
	user.RegisterMode = model.RegisterMode
	user.Age = model.Age
	user.TelPhone = model.TelPhone
	user.Gender = model.Gender

	db := conf.GetDb()
	result := db.Create(&user)
	if result.Error != nil {

		encrpytMd5 := md5.New()
		encrpytMd5.Write([]byte(model.EncryptPassword))
		// 密码的加密密文
		encrptyPwd := hex.EncodeToString(encrpytMd5.Sum(nil))

		model.Id = user.ID
		model.EncryptPassword = encrptyPwd
		userPassword := new(UserPassword)
		userPassword.UserRegister(model)
	}
}

func (pwd *UserPassword) UserRegister(model *model.UserModel) {
	db := conf.GetDb()
	pwd.UserId = model.Id
	pwd.EncryptPassword = model.EncryptPassword
	db.Create(pwd)
}
