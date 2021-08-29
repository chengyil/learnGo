package main

import (
	"fmt"
	"math"
)

var packlevel, variable, yousee = true, false, "love"
var hello = 1

func main() {
	test()
	moreSlice()
	obj()
	tour3()
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func swap(a, b string) (string, string) {
	return b, a
}

func add(a, b int) int {
	return a + b
}

func tour3() {
	var (
		ToBe   bool = true
		MaxInt uint = 1<<64 - 1
	)
	fmt.Println(math.Pi)
	sum := func(a int, b int) int { return a + b }(1, 2)
	a, b := swap("a", "b")
	x, y := split(23)
	fmt.Println(sum, add(2, 2), a, b, x, y, packlevel, variable, yousee, hello, ToBe, MaxInt)
	var a1 int8 = 1
	var a2 int16 = 1
	var a3 int32 = 1
	var a4 int64 = 1
	var a5 int8 = 1

	fmt.Printf("Type: %T Value: %v %p\n", a1, a1, &a1)
	fmt.Printf("Type: %T Value: %v %p\n", a2, a2, &a2)
	fmt.Printf("Type: %T Value: %v %p\n", a3, a3, &a3)
	fmt.Printf("Type: %T Value: %v %p\n", a4, a4, &a4)
	fmt.Printf("Type: %T Value: %v %p\n", a5, a5, &a5)

	const Pi int = 1
	const pi = 3.14
	const (
		HELLO = iota
		WORLD
		TOO
	)
	fmt.Println("Pi", Pi, pi, HELLO, WORLD, TOO)

}

func test() {
	var i int = 1
	j := 2
	var p = &i
	fmt.Println("Hello World", i, j, p, &i)
	arr := [3]int{1, 2, 3}
	arrp := &arr[0]
	*arrp = 10
	slice := arr[:]
	slice[2] = 5
	slice = append(slice, 100)
	fmt.Println(arr, &arr[0], &arr[1], slice)
	slice[2] = 15
	fmt.Println(arr, &arr[0], &arr[1], slice)
}

func moreSlice() {
	slice := [3]int{1, 2, 3}
	newSlice := slice[:]
	moreSlice := append(newSlice, 1)
	fmt.Println(&slice[0], &moreSlice[0], &newSlice[0])
}

func obj() {
	var a []int
	b := []int{1, 2}
	// var c map[string]int
	// c["a"] = 1
	fmt.Println(a, b)
}
