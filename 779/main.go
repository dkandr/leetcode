package main

import (
	"fmt"
)

func kthGrammar(n int, k int) int {
	sq := 1
	for k > sq {
		sq *= 2
	}

	reverse := 1

	for k > 2 {
		for sq >= k {
			sq /= 2
		}

		k = k - sq
		reverse *= -1
	}

	switch {
	case k == 1 && reverse == 1:
		return 0
	case k == 1 && reverse == -1:
		return 1
	case k == 2 && reverse == 1:
		return 1
	case k == 2 && reverse == -1:
		return 0
	}

	return sq
}

func main() {
	// fmt.Println(kthGrammar(4, 6))
	fmt.Println(kthGrammar(4, 4))
	// fmt.Println(kthGrammar(30, 434991989))
}
