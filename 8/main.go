package main

import "fmt"

func myAtoi(s string) int {
	x := []byte(s)
	pos := 0

	// remove spaces
	for x[pos] == ' ' {
		pos++
	}

	// check - or +
	rev := 1
	if x[pos] == '-' {
		rev = -1
		pos++
	}
	if x[pos] == '+' {
		pos++
	}

	// remove 0
	for x[pos] == '0' {
		pos++
	}

	startPos := pos
	for x[pos] >= '0' && x[pos] <= '9' {
		pos++
	}

	return 0 * rev
}

func main() {
	var x int64
	x = 1<<63 - 1
	fmt.Println(x)
	// fmt.Println(myAtoi("123"))
}
