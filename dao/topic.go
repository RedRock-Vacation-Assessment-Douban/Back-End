package dao

import (
	"douban/model"
	"fmt"
)

// InsertTopic 向话题中插入(发布话题)
func InsertTopic(topic model.Topic) error {
	_, err := dB.Exec("INSERT INTO topic(MovieId, Name, Context, PostTime) "+"values(?, ?, ?, ?);", topic.MovieId, topic.Name, topic.Context, topic.PostTime)
	return err
}

// SelectTopicById 通过id来搜索话题
func SelectTopicById(topicId int) (model.Topic, error) {
	var topic model.Topic

	row := dB.QueryRow("SELECT id, MovieId, Name, Context, PostTime, CommentNum, Likes FROM topic WHERE id = ? ", topicId)
	if row.Err() != nil {
		return topic, row.Err()
	}

	err := row.Scan(&topic.Id, &topic.MovieId, &topic.Name, &topic.Context, &topic.PostTime, &topic.CommentNum, &topic.Likes)
	if err != nil {
		return topic, err
	}

	return topic, nil
}

// SelectNameById 通过id查找发布用户
func SelectNameById(topicId int) (string, error) {
	var topic model.Topic

	row := dB.QueryRow("SELECT Name FROM topic WHERE id = ? ", topicId)
	if row.Err() != nil {
		return topic.Name, row.Err()
	}

	err := row.Scan(&topic.Name)
	if err != nil {
		return topic.Name, err
	}

	return topic.Name, nil
}

// SelectTopic 查找话题
func SelectTopic(movieId int) ([]model.Topic, error) {
	var topics []model.Topic
	rows, err := dB.Query("SELECT id, MovieId, Name, Context, PostTime, CommentNum, Likes FROM topic where MovieId = ?", movieId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var topic model.Topic

		err = rows.Scan(&topic.Id, &topic.MovieId, &topic.Name, &topic.Context, &topic.PostTime, &topic.CommentNum, &topic.Likes)
		if err != nil {
			return nil, err
		}

		topics = append(topics, topic)
	}

	return topics, nil
}

// DeleteTopic 删除话题
func DeleteTopic(topicId int) error {

	sqlStr := `delete from topic where Id=?`
	_, err := dB.Exec(sqlStr, topicId)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return err
	}
	return err
}

// TopicLikes 给话题点赞
func TopicLikes(id int) error {
	sqlStr := `update topic set Likes=Likes+1 where id = ?`
	_, err := dB.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}
