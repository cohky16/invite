package main

import "testing"

type Regexp struct {
	arg    string
	match  string
	result bool
}

func TestCheckRegexp(t *testing.T) {
	tests := []Regexp{
		{"talk", "talk-request", true},
		{"meeting", "meeting-request", true},
		{"Talk", "talk-request", false},
		{"talk", "Talk-request", false},
		{"test-talk", "talk-request", false},
		{"meeting", "meeting", true},
		{"meeting", "meetin", false},
	}

	for _, test := range tests {
		if checkRegexp(test.arg, test.match) != test.result {
			t.Error("not Equal")
		}
	}
}
