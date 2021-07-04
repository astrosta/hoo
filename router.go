// @Title  router.go
// @Description  路由，提供路由注册方法
// @Author  astrosta 2021/07/03
// @Update  astrosta 2021/07/03

package hoo

import (
	"github.com/julienschmidt/httprouter"
)

type router struct {
	//handlers map[string]HandleFunc
	router *httprouter.Router
}

func newRouter() *router {
	return &router{
		router: httprouter.New(),
	}
}

//func (r *router) addRoute(method, pattern string, handler HandleFunc) {
//	key := method + "-" + pattern
//	r.handlers[key] = handler
//}

//func (r *router) handle(ctx *Context) {
//	key := ctx.Method + "-" + ctx.Path
//	if handler, ok := r.handlers[key]; ok {
//		handler(ctx)
//	} else {
//		fmt.Fprintf(ctx.Writer, "404 NOT FOUND:%s", ctx.Req.URL)
//	}
//}
