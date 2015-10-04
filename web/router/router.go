package router

import (
	"net/http"
)

type (
	Handler func(http.ResponseWriter, *http.Request) error

	Mapping struct {
		Method, Pattern string
		Handler         Handler
	}

	Router struct {
		mappings []Mapping
	}
)

func (r *Router) HandlerFunc() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (r *Router) Register(method, pattern string, handler Handler) {
	m := Mapping{
		Method:  method,
		Pattern: pattern,
		Handler: handler,
	}
	r.mappings = append(r.mappings, m)
}
