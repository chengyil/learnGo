package main

import (
	"fmt"
	"math"
	"runtime"
)

func myos() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Println("%S. \n", os)
	}
}

func mysqrt(x float64) float64 {
	var z float64 = 1.0
	for i := 0; i < 100; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println("z is ", z)
	}
	return z
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt((-4)))
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
	fmt.Println("mysqrt", mysqrt(50))
	myos()

}
