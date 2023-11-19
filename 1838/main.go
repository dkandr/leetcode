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
	if len(nums) == 1 {
		return 1
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[j] > nums[i]
	})

	res := len(nums)
	sum := 0
	max := nums[len(nums)-1]
	for i := 0; i < len(nums)-1; i++ {
		sum += max - nums[i]
	}

	i, j := 0, len(nums)-2
	for i <= j {
		if res == 1 {
			return 1
		}
		if sum <= k {
			return res
		}

		if nums[i+1]-nums[i] > nums[j+1]-nums[j] {
			sum -= nums[j+1] - nums[i]
			i++
		} else {
			sum -= (nums[j+1] - nums[j]) * (res - 1)
			j--
		}

		res--
	}

	return res
}
