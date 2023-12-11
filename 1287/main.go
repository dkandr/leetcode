package main

func findSpecialInteger(arr []int) int {
	max := len(arr)/4 + 1

	last := arr[0]
	count := 1

	for i := 1; i < len(arr); i++ {
		if count >= max {
			return last
		}

		if last == arr[i] {
			count++
		} else {
			last = arr[i]
			count = 1
		}
	}

	return last
}
