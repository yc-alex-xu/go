package main

func isMatch(s string, p string) bool {
	return greedy(s, p)
}

func dpMethod(s string, p string) bool {
	m, n := len(s), len(p)
	// dp[i][j] 表示字符串 s的前i个字符和模式p的前j个字符是否能匹配
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	//dp matrix 初始化
	//所有的 dp[0][j]和 dp[i][0] 都是边界条件
	dp[0][0] = true //当字符串 s 和模式 p 均为空时，匹配成功

	//dp[i][0]=False，即空模式无法匹配非空字符串
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' { //只有当模式p 的前j个字符均为星号时，dp[0][j]才为真
			dp[0][i] = true
		} else {
			break
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			//进行状态转移时，我们可以考虑模式 p 的第j个字符
			//和字符串 s 中的第 i 个字符
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

//相对于动态规划，需要的辅助空间为O(1)
func greedy(s string, p string) bool {
	for len(s) > 0 && len(p) > 0 && p[len(p)-1] != '*' {
		if charMatch(s[len(s)-1], p[len(p)-1]) {
			s = s[:len(s)-1]
			p = p[:len(p)-1]
		} else {
			return false
		}
	}
	if len(p) == 0 {
		return len(s) == 0
	}
	sIndex, pIndex := 0, 0
	sRecord, pRecord := -1, -1
	for sIndex < len(s) && pRecord < len(p) {
		if p[pIndex] == '*' {
			pIndex++
			sRecord, pRecord = sIndex, pIndex
		} else if charMatch(s[sIndex], p[pIndex]) {
			sIndex++
			pIndex++
		} else if sRecord != -1 && sRecord+1 < len(s) {
			sRecord++
			sIndex, pIndex = sRecord, pRecord
		} else {
			return false
		}
	}
	return allStars(p, pIndex, len(p))
}

func allStars(str string, left, right int) bool {
	for i := left; i < right; i++ {
		if str[i] != '*' {
			return false
		}
	}
	return true
}

func charMatch(u, v byte) bool {
	return u == v || v == '?'
}
