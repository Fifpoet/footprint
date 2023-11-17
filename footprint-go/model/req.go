package model

type LoginReq struct {
	UserName string
	Password string
}

type RegisterReq struct {
	UserName string
	Password string
	Email    string
}

type AddArticleReq struct {
	UserId     uint
	CategoryId uint
	Cover      string
	Title      string
	Content    string
	Lat        float64
	Lng        float64
	City       string
}

type UpdateArticleReq struct {
	ID         uint
	UserId     uint
	CategoryId uint
	Cover      string
	Title      string
	Content    string
	Lat        float64
	Lng        float64
	City       string
}
