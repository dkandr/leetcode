package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	head, head.Next, head.Next.Next = head.Next, head.Next.Next, head
	cur := head.Next

	for cur.Next != nil {
		if cur.Next.Next == nil {
			break
		}
		cur.Next, cur.Next.Next, cur.Next.Next.Next = cur.Next.Next, cur.Next.Next.Next, cur.Next
		cur = cur.Next.Next
	}

	return head
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{Val: 4},
			},
		},
	}

	head = swapPairs(head)
	fmt.Println(head.Val, head.Next.Val, head.Next.Next.Val, head.Next.Next.Next.Val)
}
