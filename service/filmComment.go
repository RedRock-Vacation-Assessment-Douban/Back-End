package service

import (
	"douban/dao"
	"douban/model"
)

// AddFilmComment 添加影评
func AddFilmComment(filmComment model.FilmComment) error {
	err := dao.InsertFilmComment(filmComment)
	return err
}

// DeleteFilmComment 删除影评
func DeleteFilmComment(FilmCommentId int) error {
	err := dao.DeleteFilmComment(FilmCommentId)
	return err
}

// GetFilmComments 得到影评
func GetFilmComments(FilmCommentId int) ([]model.FilmComment, error) {
	return dao.SelectFilmComment(FilmCommentId)
}

// GetFCById 通过FCid得到影评
func GetFCById(Id int) (model.FilmComment, error) {
	return dao.SelectFilmCommentById(Id)
}

// GetNameByFCId 通过id拿到用户名
func GetNameByFCId(FilmCommentId int) (string, error) {
	return dao.SelectNameByFCId(FilmCommentId)
}

// FilmCommentLikes 影评点赞
func FilmCommentLikes(FCId int) error {
	err := dao.FilmCommentLikes(FCId)
	return err
}

// FilmCommentDown 影评点踩
func FilmCommentDown(FCId int) error {
	err := dao.FilmCommentDown(FCId)
	return err
}

// GetMostPopular 得到最流行影评
func GetMostPopular() ([]model.MostPopularFC, error) {
	return dao.SelectMPFC()
}

// GetFilmCommentsByUsername 得到影评
func GetFilmCommentsByUsername(name string) ([]model.Personal, error) {
	return dao.SelectFilmCommentByUsername(name)
}

// GetMNById 通过FCid得到影评
func GetMNById(Id int) (string, error) {
	return dao.SelectMNById(Id)
}

// GetURLByMId 通过FCid得到影评
func GetURLByMId(Id int) (string, error) {
	return dao.SelectURLByMId(Id)
}
