package initialize

import (
	"fmt"
	"gorm-demo/global"
	"gorm-demo/initialize/initconfig"
	"gorm-demo/utils"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 初始化日志库方法
// GetZapCores获取日志配置数组、
// zap.New()初始化日志
// 指南：https://juejin.cn/post/7182814706757795897
func InitZap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.GVA_CONFIG.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.GVA_CONFIG.Zap.Director)
		_ = os.Mkdir(global.GVA_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := initconfig.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.GVA_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
