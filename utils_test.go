package main

import "testing"

func TestCheckRegexp(t *testing.T) {
	tests := []struct {
		Regexp string
		Arg    string
		Exp    bool
	}{
		{"talk", "talk-request", true},
		{"meeting", "meeting-request", true},
		{"Talk", "talk-request", false},
		{"talk", "Talk-request", false},
		{"test-talk", "talk-request", false},
		{"meeting", "meeting", true},
		{"meeting", "meetin", false},
	}

	for _, test := range tests {
		if checkRegexp(test.Regexp, test.Arg) != test.Exp {
			t.Error("not Equal")
		}
	}
}
