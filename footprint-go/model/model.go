package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Header  string
	Content string
}

type UserInfo struct {
	Name     string
	Email    string
	Password string
	Model    gorm.Model
}
