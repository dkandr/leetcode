package main

import "sync"

// 12%
//func getLastMoment(n int, left []int, right []int) int {
//	res := 0
//
//	if len(left) == 0 && len(right) == 0 {
//		return res
//	}
//
//	sort.Slice(left, func(i, j int) bool {
//		return left[j] > left[i]
//	})
//	sort.Slice(right, func(i, j int) bool {
//		return right[j] > right[i]
//	})
//
//	if len(left) != 0 && left[0] == 0 {
//		left = left[1:]
//	}
//	if len(right) != 0 && right[len(right)-1] == n {
//		right = right[:len(right)-1]
//	}
//
//	for len(left) != 0 && len(right) != 0 {
//		for i := 0; i < len(left); i++ {
//			left[i]--
//		}
//		for i := 0; i < len(right); i++ {
//			right[i]++
//		}
//		if left[0] == 0 {
//			left = left[1:]
//		}
//		if right[len(right)-1] == n {
//			right = right[:len(right)-1]
//		}
//
//		res++
//	}
//
//	if len(left) != 0 {
//		res += left[len(left)-1]
//	}
//	if len(right) != 0 {
//		res += n - right[0]
//	}
//
//	return res
//}

// 62%
// 86% with goroutine
func getLastMoment(n int, left []int, right []int) int {
	res := 0

	if len(left) == 0 && len(right) == 0 {
		return res
	}

	minRight, maxLeft := n, 0
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		for i := 0; i < len(right); i++ {
			if minRight > right[i] {
				minRight = right[i]
			}
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < len(left); i++ {
			if maxLeft < left[i] {
				maxLeft = left[i]
			}
		}
		wg.Done()
	}()

	wg.Wait()

	switch {
	case len(left) == 0 && len(right) != 0:
		res = n - minRight
	case len(right) == 0 && len(left) != 0:
		res = maxLeft
	default:
		if minRight <= n-maxLeft {
			res = n - minRight
		} else {
			res = maxLeft
		}
	}

	return res
}
