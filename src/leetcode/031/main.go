package main

import (
	"fmt"
)

func nextPermutation(nums []int) {
	//只需要将后面的「大数」与前面的「小数」交换
	n := len(nums)
	if n <= 1 {
		return
	}
	//增加的幅度尽可能的小-->尽可能靠右的低位进行交换，需要从后向前找到第一个减少的数目
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 { //否则也无解了
		//找到nums[i]右边，刚刚比它的数用来交换
		j := n - 1
		//根据i的产生办法：一个拐点。 所以i右边一定有这么一个j
		for nums[j] <= nums[i] {
			j--
		}
		//按i,j的产生办法： j(包括j) 以后肯定是递减的，i+1:j 之间也是递减的
		//i,j 交换后，以上的两个判断继续成立
		nums[i], nums[j] = nums[j], nums[i]
		//把i后面的数字，从降序改为升序
		for l, h := i+1, n-1; l < h; {
			nums[l], nums[h] = nums[h], nums[l]
			l++
			h--
		}
	} else {
		//如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）
		//说明原来就是降序，改为升序
		for l, h := 0, n-1; l < h; {
			nums[l], nums[h] = nums[h], nums[l]
			l++
			h--
		}
	}
}

func main() {
	for _, arr := range [][]int{{1, 2, 3}, {3, 2, 1}, {1, 1, 5}} {
		nextPermutation(arr)
		fmt.Println(arr)
	}
}
