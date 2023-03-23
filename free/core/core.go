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
	Router
	pool sync.Pool //使用连接池处理并发
}

func New() *Fhttp {
	f := &Fhttp{
		Router: Router{
			handlers: make(map[string]HandlerFunc),
		},
	}
	return f
}

// 实现ServerHttp接口
// 从连接池内取出请求示例后再放回去
func (f *Fhttp) ServeHTTP(c *Context) {
	key := c.Path
	if h, ok := f.handlers[key]; ok {
		h(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
