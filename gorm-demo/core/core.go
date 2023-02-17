package core

import (
	"net/http"
	"sync"
)

type Fhttp struct {
	pool  sync.Pool //使用连接池处理并发
	trees methodTrees
}

func New() *Fhttp {
	f := &Fhttp{
		trees: make(methodTrees, 0, 9),
	}
	f.pool.New = func() interface{} {
		return allocateContext()
	}
	return f
}

// 从连接池内取出请求示例后再放回去
func (f *Fhttp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := f.pool.Get().(*Context)
	f.HttpRequestHandle(c)
	f.pool.Put(c)
}

func (f *Fhttp) HttpRequestHandle(c *Context) {}

func (engine *Engine) allocateContext() *Context {
	v := make(Params, 0, engine.maxParams)
	return &Context{engine: engine, params: &v}
}
