package router

import (
	"gorm-demo/controller"
	"gorm-demo/core"
	"net/http"
)

func InitRouter() *core.Fhttp {
	router := core.New()
	router.POST("/register", controller.UserApi.Register)
	router.POST("/login", controller.UserApi.Register)
	router.AddRoute("GET", "/logout", controller.UserApi.Register)
	v1 := router.Group("/v1")
	{
		v1.GET("/", func(c *core.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *core.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s", c.Query("name"), c.Path)
		})
	}
	return router
}
