package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	get = http.Get
)

func main() {
	args := parseArgs(os.Args)

	if len(args) == 0 {
		fmt.Println("Needs at least 1 url")
		os.Exit(1)
	}

	all := fetchAll(args)
	for i := 0; i < len(args); i++ {
		fmt.Println(<-all)
	}
	close(all)
}

func parseArgs(args []string) []string {
	return args[1:]
}

func fetchAll(urls []string) chan string {
	buf := make(chan string)
	for _, url := range urls {
		go func(url string, buf chan string) {
			resp, _ := fetch(url)
			buf <- resp
		}(url, buf)
	}
	return buf
}

func fetch(url string) (string, error) {
	resp, err := get(url)

	if err != nil {
		return "", err
	}

	return readAll(resp.Body)
}

func readAll(readerCloser io.ReadCloser) (string, error) {
	buf := make([]byte, 0)
	for {
		if len(buf) == cap(buf) {
			buf = append(buf, 0)[:len(buf)]
		}
		in, err := readerCloser.Read(buf[len(buf):cap(buf)])
		buf = buf[:len(buf)+in]
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return string(buf), nil
		}
	}
}
