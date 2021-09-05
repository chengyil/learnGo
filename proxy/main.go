package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/textproto"
	"net/url"
)

var (
	upStream = "http://0.0.0.0:3001/proxied/"
)

type Proxy struct {
	proxy *httputil.ReverseProxy
}

func NewProxy() Proxy {
	url, err := url.Parse(upStream)

	if err != nil {
		panic("invalid upstream")
	}

	return Proxy{
		proxy: httputil.NewSingleHostReverseProxy(url),
	}
}

func (p Proxy) handle(w http.ResponseWriter, r *http.Request) {
	// Headers allow multiple Values
	// Add append new values to the Header Key
	// Add append new values to the Header Key
	r.Header.Add("Header1", "bar")
	r.Header.Add("Header1", "world")
	// Set replace value to the header Key
	r.Header.Set("Header2", "foo")
	p.proxy.ServeHTTP(w, r)
}

func main() {
	p := NewProxy()
	http.HandleFunc("/", p.handle)
	http.HandleFunc("/proxied/", proxied)
	fmt.Println(http.ListenAndServe("0.0.0.0:3001", nil))
}

func proxied(w http.ResponseWriter, r *http.Request) {
	fmt.Println("header", r.Header)
	fmt.Println("header Get", r.Header.Get("Header1"))
	fmt.Println("header Access", r.Header["Header1"])
	fmt.Println("content-type", r.Header["content-type"], "content-type")
	fmt.Println("content-type parsed", r.Header[textproto.CanonicalMIMEHeaderKey("content-type")], textproto.CanonicalMIMEHeaderKey("content-type"))
	message := []byte("world")
	w.Write(message)
}
