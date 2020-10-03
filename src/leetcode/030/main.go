package main

import "fmt"

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	//一些长度相同的单词 words
	lenWord := len(words[0])
	countWords := len(words)

	ans := make([]int, 0)
	for i := 0; i <= len(s)-countWords*lenWord; {
		m := map[string]int{}
		//切分
		for j := 0; j < countWords; j++ {
			m[s[i+j*lenWord:i+(j+1)*lenWord]]++
		}
		matched := true
		//要处理words中有重复的情况
		for _, w := range words {
			if m[w] < 1 {
				matched = false
				break
			} else {
				m[w]--
			}
		}
		if matched {
			ans = append(ans, i)
		}
		//words有极端情况，没有想出优化方法
		i++
	}
	return ans
}

func main() {

	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))

	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
}
