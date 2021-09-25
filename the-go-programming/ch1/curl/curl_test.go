package main

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestCurl(t *testing.T) {
	getter := func(url string) (*http.Response, error) {
		body := io.NopCloser(strings.NewReader(url))

		return &http.Response{
			Body: body,
		}, nil
	}
	url := "https://api.chucknorris.io/jokes/random"
	resp, err := curl(url, getter)
	if resp != url {
		t.Fatalf("Invalid Data %v", resp)
	}
	if err != nil {
		t.Fatalf("Should be nil")
	}
}
