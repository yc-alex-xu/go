package main

import "math"
import . "leetcode/comm"

func isValidBST(root *TreeNode) bool {
	return valid(root, math.MinInt64, math.MaxInt64)
}

func valid(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	if valid(root.Left, min, root.Val) && valid(root.Right, root.Val, max) {
		return true
	}
	return false
}
