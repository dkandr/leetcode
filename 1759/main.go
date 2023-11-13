package main

func countHomogenous(s string) int {
	var out int

	var count int
	var last rune
	for _, v := range s {
		if last == 0 {
			last = v
		}

		if last == v {
			count++
		} else {
			last = v
			out += count * (count + 1) / 2
			count = 1
		}
	}

	// return it modulo 1e9 + 7
	out += count * (count + 1) / 2 % (1e9 + 7)

	return out
}
