package hoo

type RouterGroup struct {
	prefix          string
	middlewareChain []middleware
	parent          *RouterGroup
	engine          *Engine
}

func (rg *RouterGroup) Group(prefix string) *RouterGroup {
	engine := rg.engine
	newGroup := &RouterGroup{
		prefix: rg.prefix + prefix,
		parent: rg,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (rg *RouterGroup) GET(path string, handler HandleFunc) {
	p := rg.prefix + path
	rg.engine.GET(p, handler)
}

func (rg *RouterGroup) POST(path string, handler HandleFunc) {
	p := rg.prefix + path
	rg.engine.POST(p, handler)
}

func (rg *RouterGroup) Use(m middleware) {
	rg.middlewareChain = append(rg.middlewareChain, m)
}
