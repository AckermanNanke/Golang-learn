package main

import (
	"gorm-demo/global"
	"gorm-demo/initialize"
	"gorm-demo/server"

	"go.uber.org/zap"
)

func main() {
	global.GVA_VP = initialize.InitViper()

	global.GVA_LOG = initialize.InitZap()
	zap.ReplaceGlobals(global.GVA_LOG)

	// global.MY_SQL = initialize.InitMySQL()
	// if global.MY_SQL != nil {
	// 	initialize.RegisterTables(global.MY_SQL)
	// 	mysql, _ := global.MY_SQL.DB()
	// 	defer mysql.Close()
	// }

	server.RunServer()
}
