package main

import "fmt"

// time limit exceeded
// func maxArea(height []int) int {
// 	var max, mh int
// 	for i := 0; i < len(height); i++ {
// 		for j := i + 1; j < len(height); j++ {
// 			if height[i] <= height[j] {
// 				mh = height[i]
// 			} else {
// 				mh = height[j]
// 			}
//
// 			if max < mh*(j-i) {
// 				max = mh * (j - i)
// 			}
// 		}
//
// 	return max
// }

// 5%
//
//	func maxArea(height []int) int {
//		var max, mh, tmp int
//
//		for i := 0; i < len(height)-1; i++ {
//			for j := len(height) - 1; j > i; j-- {
//				if height[i] <= height[j] {
//					mh = height[i]
//				} else {
//					mh = height[j]
//				}
//
//				tmp = mh * (j - i)
//				if max < tmp {
//					max = tmp
//				}
//
//				if height[i] <= height[j] {
//					break
//				}
//			}
//		}
//
//		return max
//	}
func maxArea(height []int) int {
	left, right := 0, len(height)-1

	max, h := 0, 0
	for left < right {
		if height[left] <= height[right] {
			h = height[left]
		} else {
			h = height[right]
		}

		if max < h*(right-left) {
			max = h * (right - left)
		}

		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return max
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
}
