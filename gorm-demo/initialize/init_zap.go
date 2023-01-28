package initialize

import (
	"fmt"
	"gorm-demo/global"
	"gorm-demo/utils"
	"os"

	"go.uber.org/zap"
)

// 初始化日志库方法
// 1.创建日志存放文件夹
// 指南：https://juejin.cn/post/7182814706757795897
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.GVA_CONFIG.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.GVA_CONFIG.Zap.Director)
		_ = os.Mkdir(global.GVA_CONFIG.Zap.Director, os.ModePerm)
	}
}
