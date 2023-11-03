package main

import "fmt"

func buildArray(target []int, n int) []string {
	if n < target[len(target)-1] {
		return nil
	}

	curSteam := 1
	out := make([]string, 0, target[len(target)-1]*2)

	for i := 0; i < len(target); i++ {
		for target[i] > curSteam {
			out = append(out, "Push")
			out = append(out, "Pop")
			curSteam++
		}

		out = append(out, "Push")
		curSteam++
	}

	return out
}

func main() {
	fmt.Println(buildArray([]int{1, 3}, 3))
	fmt.Println(buildArray([]int{1, 2, 3}, 3))
	fmt.Println(buildArray([]int{1, 2}, 4))
}
