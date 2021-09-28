package routes

import "net/http"

type Route struct {
	path    string
	handler http.Handler
}

type routeRegexp struct {
}

func makeRouteRegexp(path string) *routeRegexp {
	return &routeRegexp{}
}

func findBraceIndices(s string) ([]int, error) {
	idx := make([]int, 0)

	return idx, nil

}

func (route *Route) Match(r *http.Request) bool {
	return route.path == r.URL.Path
}
