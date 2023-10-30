package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	count := 0

	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	for _, v := range m {
		count += v * (v - 1) / 2
	}

	return count
}

func main() {
	fmt.Println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
	fmt.Println(numIdenticalPairs([]int{1, 1, 1, 1}))
}
