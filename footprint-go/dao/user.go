package dao

import (
	"github.com/fifpoet/footprint/global"
	"github.com/fifpoet/footprint/model"
)

var us = []model.User{
	{
		Name:     "admin",
		Password: "123",
	},
	{
		Name:     "11",
		Password: "11",
	},
}
var db = global.FP_DB

type UserRepo struct {
	Users []model.User
}

func (r UserRepo) FindById(id uint) (UserRepo, error) {
	var res UserRepo
	if err := db.First(&res, id).Error; err != nil {
		return res, err
	}
	return res, nil
}
