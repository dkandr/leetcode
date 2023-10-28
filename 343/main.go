package leetcode

import "math"

func integerBreak(n int) int {
	if n <= 3 {
		return n - 1
	}

	x := n / 3

	switch n % 3 {
	case 0:
		return int(math.Pow(3, float64(x)))
	case 1:
		return int(math.Pow(3, float64(x-1))) * 4
	case 2:
		return int(math.Pow(3, float64(x))) * 2
	}

	return 0
}
