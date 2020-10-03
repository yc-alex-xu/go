package leetcode

import . "leetcode/comm"

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxSum := root.Val

	//由于涉及递归调用，所以得申明
	var maxGain func(*TreeNode) int
	//内嵌函数，以利于访问外部变量
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 递归计算左右子节点的最大贡献值
		// 只有在最大贡献值大于 0 时，才会选取对应子节点
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)
		subTreeGain := node.Val + leftGain + rightGain

		// 选择gain 最大的subtree
		maxSum = max(maxSum, subTreeGain)

		// 返回左支或右支
		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
