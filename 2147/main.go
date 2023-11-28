package main

func numberOfWays(corridor string) int {
	seats := 0
	rightPos := 0

	for i := len(corridor) - 1; i >= 0; i-- {
		if corridor[i] == 'S' {
			seats++
			if seats == 2 {
				rightPos = i
			}
		}
	}

	if seats == 0 || seats%2 == 1 {
		return 0
	}

	seats = 0
	count := 1
	plants := 1

	for i := 0; i < rightPos; i++ {
		if corridor[i] == 'S' {
			if seats < 2 {
				seats++
			} else {
				seats = 1
				count = count * plants % int(1e9+7)
				plants = 1
			}
		} else {
			if seats == 2 {
				plants++
			}
		}
	}

	return count * plants % int(1e9+7)
}
