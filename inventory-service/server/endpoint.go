package server

import (
	"net/http"
)

var (
	Handle = http.Handle
)

type middleware func(http.Handler) http.Handler

func registerEndpoint(path string, handler http.Handler, middlewares ...middleware) {
	h := handler

	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}

	Handle(path, h)
}
