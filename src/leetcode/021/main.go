package main

import . "leetcode/comm"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := ListNode{}
	pre := &dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pre.Next = l1
			pre = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			pre = l2
			l2 = l2.Next
		}
	}
	if l1 == nil {
		pre.Next = l2
	} else if l2 == nil {
		pre.Next = l1
	}
	return dummy.Next
}

func main() {

	tail := &ListNode{Val: 4, Next: nil}
	tail = &ListNode{Val: 2, Next: tail}
	l1 := &ListNode{Val: 1, Next: tail}

	tail = &ListNode{Val: 4, Next: nil}
	tail = &ListNode{Val: 3, Next: tail}
	l2 := &ListNode{Val: 1, Next: tail}

	mergeTwoLists(l1, l2).PrintList()
}
