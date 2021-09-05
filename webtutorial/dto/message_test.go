package dto

import (
	"fmt"
	"testing"
)

func TestMessage(t *testing.T) {
	todo := Todo{
		Author: "CY",
		Title:  "Buy Milk",
	}

	msg, _ := todo.Marshal()
	fmt.Println(string(msg))

	var empty Todo
	fmt.Println(empty)
	empty.Unmarshal([]byte(`{
		"author": "Huey Wen",
		"title": "Buy Choco"
	}`))
	fmt.Printf("%T, %v, %#v\n", empty, empty, empty)
}
