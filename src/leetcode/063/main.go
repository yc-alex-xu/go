package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	return dpWay(obstacleGrid)
}
func dpWay(obstacleGrid [][]int) int {
	//m行*n列
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	//用 dp(i, j)来表示从坐标 (0, 0)到坐标 (i, j) 的路径总数
	//由于状态转移方程: dp(i, j)只与 dp(i−1,j) 和 dp(i,j−1) 相关，
	//我们可以运用「滚动数组思想」把空间复杂度优化称 O(m)。
	//由于逐行迭代，所以省略了i
	dp := make([]int, n)
	//初始化值
	if obstacleGrid[0][0] == 0 {
		dp[0] = 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				continue
			}
			//即 dp(i, j)= f(i-1, j1)+dp(i,j−1) 转移得到。
			if j-1 >= 0 && obstacleGrid[i][j-1] == 0 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n-1]
}
