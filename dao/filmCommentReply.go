package dao

import (
	"douban/model"
	"fmt"
)

// InsertFilmCommentReply 插入回复
func InsertFilmCommentReply(filmCommentReply model.FilmCommentReply) error {
	sqlStr1 := `update filmComment set CommentNum=CommentNum+1 where id = ?`
	_, err := dB.Exec(sqlStr1, filmCommentReply.FilmCommentId)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	sqlStr2 := "insert into filmCommentReply(Name,FilmCommentId,Context,CommentTime)values (?,?,?,?)"
	_, err = dB.Exec(sqlStr2, filmCommentReply.Name, filmCommentReply.FilmCommentId, filmCommentReply.Context, filmCommentReply.CommentTime)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

// SelectCommentByFCId 查找评论
func SelectCommentByFCId(FCId int) ([]model.FilmCommentReply, error) {
	var filmCommentReplys []model.FilmCommentReply

	rows, err := dB.Query("SELECT id, FilmCommentId, Context, Name, CommentTime FROM filmCommentReply WHERE FilmCommentId = ?", FCId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var filmCommentReply model.FilmCommentReply

		err = rows.Scan(&filmCommentReply.Id, &filmCommentReply.FilmCommentId, &filmCommentReply.Context, &filmCommentReply.Name, &filmCommentReply.CommentTime)
		if err != nil {
			return nil, err
		}

		filmCommentReplys = append(filmCommentReplys, filmCommentReply)
	}

	return filmCommentReplys, nil
}

// DeleteFilmCommentReply 删除评论
func DeleteFilmCommentReply(Id int) error {
	sqlStr := `delete from filmCommentReply where Id=?`
	_, err := dB.Exec(sqlStr, Id)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return err
	}
	return err
}

// SelectNameByFCRId 通过id找到用户名
func SelectNameByFCRId(filmCommentReplyId int) (string, error) {
	var filmCommentReply model.FilmCommentReply

	row := dB.QueryRow("SELECT Name FROM filmCommentReply WHERE id = ? ", filmCommentReplyId)
	if row.Err() != nil {
		return filmCommentReply.Name, row.Err()
	}

	err := row.Scan(&filmCommentReply.Name)
	if err != nil {
		return filmCommentReply.Name, err
	}

	return filmCommentReply.Name, nil
}
