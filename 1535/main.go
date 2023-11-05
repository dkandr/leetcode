package main

import "sort"

// 21%
//func getWinner(arr []int, k int) int {
//	x := 0
//	for x != k {
//		if len(arr) == 1 {
//			return arr[0]
//		}
//
//		if arr[0] > arr[1] {
//			arr = append(arr[:1], arr[2:]...)
//			x++
//		} else {
//			arr = arr[1:]
//			x = 1
//		}
//	}
//
//	return arr[0]
//}

// 78%
func getWinner(arr []int, k int) int {
	if k >= len(arr) {
		sort.Slice(arr, func(i, j int) bool {
			return arr[j] < arr[i]
		})
		return arr[0]
	}

	winner := arr[0]
	x := 0
	for i := 1; i < len(arr); i++ {
		if x == k {
			return winner
		}

		if winner < arr[i] {
			winner = arr[i]
			x = 1
		} else {
			x++
		}
	}

	return winner
}
