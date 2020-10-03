package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
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
	//candidates 中的数字可以无限制重复被选取
	//所以选取t后，l保持不变
	result := backtrace(candidates, target-t, l, h)
	for _, v := range result {
		ans = append(ans, append([]int{t}, v...))
	}
	result = backtrace(candidates, target, l+1, h)
	for _, v := range result {
		ans = append(ans, v)
	}
	return ans

}

func main() {
	fmt.Println("\ntc1:")
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println("\ntc2:")
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}
