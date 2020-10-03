package leetcode

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n) //第n层有n个元素
	//dp[i][j]表示从顶部到triangle[i][j]的最小路径，
	//不过由于推导dp[i][]时，只涉及dp[i-1][],所以空间优化了一下
	dp[0] = triangle[0][0]
	for i := 1; i < n; i++ { //逐层循环
		//每层最后一个结点只有一个选项
		dp[i] = dp[i-1] + triangle[i][i]
		//如果题目改成“相邻的结点 在这里指的是 下标 与相同或者等于上一层结点下标-1 的两个结点”
		//那么循环就得改成递增了
		for j := i - 1; j > 0; j-- {
			//其他结点有选择
			dp[j] = min(dp[j-1], dp[j]) + triangle[i][j]
		}
		//每层第一个结点也只有一个选项
		dp[0] += triangle[i][0]
	}
	ans := dp[0]
	for _, v := range dp[1:] {
		if v < ans {
			ans = v
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
