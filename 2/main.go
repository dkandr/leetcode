package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{
		Val: (l1.Val + l2.Val) % 10,
	}
	add := (l1.Val + l2.Val) / 10
	cur := root

	l1 = l1.Next
	l2 = l2.Next

	for l1 != nil || l2 != nil {
		var tmp *ListNode

		switch {
		case l1 != nil && l2 != nil:
			tmp = &ListNode{
				Val: (l1.Val + l2.Val + add) % 10,
			}
			add = (l1.Val + l2.Val + add) / 10
			l1 = l1.Next
			l2 = l2.Next
		case l1 != nil:
			tmp = &ListNode{
				Val: (l1.Val + add) % 10,
			}
			add = (l1.Val + add) / 10
			l1 = l1.Next
		case l2 != nil:
			tmp = &ListNode{
				Val: (l2.Val + add) % 10,
			}
			add = (l2.Val + add) / 10
			l2 = l2.Next

		}
		cur.Next = tmp
		cur = tmp
	}

	if add == 1 {
		tmp := &ListNode{
			Val: add,
		}
		cur.Next = tmp
	}

	return root
}

func main() {
	// fmt.Println(addTwoNumbers(&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}, &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}))
	root := addTwoNumbers(&ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}, &ListNode{Val: 9})
	_ = root
}
