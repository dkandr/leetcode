package main

func rotate(matrix [][]int) {
	l := len(matrix) - 1
	for i := 0; i < (l+1)/2; i++ {
		for j := 0; j < (l+2)/2; j++ {
			matrix[j][l-i], matrix[l-i][l-j], matrix[l-j][i], matrix[i][j] = matrix[i][j], matrix[j][l-i], matrix[l-i][l-j], matrix[l-j][i]
		}
	}
}

func main() {
	rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	rotate([][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}})
}
