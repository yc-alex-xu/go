package main

import "fmt"

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	ans = backtracking(n, n, "", ans)
	return ans
}

/* left,right 未使用的左右括号数目,
相比官方的 0,0，n 可以少用一个param */
func backtracking(left, right int, tmp string, ans []string) []string {
	/*
		回溯跳出条件，
		并不需要判断左括号是否用完，因为右括号生成的条件 right > left ，
		所以右括号用完了就意味着左括号必定用完了
	*/
	if right == 0 {
		ans = append(ans, tmp)
		return ans
	}

	// 生成左括号
	if left > 0 {
		ans = backtracking(left-1, right, tmp+"(", ans)
		/* 传的就是 tmp+"(",而不是先
		tmp += "("，然后传tmp
		所以不需要回嗍这一步了
		*/
	}

	// 括号成对存在，有左括号才会有右括号
	if right > left {
		ans = backtracking(left, right-1, tmp+")", ans)
		//同理，因为param用的是表达式，所以不用回嗍
	}
	return ans
}

func main() {
	fmt.Println("when n =3 ")
	for _, v := range generateParenthesis(3) {
		fmt.Println(v)
	}

	fmt.Println("when n =4 ")
	for _, v := range generateParenthesis(4) {
		fmt.Println(v)
	}
}
