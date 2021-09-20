package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	data := dup(os.Stdin)
	data.format(data)
}

type WordMap map[string]int

func dup(reader io.Reader) WordMap {
	scanner := bufio.NewScanner(reader)
	data := WordMap{}
	for scanner.Scan() {
		data[scanner.Text()] += 1
	}
	return data
}

func (w WordMap) format(data WordMap) {
	for line, n := range data {
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
