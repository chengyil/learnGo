package controller

import (
	"fmt"
	"net/http"
	"time"
)

type HelloController struct{}

func (c *HelloController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s\n", time.Now())
}
