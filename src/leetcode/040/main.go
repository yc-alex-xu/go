package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	l, h := 0, len(candidates)
	return backtrace(candidates, target, l, h)
}

func backtrace(candidates []int, target int, l, h int) [][]int {
	ans := [][]int{}
	if l >= h {
		return ans
	}
	t := candidates[l]
	if t > target {
		return ans
	}
	if t == target {
		ans = append(ans, []int{t})
		return ans
	}

	//考虑多个元素的情况
	result := backtrace(candidates, target-t, l+1, h)
	for _, v := range result {
		ans = append(ans, append([]int{t}, v...))
	}

	//解集不能包含重复的组合
	for l < h && candidates[l] == t {
		l++
	}
	result = backtrace(candidates, target, l, h)
	for _, v := range result {
		ans = append(ans, v)
	}
	return ans
}

func main() {
	fmt.Println("\ntc1:")
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println("\ntc2:")
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}
