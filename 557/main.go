package main

import (
	"fmt"
	"strings"
)

func reverseWord(w string) string {
	rw := []rune(w)

	for i, j := 0, len(rw)-1; i < j; i, j = i+1, j-1 {
		rw[i], rw[j] = rw[j], rw[i]
	}

	return string(rw)
}

func reverseWords(s string) string {
	words := strings.Split(s, " ")

	var out string
	for i := 0; i < len(words)-1; i++ {
		out += reverseWord(words[i])
		out += " "
	}

	out += reverseWord(words[len(words)-1])

	return out
}

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}
