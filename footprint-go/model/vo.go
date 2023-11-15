package model

type LoginVO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Role     []int  `json:"role"`
	ID       uint
}
