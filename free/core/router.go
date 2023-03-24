package core

import "net/http"

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

// 绑定 handser 请求
func (r *Router) handle(c *Context) {
	key := c.Path
	if h, ok := r.handlers[key]; ok {
		h(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
