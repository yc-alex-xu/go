/*
some data structure/func common for the leetcode

*/

package leetcode

import "fmt"

//ListNode the linked list node
type ListNode struct {
	Val  int
	Next *ListNode
}

//InitList4 init a ListNode of 1->2-3->4
func InitList4() *ListNode {
	tail := &ListNode{Val: 4}
	tail = &ListNode{Val: 3, Next: tail}
	tail = &ListNode{Val: 2, Next: tail}
	tail = &ListNode{Val: 1, Next: tail}
	return tail
}

//InitList5 init a ListNode of 1->2-3->4->5
func InitList5() *ListNode {
	tail := &ListNode{Val: 5}
	tail = &ListNode{Val: 4, Next: tail}
	tail = &ListNode{Val: 3, Next: tail}
	tail = &ListNode{Val: 2, Next: tail}
	tail = &ListNode{Val: 1, Next: tail}
	return tail
}

//PrintList Print the list of h
func (h *ListNode) PrintList() {
	for h != nil {
		fmt.Printf("%d->", h.Val)
		h = h.Next
	}
	fmt.Println()
}

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//FirstOrder  traverse the tree in firstOrder
func (root *TreeNode) FirstOrder() {
	if root == nil {
		return
	}
	fmt.Printf("%d", root.Val)
	if root.Left != nil {
		fmt.Printf("\t->%d", root.Left.Val)
	} else {
		fmt.Printf("\t   ")
	}
	if root.Right != nil {
		fmt.Printf("\t->%d", root.Right.Val)
	}
	fmt.Println()
	root.Left.FirstOrder()
	root.Right.FirstOrder()
}
