package web

import (
	"fmt"
	"net/http"
)

type Handler struct {
	router *Router
}

func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request received", r.RequestURI)

	// Find matching route
	route := handler.router.MatchURI(r.Method, r.RequestURI)
	// Create response object
	res := new(Response)
	res.writer = w // set writer
	// Create request object
	req := new(Request)
	req.source = r // set original request

	// If a matching route is found
	if route != nil {
		// Invoke route callback
		req.params = route.params
		route.Callback(res, req)
	} else {
		// Return a 404 if no matching route was found
		res.NotFound()
	}

	// Request finished. On to the next one.
	return
}

func (handler *Handler) setRouter(router *Router) {
	handler.router = router
}
