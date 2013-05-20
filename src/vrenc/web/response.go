package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	writer http.ResponseWriter
}

func (res *Response) Send(text string) {
	fmt.Fprintf(res.writer, text)
}

func (res *Response) Json(data interface{}) {
	encoder := json.NewEncoder(res.writer)

	encoder.Encode(data)
}

func (res *Response) NotFound() {
	res.writer.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(res.writer, "404 Not found")
}

func (res *Response) SetStatusCode(status int) {
	res.writer.WriteHeader(status)
}

func (res *Response) Headers() http.Header {
	return res.writer.Header()
}
