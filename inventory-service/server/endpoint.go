package server

import (
	"net/http"
)

var (
	Handle = http.Handle
)

type Middleware http.Handler

type Routes struct {
	middlewares []*Middleware
	routes      []*Route
}

type Route struct {
	path    string
	handler http.Handler
}

func (route *Route) Match(r *http.Request) bool {
	return route.path == r.URL.Path
}

func NewRoutes() *Routes {
	return &Routes{}
}

func (routes *Routes) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var route *Route

	for _, current := range routes.routes {
		if current.Match(r) {
			route = current
		}
	}

	if route == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for _, middleware := range routes.middlewares {
		m := *middleware
		m.ServeHTTP(w, r)
	}
}

func (r *Routes) registerEndpoint(path string, handler http.Handler) {
	route := &Route{
		path:    path,
		handler: handler,
	}

	r.routes = append(r.routes, route)
}
