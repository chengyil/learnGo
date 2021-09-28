package routes

import (
	"fmt"
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
		fmt.Println(r.URL)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for _, middleware := range routes.middlewares {
		m := *middleware
		m.ServeHTTP(w, r)
	}

	route.ServerHTTP(w, r)
}

func (r *Routes) RegisterEndpoint(path string, handler http.Handler) {
	route := &Route{
		path:    path,
		handler: handler,
	}
	r.routes = append(r.routes, route)
}

func (route *Route) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	route.handler.ServeHTTP(w, r)
}
