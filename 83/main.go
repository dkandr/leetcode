// 83. Remove Duplicates from Sorted List

package main

import . "github.com/dkandr/leetcode/list"

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	tmp := cur.Next

	for tmp != nil {
		if cur.Val != tmp.Val {
			cur.Next = tmp
			cur = tmp
		}
		tmp = tmp.Next
	}

	if cur != tmp {
		cur.Next = nil
	}

	return head
}
