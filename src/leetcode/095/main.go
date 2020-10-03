package main

func numTrees(n int) int {
	//用递推来代替递归，否者就是 n=0或1返回1
	G := make([]int, n+1)
	G[0], G[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ { //root选第一个点---到选最后一个点
			G[i] += G[j-1] * G[i-j] //左子树个数 × 右子树个数
		}
	}
	return G[n]
}
