package main

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	b := []byte(strconv.Itoa(x))

	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		if b[i] != b[j] {
			return false
		}
	}

	return true
}

func main() {
	isPalindrome(12121)
}
