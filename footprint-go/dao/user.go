package dao

import (
	"github.com/fifpoet/footprint/global"
	"github.com/fifpoet/footprint/model"
	"gorm.io/gorm"
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

func Add(user model.RegisterReq) error {
	tx := global.FP_DB.Create(&model.User{
		Name:     user.UserName,
		Email:    user.Email,
		Password: user.Password,
		Model:    gorm.Model{},
	})
	return tx.Error
}
