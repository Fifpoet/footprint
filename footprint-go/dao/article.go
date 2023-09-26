package dao

import (
	"github.com/fifpoet/footprint/global"
	"github.com/fifpoet/footprint/model"
)

type ArticleRepo struct{}

func (r ArticleRepo) Update(art model.Article) error {
	tx := global.FP_DB.Create(&art)
	return tx.Error
}

func (r ArticleRepo) FindById(id uint) ([]model.Article, error) {
	as := make([]model.Article, 10)
	tx := global.FP_DB.Where("user_id = ?", id).Find(&as)
	return as, tx.Error
}
