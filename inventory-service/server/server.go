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
	routes := registerRoutes()
	ListenAndServe("0.0.0.0:3000", routes)
}

func registerRoutes() *Routes {
	routes := NewRoutes()
	routes.registerEndpoint("/products", handler.HandleProducts())
	return routes
}
