package core

import (
	"gorm-demo/global"

	"go.uber.org/zap"
)

// 路由分组
type RouterGroup struct {
	prefix      string        //路由前缀
	middlewares []HandlerFunc //中间件方法数组
	f           *Fhttp
}

// 创建路由分组
func (g *RouterGroup) Group(prefix string) *RouterGroup {
	f := g.f
	newGroup := &RouterGroup{
		prefix: g.prefix + prefix,
		f:      f,
	}
	f.groups = append(f.groups, newGroup)
	global.GVA_LOG.Info("newGroup", zap.Any("newGroup", newGroup.prefix))
	return newGroup
}

// 添加路由
func (g *RouterGroup) addRoute(method string, prefix string, handle HandlerFunc) {
	pattern := g.prefix + prefix
	global.GVA_LOG.Info("pattern", zap.Any("pattern", g.prefix))
	g.f.router.addRoute(method, pattern, handle)
}

// 添加路由
func (g *RouterGroup) GET(prefix string, handle HandlerFunc) {
	g.addRoute("GET", prefix, handle)
}

// 添加路由
func (g *RouterGroup) POST(prefix string, handle HandlerFunc) {
	g.addRoute("POST", prefix, handle)
}
