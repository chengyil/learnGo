package main

import (
	"bufio"
	"fmt"
	"io"
)

func main() {
	fmt.Println("Hello world")
}

type LineMap map[string]int

func NewLineMap() LineMap {
	return LineMap{}
}

func (count LineMap) CountMap(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		count[scanner.Text()]++
	}
}
