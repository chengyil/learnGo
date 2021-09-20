package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestDup(t *testing.T) {
	reader := bytes.NewBufferString("Hello\nWorld")
	exp := dup(reader)

	if !reflect.DeepEqual(exp, map[string]int{
		"Hello": 1,
		"World": 1,
	}) {
		t.Fatalf("Dup Failed %v", exp)
	}

}
