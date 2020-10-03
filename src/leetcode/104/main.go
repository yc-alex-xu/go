package main

import . "leetcode/comm"

func maxDepthOld(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if l < r {
		l, r = r, l
	}
	return l + 1
}

//bfs
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	ans := 0
	for len(queue) > 0 {
		ans++
		sz := len(queue)          //当前的层node数目
		for i := 0; i < sz; i++ { //append当前层node的left/right
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[sz:] //当前的层node出队列
	}
	return ans
}
