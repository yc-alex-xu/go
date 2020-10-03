package main

import . "leetcode/comm"

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	//加了一个哨兵结点
	dummy := ListNode{Next: head}
	//它总是出现在下个k元组之前
	pre := &dummy
	//p,q 分别为k个结点的头尾
	for p := head; p != nil; {
		q := pre
		for i := 0; i < k; i++ {
			q = q.Next
			if q == nil {
				return dummy.Next
			}
		}
		t := q.Next

		//k个结点交换，得到新的p,q
		p, q = myReverse(p, q)

		pre.Next = p
		q.Next = t
		pre = q
		p = t
	}
	return dummy.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev, p := head, head.Next
	for prev != tail {
		t := p.Next
		p.Next = prev
		prev = p
		p = t
	}
	return tail, head
}

func main() {

	reverseKGroup(InitList5(), 3).PrintList()
}
