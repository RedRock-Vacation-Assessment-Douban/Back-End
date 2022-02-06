package main

import (
	"douban/api"
	"douban/dao"
	"fmt"
)

func main() {
	err := dao.InitDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
	} else {
		fmt.Println("连接数据库成功!")
	}
	api.InitEngine()
}
