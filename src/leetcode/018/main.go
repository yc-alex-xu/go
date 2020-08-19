package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][4]int {
	ans := make([][4]int, 0)
	n := len(nums)
	if n < 4 {
		return ans
	}
	//排序的目的：1，跳过重复元素 2，可以用双指针法，注意它的解不是唯一的。
	sort.Ints(nums)

	for i := 0; i < n-3; i++ {
		//上一轮已经用过的就不用了
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			//在数组的剩余部分用双指针求和
			target2 := target - nums[i] - nums[j]
			for l, h := j+1, len(nums)-1; l < h; {
				if t := nums[l] + nums[h]; t == target2 {
					ans = append(ans, [4]int{nums[i], nums[j], nums[l], nums[h]})
					l++ //模仿i的写法，先变化，然后看是不是要跳过更多
					for l < h && nums[l] == nums[l-1] {
						l++
					}
					h--
					for l < h && nums[h] == nums[h+1] {
						h--
					}

				} else if t < target2 {
					l++
				} else {
					h--
				}
			} //for l,h
		} //for j
	} //for i
	return ans
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}
