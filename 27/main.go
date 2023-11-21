package main

func removeElement(nums []int, val int) int {
	res := len(nums)
	tmp := make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		if val == nums[i] {
			res--
		} else {
			tmp = append(tmp, nums[i])
		}
	}

	copy(nums, tmp)

	return res
}
