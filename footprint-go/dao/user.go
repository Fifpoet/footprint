package dao

import (
	"github.com/fifpoet/footprint/global"
	"github.com/fifpoet/footprint/model"
)

type UserRepo struct{}

func FindByName(name string) (*model.User, error) {
	var res model.User
	tx := global.FP_DB.Where("name = ?", name).First(&res)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &res, nil
}
