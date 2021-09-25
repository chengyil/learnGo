package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("need url")
		os.Exit(-1)
	}
	url := os.Args[1]
	run(url)
}

func run(url string) {
	res, _ := curl(url, http.Get)
	fmt.Println(res)
}

type GetFunc func(string) (*http.Response, error)

func curl(url string, get GetFunc) (string, error) {
	res, err := get(url)
	if err != nil {
		return "", err
	}

	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(out), nil
}
