package leetcode

import . "leetcode/comm"

func postorderTraversal(root *TreeNode) []int {
	return post(root)
}

func post(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	if root.Left != nil {
		result = post(root.Left)
	}
	if root.Right != nil {
		r := post(root.Right)
		result = append(result, r...)
	}
	result = append(result, root.Val)
	return result
}
