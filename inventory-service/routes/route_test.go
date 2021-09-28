package routes

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestRouteMatch(t *testing.T) {
	scenarios := []struct {
		name     string
		route    *Route
		r        *http.Request
		expMatch bool
	}{
		{
			name: "Should Match Static Path",
			route: &Route{
				path: "/hello",
				handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				}),
			},
			r: &http.Request{
				URL: &url.URL{
					Path: "/hello",
				},
			},
			expMatch: true,
		},
		{
			name: "Should Match Route with Path Variable",
			route: &Route{
				path: "/hello/{name}",
				handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				}),
			},
			r: &http.Request{
				URL: &url.URL{
					Path: "/hello",
				},
			},
			expMatch: true,
		},
	}

	for _, scenario := range scenarios {
		if isMatch := scenario.route.Match(scenario.r); isMatch != scenario.expMatch {
			t.Fatalf("%v Failed, expected %v, but got %v", scenario.name, scenario.expMatch, isMatch)
		}
	}
}

func TestMakeRouteRegexp(t *testing.T) {
	routeExp := makeRouteRegexp("/helloworld")

	if routeExp == nil {
		t.Fatalf("Failed to build route regexp")
	}
}

func TestFindBraceIndices(t *testing.T) {
	scenarios := []struct {
		name   string
		path   string
		expIdx []int
		expErr bool
	}{
		{
			name:   "Empty Path",
			path:   "/",
			expIdx: make([]int, 0),
		},
		{
			name:   "Has 1 Bucket",
			path:   "/{ID}",
			expIdx: make([]int, 1, 4),
		},
	}
	for _, scenario := range scenarios {
		idx, err := findBraceIndices(scenario.path)
		if !scenario.expErr && err != nil {
			t.Fatalf("%v: Not expected error but %v", scenario.name, err)
		} else if scenario.expErr && err == nil {
			t.Fatalf("%v:expected error but no error", scenario.name)
		}
		if !Equal(idx, scenario.expIdx) {
			t.Fatalf("%v: expected %v, but got %v", scenario.name, scenario.expIdx, idx)
		}

	}
}

func Equal(a, b interface{}) bool {
	aString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), ","), "[]")
	bString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(b)), ","), "[]")
	fmt.Println(aString, bString)
	return aString == bString
}

// func TestRegisterEndpoint(t *testing.T) {
// 	defer testingtool.Restore(testingtool.Pairs(&Handle))
// 	path := ""
// 	var handler http.Handler
// 	Handle = func(p string, h http.Handler) {
// 		path = p
// 		handler = h
// 	}
// 	trace := []string{}
// 	endpoint := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		trace = append(trace, "in endpoint")
// 		trace = append(trace, "out endpoint")
// 	})

// 	scenarios := []struct {
// 		name        string
// 		trace       string
// 		path        string
// 		middlewares []middleware
// 	}{
// 		{
// 			name:  "No Middleware",
// 			trace: "in endpoint,out endpoint",
// 			path:  "/path",
// 		},
// 		{
// 			name:  "One Middleware",
// 			trace: "in one,in endpoint,out endpoint,out one",
// 			path:  "/path",
// 			middlewares: []middleware{
// 				func(h http.Handler) http.Handler {
// 					return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 						trace = append(trace, "in one")
// 						h.ServeHTTP(w, r)
// 						trace = append(trace, "out one")
// 					})
// 				},
// 			},
// 		},
// 		{
// 			name:  "Three Middleware",
// 			trace: "in one,in two,in three,in endpoint,out endpoint,out three,out two,out one",
// 			path:  "/path",
// 			middlewares: []middleware{
// 				func(h http.Handler) http.Handler {
// 					return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 						trace = append(trace, "in one")
// 						h.ServeHTTP(w, r)
// 						trace = append(trace, "out one")
// 					})
// 				},
// 				func(h http.Handler) http.Handler {
// 					return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 						trace = append(trace, "in two")
// 						h.ServeHTTP(w, r)
// 						trace = append(trace, "out two")
// 					})
// 				},
// 				func(h http.Handler) http.Handler {
// 					return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 						trace = append(trace, "in three")
// 						h.ServeHTTP(w, r)
// 						trace = append(trace, "out three")
// 					})
// 				},
// 			},
// 		},
// 	}

// 	for _, scenario := range scenarios {
// 		registerEndpoint(scenario.path, endpoint, scenario.middlewares...)
// 		handler.ServeHTTP(nil, nil)
// 		if path != scenario.path {
// 			t.Errorf("Scenario:%v\nExpecting path to be %v but %v\n", scenario.name, scenario.path, path)
// 		}
// 		if expected := strings.Join(trace, ","); expected != scenario.trace {
// 			t.Errorf("Scenario:%v\nExpecting trace to be %v but %v\n", scenario.name, scenario.trace, expected)
// 		}
// 		trace = make([]string, 0)
// 	}
// }
