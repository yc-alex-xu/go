package main

import "fmt"
import . "leetcode/comm"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	c := 0 //carry
	head, p := (*ListNode)(nil), (*ListNode)(nil)
	for ; l1 != nil || l2 != nil || c != 0; l1, l2 = next(l1), next(l2) {
		a, b := val(l1), val(l2)
		t := new(ListNode)
		t.Val = (a + b + c) % 10
		t.Next = nil
		c = (a + b + c) / 10
		if p == nil { //first node
			p, head = t, t
		} else {
			p.Next = t
			p = t
		}
	}
	return head
}

func next(p *ListNode) *ListNode {
	if p == nil {
		return nil
	}
	return p.Next

}

func val(p *ListNode) int {
	if p == nil {
		return 0
	}
	return p.Val
}

func main() {
	l1 := &ListNode{Val: 3, Next: nil}
	l1 = &ListNode{Val: 4, Next: l1}
	l1 = &ListNode{Val: 2, Next: l1}

	l2 := &ListNode{Val: 4, Next: nil}
	l2 = &ListNode{Val: 6, Next: l2}
	l2 = &ListNode{Val: 5, Next: l2}

	l := addTwoNumbers(l1, l2)
	for l != nil {
		fmt.Printf("%d->", l.Val)
		l = l.Next
	}
	fmt.Printf("\n")

}
