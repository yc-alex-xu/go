package main

func minPathSum(grid [][]int) int {
	m, n := len(grid[0]), len(grid)
	//将grid 作为dp矩阵
	for j := 1; j < m; j++ {
		grid[0][j] += grid[0][j-1]
	}
	for i := 1; i < n; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[n-1][m-1]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
