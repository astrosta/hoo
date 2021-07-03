package hoo

import (
	"fmt"
	"net/http"
)

//HandleFunc defines the request handler
type HandleFunc func(http.ResponseWriter, *http.Request)

//RouterEngine implement the interface of ServerHttp
type RouterEngine struct {
	router map[string]HandleFunc
}

func New() *RouterEngine {
	return &RouterEngine{
		router: make(map[string]HandleFunc),
	}
}

func (engine *RouterEngine) addRoute(method, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

func (engine *RouterEngine) Get(pattern string, handler HandleFunc) {
	engine.addRoute("Get", pattern, handler)
}

func (engine *RouterEngine) Post(pattern string, handler HandleFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *RouterEngine) Run(addr string) {
	http.ListenAndServe(addr, engine)
}

func (engine *RouterEngine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND:%s", r.URL)
	}
}
