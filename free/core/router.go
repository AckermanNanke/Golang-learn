package core

import (
	"net/http"
	"strings"
)

type Router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// 路由地址切割
// part != "" 判断时不把最开始 / 符号前面的空内容添加进去
func (r *Router) parsePattern(pattern string) []string {
	partsOld := strings.Split(pattern, "/")
	partsNew := make([]string, 0)

	for _, part := range partsOld {
		if part != "" {
			partsNew = append(partsNew, part)
			if part[0] == '*' {
				break
			}
		}
	}

	return partsNew
}

// 添加路由 添加方法分类 ——》 插入树节点 ——》添加对应方法
func (r *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := r.parsePattern(pattern)

	key := pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
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
