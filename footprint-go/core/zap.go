package core

import (
	"fmt"
	"github.com/fifpoet/footprint/core/internal"
	"github.com/fifpoet/footprint/global"
	utils "github.com/fifpoet/footprint/util"

	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitializeZap Zap 获取 zap.Logger
func InitializeZap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.FP_CONFIG.Zap.Director); !ok {
		// 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.FP_CONFIG.Zap.Director)
		_ = os.Mkdir(global.FP_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.FP_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	fmt.Println("====2-zap====: zap log init success")
	return logger
}
