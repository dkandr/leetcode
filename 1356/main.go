package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		if strings.Count(strconv.FormatInt(int64(arr[j]), 2), "1") > strings.Count(strconv.FormatInt(int64(arr[i]), 2), "1") {
			return true
		}
		if strings.Count(strconv.FormatInt(int64(arr[i]), 2), "1") == strings.Count(strconv.FormatInt(int64(arr[j]), 2), "1") {
			return arr[j] > arr[i]
		}
		return false
	})

	return arr
}

func main() {
	fmt.Println(sortByBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
}
