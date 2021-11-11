package gee

import (
	"net/http"
	"strings"
)

type router struct {
	root     map[string]*node
	handlers map[string]HandlerFunc
}

func NewRouter() *router {
	return &router{
		root:     make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	var parts []string
	for _, part := range vs {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern

	_, ok := r.root[method]
	if !ok {
		r.root[method] = &node{}
	}

	r.root[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRouter(method string, pattern string) (*node, map[string]string) {
	oriParts := parsePattern(pattern)
	params := make(map[string]string)

	_, ok := r.root[method]
	if !ok {
		return nil, nil
	}

	n := r.root[method].search(oriParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = oriParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(oriParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) Handler(c *Context) {
	n, params := r.getRouter(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + c.Path
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
