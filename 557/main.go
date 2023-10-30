package main

import (
	"fmt"
)

func reverseWords(s string) string {
	var out string
	var rw string

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			out = rw + out
			rw = ""
			out = " " + out
		} else {
			rw += string(s[i])
		}
	}

	out = rw + out

	return out
}

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}
