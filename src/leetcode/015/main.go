package main

import "sort"

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	l := len(nums)
	if l < 3 {
		return ans
	}
	sort.Ints(nums)

	//i point to a
	for i := 0; i < l-2; i++ {
		a := nums[i]
		if a > 0 {
			//因为后面的肯定比a大
			return ans
		}
		//跳过重复的a
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		//j,k point to b,c
		//用双指针法
		for j, k := i+1, l-1; j < k; {
			b, c := nums[j], nums[k]
			if a+b+c == 0 {
				ans = append(ans, []int{a, b, c})
				//跳过重复的b
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				//跳过重复的c
				for j < k && nums[k] == nums[k-1] {
					k--
				}
			} //if

			//对于a+b+c==target的情况,以k--或j++开始皆可，因为两步都是必然的
			if a+b+c > 0 {
				k--
			} else {
				j++
			}
		} //for j,k
	} //for i
	return ans
}
