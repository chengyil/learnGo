package main

import (
	"fmt"
	"math"
	"time"
)

type Node struct {
	node  *Node
	value string
}

func (n Node) isEmpty() bool {
	return n.value == ""
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it does not work",
	}
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func main() {
	root := Node{}
	fmt.Println("root is empty", root.isEmpty())
	root.value = "hello"
	fmt.Println("root is empty", root.isEmpty(), root.value)
	root.node = &Node{}
	fmt.Println("root", root.isEmpty(), root)
	fmt.Println("root", root.isEmpty(), root.node.isEmpty(), root.node)
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	err := run()
	if err != nil {
		fmt.Println(err)
	}

	if v, err := Sqrt(-2); err != nil {
		fmt.Println(v, err)
	}

}
