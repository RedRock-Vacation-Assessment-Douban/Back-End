package service

import (
	"douban/dao"
	"douban/model"
)

// AddShortComment 添加短评
func AddShortComment(shortComment model.ShortComment) error {
	err := dao.InsertShortComment(shortComment)
	return err
}

// DeleteShortComment 删除短评
func DeleteShortComment(shortCommentId int) error {
	err := dao.DeleteShortComment(shortCommentId)
	return err
}

// GetShortComment 得到短评
func GetShortComment(movieId int) ([]model.ShortComment, error) {
	return dao.SelectShortComment(movieId)
}

// GetNameBySCId 通过id拿到用户名
func GetNameBySCId(SCId int) (string, error) {
	return dao.SelectNameBySCId(SCId)
}

// ShortCommentLikes 话题点赞
func ShortCommentLikes(SCId int) error {
	err := dao.ShortCommentLikes(SCId)
	return err
}

// GetShortCommentByUsername 得到短评
func GetShortCommentByUsername(name string) ([]model.SCPersonal, error) {
	return dao.SelectShortCommentByUsername(name)
}
