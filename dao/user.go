package dao

import (
	"douban/model"
	"fmt"
)

// UpdatePassword 更新密码操作
func UpdatePassword(Name string, newPassword string) error {
	sqlStr := `update user set Password=? where Name = ?`
	_, err := dB.Exec(sqlStr, newPassword, Name)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

// SelectUserByUsername 查找用户
func SelectUserByUsername(Name string) (model.User, error) {
	user := model.User{}

	row := dB.QueryRow("SELECT id, password FROM user WHERE Name = ? ", Name)
	if row.Err() != nil {
		return user, row.Err()
	}

	err := row.Scan(&user.Id, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Insert 注册时插入数据
func Insert(user model.User) error {

	sqlStr := "insert into user(Name,Password,Question,Answer,RegisterTime)values (?,?,?,?,?)"
	_, err := dB.Exec(sqlStr, user.Name, user.Password, user.Question, user.Answer, user.RegisterTime)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

// SelectAnswerByUsername 通过用户名来找到密保的答案
func SelectAnswerByUsername(Name string) string {
	user := model.User{}
	sqlStr := `select answer from user where name=?;`
	dB.QueryRow(sqlStr, Name).Scan(&user.Answer)
	return user.Answer
}

//// SelectIdByName 通过用户名查Id
//func SelectIdByName(Name string) int {
//	user := model.User{}
//	sqlStr := `select Id from user where Name=?;`
//	dB.QueryRow(sqlStr, Name).Scan(&user.Id)
//	return user.Id
//}

// SelectQuestionByUsername 通过用户名来找到密保问题
func SelectQuestionByUsername(Name string) string {
	user := model.User{}
	sqlStr := `select Question from user where name=?;`
	dB.QueryRow(sqlStr, Name).Scan(&user.Question)
	if user.Question == "" {
		return ""
	}
	return user.Question
}

// UpdateSI 更新自我介绍
func UpdateSI(Name string, newSelfIntroduction string) error {
	sqlStr := `update user set SelfIntroduction=? where Name = ?`
	_, err := dB.Exec(sqlStr, newSelfIntroduction, Name)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func SelectUser(username string) ([]model.User2, error) {
	var users []model.User2
	rows, err := dB.Query("SELECT SelfIntroduction,RegisterTime FROM user WHERE Name = ?", username)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var user model.User2

		err = rows.Scan(&user.SelfIntroduction, &user.RegisterTime)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// UpdateWTWId 更新想看
func UpdateWTWId(Name string, id int) error {
	sqlStr := `update user set WantToWatchId=CONCAT_WS(',', WantToWatchId, ?) where Name = ?`
	_, err := dB.Exec(sqlStr, id, Name)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

// UpdateWTWURL 更新想看
func UpdateWTWURL(Name string, URL string) error {
	sqlStr := `update user set WantToWatchURL=CONCAT_WS(',', WantToWatchURL, ?) where Name = ?`
	_, err := dB.Exec(sqlStr, URL, Name)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

// UpdateHWId 更新看过
func UpdateHWId(Name string, id int) error {
	sqlStr := `update user set HaveWatchedId=CONCAT_WS(',', HaveWatchedId, ?) where Name = ?`
	_, err := dB.Exec(sqlStr, id, Name)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

// UpdateHWURL 更新看过
func UpdateHWURL(Name string, URL string) error {
	sqlStr := `update user set HaveWatchedURL=CONCAT_WS(',', HaveWatchedURL, ?) where Name = ?`
	_, err := dB.Exec(sqlStr, URL, Name)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func SelectUserMovie(username string) ([]model.UserMovie, error) {
	var users []model.UserMovie
	rows, err := dB.Query("SELECT WantToWatchId, WantToWatchURL, HaveWatchedId, HaveWatchedURL FROM user WHERE Name = ?", username)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var user model.UserMovie

		err = rows.Scan(&user.WantToWatchId, &user.WantToWatchURL, &user.HaveWatchedId, &user.HaveWatchedURL)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
