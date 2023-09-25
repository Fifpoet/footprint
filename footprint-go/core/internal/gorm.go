package internal

import (
	"github.com/fifpoet/footprint/global"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type DBBASE interface {
	GetLogMode() string
}

// Gorm 临时实例, 便于引用
var Gorm = new(_gorm)

type _gorm struct{}

// Config gorm 自定义配置
func (g *_gorm) Config(prefix string, singular bool) *gorm.Config {

	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,   // 表名前缀
			SingularTable: singular, // 是否使用单数形式的表名
		},
		// 是否在迁移时禁用外键约束，默认为 false，表示会根据模型之间的关联自动生成外键约束语句
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	_default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})

	var logMode DBBASE
	switch global.FP_CONFIG.App.DbType {
	case "mysql":
		logMode = &global.FP_CONFIG.MySQL
	default:
		logMode = &global.FP_CONFIG.MySQL
	}

	switch logMode.GetLogMode() {
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Info)
	}
	return config
}
