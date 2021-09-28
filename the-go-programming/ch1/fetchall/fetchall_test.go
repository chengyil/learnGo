package main

import (
	"errors"
	"io/ioutil"
	"learn/testingtool"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseArgs(t *testing.T) {
	scenarios := []struct {
		name    string
		args    []string
		expArgs []string
	}{
		{
			name: "Empty Args",
			args: []string{
				"./app",
			},
			expArgs: []string{},
		},
		{
			name: "Has 1 arg",
			args: []string{
				"./app",
				"google.com",
			},
			expArgs: []string{"google.com"},
		},
		{
			name: "Has 2 arg",
			args: []string{
				"./app",
				"google.com",
				"yahoo.com",
			},
			expArgs: []string{"google.com", "yahoo.com"},
		},
	}

	for _, scenario := range scenarios {
		scenario := scenario
		t.Run(scenario.name, func(t *testing.T) {
			expArgs := parseArgs(scenario.args)
			assert.Equal(t, expArgs, scenario.expArgs, scenario.name)
		})
	}
}

func TestFetch(t *testing.T) {
	defer testingtool.Restore(testingtool.Pairs(&get))

	scenarios := []struct {
		name        string
		argUrl      string
		expResponse string
		expError    error
	}{
		{
			name:     "Get Error -> error",
			argUrl:   "google.com",
			expError: errors.New("500"),
		},
		{
			name:        "Happy Path",
			argUrl:      "google.com",
			expResponse: "This Is From Google",
		},
	}

	for _, scenario := range scenarios {
		scenario := scenario
		t.Run(scenario.name, func(t *testing.T) {
			get = func(url string) (resp *http.Response, err error) {
				return &http.Response{
					Body: ioutil.NopCloser(strings.NewReader(scenario.expResponse)),
				}, scenario.expError
			}

			response, err := fetch(scenario.argUrl)
			assert.Equal(t, response, scenario.expResponse, scenario.name)
			assert.Equal(t, err, scenario.expError, scenario.name)
		})
	}
}

func TestReadAll(t *testing.T) {
	scenarios := []struct {
		name     string
		argData  string
		exp      string
		hasError bool
	}{
		{
			name:    "Happy Path",
			argData: "Hello World",
			exp:     "Hello World",
		},
		{
			name:    "Happy Path",
			argData: "Hello World Hello World Hello World Hello World Hello World Hello World Hello World",
			exp:     "Hello World Hello World Hello World Hello World Hello World Hello World Hello World",
		},
	}

	for _, scenario := range scenarios {
		scenario := scenario
		t.Run(scenario.name, func(t *testing.T) {
			readerClose := ioutil.NopCloser(strings.NewReader(scenario.argData))
			exp, err := readAll(readerClose)

			assert.Equal(t, exp, scenario.exp, scenario.name)
			if scenario.hasError {
				assert.NotNil(t, err, scenario.name)
			} else {
				assert.Nil(t, err, scenario.name)
			}
		})
	}
}
