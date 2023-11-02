package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getTreeData(root *TreeNode, nodes *int) (int, int) {
	if root == nil {
		return 0, 0
	}

	var sum, count, s, c int

	s, c = getTreeData(root.Left, nodes)
	sum += root.Val + s
	count += 1 + c
	s, c = getTreeData(root.Right, nodes)
	sum += s
	count += c

	if root.Val == sum/count {
		(*nodes)++
	}

	return sum, count
}

func averageOfSubtree(root *TreeNode) int {
	nodes := new(int)

	getTreeData(root, nodes)

	return *nodes
}

func main() {
	// head := &TreeNode{
	// 	Val: 4,
	// 	Left: &TreeNode{
	// 		Val: 8,
	// 		Left: &TreeNode{
	// 			Val: 0,
	// 		},
	// 		Right: &TreeNode{
	// 			Val: 1,
	// 		},
	// 	},
	// 	Right: &TreeNode{
	// 		Val: 5,
	// 		Right: &TreeNode{
	// 			Val: 6,
	// 		},
	// 	},
	// }
	head := &TreeNode{
		Val: 1,
	}

	fmt.Println(averageOfSubtree(head))
}
