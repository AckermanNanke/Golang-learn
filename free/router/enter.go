package router

import (
	"gorm-demo/controller"
	"gorm-demo/core"
)

func InitRouter() *core.Fhttp {
	router := core.New()
	router.POST("/register", controller.UserApi.Register)
	router.POST("/login", controller.UserApi.Register)
	router.AddRoute("GET", "/logout", controller.UserApi.Register)
	return router
}
