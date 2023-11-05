package main

func getWinner(arr []int, k int) int {
	x := 0
	for x != k {
		if len(arr) == 1 {
			return arr[0]
		}

		if arr[0] > arr[1] {
			arr = append(arr[:1], arr[2:]...)
			x++
		} else {
			arr = arr[1:]
			x = 1
		}
	}

	return arr[0]
}
