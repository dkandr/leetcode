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

func maxArea(height []int) int {
	m := make(map[int]int, len(height)-1)

	return 0
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
