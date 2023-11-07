package main

func eliminateMaximum(dist []int, speed []int) int {
	var kills int
	var maxDay int
	mDays := make(map[int]int)

	for i := 0; i < len(dist); i++ {
		if dist[i]%speed[i] == 0 {
			if maxDay < dist[i]/speed[i]-1 {
				maxDay = dist[i]/speed[i] - 1
			}
			mDays[dist[i]/speed[i]-1]++
		} else {
			if maxDay < dist[i]/speed[i] {
				maxDay = dist[i] / speed[i]
			}
			mDays[dist[i]/speed[i]]++
		}

	}

	var x int
	if v, ok := mDays[0]; ok {
		kills = 1
		if v > kills {
			return kills
		}
	} else {
		x++
	}
	x++

	for i := 1; i <= maxDay; i++ {
		if v, ok := mDays[i]; ok {
			if x-v < 0 {
				kills += x
				return kills
			} else {
				kills += v
				x = x - v + 1
			}
		} else {
			x++
		}
	}

	return kills
}
