package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	var res = make([]int, 0, 32)

	nodes := []*TreeNode{root}

	finished := false

	for !finished {
		var max int
		finished = true

		tmpNodes := make([]*TreeNode, 0, len(nodes)*2)

		for i := 0; i < len(nodes); i++ {
			if nodes[i] != nil {

				// get max in line
				if finished {
					max = nodes[i].Val
				} else {
					if nodes[i].Val > max {
						max = nodes[i].Val
					}
				}

				finished = false
				tmpNodes = append(tmpNodes, nodes[i].Left)
				tmpNodes = append(tmpNodes, nodes[i].Right)
			}
		}
		// only nil slice
		if finished {
			continue
		}

		res = append(res, max)
		nodes = tmpNodes
	}

	return res
}
