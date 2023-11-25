package main

// в лоб (time limit)
//func getSumAbsoluteDifferences(nums []int) []int {
//	res := make([]int, len(nums))
//
//	var sum int
//	for i := 0; i < len(nums); i++ {
//		sum = 0
//		for j := 0; j < len(nums); j++ {
//			if nums[i]-nums[j] > 0 {
//				sum += nums[i] - nums[j]
//			} else {
//				sum -= nums[i] - nums[j]
//			}
//		}
//
//		res[i] = sum
//	}
//
//	return res
//}

func getSumAbsoluteDifferences(nums []int) []int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	l := 0
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sum -= nums[i]

		res[i] = nums[i]*i - l + sum - nums[i]*(len(nums)-i-1)

		l += nums[i]
	}

	return res
}
