package model

import "gorm.io/gorm"

type UserInfo struct {
	Name     string
	Email    string
	Password string
	Model    gorm.Model
}
