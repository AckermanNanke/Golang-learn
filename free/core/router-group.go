package core

// // 路由分组
// type RouterGroup struct {
// 	prefix      string        //路由前缀
// 	middlewares []HandlerFunc //中间件方法数组
// 	parent      *RouterGroup  // 父路由
// 	f           *Fhttp
// }

// // 创建路由分组
// func (g *RouterGroup) Group(prefix string) *RouterGroup {
// 	f := g.f
// 	newGroup := &RouterGroup{
// 		prefix: g.prefix + f.prefix,
// 		parent: g,
// 		f:      f,
// 	}
// 	f.Groups = append(f.Groups, newGroup)
// 	return newGroup
// }

// // 添加路由
// func (g *RouterGroup) addRoute(method string, prefix string, handle HandlerFunc) {
// 	pattern := g.prefix + prefix
// 	g.f.addRoute(method, pattern, handle)
// }

// // 添加路由
// func (g *RouterGroup) GET(prefix string, handle HandlerFunc) {
// 	g.f.addRoute("GET", prefix, handle)
// }

// // 添加路由
// func (g *RouterGroup) POST(prefix string, handle HandlerFunc) {
// 	g.f.addRoute("POST", prefix, handle)
// }
