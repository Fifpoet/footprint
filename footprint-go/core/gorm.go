package core

import (
	"fmt"
	"github.com/fifpoet/footprint/core/internal"
	"github.com/fifpoet/footprint/global"
	"github.com/fifpoet/footprint/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitializeGorm Gorm 初始化数据库并产生数据库全局变量
func InitializeGorm() *gorm.DB {
	var db *gorm.DB
	switch global.FP_CONFIG.App.DbType {
	case "mysql":
		db = GormMysql()
	default:
		db = GormMysql()
	}
	CreateAllTable(db)
	return db
}

// GormMysql 初始化Mysql数据库
func GormMysql() *gorm.DB {
	m := global.FP_CONFIG.MySQL
	if m.Dbname == "" {
		return nil
	}

	// MySQL连接配置
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         255,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular))

	// 将引擎设置为我们配置的引擎，并设置每个连接的最大空闲数和最大连接数。
	if err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)

		fmt.Println("====3-gorm====: gorm link mysql success")
		return db
	}
}

func CreateAllTable(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Article{})
}
