package gee

import (
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: NewRouter()}
}

func (engine *Engine) addRouter(method string, pattern string, handler HandlerFunc) {
	engine.router.addRouter(method, pattern, handler)
}

func (engine *Engine) Get(pattern string, handler HandlerFunc) {
	engine.addRouter("Get", pattern, handler)
}

func (engine *Engine) Post(pattern string, handler HandlerFunc) {
	engine.addRouter("Post", pattern, handler)
}

func (engine *Engine) Run(add string) (err error) {
	return http.ListenAndServe(add, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	engine.router.Handler(NewContext(w, req))
}
