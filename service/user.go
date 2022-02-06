package service

import (
	"database/sql"
	"douban/dao"
	"douban/model"
)

// ChangePassword 修改密码服务
func ChangePassword(username, newPassword string) error {
	err := dao.UpdatePassword(username, newPassword)
	return err
}

// IsPasswordCorrect 判断密码是否正确服务
func IsPasswordCorrect(username, password string) (bool, error) {
	user, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	if user.Password != password {
		return false, nil
	}

	return true, nil
}

// IsRepeatUsername 判断用户名是否重复
func IsRepeatUsername(username string) (bool, error) {
	_, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

// Register 注册服务
func Register(user model.User) error {
	err := dao.Insert(user)
	return err
}

//SelectAnswerByUsername 通过昵称查找答案服务
func SelectAnswerByUsername(username string) string {
	answer := dao.SelectAnswerByUsername(username)
	return answer
}

//// SelectIdByUsername 通过昵称查找Id服务
//func SelectIdByUsername(username string) int {
//	id := dao.SelectIdByName(username)
//	return id
//}

//SelectQuestionByUsername 通过昵称查找问题
func SelectQuestionByUsername(username string) string {
	question := dao.SelectQuestionByUsername(username)
	return question
}

// ChangeSI 修改自我介绍
func ChangeSI(username, newSI string) error {
	err := dao.UpdateSI(username, newSI)
	return err
}

// GetUser 得到用户页
func GetUser(username string) ([]model.User2, error) {
	return dao.SelectUser(username)
}

func ChangeWTWId(name string, id int) error {
	err := dao.UpdateWTWId(name, id)
	return err
}

func ChangeWTWURL(name string, URL string) error {
	err := dao.UpdateWTWURL(name, URL)
	return err
}

func ChangeHWId(name string, id int) error {
	err := dao.UpdateHWId(name, id)
	return err
}

func ChangeHWURL(name string, URL string) error {
	err := dao.UpdateHWURL(name, URL)
	return err
}

func GetUserMovie(username string) ([]model.UserMovie, error) {
	return dao.SelectUserMovie(username)
}
