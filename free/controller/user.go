package controller

import (
	"gorm-demo/core"
	"net/http"
)

type User struct{}

func (u User) Register(c *core.Context) {
	c.JSON(http.StatusOK, core.FJSON{
		"username": c.PostForm("username"),
		"password": c.PostForm("password"),
	})
}

// func (u User) Register(w http.ResponseWriter, r *http.Request) {
// 	global.GVA_LOG.Info("请求数据", zap.Any("Header", r.Header))
// 	global.GVA_LOG.Info("请求数据", zap.Any("Method", r.Method))
// 	global.GVA_LOG.Info("请求数据", zap.Any("PostForm", req))
// 	user := &freedb.FreeUsers{UserName: req.UserName, Phone: req.Phone, Password: req.Password}
// }
