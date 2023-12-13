package main

func numSpecial(mat [][]int) int {
	var count int
	mm := make(map[int]int, len(mat))
	mn := make(map[int]int, len(mat[0]))

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 {
				mm[i]++
				mn[j]++
			}
		}
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 {
				if mm[i] == 1 && mn[j] == 1 {
					count++
				}
			}
		}
	}

	return count
}
