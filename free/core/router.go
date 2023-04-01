package core

import (
	"net/http"
	"strings"
)

type Route struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *Route {
	return &Route{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// 路由地址切割
// part != "" 判断时不把最开始 / 符号前面的空内容添加进去
func (r *Route) parsePattern(pattern string) []string {
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
func (r *Route) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := r.parsePattern(pattern)

	key := pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

// 获取路由
func (r *Route) GetRoute(method string, path string) (*node, map[string]string) {
	pathArr := r.parsePattern(path)
	params := make(map[string]string)

	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	n := root.search(pathArr, 0)
	if n != nil {
		parts := r.parsePattern(n.pattern)
		for i, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = pathArr[i]
			}
			if part[0] == '*' && len(part) > 0 {
				params[part[1:]] = strings.Join(pathArr[i:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

// 添加 GET 请求
func (r *Route) GET(pattern string, handler HandlerFunc) {
	r.addRoute("GET", pattern, handler)
}

// 添加 POST 请求
func (r *Route) POST(pattern string, handler HandlerFunc) {
	r.addRoute("POST", pattern, handler)
}

// 绑定 handser 请求
func (r *Route) handle(c *Context) {
	key := c.Path
	if h, ok := r.handlers[key]; ok {
		h(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
