package dao

import (
	"github.com/fifpoet/footprint/global"
	"github.com/fifpoet/footprint/model"
	utils "github.com/fifpoet/footprint/util"
	"gorm.io/gorm"
)

type ArticleRepo struct{}

func (r ArticleRepo) Add(art model.Article) error {
	tx := global.FP_DB.Create(&art)
	return tx.Error
}

func (r ArticleRepo) Update(art model.Article) error {
	// Save更新零值
	tx := global.FP_DB.Where("id = ?", art.Model.ID).Updates(&art)
	return tx.Error
}

func (r ArticleRepo) FindByPage(pageNo, pageSize int) ([]model.Article, error) {
	as := make([]model.Article, 10)
	tx := global.FP_DB.Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&as)
	return as, tx.Error
}

func (r ArticleRepo) DelById(id string) error {
	uid, err := utils.String2Uint(id)
	if err != nil {
		return err
	}
	user := model.Article{Model: gorm.Model{ID: uid}}
	tx := global.FP_DB.Where("id = ?", id).Delete(&user)
	return tx.Error
}
