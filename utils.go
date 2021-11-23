package main

import "regexp"

func checkRegexp(reg, str string) bool {
	return regexp.MustCompile("^" + reg + ".*$").Match([]byte(str))
}
