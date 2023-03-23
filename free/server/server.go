package server

import (
	"fmt"
	"gorm-demo/global"
	"gorm-demo/initialize"
	"gorm-demo/router"
)

// 初始化路由
// 初始化服务配置
// 启动服务
func RunServer() {
	// int -> string
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)

	router := router.InitRouter()
	s := initialize.InitServer(address, router)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
