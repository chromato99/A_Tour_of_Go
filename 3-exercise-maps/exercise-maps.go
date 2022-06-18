package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	arr := strings.Fields(s)
	m := make(map[string]int)
	for _, str := range arr {
		if _, ok := m[str]; ok {
			m[str]++
		} else {
			m[str] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
