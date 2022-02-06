package service

import (
	"douban/dao"
	"douban/model"
)

// AddTopic 添加话题
func AddTopic(topic model.Topic) error {
	err := dao.InsertTopic(topic)
	return err
}

// DeleteTopic 删除话题
func DeleteTopic(topicId int) error {
	err := dao.DeleteTopic(topicId)
	return err
}

// GetTopics 得到话题
func GetTopics(movieId int) ([]model.Topic, error) {
	return dao.SelectTopic(movieId)
}

// GetTopicById 通过id得到话题
func GetTopicById(topicId int) (model.Topic, error) {
	return dao.SelectTopicById(topicId)
}

// GetNameById 通过id拿到用户名
func GetNameById(topicId int) (string, error) {
	return dao.SelectNameById(topicId)
}

// TopicLikes 话题点赞
func TopicLikes(topicId int) error {
	err := dao.TopicLikes(topicId)
	return err
}
