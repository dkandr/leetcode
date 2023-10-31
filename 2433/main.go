package main

import "fmt"

func findArray(pref []int) []int {
	res := make([]int, len(pref))
	res[0] = pref[0]

	for i := 1; i < len(pref); i++ {
		res[i] = pref[i-1] ^ pref[i]
	}

	return res
}

func main() {
	fmt.Println(findArray([]int{5, 2, 0, 3, 1}))
}
