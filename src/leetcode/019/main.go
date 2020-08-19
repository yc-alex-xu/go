package main

//ListNode  node of linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p, pp := head, head
	for i := 0; i < n && pp != nil; i++ {
		pp = pp.Next
	}
	pre := p
	for pp != nil {
		pp = pp.Next
		pre = p
		p = p.Next
	}
	//pre did not move,这意味需要删除的是head
	//可以引入一个 sentinel node,让sentinel.Next = head;这样投指针就没有特殊性了
	if pre == p {
		return head.Next
	}
	pre.Next = p.Next
	return head
}
