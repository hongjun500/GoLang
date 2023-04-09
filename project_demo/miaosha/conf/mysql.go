package conf

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

/**
 * @author hongjun500
 * @date 2022/5/29 19:58
 * @tool ThinkPadX1隐士
 * Created with GoLand 2021.2
 * Description:
 */

func GetDb() *gorm.DB {
	dsn := "root:hongjun500@tcp(127.0.0.1:3306)/miaosha_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Println("数据库连接失败", err)

	}
	return db
}
