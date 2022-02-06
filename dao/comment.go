package dao

import (
	"douban/model"
	"fmt"
)

// InsertComment 插入回复
func InsertComment(comment model.Comment) error {
	sqlStr1 := `update topic set CommentNum=CommentNum+1 where id = ?`
	_, err := dB.Exec(sqlStr1, comment.TopicId)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	sqlStr2 := "insert into comment(Name,TopicId,Context,CommentTime)values (?,?,?,?)"
	_, err = dB.Exec(sqlStr2, comment.Name, comment.TopicId, comment.Context, comment.CommentTime)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

// SelectCommentByTopicId 查找评论
func SelectCommentByTopicId(topicId int) ([]model.Comment, error) {
	var comments []model.Comment

	rows, err := dB.Query("SELECT id, TopicId, Context, Name, CommentTime FROM comment WHERE TopicId = ?", topicId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var comment model.Comment

		err = rows.Scan(&comment.Id, &comment.TopicId, &comment.Context, &comment.Name, &comment.CommentTime)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

// DeleteComment 删除评论
func DeleteComment(Id int) error {
	sqlStr := `delete from comment where Id=?`
	_, err := dB.Exec(sqlStr, Id)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return err
	}
	return err
}

// SelectNameById2 通过id找到用户名
func SelectNameById2(commentId int) (string, error) {
	var comment model.Comment

	row := dB.QueryRow("SELECT Name FROM comment WHERE id = ? ", commentId)
	if row.Err() != nil {
		return comment.Name, row.Err()
	}

	err := row.Scan(&comment.Name)
	if err != nil {
		return comment.Name, err
	}

	return comment.Name, nil
}

// CommentLikes 评论点赞
func CommentLikes(id int) error {
	sqlStr := `update comment set Likes=Likes+1 where id = ?`
	_, err := dB.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}
