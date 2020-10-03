package main

func longestValidParentheses(s string) int {
	return dpWay(s)
}

func dpWay(s string) int {
	max, l := 0, len(s)
	//dp[i]的值表示，到index i为止的，最长有效子串长度
	dp := make([]int, l)
	for i := 1; i < l; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' { //匹配到最近的"("
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				//匹配到远处的"("，所以还包含有效子串
				dp[i] = dp[i-1] + 2
				if prev := i - dp[i-1] - 2; prev >= 0 { //本子串前还有子串
					dp[i] += dp[prev]
				}
			}
			if dp[i] > max {
				max = dp[i]
			}
		}
	}
	return max
}

func stackWay(s string) int {
	max := 0
	stack := []int{}
	//为什么要push idx=-1 是一难点
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
			continue
		}
		//一个只包含 '(' 和 ')' 的字符串
		//先弹出栈顶元素表示匹配了当前右括号
		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			//如果栈为空，说明当前的右括号为没有被匹配的右括号，我们将其下标放入栈中
			stack = append(stack, i)
		} else if l := i - stack[len(stack)-1]; l > max {
			//因为所有的下标idx都入过栈，
			max = l
		}
	}
	return max
}
