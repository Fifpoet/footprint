package model

import "gorm.io/gorm"

type Article struct {
	UserId     uint
	CategoryId uint
	Cover      string
	Header     string
	Content    string
	gorm.Model
}

type User struct {
	Name     string
	Email    string
	Password string
	gorm.Model
}
