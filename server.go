package web

import (
	"net/http"
	"strconv"
)

type Server struct {
	router  *Router
	handler *Handler
}

func (server *Server) Run(port int) {
	http.ListenAndServe(":"+strconv.Itoa(port), server.handler)
}

func (server *Server) Router() *Router {
	return server.router
}

func CreateServer() *Server {
	server := new(Server)
	server.router = new(Router)
	server.handler = new(Handler)
	server.handler.setRouter(server.Router())

	return server
}
