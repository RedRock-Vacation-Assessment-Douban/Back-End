package model

import "time"

type FilmComment struct {
	Id         int       `json:"id"`
	MovieId    int       `json:"MovieId" gorm:"column:MovieId"`
	Context    string    `json:"context"`
	Name       string    `json:"name"`
	PostTime   time.Time `json:"post_time" gorm:"column:PostTime"`
	CommentNum int       `json:"comment_num" gorm:"column:CommentNum"`
	StarNum    int       `json:"star_num" gorm:"column:StarNum"`
	Likes      int       `json:"likes"`
	Down       int       `json:"down"`
	MovieName  string    `json:"movie_name" gorm:"column:MovieName"`
	URL        string    `json:"URL"`
}

func (FilmComment) TableName() string {
	return "filmComment"
}

type FilmCommentDetail struct {
	FilmComment
	FilmCommentReplys []FilmCommentReply
}

type MostPopularFC struct {
	Id        int
	Name      string
	Context   string
	StarNum   string
	MovieId   int
	URL       string
	MovieName string
}

type Personal struct {
	Id         int
	MovieId    string
	MovieName  string
	Name       string
	Context    string
	StarNum    string
	Likes      int
	CommentNum int
	PostTime   time.Time
	URL        string
}
