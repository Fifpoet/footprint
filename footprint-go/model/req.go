package model

import "github.com/fifpoet/footprint/model/internal"

type LoginReq struct {
	UserName string
	Password string
	Email    string
}

type UploadArticleReq struct {
	Title    string
	Detail   string
	Lat      float64
	Lng      float64
	Location internal.Location
	Cover    string
}
