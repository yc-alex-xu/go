package main

import . "leetcode/comm"

func swapPairs(head *ListNode) *ListNode {
	//为空或只有一个结点，直接返回
	if head == nil || head.Next == nil {
		return head
	}
	prev, p := head, head.Next
	head = p
	for {
		//先保存会修改的值
		t := p.Next

		//交换pre,p
		p.Next = prev

		//如果后面只有0或1个结点
		if t == nil || t.Next == nil {
			prev.Next = t
			break
		}
		prev.Next = t.Next
		prev, p = t, t.Next
	}
	return head
}

func main() {
	swapPairs(InitList4()).PrintList()
}
