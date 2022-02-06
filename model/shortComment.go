package model

import (
	"time"
)

type ShortComment struct {
	Id          int       `json:"id"`
	MovieId     int       `json:"MovieId"`
	Name        string    `json:"Name"`
	Status      string    `json:"Status"`
	StarNum     int       `json:"StarNum"`
	CommentTime time.Time `json:"CommentTime"`
	Likes       int       `json:"Likes"`
	Context     string    `json:"Context"`
}

type SCPersonal struct {
	Id          int
	Name        string
	MovieName   string
	StarNum     int
	Context     string
	CommentTime time.Time
	Likes       int
}
