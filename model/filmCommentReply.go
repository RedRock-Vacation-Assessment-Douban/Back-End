package model

import "time"

type FilmCommentReply struct {
	Id            int
	FilmCommentId int
	Context       string
	Name          string
	CommentTime   time.Time
}
