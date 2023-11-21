package main

// time limit
//func countNicePairs(nums []int) int {
//	rev := func(x int) int {
//		sl := make([]int, 0, 10)
//
//		for x > 0 {
//			sl = append(sl, x%10)
//			x = x / 10
//		}
//
//		var res int
//		p := 1
//		for i := len(sl) - 1; i >= 0; i-- {
//			res += sl[i] * p
//			p *= 10
//		}
//
//		return res
//	}
//
//	var res int
//	for i := 0; i < len(nums)-1; i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[i]+rev(nums[j]) == nums[j]+rev(nums[i]) {
//				res++
//			}
//		}
//	}
//
//	return res % int(math.Pow(10, 9)+7)
//}

func countNicePairs(nums []int) int {
	rev := func(x int) int {
		var res int
		for x > 0 {
			res = res*10 + x%10
			x /= 10
		}

		return res
	}

	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		s := nums[i] - rev(nums[i])
		m[s]++
	}

	var res int
	for _, v := range m {
		res += v * (v - 1) / 2
	}

	mod := 1000000007
	return res % mod
}
