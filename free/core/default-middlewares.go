package core

import (
	"gorm-demo/global"

	"go.uber.org/zap"
)

// 日志打印
func logger() HandlerFunc {
	return func(c *Context) {
		global.GVA_LOG.Info("请求日志", zap.Any("fhttp", c.Path))
		c.Next()
	}
}

// panic恢复
func recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				global.GVA_LOG.Error("请求失败", zap.Any("fhttp", c.Path))
				c.DefaultError(500)
			}
		}()

		c.Next()
	}
}
