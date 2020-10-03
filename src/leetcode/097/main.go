package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m, t := len(s1), len(s2), len(s3)
	if (n + m) != t {
		return false
	}
	f := make([]bool, m+1)
	//f[i][j]表示s3前k=i+j个字符是否为s1,s2 前i,j个字符的交错
	//滚动数组优化空间,只用一行，所以f[j]表示即可
	f[0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			k := i + j
			//右边的f[j]实际是上一行的dp结果
			if i > 0 {
				f[j] = f[j] && s1[i-1] == s3[k-1]
			}
			if j > 0 {
				f[j] = f[j] || f[j-1] && s2[j-1] == s3[k-1]
			}
		}
	}
	return f[m]
}
