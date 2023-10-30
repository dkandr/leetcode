package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	l := len(s)
	n := 2*numRows - 2
	numCols := l / (n - 1)

	tab := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		tab[i] = make([]byte, 0, numCols*2)
	}

	for i := 0; i < l; i++ {
		switch {
		case i%n == 0:
			tab[0] = append(tab[0], s[i])
		case i%n <= numRows-1:
			tab[i%n] = append(tab[i%n], s[i])
		case i%n > numRows-1:
			tab[n-i%n] = append(tab[n-i%n], s[i])
		}
	}

	out := ""

	for i := 0; i < numRows; i++ {
		out += string(tab[i])
	}

	return out
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}
