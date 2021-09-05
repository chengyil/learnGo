package rexp

import (
	"regexp"
	"strconv"
	"testing"
)

func TestGetPathInput(t *testing.T) {
	pathPattern := `/todo/(?P<todoID>\d+)/tasks/(?P<taskID>\d+)`
	re, _ := regexp.Compile(pathPattern)
	path := "/todo/10/tasks/100"

	scenarios := []struct {
		name    string
		pattern string
		result  int
	}{
		{
			name:    "todoID",
			pattern: "${todoID}",
			result:  10,
		},
		{
			name:    "taskID",
			pattern: "${taskID}",
			result:  100,
		},
	}

	for _, scenario := range scenarios {
		expected, err := strconv.Atoi(re.ReplaceAllString(path, scenario.pattern))
		if err != nil || expected != scenario.result {
			t.Errorf("should be %v but %v ", expected, scenario.result)
		}
	}
}
