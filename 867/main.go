package main

func transpose(matrix [][]int) [][]int {
	res := make([][]int, len(matrix[0]))

	for i := 0; i < len(matrix[0]); i++ {
		res[i] = make([]int, len(matrix))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			res[j][i] = matrix[i][j]
		}
	}

	return res
}
