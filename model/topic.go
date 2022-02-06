package model

import "time"

type Topic struct {
	Id         int       `json:"id"`
	MovieId    int       `json:"MovieId"`
	Context    string    `json:"context"`
	Name       string    `json:"name"`
	PostTime   time.Time `json:"post_time"`
	CommentNum int       `json:"comment_num"`
	Likes      int       `json:"likes"`
}

type TopicDetail struct {
	Topic
	Comments []Comment
}
