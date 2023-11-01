package main

import (
	"fmt"
)

func checkMax(x []byte, revers int) bool {
	var check []byte
	if revers == -1 {
		check = []byte{'2', '1', '4', '7', '4', '8', '3', '6', '4', '8'}
	} else {
		check = []byte{'2', '1', '4', '7', '4', '8', '3', '6', '4', '7'}
	}

	if len(x) > len(check) {
		return false
	}

	if len(x) < len(check) {
		return true
	}

	for i := 0; i < len(x); i++ {
		if x[i] > check[i] {
			return false
		}
		if x[i] == check[i] {
			continue
		}
		if x[i] < check[i] {
			return true
		}
	}

	return true
}

func myAtoi(s string) int {
	if s == "" {
		return 0
	}

	x := []byte(s)
	pos := 0

	// remove spaces
	for pos < len(x) && x[pos] == ' ' {
		pos++
	}

	// check - or +
	rev := 1
	if pos < len(x) && x[pos] == '-' {
		rev = -1
		pos++
	} else if pos < len(x) && x[pos] == '+' {
		pos++
	}

	// remove 0
	for pos < len(x) && x[pos] == '0' {
		pos++
	}

	startPos := pos
	for pos < len(x) && x[pos] >= '0' && x[pos] <= '9' {
		pos++
	}

	x = x[startPos:pos]

	if !checkMax(x, rev) {
		if rev == -1 {
			return 1 << 31 * rev
		}
		return 1<<31 - 1
	}

	var out int
	for i := 0; i < len(x); i++ {
		out = out*10 + int(x[i]-'0')
	}

	return out * rev
}

func main() {
	// fmt.Println(myAtoi("123"))
	// fmt.Println(myAtoi("-9128347233"))
	fmt.Println(myAtoi(" "))
}
