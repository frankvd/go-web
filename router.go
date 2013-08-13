package web

import (
	"fmt"
)

type Router struct {
	routes []*Route
}

func (r *Router) Get(pattern string, callback func(res *Response, req *Request)) {
	r.Register("GET", pattern, callback)
}

func (r *Router) Register(method string, pattern string, callback func(res *Response, req *Request)) {
	fmt.Println("route registered")
	route := new(Route)
	route.method = method
	route.pattern = pattern
	route.Callback = callback

	r.routes = append(r.routes, route)
}

func (r *Router) MatchURI(method string, uri string) *Route {
	for _, route := range r.routes {
		if route.matches(uri) {
			return route
		}
	}

	return nil
}
