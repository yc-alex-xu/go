package main

func removeElement(nums []int, val int) int {
	i, n := 0, len(nums)
	//题目要求：元素的顺序可以改变
	//这样数组元素就没有必要顺序前移了
	//而是直接把不符合条件的元素剔除（覆盖）即可
	for i < n {
		if nums[i] == val {
			//也许此时i == n-1,不过这种特例就一次
			nums[i] = nums[n-1]
			n--
		} else { //元素i已经符合条件，下一个
			i++
		}
	}
	return i
}
