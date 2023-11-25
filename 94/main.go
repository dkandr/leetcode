package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int

	res = getNodes(root, res)

	return res
}

func getNodes(root *TreeNode, sl []int) []int {
	if root == nil {
		return sl
	}

	sl = getNodes(root.Left, sl)
	sl = append(sl, root.Val)
	sl = getNodes(root.Right, sl)

	return sl
}

func main() {
	root := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}

	fmt.Println(inorderTraversal(root))
}
