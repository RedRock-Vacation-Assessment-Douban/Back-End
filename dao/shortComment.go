package dao

import (
	"douban/model"
	"fmt"
)

// InsertShortComment 向话题中插入(发布话题)
func InsertShortComment(shortComment model.ShortComment) error {
	_, err := dB.Exec("INSERT INTO shortComment(MovieId, Name, Status, StarNum, CommentTime, Likes, Context, MovieName) "+"values(?, ?, ?, ?, ?, ?, ?, ?);", shortComment.MovieId, shortComment.Name, shortComment.Status, shortComment.StarNum, shortComment.CommentTime, shortComment.Likes, shortComment.Context, shortComment.MovieName)
	return err
}

//// SelectShortCommentById 通过id来搜索短评
//func SelectShortCommentById(shortCommentId int) (model.ShortComment, error) {
//	var shortComment model.ShortComment
//
//	row := dB.QueryRow("SELECT id, MovieId, Name, Status, StarNum, CommentTime, Likes, Context FROM shortComment WHERE id = ? ", shortCommentId)
//	if row.Err() != nil {
//		return shortComment, row.Err()
//	}
//
//	err := row.Scan(&shortComment.Id, &shortComment.MovieId, &shortComment.Name, &shortComment.Status, &shortComment.StarNum, &shortComment.CommentTime, &shortComment.Likes, &shortComment.Context)
//	if err != nil {
//		return shortComment, err
//	}
//
//	return shortComment, nil
//}

// SelectNameBySCId 通过id查找发布用户
func SelectNameBySCId(SCId int) (string, error) {
	var shortComment model.ShortComment

	row := dB.QueryRow("SELECT Name FROM shortComment WHERE id = ? ", SCId)
	if row.Err() != nil {
		return shortComment.Name, row.Err()
	}

	err := row.Scan(&shortComment.Name)
	if err != nil {
		return shortComment.Name, err
	}

	return shortComment.Name, nil
}

// SelectShortComment 查找短评
func SelectShortComment(movieId int) ([]model.ShortComment, error) {
	var shortComments []model.ShortComment
	rows, err := dB.Query("SELECT id, MovieId, Name, Status, StarNum, CommentTime, Likes, Context FROM shortComment where MovieId = ?", movieId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var shortComment model.ShortComment

		err = rows.Scan(&shortComment.Id, &shortComment.MovieId, &shortComment.Name, &shortComment.Status, &shortComment.StarNum, &shortComment.CommentTime, &shortComment.Likes, &shortComment.Context)
		if err != nil {
			return nil, err
		}

		shortComments = append(shortComments, shortComment)
	}

	return shortComments, nil
}

// SelectShortCommentByUsername 查找短评
func SelectShortCommentByUsername(name string) ([]model.SCPersonal, error) {
	var shortComments []model.SCPersonal
	rows, err := dB.Query("SELECT id, MovieName, Name, StarNum, CommentTime, Likes, Context FROM shortComment where Name = ?", name)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var shortComment model.SCPersonal

		err = rows.Scan(&shortComment.Id, &shortComment.MovieName, &shortComment.Name, &shortComment.StarNum, &shortComment.CommentTime, &shortComment.Likes, &shortComment.Context)
		if err != nil {
			return nil, err
		}

		shortComments = append(shortComments, shortComment)
	}

	return shortComments, nil
}

// DeleteShortComment 删除短评
func DeleteShortComment(SCId int) error {

	sqlStr := `delete from shortComment where Id=?`
	_, err := dB.Exec(sqlStr, SCId)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return err
	}
	return err
}

// ShortCommentLikes 给话题点赞
func ShortCommentLikes(id int) error {
	sqlStr := `update shortComment set Likes=Likes+1 where id = ?`
	_, err := dB.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}
