package main

import (
	"gorm-demo/global"
	"gorm-demo/initialize"
)

func main() {
	global.GVA_VP = initialize.InitViper()
	global.MY_SQL = initialize.InitMySQL()
	if global.MY_SQL != nil {
		initialize.RegisterTables(global.MY_SQL)
		mysql, _ := global.MY_SQL.DB()
		defer mysql.Close()
	}

	global.SERVE_ADDR = ":8999"
	myserve := initialize.InitServer(global.SERVE_ADDR)
	myserve.ListenAndServe()
}
