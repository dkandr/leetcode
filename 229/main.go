package main

func majorityElement(nums []int) []int {
	t := len(nums) / 3

	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	res := make([]int, 0, len(m))

	for k, v := range m {
		if v > t {
			res = append(res, k)
		}
	}

	return res
}
