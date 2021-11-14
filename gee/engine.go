package gee

import (
	"net/http"
	"strings"
)

type HandlerFunc func(c *Context)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

func New() *Engine {
	engine := &Engine{router: NewRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.RouterGroup.middlewares = append(engine.RouterGroup.middlewares, Myrecover())
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) Group(pattern string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + pattern,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRouter(method string, path string, handler HandlerFunc) {
	pattern := group.prefix + path
	group.engine.router.addRouter(method, pattern, handler)
}

func (group *RouterGroup) Get(pattern string, handler HandlerFunc) {
	group.addRouter("Get", pattern, handler)
}

func (group *RouterGroup) Post(pattern string, handler HandlerFunc) {
	group.addRouter("Post", pattern, handler)
}

func (group *RouterGroup) Use(middleware ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middleware...)
}

func (engine *Engine) Run(add string) (err error) {
	return http.ListenAndServe(add, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	var middleware []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middleware = append(middleware, group.middlewares...)
		}
	}
	c.handler = middleware
	engine.router.Handler(c)
}
