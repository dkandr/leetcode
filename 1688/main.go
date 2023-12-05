package main

func numberOfMatches(n int) int {
	var res int

	for n > 1 {
		res += n / 2
		if n%2 == 1 {
			n++
		}
		n /= 2
	}

	return res
}
