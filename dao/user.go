package dao

import (
	"douban/model"
	"fmt"
)

// UpdatePassword 更新密码操作
//db.Model(&model.User{}).Where("Name = ", "yxh").Update("Name", "test")
func UpdatePassword(Name string, newPassword string) error {
	//sqlStr := `update user set Password=? where Name = ?`
	//_, err := dB.Exec(sqlStr, newPassword, Name)
	deRes := db.Model(&model.User{}).Where("Name = ?", Name).Update("Password", newPassword)
	err := deRes.Error
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

// SelectUserByUsername 查找用户
//var U []model.UserInfo
//db.Model(&model.User{}).Select("Name").Where("Name LIKE ?", "%袁%").Find(&U)
func SelectUserByUsername(Name string) (model.UserInfo, error) {
	//var user model.User
	var user model.UserInfo
	//row := dB.QueryRow("SELECT id, password FROM user WHERE Name = ? ", Name)
	//if row.Err() != nil {
	//	return user, row.Err()
	//}
	//err := row.Scan(&user.Id, &user.Password)
	dbRes := db.Model(&model.User{}).Select("id", "password").Where("Name = ?", Name).First(&user)
	err := dbRes.Error
	if err != nil {
		return user, err
	}
	fmt.Println(user)
	return user, nil
}

// Insert 注册时插入数据
func Insert(user model.User) error {
	//sqlStr := "insert into user(Name,Password,Question,Answer,RegisterTime)values (?,?,?,?,?)"
	//_, err := dB.Exec(sqlStr, user.Name, user.Password, user.Question, user.Answer, user.RegisterTime)
	deres := db.Select("Name", "Password", "Question", "Answer", "RegisterTime").Create(&model.User{Name: user.Name, Password: user.Password, Question: user.Question, Answer: user.Answer, RegisterTime: user.RegisterTime})
	err := deres.Error
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

// SelectAnswerByUsername 通过用户名来找到密保的答案
func SelectAnswerByUsername(Name string) string {
	user := model.User{}
	//sqlStr := `select answer from user where name=?;`
	//dB.QueryRow(sqlStr, Name).Scan(&user.Answer)
	db.Model(&model.User{}).Select("answer").Where("Name = ?", Name).Find(&user)
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
	//sqlStr := `select Question from user where name=?;`
	//dB.QueryRow(sqlStr, Name).Scan(&user.Question)
	db.Model(&model.User{}).Select("question").Where("Name = ?", Name).Find(&user)
	if user.Question == "" {
		return ""
	}
	return user.Question
}

// UpdateSI 更新自我介绍
func UpdateSI(Name string, newSelfIntroduction string) error {
	//sqlStr := `update user set SelfIntroduction=? where Name = ?`
	//_, err := dB.Exec(sqlStr, newSelfIntroduction, Name)
	dbRes := db.Model(&model.User{}).Where("Name = ?", Name).Update("SelfIntroduction", newSelfIntroduction)
	err := dbRes.Error
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func SelectUser(username string) ([]model.User2, error) {
	var users []model.User2
	dbRes := db.Model(&model.User{}).Select("RegisterTime", "SelfIntroduction").Where("Name = ?", username).Find(&users)
	err := dbRes.Error
	if err != nil {
		fmt.Printf("select failed, err:%v\n", err)
		return users, err
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
	dbRes := db.Model(&model.User{}).Select("WantToWatchId", "WantToWatchURL", "HaveWatchedId", "HaveWatchedURL").Where("Name = ?", username).Find(&users)
	err := dbRes.Error
	if err != nil {
		fmt.Printf("select failed, err:%v\n", err)
		return users, err
	}
	return users, nil
}

//func UserRecommend(user string, movie string) error {
//	_, err := rdb.ZIncrBy(user, 1, movie).Result()
//	if err != nil {
//		fmt.Printf("zincrby failed, err:%v\n", err)
//		return err
//	}
//	return err
//}

//func Recommend(user string) ([]redis.Z, error) {
//	ret, err := rdb.ZRevRangeWithScores(user, 0, 2).Result()
//	if err != nil {
//		fmt.Printf("zrevrange failed, err:%v\n", err)
//		return _, err
//	}
//
//	return ret, err
//}
