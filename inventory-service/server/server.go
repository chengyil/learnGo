package server

import (
	"inventory-service/routes"
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

func registerRoutes() *routes.Routes {
	myRoutes := routes.NewRoutes()
	myRoutes.RegisterEndpoint("/products", handler.HandleProducts())
	return myRoutes
}
