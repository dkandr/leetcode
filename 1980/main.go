package main

import (
	"fmt"
)

func findDifferentBinaryString(nums []string) string {
	m := make(map[string]struct{})

	for _, v := range nums {
		m[v] = struct{}{}
	}

	l := len(nums[0])
	format := fmt.Sprintf("%%0%db", l)

	for i := 0; i < 2<<l; i++ {
		s := fmt.Sprintf(format, i)
		if _, ok := m[s]; !ok {
			return s
		}
	}

	return ""
}
