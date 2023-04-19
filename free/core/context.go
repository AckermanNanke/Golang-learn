package core

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 返回数据体
type FJSON map[string]interface{}

type Context struct {
	R          *http.Request
	W          http.ResponseWriter
	Path       string
	Method     string
	Params     map[string]string
	StatusCode int
	handlers   HandlersChain
	index      int
}

// Context创建
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		R:      r,
		W:      w,
		Path:   r.URL.Path,
		Method: r.Method,
		index:  -1,
	}
}

// 执行当前顺序（index）中间件
func (c *Context) Next() {
	c.index++
	for ; c.index < len(c.handlers); c.index++ {
		c.handlers[c.index](c)
	}
}

// 获取访问参数
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// 获取 POST 请求的参数
func (c *Context) PostForm(key string) string {
	return c.R.FormValue(key)
}

// 获取 GET 请求 URL 的参数
func (c *Context) Query(key string) string {
	return c.R.URL.Query().Get(key)
}

// 设置返回体的状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.W.WriteHeader(code)
}

// 设置返回体的 Header
func (c *Context) SetHeader(key string, value string) {
	c.W.Header().Set(key, value)
}

// 设置返回体的 body 为JSON格式
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.W)
	// 回码一经赋值将不会再更改
	if err := encoder.Encode(obj); err != nil {
		panic(err)
	}
}

// 设置返回体的 body 为HTML格式
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.W.Write([]byte(html))
}

// 设置返回体的 body 为字节格式
func (c *Context) ByteData(code int, data []byte) {
	c.Status(code)
	c.W.Write(data)
}

// 设置返回体的 body 为格式化后的字符串
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.W.Write([]byte(fmt.Sprintf(format, values)))
}
