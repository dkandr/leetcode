package main

func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	if sx == fx && sy == fy && t == 1 {
		return false
	}

	x := fx - sx
	y := fy - sy

	if x < 0 {
		x *= -1
	}
	if y < 0 {
		y *= -1
	}

	d := x
	if d > y {
		d = y
	}

	if x+y-d > t {
		return false
	}

	return true
}
