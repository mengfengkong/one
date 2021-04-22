package frame

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type RouterGroup struct {
	middlewares []HandlerFunc
	parent      *RouterGroup
	prefix      string
	engine      *Engine
}

type Engine struct {
	route *router
	*RouterGroup
	groups []*RouterGroup
}

func New() *Engine {
	engine := &Engine{route: newRoute()}
	engine.RouterGroup = &RouterGroup{
		engine: engine,
	}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		parent: group,
		prefix: group.prefix + prefix,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)

	return newGroup
}

func (group *RouterGroup) addRoutes(method string, pattern string, handler HandlerFunc) {
	completePattern := group.prefix + pattern
	group.engine.route.addRoute(method, completePattern, handler)
}

func (group *RouterGroup) Get(pattern string, handler HandlerFunc) {
	group.addRoutes("GET", pattern, handler)
}

func (group *RouterGroup) Post(pattern string, handler HandlerFunc) {
	group.addRoutes("POST", pattern, handler)
}

func (e *Engine) addRoutes(method string, pattern string, handler HandlerFunc) {
	e.route.addRoute(method, pattern, handler)
}

func (e *Engine) Get(pattern string, handler HandlerFunc) {
	e.addRoutes("GET", pattern, handler)
}

func (e *Engine) Post(pattern string, handler HandlerFunc) {
	e.addRoutes("POST", pattern, handler)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	e.route.handle(c)
}
