package router

import (
	"gorm-demo/controller"
	"gorm-demo/core"
	"gorm-demo/global"
	"net/http"

	"go.uber.org/zap"
)

func logger() core.HandlerFunc {
	return func(c *core.Context) {
		global.GVA_LOG.Info("请求日志", zap.Any("fhttp", c.Path))
		c.Next()
	}
}

func InitRouter() *core.Fhttp {
	r := core.New()
	r.Use(logger())
	r.POST("/register", controller.UserApi.Register)
	r.POST("/login", controller.UserApi.Register)
	r.AddRoute("GET", "/logout", controller.UserApi.Register)
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
