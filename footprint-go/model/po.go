package model

import "gorm.io/gorm"

type Article struct {
	UserId     uint    `json:"user-id"`
	CategoryId uint    `json:"category-id"`
	Cover      string  `json:"cover"`
	Title      string  `json:"title"`
	Content    string  `json:"content"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
	City       string  `json:"city"`
	gorm.Model
}

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Role     string `json:"role"`
	gorm.Model
}
