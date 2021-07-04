package hoo

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//HandleFunc defines the request handler
type HandleFunc func(ctx *Context)

func handleFuncConvert(handleFunc HandleFunc) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		ctx := newContext(writer, request)
		for _, value := range params {
			ctx.Params[value.Key] = value.Value
		}
		handleFunc(ctx)
	}
}

//Engine implement the interface of ServerHttp
type Engine struct {
	router *httprouter.Router
}

func New() *Engine {
	return &Engine{
		router: httprouter.New(),
	}
}

func (engine *Engine) GET(path string, handler HandleFunc) {
	engine.router.GET(path, handleFuncConvert(handler))
}

func (engine *Engine) POST(path string, handler HandleFunc) {
	engine.router.POST(path, handleFuncConvert(handler))
}

func (engine *Engine) Run(addr string) {
	http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	engine.router.ServeHTTP(w, r)
}
