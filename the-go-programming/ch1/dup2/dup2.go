package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	readers := makeReaders(os.Args)
	run(readers)
}

type readerFunc func() (io.Reader, func())

var emptyCloseFunc = func() {}

func makeReaders(args []string) []readerFunc {
	readerFuncs := make([]readerFunc, 0)
	if len(args) == 1 {
		readerFuncs = append(readerFuncs, func() (io.Reader, func()) {
			return os.Stdin, emptyCloseFunc
		})
		return readerFuncs
	}

	files := args[1:]

	for _, filename := range files {
		filename := filename
		readerFuncs = append(readerFuncs, func() (io.Reader, func()) {
			file, err := os.Open(filename)
			if err != nil {
				return bytes.NewBufferString(""), emptyCloseFunc
			}
			return file, func() {
				file.Close()
			}
		})
	}

	return readerFuncs
}

func run(readers []readerFunc) {
	lineMap := NewLineMap()
	for _, readerFunc := range readers {
		reader, close := readerFunc()
		lineMap.CountMap(reader)
		close()
	}
	fmt.Print(lineMap)
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

func (lineMap LineMap) String() string {
	buffer := bytes.NewBufferString("")
	for line, n := range lineMap {
		if n > 0 {
			buffer.WriteString(fmt.Sprintf("%d\t%s\n", n, line))
		}
	}
	return buffer.String()
}
