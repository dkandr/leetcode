package main

import "fmt"

var m map[int]int

func mmm() {
	m = make(map[int]int, 16)

	m[0] = 1

	for i := 1; i < 16; i++ {
		m[i] = m[i-1] * 4
	}
}

func isPowerOfFour(n int) bool {
	if len(m) == 0 {
		mmm()
	}

	for _, v := range m {
		if v == n {
			return true
		}
	}

	return false
}

func main() {

	fmt.Println(isPowerOfFour(1))
	fmt.Println(isPowerOfFour(2))
	fmt.Println(isPowerOfFour(4))
	fmt.Println(isPowerOfFour(5))
	fmt.Println(isPowerOfFour(256))
}
