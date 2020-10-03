package main

import (
	"fmt"
	. "leetcode/comm"
)

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	l := inorderTraversal(root.Left)
	if l != nil {
		result = append(result, l...)
	}
	result = append(result, root.Val)
	r := inorderTraversal(root.Right)
	if r != nil {
		result = append(result, r...)
	}
	return result
}

func main() {
	root := &TreeNode{Val: 3}
	root = &TreeNode{Val: 2, Left: root}
	root = &TreeNode{Val: 1, Right: root}
	fmt.Println(inorderTraversal(root))
}
