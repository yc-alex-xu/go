package leetcode

import . "leetcode/comm"

func isBalanced(root *TreeNode) bool {
	_, b := depth(root)
	return b
}

func depth(root *TreeNode) (l int, b bool) {
	if root == nil {
		return 0, true
	}
	l, bl := depth(root.Left)
	if !bl {
		b = false
		return
	}
	r, br := depth(root.Right)
	if !br {
		b = false
		return
	}
	if l < r {
		l, r = r, l
	}
	b = ((l - r) < 2)
	return 1 + l, b
}
