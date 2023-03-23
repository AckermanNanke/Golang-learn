package router

import (
	"gorm-demo/controller"
	"gorm-demo/core"
)

func InitRouter() *core.Fhttp {
	router := core.New()
	router.POST("/register", controller.UserApi.Register)
	return router
}
