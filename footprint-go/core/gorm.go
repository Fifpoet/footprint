package core

import (
	"github.com/fifpoet/footprint/global"
	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB {
	switch global.FP_CONFIG.App.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}
