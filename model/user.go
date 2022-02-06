package model

import "time"

type User struct {
	Id               int
	Name             string
	Password         string
	Question         string
	Answer           string
	SelfIntroduction string
	RegisterTime     time.Time
}

type User2 struct {
	SelfIntroduction string
	RegisterTime     time.Time
}

type UserMovie struct {
	WantToWatchId  string
	WantToWatchURL string
	HaveWatchedId  string
	HaveWatchedURL string
}
