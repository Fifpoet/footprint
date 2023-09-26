package model

import "gorm.io/gorm"

type Article struct {
	UserId       uint
	CategoryId   uint
	Cover        string
	Title        string
	Content      string
	Lat          float64
	Lng          float64
	City         string
	District     string
	Province     string
	street       string
	streetNumber string
	gorm.Model
}

type User struct {
	Name     string
	Email    string
	Password string
	gorm.Model
}
