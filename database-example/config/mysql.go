package config

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"time"
)

type DbConn struct {
	Db *sql.DB
	Tx *sql.Tx
}

func GetMySQLConn() (DbConn, error) {
	// 设置时区为东八区
	loc, _ := time.LoadLocation("Asia/Shanghai")
	// 数据库配置
	config := mysql.Config{
		User:   "root",
		Passwd: "hongjun500",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "recordings",
		Loc:    loc,
		// 允许本地密码认证
		AllowNativePasswords: true,
	}
	var err error
	var dbconn DbConn

	dbconn.Db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return dbconn, fmt.Errorf("MySQL Connected Fail, ERR = %v", err)
	}
	err = dbconn.Db.Ping()
	if err != nil {
		return dbconn, fmt.Errorf("MySQL PingErr Fail, ERR = %v", err)
	}
	fmt.Println("MySQL Connected!")
	return dbconn, nil
}
