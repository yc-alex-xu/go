package leetcode

import . "leetcode/comm"

func hasPathSum(root *TreeNode, sum int) bool {
	return bfs(root, sum)
}

func dummy(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	sum -= root.Val
	r1 := root.Left != nil && hasPathSum(root.Left, sum)
	r2 := root.Right != nil && hasPathSum(root.Right, sum)
	return r1 || r2
}

//消耗内存比简单的递归还多
func bfs(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	queNode := []*TreeNode{}
	queVal := []int{}
	queNode = append(queNode, root)
	queVal = append(queVal, root.Val)
	for len(queNode) != 0 {
		now := queNode[0]
		queNode = queNode[1:]
		temp := queVal[0]
		queVal = queVal[1:]
		if now.Left == nil && now.Right == nil {
			if temp == sum {
				return true
			}
			continue
		}
		if now.Left != nil {
			queNode = append(queNode, now.Left)
			queVal = append(queVal, now.Left.Val+temp)
		}
		if now.Right != nil {
			queNode = append(queNode, now.Right)
			queVal = append(queVal, now.Right.Val+temp)
		}
	}
	return false
}
