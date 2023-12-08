package main

import (
	"strconv"

	. "github.com/dkandr/leetcode/tree"
)

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var out string

	out += strconv.Itoa(int(root.Val))
	out += getNodes(root.Left)
	if root.Left == nil && root.Right != nil {
		out += "()"
	}
	out += getNodes(root.Right)

	return out
}

func getNodes(node *TreeNode) string {
	if node == nil {
		return ""
	}

	var out string

	out += "("
	out += strconv.Itoa(int(node.Val))
	out += getNodes(node.Left)
	if node.Left == nil && node.Right != nil {
		out += "()"
	}
	out += getNodes(node.Right)
	out += ")"

	return out
}
