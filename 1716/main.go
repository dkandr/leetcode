package main

// func totalMoney(n int) int {
// 	var res int
//
// 	m := make(map[int]int, 8)
// 	m[0] = 0
// 	for i := 1; i <= 7; i++ {
// 		m[i] = m[i-1] + i
// 	}
//
// 	w := n / 7
// 	d := n % 7
//
// 	res = m[7]*w + 7*(w*(w-1)/2) + m[d] + d*w
//
// 	return res
// }

func totalMoney(n int) int {
	var res int

	sl := make([]int, 8)
	for i := 1; i <= 7; i++ {
		sl[i] = sl[i-1] + i
	}

	w := n / 7
	d := n % 7

	res = sl[7]*w + 7*(w*(w-1)/2) + sl[d] + d*w

	return res
}
