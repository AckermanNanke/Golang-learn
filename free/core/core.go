package core

import (
	"net/http"
	"strings"
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
	*RouterGroup                //相当于继承RouterGroup
	router       *router        //路由树
	groups       []*RouterGroup //存放所有路由分组
	Pool         sync.Pool      //使用连接池处理并发
}

func New() *Fhttp {
	f := &Fhttp{
		router: newRouter(),
	}
	f.RouterGroup = &RouterGroup{f: f}
	f.groups = []*RouterGroup{
		f.RouterGroup,
	}
	return f
}

// 默认模板
// 使用日志打印功能与panic恢复功能
func Default() *Fhttp {
	f := New()
	f.Use(logger(), recovery())
	return f
}

// 实现ServerHttp接口
// 从连接池内取出请求示例后再放回去
func (f *Fhttp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares HandlersChain
	for _, g := range f.groups {
		if strings.HasPrefix(r.URL.Path, g.prefix) {
			middlewares = append(middlewares, g.middlewares...)
		}
	}
	c := NewContext(w, r)
	c.handlers = middlewares
	f.router.handle(c)
}

// 调用路由模块添加方法 - 添加路由
func (f *Fhttp) AddRoute(method string, pattern string, handler HandlerFunc) {
	f.router.addRoute(method, pattern, handler)
}

// 添加 GET 请求
func (f *Fhttp) GET(pattern string, handler HandlerFunc) {
	f.router.GET(pattern, handler)
}

// 添加 POST 请求
func (f *Fhttp) POST(pattern string, handler HandlerFunc) {
	f.router.POST(pattern, handler)
}
