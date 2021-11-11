package main

import (
	"gee/gee"
	"net/http"
)

func main() {
	g := gee.New()
	g.Get("http", func(w http.ResponseWriter, r *http.Request) {

	})
}
