package core

import (
	"net/http"
	"sync"
)

// 请求处理函数
type HandlerFunc func(*Context)

// 请求处理函数列表
type HandlersChain []HandlerFunc

// 参数对象
type Param struct {
	Key   string
	Value string
}

type Params []Param

// 路由体对象
type Fhttp struct {
	Route
	pool sync.Pool //使用连接池处理并发
}

func New() *Fhttp {
	f := &Fhttp{
		Route: Route{
			handlers: make(map[string]HandlerFunc),
		},
	}
	return f
}

// 实现ServerHttp接口
// 从连接池内取出请求示例后再放回去
func (f *Fhttp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w, r)
	f.Route.handle(c)
}

// 调用路由模块添加方法 - 添加路由
func (f *Fhttp) AddRoute(method string, pattern string, handler HandlerFunc) {
	f.Route.addRoute(method, pattern, handler)
}

// 添加 GET 请求
func (f *Fhttp) GET(pattern string, handler HandlerFunc) {
	f.Route.GET(pattern, handler)
}

// 添加 POST 请求
func (f *Fhttp) POST(pattern string, handler HandlerFunc) {
	f.Route.POST(pattern, handler)
}
