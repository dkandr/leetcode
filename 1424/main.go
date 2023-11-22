package main

//func findDiagonalOrder(nums [][]int) []int {
//	var resLen, m int
//	for i := 0; i < len(nums); i++ {
//		resLen += len(nums[i])
//		if m < len(nums[i]) {
//			m = len(nums[i])
//		}
//	}
//
//	m += len(nums)
//
//	res := make([]int, 0, resLen)
//
//	for i := 0; i < m; i++ {
//		for h, v := 0, i; h <= i && v >= 0; h, v = h+1, v-1 {
//			if v < len(nums) {
//				if h < len(nums[v]) {
//					res = append(res, nums[v][h])
//					if len(res) == resLen {
//						return res
//					}
//				}
//			}
//		}
//	}
//
//	return res
//}

// подсмотрел у индуса
//func findDiagonalOrder(nums [][]int) []int {
//	resSize := 0
//	dLineCount := 0
//	m := make(map[int][]int, int(1e5+1))
//
//	for i := 0; i < len(nums); i++ {
//		resSize += len(nums[i])
//
//		for j := 0; j < len(nums[i]); j++ {
//			m[i+j] = append(m[i+j], nums[i][j])
//			if dLineCount < i+j {
//				dLineCount = i + j
//			}
//		}
//	}
//
//	res := make([]int, 0, resSize)
//
//	for i := 0; i <= dLineCount; i++ {
//		for j := len(m[i]) - 1; j >= 0; j-- {
//			res = append(res, m[i][j])
//		}
//	}
//
//	return res
//}

func findDiagonalOrder(nums [][]int) []int {
	var resSize, maxLen int
	for i := 0; i < len(nums); i++ {
		resSize += len(nums[i])
		if maxLen < len(nums[i]) {
			maxLen = len(nums[i])
		}
	}

	dLineCount := 0
	sl := make([][]int, len(nums)+maxLen-1)

	for i := len(nums) - 1; i >= 0; i-- {
		for j := len(nums[i]) - 1; j >= 0; j-- {
			sl[i+j] = append(sl[i+j], nums[i][j])
			if dLineCount < i+j {
				dLineCount = i + j
			}
		}
	}

	res := make([]int, 0, resSize)

	for i := 0; i < len(sl); i++ {
		res = append(res, sl[i]...)
	}

	return res
}
