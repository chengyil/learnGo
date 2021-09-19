package server

import (
	"inventory-service/server/handler"
	"inventory-service/storage"
	"net/http"
)

var (
	ListenAndServe = http.ListenAndServe
)

func Serve() {
	storage.Init()
	serve()
}

func serve() {
	registerEndpoints()
	ListenAndServe("0.0.0.0:3000", nil)
}

func registerEndpoints() {
	registerEndpoint("/products", handler.HandleProducts())
}
