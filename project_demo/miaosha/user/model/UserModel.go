package model

/**
 * @author hongjun500
 * @date 2022/5/29 18:44
 * @tool ThinkPadX1隐士
 * Created with GoLand 2021.2
 * Description: 用户模型
 */

type UserModel struct {
	// 主键
	Id uint `json:"id"`
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

	// 密文密码
	EncryptPassword string `json:"encrypt_password"`
}
