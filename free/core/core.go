package core

import (
	"net/http"
	"sync"
)

// 请求处理函数
type HandlerFunc func(*Context)

// 请求处理函数列表
type HandlersChain []HandlerFunc

// 处理数据组合
type Context struct{}

// 参数对象
type Param struct {
	Key   string
	Value string
}
type Params []Param

// 路由体对象
type Fhttp struct {
	pool  sync.Pool   //使用连接池处理并发
	trees methodTrees //路由树
}

func New() *Fhttp {
	f := &Fhttp{
		trees: make(methodTrees, 0, 9),
	}
	f.pool.New = func() interface{} {
		return nil
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
