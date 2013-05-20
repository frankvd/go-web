package web

import (
	"net/http"
)

type Request struct {
	source *http.Request
	params map[string]string
}

func (r *Request) Param(key string) string {
	return r.params[key]
}
