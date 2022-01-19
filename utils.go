package main

import (
	"regexp"
	"strconv"
)

func checkRegexp(reg, str string) bool {
	return regexp.MustCompile("^" + reg + "(?s).*$").MatchString(str)
}

func makeNumber(arg string) int {
	if checkRegexp("!invite \\d <@", arg) {
		str := regexp.MustCompile("\\d").FindString(arg)
		num, err := strconv.Atoi(str)

		if err != nil {
			return 0
		}

		return num
	}

	return 0
}
