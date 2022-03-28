package dao

import (
	"database/sql"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

var dB *sql.DB
var rdb *redis.Client

func InitDB() (err error) {
	dsn := "lance:yxh030714@tcp(1.14.43.76:3306)/douban?charset=utf8&parseTime=True"
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

func InitRedis() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
