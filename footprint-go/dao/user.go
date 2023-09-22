package dao

import (
	"github.com/fifpoet/footprint/model"
)

var us = []model.UserInfo{
	{
		Name:     "admin",
		Password: "123",
	},
	{
		Name:     "11",
		Password: "11",
	},
}

type UserRepo struct {
	Users []model.UserInfo
}

func (r UserRepo) FindById(id uint) (model.UserInfo, error) {
	return us[0], nil
}
