package main

import (
	"fmt"
)

//func merge(nums1 []int, m int, nums2 []int, n int) {
//	nums1 = append(nums1[:m], nums2...)
//	sort.Slice(nums1, func(i, j int) bool {
//		return nums1[j] > nums1[i]
//	})
//}

func merge(nums1 []int, m int, nums2 []int, n int) {
	p := len(nums1) - 1
	p1 := m - 1
	p2 := n - 1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	for i := 0; i <= p2; i++ {
		nums1[i] = nums2[i]
	}
}

func main() {
	var nums1 = []int{2, 2, 3, 0, 0, 0}
	var nums2 = []int{1, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}
