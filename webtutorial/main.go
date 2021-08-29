package main

import (
	"fmt"
	"net/http"
)

func main() {
	startServer()
}

func startServer() {
	registerEndpoint()
	err := http.ListenAndServe("0.0.0.0:3000", nil)
	if err == nil {
		fmt.Printf("Error while serving %v", err)
	}
}

type worldHandler struct{}

func (t *worldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func registerEndpoint() {
	http.Handle("/world", new(worldHandler))
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		message := []byte("world")
		w.Write(message)
	})
}
