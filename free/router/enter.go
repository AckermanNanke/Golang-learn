package router

import (
	"gorm-demo/controller"
	"gorm-demo/core"
	"net/http"
)

func InitRouter() *core.Fhttp {
	r := core.Default()
	r.POST("/register", controller.UserApi.Register)
	r.POST("/login", controller.UserApi.Register)
	r.AddRoute("GET", "/logout", controller.UserApi.Register)
	r.GET("/panic", func(c *core.Context) {
		arr := []string{
			"123456",
		}
		c.String(http.StatusOK, arr[100])
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *core.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *core.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s", c.Query("name"), c.Path)
		})
	}
	return r
}
