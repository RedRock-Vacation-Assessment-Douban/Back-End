package service

import (
	"douban/dao"
	"douban/model"
)

// AddFilmCommentReply 添加评论
func AddFilmCommentReply(FilmCommentReply model.FilmCommentReply) error {
	return dao.InsertFilmCommentReply(FilmCommentReply)
}

// GetFilmCommentReply 得到影评评论
func GetFilmCommentReply(FCId int) ([]model.FilmCommentReply, error) {
	return dao.SelectCommentByFCId(FCId)
}

// DeleteFilmCommentReply 删除评论
func DeleteFilmCommentReply(Id int) error {
	err := dao.DeleteFilmCommentReply(Id)
	return err
}

// GetNameByFCRId 通过id得到用户名
func GetNameByFCRId(FCRId int) (string, error) {
	return dao.SelectNameByFCRId(FCRId)
}
