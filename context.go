// @Title  context.go
// @Description  上下文(Context)，封装 Request 和 Response ，提供对 JSON、HTML 等返回类型的支持
// @Author  astrosta 2021/07/03
// @Update  astrosta 2021/07/03

package hoo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	StatusCode int //response info
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

func (c *Context) PostForm(key string) string {
	//fmt.Println(c.Req.Form)
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) GetRawData() ([]byte, error) {
	return ioutil.ReadAll(c.Req.Body)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain; charset=utf-8")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json; charset=utf-8")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		panic(err)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html; charset=utf-8")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
