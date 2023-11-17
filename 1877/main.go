package main

//	func minPairSum(nums []int) int {
//		sort.Slice(nums, func(i, j int) bool {
//			return nums[j] > nums[i]
//		})
//
//		res := nums[0] + nums[len(nums)-1]
//
//		for i, j := 1, len(nums)-2; i < j; i, j = i+1, j-1 {
//			if res < nums[i]+nums[j] {
//				res = nums[i] + nums[j]
//			}
//		}
//
//		return res
//	}

func minPairSum(nums []int) int {
	min, max := nums[0], nums[0]

	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++

		if min > nums[i] {
			min = nums[i]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}

	res := 0
	delta := 0

	i, j := min, max

	for i <= j {
		if _, ok := m[i]; !ok {
			i++
			continue
		}
		if _, ok := m[j]; !ok {
			j--
			continue
		}

		if res < i+j {
			res = i + j
		}

		delta = m[i] - m[j]

		switch {
		case delta < 0:
			m[j] -= m[i]
			i++
		case delta > 0:
			m[i] -= m[j]
			j--
		default:
			i++
			j--
		}
	}

	return res
}
