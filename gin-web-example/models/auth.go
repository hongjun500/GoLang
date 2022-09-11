package models

/**
 * @author hongjun500
 * @date 2022/9/11 15:29
 * @tool ThinkPadX1隐士
 * Created with GoLand 2021.2
 * Description:鉴权模型及接口
 */

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckAuth 查询数据库是否存在
func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{
		Username: username,
		Password: password,
	}).Find(&auth)
	return auth.ID > 0
}
