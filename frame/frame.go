package frame

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	route *router
}

func New() *Engine {
	return &Engine{route: newRoute()}
}

func (e *Engine) addRoutes(method string, pattern string, handler HandlerFunc)  {
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
