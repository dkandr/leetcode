package main

func maxProduct(nums []int) int {
	max1 := nums[0]
	max2 := nums[1]

	if max1 > max2 {
		max1, max2 = max2, max1
	}

	for i := 2; i < len(nums); i++ {
		if nums[i] > max2 {
			max1 = max2
			max2 = nums[i]
			continue
		}
		if nums[i] > max1 {
			max1 = nums[i]
		}
	}

	return (max1 - 1) * (max2 - 1)
}
