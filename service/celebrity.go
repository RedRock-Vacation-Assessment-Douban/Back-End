package service

import (
	"douban/dao"
	"douban/model"
)

// GetCelebrityById 通过id得到影人
func GetCelebrityById(Id int, mid int) (model.Celebrity, error) {
	return dao.SelectCelebrityById(Id, mid)
}

func GetSearch2(context string) ([]model.Search2, error) {
	return dao.Search2(context)
}
