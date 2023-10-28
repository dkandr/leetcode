package tree

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SliceToTree(sl []string) *TreeNode {
	if len(sl) == 0 || sl[0] == "nil" || sl[0] == "null" {
		return nil
	}

	val, err := strconv.Atoi(sl[0])
	if err != nil {
		panic("bad first element")
	}

	root := &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}

	i := 1
	nodesInLine := 2
	nodes := make([]*TreeNode, 1)
	nodes[0] = root

	for i < len(sl) {
		tmpNodes := make([]*TreeNode, 0, nodesInLine)

		for j := 0; j < len(nodes); j = j + 2 {
			var tmpNodeL *TreeNode
			var tmpNodeR *TreeNode

			// left
			if sl[i+j] == "nil" || sl[i+j] == "null" {
				tmpNodeL = nil
			} else {
				val, err := strconv.Atoi(sl[i+j])
				if err != nil {
					panic("bad int")
				}
				tmpNodeL = &TreeNode{
					Val:   val,
					Left:  nil,
					Right: nil,
				}
			}

			nodes[j].Left = tmpNodeL

			// right
			if sl[i+j+1] == "nil" || sl[i+j+1] == "null" {
				tmpNodeR = nil
			} else {
				val, err := strconv.Atoi(sl[i+j+1])
				if err != nil {
					panic("bad int")
				}
				tmpNodeR = &TreeNode{
					Val:   val,
					Left:  nil,
					Right: nil,
				}
			}

			nodes[j].Right = tmpNodeR

			tmpNodes = append(tmpNodes, tmpNodeL)
			tmpNodes = append(tmpNodes, tmpNodeR)
		}

		nodes = tmpNodes
		i += nodesInLine
		nodesInLine *= 2
	}

	return root
}
