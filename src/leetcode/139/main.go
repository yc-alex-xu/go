package leetcode

func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for _, v := range wordDict {
		dict[v] = true
	}
	//dp[i]=true mean, s[:i] had be splitted and found in the dict
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for h := 1; h <= len(s); h++ {
		for l := 0; l < h; l++ {
			if dp[l] && dict[s[l:h]] {
				dp[h] = true
				break // a feasible split found
			}
		}
	}
	return dp[len(s)]
}
