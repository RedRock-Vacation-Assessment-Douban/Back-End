package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var dB *sql.DB

func InitDB() (err error) {
	dsn := "lance:yxh030714@tcp(42.192.155.29:3306)/douban?charset=utf8&parseTime=True"
	// 连接数据库
	dB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = dB.Ping()
	if err != nil {
		return err
	}
	return nil
}
