package testingtool

import (
	"fmt"
	"testing"
)

var mockFn = func() {
	fmt.Println("Original Test")
}

func TestMock(t *testing.T) {
	fn := Pairs(&mockFn)
	mockFn = func() {
		fmt.Println("Mocked Test")
	}
	mockFn()
	Restore(fn)
	mockFn()
}
