package leetcode

import . "leetcode/comm"

func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}
	mid := l / 2
	root := TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return &root
}

func main() {
	root := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	root.FirstOrder()
}
