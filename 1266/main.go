package main

func minTimeToVisitAllPoints(points [][]int) int {
	var x, y, res int

	for i := 1; i < len(points); i++ {
		x = points[i][0] - points[i-1][0]
		y = points[i][1] - points[i-1][1]
		if x < 0 {
			x *= -1
		}
		if y < 0 {
			y *= -1
		}

		if x > y {
			res += x
		} else {
			res += y
		}
	}

	return res
}
