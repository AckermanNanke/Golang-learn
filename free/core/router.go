package core

type Router struct {
	handlers map[string]HandlerFunc
}

// 添加路由
func (r *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := pattern
	r.handlers[key] = handler
}

// 添加 GET 请求
func (r *Router) GET(pattern string, handler HandlerFunc) {
	r.addRoute("GET", pattern, handler)
}

// 添加 POST 请求
func (r *Router) POST(pattern string, handler HandlerFunc) {
	r.addRoute("POST", pattern, handler)
}
