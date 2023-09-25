package global

import (
	"github.com/fifpoet/footprint/config"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	FP_CONFIG config.Configuration
	FP_VIPER  *viper.Viper
	FP_LOG    *zap.Logger
	FP_DB     *gorm.DB
	FP_REDIS  *redis.Client
)
