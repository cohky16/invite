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
		{"!invite", "!invite @hoge", true},
		{"!invite", "!invite @hoge\n message", true},
		{"!invite \\d", "!invite 1 @hoge", true},
		{"!invite \\d", "!invite10 @hoge", false},
		{"!invite \\d", "!invite　1 @hoge", false},
		{"!invite \\d", "!invite @hoge", false},
		{"!invite \\d <@", "!invite 1 <@hoge", true},
		{"!invite \\d <@", "!invite 2 hoge", false},
		{"!invite \\d <@", "!invite 3 <@hoge\n message", true},
		{"Talk.*1", "Talk Room1", true},
		{"Talk.*1", "Talk Room3", false},
		{"Talk.*2", "Talk Room1", false},
	}

	for _, test := range tests {
		if checkRegexp(test.Regexp, test.Arg) != test.Exp {
			t.Error("not Equal")
		}
	}
}

func TestMakeNumber(t *testing.T) {
	tests := []struct {
		Arg    string
		Result int
	}{
		{"!invite 1 <@hoge>", 1},
		{"!invite 3 <@hoge>", 3},
		{"!invite3 <@hoge>", 0},
		{"!invite　3 <@hoge>", 0},
		{"!invite 3 <@hoge> 11111222", 3},
		{"!invite 1 <@hoge> 11111", 1},
		{"!invite 1 @hoge 11111", 0},
	}

	for _, test := range tests {
		if makeNumber(test.Arg) != test.Result {
			t.Error("not equal")
		}
	}
}
