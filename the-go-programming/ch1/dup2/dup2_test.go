package main

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestCountMap(t *testing.T) {
	scenarios := []struct {
		name      string
		argInputs []string
		expMap    LineMap
	}{
		{
			name: "One Buffered Input",
			argInputs: []string{
				"Hello\nWorld",
			},
			expMap: LineMap{
				"World": 1,
				"Hello": 1,
			},
		},
		{
			name: "More than one Buffered Input",
			argInputs: []string{
				"Hello\nWorld",
				"Hello\nSingapore",
			},
			expMap: LineMap{
				"World":     1,
				"Hello":     2,
				"Singapore": 1,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario := scenario
		t.Run(scenario.name, func(t *testing.T) {
			count := NewLineMap()

			for _, input := range scenario.argInputs {
				buf := bytes.NewBufferString(input)
				count.CountMap(buf)
			}

			if !reflect.DeepEqual(count, scenario.expMap) {
				t.Fatalf("Invalid Count Map %v", count)
			}
		})
	}

}

func TestLineMapString(t *testing.T) {
	scenarios := []struct {
		name     string
		argInput io.Reader
		exp      string
	}{
		{
			name:     "One Buffered Input",
			argInput: bytes.NewBufferString("Hello"),
			exp:      "1\tHello\n",
		},
	}

	for _, scenario := range scenarios {
		scenario := scenario
		t.Run(scenario.name, func(t *testing.T) {
			lineMap := NewLineMap()
			lineMap.CountMap(scenario.argInput)
			toString := lineMap.String()

			if toString != scenario.exp {
				t.Fatalf("Invalid String format %v", toString)
			}
		})
	}

}

func TestMakeReaders(t *testing.T) {
	scenarios := []struct {
		name   string
		args   []string
		expLen int
	}{
		{
			name:   "Only name, read from STDIN",
			args:   []string{"./dup"},
			expLen: 1,
		},
		{
			name:   "Name with 1 Input",
			args:   []string{"./dup", "input"},
			expLen: 1,
		},
	}
	for _, scenario := range scenarios {
		readers := makeReaders(scenario.args)
		if len(readers) != scenario.expLen {
			t.Fatalf("Expected %v, but got %v", 1, len(readers))
		}
	}
}
