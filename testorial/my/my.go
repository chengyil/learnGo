package my

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Concat() string {
	return "hello " + "world"
}

func Format() string {
	return fmt.Sprintf("hello %v", "world")
}
