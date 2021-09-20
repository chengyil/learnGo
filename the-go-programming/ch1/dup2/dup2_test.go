package main

import (
	"bytes"
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
