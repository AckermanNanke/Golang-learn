package initialize

import (
	"gorm-demo/dao"
	"gorm-demo/global"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化MYSQL表
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(&dao.FreeApi{}, &dao.User{})
	if err != nil {
		panic("初始化表失败")
		os.Exit(0)
	}
}

// 初始化数据库，失败时触发panic
func InitMySQL() *gorm.DB {
	m := global.GVA_CONFIG.Mysql
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         256,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		return nil
	}
	db_DB, _ := db.DB()
	db_DB.SetMaxIdleConns(m.MaxIdleConns)
	db_DB.SetMaxOpenConns(m.MaxOpenConns)
	return db
}
