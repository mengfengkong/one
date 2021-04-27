package frame

import (
	"net/http"
	"strings"
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

func (group *RouterGroup) Use(middwares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middwares...)
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

//
//func (e *Engine) Use(middleware ...HandlerFunc) {
//	e.middlewares = append(e.middlewares, middleware...)
//}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, r)
	c.handlers = middlewares
	c.engine = e
	e.route.handle(c)
}

func Default() *Engine {
	f := New()
	f.Use(Logger(), Recovery())
	return f
}
