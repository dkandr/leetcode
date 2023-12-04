package main

func largestGoodInteger(num string) string {
	if len(num) == 0 {
		return ""
	}

	var to, count int
	var max byte
	last := num[0]
	count = 1

	for i := 1; i < len(num); i++ {
		if num[i] <= max {
			last = num[i]
			count = 1
			continue
		}

		if num[i] == last && count < 3 {
			count++

			if count > 2 {
				max = last
				to = i
			}
		} else {
			last = num[i]
			count = 1
		}
	}

	if max > 0 {
		return string(num[to-2 : to+1])
	}

	return ""
}
