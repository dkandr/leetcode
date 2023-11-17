package main

import "sort"

func minPairSum(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[j] > nums[i]
	})

	res := nums[0] + nums[len(nums)-1]

	for i, j := 1, len(nums)-2; i < j; i, j = i+1, j-1 {
		if res < nums[i]+nums[j] {
			res = nums[i] + nums[j]
		}
	}

	return res
}
