package gee

import "net/http"

type router struct {
	handler map[string]HandlerFunc
}

func NewRouter() *router {
	return &router{handler: make(map[string]HandlerFunc)}
}

func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	r.handler[key] = handler
}

func (r *router) Handler(c *Context) {
	key := c.Req.Method + "-" + c.Req.URL.Path
	if hander, ok := r.handler[key]; ok {
		hander(c.Writer, c.Req)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
