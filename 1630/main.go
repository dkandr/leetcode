package main

import "sort"

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	res := make([]bool, len(l))
	for i := 0; i < len(l); i++ {
		res[i] = checkSlice(nums[l[i] : r[i]+1])
	}
	return res
}

func checkSlice(s []int) bool {
	if len(s) == 0 {
		return true
	}

	sl := make([]int, len(s))
	copy(sl, s)

	sort.Slice(sl, func(i, j int) bool {
		return sl[j] > sl[i]
	})

	d := sl[1] - sl[0]
	for i := 1; i < len(sl)-1; i++ {
		if sl[i+1]-sl[i] != d {
			return false
		}
	}

	return true
}
