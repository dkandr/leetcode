package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func fillMap(root *TreeNode, m map[int]int) {
	if root == nil {
		return
	}

	fillMap(root.Left, m)
	m[root.Val]++
	fillMap(root.Right, m)
}

func findMode(root *TreeNode) []int {
	m := make(map[int]int)
	fillMap(root, m)

	var max int
	for _, v := range m {
		if max < v {
			max = v
		}
	}

	var res []int
	for k, v := range m {
		if v == max {
			res = append(res, k)
		}
	}

	return res
}

func main() {
	// head := &TreeNode{
	// 	Val:  1,
	// 	Left: nil,
	// 	Right: &TreeNode{
	// 		Val: 2,
	// 		Left: &TreeNode{
	// 			Val:   2,
	// 			Left:  nil,
	// 			Right: nil,
	// 		},
	// 		Right: nil,
	// 	},
	// }

	// head := &TreeNode{
	// 	Val: 0,
	// }

	head := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
		},
	}

	fmt.Println(findMode(head))
}
