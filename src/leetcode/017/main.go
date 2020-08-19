package main

import "fmt"

var m = [...]string{
	"abc",
	"def",
	"ghi",
	"jkl",
	"mno",
	"pqrs",
	"tuv",
	"wxyz",
}

func letterCombinations(digits string) []string {
	ans := make([]string, 0)
	if len(digits) > 0 {
		backtrack(&ans, "", []byte(digits))
	}
	return ans
}

//感觉可以改为一个多叉树的遍历问题
func backtrack(ans *[]string, combination string, nextDigits []byte) {
	if len(nextDigits) == 0 {
		*ans = append(*ans, combination)
		return
	}

	//注意 1 不对应任何字母
	letters := m[nextDigits[0]-'2']
	for _, v := range letters {
		backtrack(ans, combination+string(v), nextDigits[1:])
	}
	return
}

func main() {
	fmt.Println(letterCombinations("23"))
}
