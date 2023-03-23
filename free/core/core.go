package core

import (
	"net/http"
	"sync"
)

// 请求处理函数
type HandlerFunc func(http.ResponseWriter, *http.Request)

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
	Router map[string]HandlerFunc
	pool   sync.Pool //使用连接池处理并发
}

// 实现ServerHttp接口
// 从连接池内取出请求示例后再放回去
func (f *Fhttp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path
	if h, ok := f.Router[key]; ok {
		h(w, r)
	} else {
		panic("404 NOT FOUND: %s\n")
	}
}

func New() *Fhttp {
	f := &Fhttp{
		Router: make(map[string]HandlerFunc),
	}
	return f
}

// 添加路由
func (f *Fhttp) addRoute(method string, pattern string, handler HandlerFunc) {
	key := pattern
	f.Router[key] = handler
}

// 添加 GET 请求
func (f *Fhttp) GET(pattern string, handler HandlerFunc) {
	f.addRoute("GET", pattern, handler)
}

// 添加 POST 请求
func (f *Fhttp) POST(pattern string, handler HandlerFunc) {
	f.addRoute("POST", pattern, handler)
}
