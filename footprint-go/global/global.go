package global

import (
	"github.com/fifpoet/footprint/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	FP_CONFIG config.Configuration
	FP_VIPER  *viper.Viper
	FP_LOG    *zap.Logger
)
