package main

import "sort"

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

//func maxFrequency(nums []int, k int) int {
//	sort.Slice(nums, func(i, j int) bool {
//		return nums[j] > nums[i]
//	})
//
//	sum := 0
//
//	for i := 0; i < len(nums)-1; i++ {
//		sum += nums[len(nums)-1] - nums[i]
//	}
//
//	if sum <= k {
//		return len(nums)
//	}
//
//	j := len(nums) - 1
//	m := 1
//	for {
//		if m > j {
//			return m
//		}
//
//		s := sum
//		for i := 0; i < j; i++ {
//			s -= nums[j] - nums[i]
//			if s <= k {
//				if m < j-i {
//					m = j - i
//				}
//				break
//			}
//		}
//
//		sum -= (nums[j] - nums[j-1]) * j
//		if sum <= k {
//			return j
//		}
//
//		j--
//	}
//}

func maxFrequency(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[j] > nums[i]
	})

	sum := 0

	for i := 0; i < len(nums)-1; i++ {
		sum += nums[len(nums)-1] - nums[i]
	}

	if sum <= k {
		return len(nums)
	}

	s1 := sum - (nums[len(nums)-1] - nums[0])
	s2 := sum - (nums[len(nums)-1]-nums[len(nums)-2])*(len(nums)-1)

	r1 := fff(nums, 1, len(nums)-1, s1, k)
	r2 := fff(nums, 0, len(nums)-2, s2, k)

	if r1 <= r2 {
		return r2
	}

	return r1
}

func fff(nums []int, i, j, sum, k int) int {
	if i == j {
		return 1
	}

	if sum <= k {
		return j - i + 1
	}

	s1 := sum - (nums[j] - nums[i])
	s2 := sum - (nums[j]-nums[j-1])*(j-i)

	r1 := fff(nums, i+1, j, s1, k)
	r2 := fff(nums, i, j-1, s2, k)

	if r1 <= r2 {
		return r2
	}

	return r1
}
