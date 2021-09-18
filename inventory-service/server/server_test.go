package server

import (
	"inventory-service/testingtool"
	"net/http"
	"testing"
)

func TestServe(t *testing.T) {
	defer testingtool.Restore(testingtool.Pairs(&ListenAndServe))

	callTimes := 0

	ListenAndServe = func(s string, h http.Handler) error {
		callTimes += 1
		return nil
	}

	Serve()

	if callTimes != 1 {
		t.Errorf("Should have called once, but instead call times is %v", callTimes)
	}
}
