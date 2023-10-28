package main

import "fmt"

func checkPalindrome(b []byte) bool {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		if b[i] != b[j] {
			return false
		}
	}

	return true
}

func longestPalindrome(s string) string {
	maxLen := 0
	var out string

	for i := 0; i < len(s)-1; i++ {
		for j := len(s) - 1; j > i; j-- {
			if s[i] == s[j] {
				if ok := checkPalindrome([]byte(s[i : j+1])); ok {
					if maxLen < j-i+1 {
						maxLen = j - i + 1
						out = string(s[i : j+1])
					}
				}
			}
		}
	}

	if maxLen > 0 {
		return out
	}

	return string(s[0])
}

func main() {
	fmt.Println(longestPalindrome("aacabdkacaa"))
}
