package core

// 路由分组
type RouterGroup struct {
	prefix      string        //路由前缀
	middlewares HandlersChain //中间件方法数组
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
	return newGroup
}

// 添加路由
func (g *RouterGroup) addRoute(method string, prefix string, handle HandlerFunc) {
	pattern := g.prefix + prefix
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

// 添加中间件
func (g *RouterGroup) Use(middleware ...HandlerFunc) {
	g.middlewares = append(g.middlewares, middleware...)
}
