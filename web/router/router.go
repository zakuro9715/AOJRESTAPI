package router

import (
	"net/http"
)

type Router struct {
}

func (r *Router) HandlerFunc() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}
