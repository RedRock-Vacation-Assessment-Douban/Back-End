package service

import (
	"douban/dao"
	"douban/model"
)

func GetVideo1() ([]model.Video1, error) {
	return dao.SelectVideo1()
}

func GetVideo2(id int) ([]model.Video2, error) {
	return dao.SelectVideo2(id)
}
