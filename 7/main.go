package main

import "fmt"

func reverse(x int) int {
	r := 1

	if x < 0 {
		r = -1
		x *= r
	}

	if x == 0 {
		return 0
	}

	for x%10 == 0 {
		x = x / 10
	}

	f := func(x int) []int {
		res := make([]int, 0, 0)

		for x > 0 {
			res = append(res, x%10)
			x = x / 10
		}

		return res
	}

	maxSl := []int{2, 1, 4, 7, 4, 8, 3, 6, 4, 7}

	xSl := f(x)

	if len(xSl) > len(maxSl) {
		return 0
	}

	if len(xSl) == len(maxSl) {
		for i := 0; i < len(maxSl); i++ {
			if xSl[i] > maxSl[i] {
				return 0
			}

			if xSl[i] == maxSl[i] {
				continue
			}

			break
		}
	}

	var res int
	for _, v := range xSl {
		res = res*10 + v
	}

	return res * r
}

func main() {
	fmt.Println(reverse(-123))
}
