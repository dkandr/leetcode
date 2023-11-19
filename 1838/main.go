package main

import (
	"sort"
)

// time limit
//func maxFrequency(nums []int, k int) int {
//	sort.Slice(nums, func(i, j int) bool {
//		return nums[j] > nums[i]
//	})
//
//	res := 1
//	count := 1
//	delta := k
//
//	for i := len(nums) - 1; i >= 0; i-- {
//		for j := i - 1; j >= 0; j-- {
//			delta -= nums[i] - nums[j]
//
//			if delta < 0 {
//				break
//			} else {
//				count++
//
//				if res < count {
//					res = count
//				}
//			}
//		}
//		count = 1
//		delta = k
//	}
//
//	return res
//}

func maxFrequency(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[j] > nums[i]
	})

	res := 1

	return res
}
