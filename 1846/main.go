package main

import "sort"

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[j] > arr[i]
	})

	arr[0] = 1

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1]+1 {
			arr[i] = arr[i-1] + 1
		}
	}

	return arr[len(arr)-1]
}
