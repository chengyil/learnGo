package server

import (
	"inventory-service/testingtool"
	"net/http"
	"strings"
	"testing"
)

func TestRegisterEndpoint(t *testing.T) {
	defer testingtool.Restore(testingtool.Pairs(&Handle))
	path := ""
	var handler http.Handler
	Handle = func(p string, h http.Handler) {
		path = p
		handler = h
	}
	trace := []string{}
	endpoint := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		trace = append(trace, "in endpoint")
		trace = append(trace, "out endpoint")
	})

	scenarios := []struct {
		name        string
		trace       string
		path        string
		middlewares []middleware
	}{
		{
			name:  "No Middleware",
			trace: "in endpoint,out endpoint",
			path:  "/path",
		},
		{
			name:  "One Middleware",
			trace: "in one,in endpoint,out endpoint,out one",
			path:  "/path",
			middlewares: []middleware{
				func(h http.Handler) http.Handler {
					return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						trace = append(trace, "in one")
						h.ServeHTTP(w, r)
						trace = append(trace, "out one")
					})
				},
			},
		},
		{
			name:  "Three Middleware",
			trace: "in one,in two,in three,in endpoint,out endpoint,out three,out two,out one",
			path:  "/path",
			middlewares: []middleware{
				func(h http.Handler) http.Handler {
					return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						trace = append(trace, "in one")
						h.ServeHTTP(w, r)
						trace = append(trace, "out one")
					})
				},
				func(h http.Handler) http.Handler {
					return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						trace = append(trace, "in two")
						h.ServeHTTP(w, r)
						trace = append(trace, "out two")
					})
				},
				func(h http.Handler) http.Handler {
					return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						trace = append(trace, "in three")
						h.ServeHTTP(w, r)
						trace = append(trace, "out three")
					})
				},
			},
		},
	}

	for _, scenario := range scenarios {
		registerEndpoint(scenario.path, endpoint, scenario.middlewares...)
		handler.ServeHTTP(nil, nil)
		if path != scenario.path {
			t.Errorf("Scenario:%v\nExpecting path to be %v but %v\n", scenario.name, scenario.path, path)
		}
		if expected := strings.Join(trace, ","); expected != scenario.trace {
			t.Errorf("Scenario:%v\nExpecting trace to be %v but %v\n", scenario.name, scenario.trace, expected)
		}
		trace = make([]string, 0)
	}
}
