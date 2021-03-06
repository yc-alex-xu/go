package main

import "fmt"

func isMatch(s string, p string) bool {
	lenStr, lenPattern := len(s), len(p)

	//判断：s[i-1] 与 p[j-1]是否匹配
	//内部函数，可以用外边函数定义的变量
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' { //'.' 匹配任意单个字符
			return true
		}
		return s[i-1] == p[j-1]
	}

	//dp[i][j]表示s[:i]与p[:j]时候匹配，最终要的是dp[lenStr][lenPattern]
	//把dp矩阵压缩成两行,dp为已知结果行
	dp := make([]bool, lenPattern+1)

	for i := 0; i <= lenStr; i++ {
		//把dp矩阵压缩成两行,dpNew为新的一行
		dpNew := make([]bool, lenPattern+1)
		if i == 0 { //因为p[:0]为nil,本来也不会用到
			dpNew[0] = true
		}
		for j := 1; j <= lenPattern; j++ {
			for {
				//*,可以对 p 的第 j-1 个字符(假设为'x')匹配任意自然数次
				if p[j-1] == '*' {
					if dpNew[j-2] { //用0个'x
						dpNew[j] = true
						break
					}
					//'x'与s的的第i个字符匹配　e.g. {s: "aa", p: "a*"}
					if matches(i, j-1) {
						//这里dp[j]即dp[i-1][j]
						//s[:i-1]与p[:j]匹配 &&'x'与s的的第i个字符匹配，这意味着＇x'多重复一次就达到s[:i]与p[:j]匹配的要求
						//s[:i-1]与p[:j]不匹配 &&'x'与s的的第i个字符匹配，这意味着＇x'多重复一次只达到了匹配s[:i]最后一个字符的目的
						dpNew[j] = dp[j]
						break
					}
				}
				if matches(i, j) {
					dpNew[j] = dp[j-1]
				} else {
					dpNew[j] = false
				}
				break
			} //for
		} //for j
		dp = dpNew
	} //for i

	return dp[lenPattern]
}

func main() {
	s := "aa"
	p := "a*"
	fmt.Println(isMatch(s, p))
}
