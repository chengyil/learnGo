package main

import (
	"fmt"
	"net/http"

	"github.com/chengyil/hello-world/controller"
)

func main() {
	fmt.Println("Start Server")
	startServer()
	fmt.Println("Closing Server")
}

func startServer() {
	c := &controller.HelloController{}
	http.Handle("/", c)
	http.ListenAndServe(":8080", nil)
}
