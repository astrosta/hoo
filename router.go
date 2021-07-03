// @Title  router.go
// @Description  路由，提供路由注册方法
// @Author  astrosta 2021/07/03
// @Update  astrosta 2021/07/03

package hoo

import "fmt"

type router struct {
	handlers map[string]HandleFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandleFunc),
	}
}

func (r *router) addRoute(method, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(ctx *Context) {
	key := ctx.Method + "-" + ctx.Path
	if handler, ok := r.handlers[key]; ok {
		handler(ctx)
	} else {
		fmt.Fprintf(ctx.Writer, "404 NOT FOUND:%s", ctx.Req.URL)
	}
}
