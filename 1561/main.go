package main

import "sort"

func maxCoins(piles []int) int {
	sort.Slice(piles, func(i, j int) bool {
		return piles[j] > piles[i]
	})

	res := 0
	count := len(piles) / 3

	for i := 1; i <= count; i++ {
		res += piles[len(piles)-2*i]
	}

	return res
}
