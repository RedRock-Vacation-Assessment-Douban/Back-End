package model

import "time"

type User struct {
	Id               int
	Name             string
	Password         string
	Question         string
	Answer           string
	SelfIntroduction string    `gorm:"column:SelfIntroduction"`
	RegisterTime     time.Time `gorm:"column:RegisterTime"`
}

type UserInfo struct {
	Id       int
	Password string
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

type Recommend struct {
	Movie1 string
	Movie2 string
	Movie3 string
}
