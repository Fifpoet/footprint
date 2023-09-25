package main

import (
	"fmt"
	"github.com/fifpoet/footprint/core"
	"github.com/fifpoet/footprint/global"
	"github.com/fifpoet/footprint/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const AppMode = "debug"

func main() {

	gin.SetMode(AppMode)

	//	1.配置初始化
	global.FP_VIPER = core.InitializeViper()
	//	2.日志
	global.FP_LOG = core.InitializeZap()
	zap.ReplaceGlobals(global.FP_LOG)
	global.FP_LOG.Info("server run success on ", zap.String("zap_log", "zap_log"))
	//  3.数据库连接
	global.FP_DB = core.InitializeGorm()
	global.FP_REDIS = core.InitializeRedis()
	//	4.其他初始化
	var user model.User
	fmt.Println(global.FP_DB.First(&user, 1))
	//	5.启动服务
	core.RunServer()

}
