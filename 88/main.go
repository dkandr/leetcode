package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2...)
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[j] > nums1[i]
	})
}

func main() {
	var nums1 = []int{1, 2, 3, 0, 0, 0}
	var nums2 = []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}
