package hoo

import (
	"net/http"
)

//HandleFunc defines the request handler
type HandleFunc func(ctx *Context)

//Engine implement the interface of ServerHttp
type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (engine *Engine) GET(pattern string, handler HandleFunc) {
	engine.router.addRoute(http.MethodGet, pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandleFunc) {
	engine.router.addRoute(http.MethodPost, pattern, handler)
}

func (engine *Engine) Run(addr string) {
	http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(w, r)
	engine.router.handle(ctx)
}
