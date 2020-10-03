package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	list := make([]int, 0)
	sort.Ints(nums)
	return backtrack(nums, 0, list, [][]int{})
}

// pos 下次添加到集合中的元素位置索引
// list 临时结果集合(每次需要复制保存)
func backtrack(nums []int, pos int, list []int, result [][]int) [][]int {
	// 把临时结果复制出来保存到最终结果
	ans := make([]int, len(list))
	copy(ans, list)
	result = append(result, ans)
	// 选择时需要剪枝、处理、撤销选择
	for i := pos; i < len(nums); i++ {
		// 排序之后，如果再遇到重复元素，则不选择此元素
		if i != pos && nums[i] == nums[i-1] {
			continue
		}
		list = append(list, nums[i])
		result = backtrack(nums, i+1, list, result)
		list = list[0 : len(list)-1]
	}
	return result
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
