package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	//return recurse(nums)
	return backtrack(nums, 0, []int{}, [][]int{})
}

func recurse(nums []int) [][]int {
	//递归法，但是跟题解中的递归法又不太一样，由于slice的使用，看起来简洁一些
	if len(nums) > 0 {
		r := subsets(nums[1:])
		for _, e := range r {
			t := make([]int, len(e)+1, len(e)+1)
			t[0] = nums[0]
			copy(t[1:], e)
			r = append(r, t)
		}
		return r
	}
	//不用担心多次返回空集的情况
	//这段话，只有两种情况会被调用
	//1.第一次调用的参数就是空集，2.递归调用至空集的情况
	//两者都只会发生一次
	return [][]int{[]int{}}
}

// nums: 给定的集合
// pos:下次添加到集合中的元素位置索引
// curr:nums 中已经选中的元素
// result 最终结果
func backtrack(nums []int, pos int, curr []int, result [][]int) [][]int {
	//if the combination is done
	//需要复制一下，因为curr 在本func中还会被多次修改
	t := make([]int, len(curr))
	copy(t, curr)
	result = append(result, t)

	for i := pos; i < len(nums); i++ {
		//add nums[i] into the current combination
		t := append(curr, nums[i])
		//use next integers to complete the combination
		result = backtrack(nums, i+1, t, result)
		//since we use the temporary varialbe, no need to pop the nums[i]
	}
	return result
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
