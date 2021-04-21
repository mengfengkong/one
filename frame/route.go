package frame

import "net/http"

type router struct {
	handlers map[string]HandlerFunc
}

func newRoute() *router {
	return &router{handlers: map[string]HandlerFunc{}}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Req.Method + "-" + c.Req.URL.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 Not Found:%s\n", c.Path)
	}
}
