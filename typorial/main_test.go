package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	data := bytes.NewBufferString("Hello World!")
	buf := make([]byte, 0, 1)

	for {
		if len(buf) == cap(buf) {
			buf = append(buf, 1)
		}

		length, err := data.Read(buf[len(buf):])
		if err != nil {
			break
		}
		fmt.Printf("%s, length : %v, len %v, cap %v\n", string(buf), length, len(buf), cap(buf))
	}

	fmt.Println(string(buf))

}

func TestInt(t *testing.T) {
	var n int
	display(n)
	var n1 *int
	display(n1)
	n1 = new(int)
	*n1 = 100
	display(n1)
	display(*n1)
}

func TestStr(t *testing.T) {
	var s string
	display(s)
	s1 := new(string)
	display(s1)
	display(*s1)
	var s2 *string
	display(s2)
}

func TestBool(t *testing.T) {
	var b bool
	display(b)
	var b1 *bool
	display(b1)
	b2 := new(bool)
	display(b2)
	display(*b2)
}

type myError struct{}

func (my *myError) Error() string {
	return "ERROR"
}

type myString string

func (m *myString) String() string {
	if m == nil {
		return "myString"
	}
	var i interface{}
	i = m
	return i.(string)
}

func TestStruct(t *testing.T) {
	var my *myError
	var e error
	e = my
	display(e)

	var ms myString = "HELLO"
	display(ms)

	var ms1 *myString
	display(ms1)
}

func display(i interface{}) {
	fmt.Printf("%T, %#v %v\n", i, i, i)
}
