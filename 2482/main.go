package main

func onesMinusZeros(grid [][]int) [][]int {
	mi := make(map[int]int, len(grid))
	mj := make(map[int]int, len(grid[0]))

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				mi[i]++
				mj[j]++
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			grid[i][j] = 2*(mi[i]+mj[j]) - len(grid) - len(grid[0])
		}
	}

	return grid
}
