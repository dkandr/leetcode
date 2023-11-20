package main

import (
	"strings"
	"sync"
)

//func garbageCollection(garbage []string, travel []int) int {
//	minutes := 0
//	var g, m, p int
//	var gf, mf, pf bool
//
//	for i := len(garbage) - 1; i >= 0; i-- {
//		g = strings.Count(garbage[i], "G")
//		m = strings.Count(garbage[i], "M")
//		p = strings.Count(garbage[i], "P")
//
//		if g != 0 {
//			gf = true
//		}
//		if m != 0 {
//			mf = true
//		}
//		if p != 0 {
//			pf = true
//		}
//
//		if i != 0 && gf {
//			minutes += g + travel[i-1]
//		}
//		if i != 0 && mf {
//			minutes += m + travel[i-1]
//		}
//		if i != 0 && pf {
//			minutes += p + travel[i-1]
//		}
//		if i == 0 {
//			minutes += g + m + p
//		}
//	}
//
//	return minutes
//}

func garbageCollection(garbage []string, travel []int) int {
	var wg sync.WaitGroup
	ch := make(chan int, 3)

	wg.Add(3)
	go func() {
		findCount(ch, garbage, travel, "G")
		wg.Done()
	}()
	go func() {
		findCount(ch, garbage, travel, "M")
		wg.Done()
	}()
	go func() {
		findCount(ch, garbage, travel, "P")
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	minutes := 0
	for v := range ch {
		minutes += v
	}

	return minutes
}

func findCount(ch chan<- int, garbage []string, travel []int, sub string) {
	var count, out int
	var find bool
	for i := len(garbage) - 1; i >= 0; i-- {
		count = strings.Count(garbage[i], sub)
		if count != 0 {
			find = true
		}

		if i != 0 && find {
			out += count + travel[i-1]
		}
		if i == 0 {
			out += count
		}
	}

	ch <- out
}
