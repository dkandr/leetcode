package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func SliceToList(sl []int) *ListNode {
	if sl == nil {
		return nil
	}

	head := &ListNode{
		Val:  sl[0],
		Next: nil,
	}

	cur := head

	for i := 1; i < len(sl); i++ {
		tmp := &ListNode{
			Val:  sl[i],
			Next: nil,
		}

		cur.Next = tmp
		cur = tmp
	}

	return head
}

func ListToSlice(head *ListNode) []int {
	if head == nil {
		return nil
	}

	sl := make([]int, 0)

	sl = append(sl, head.Val)

	for head.Next != nil {
		head = head.Next
		sl = append(sl, head.Val)
	}

	return sl
}
