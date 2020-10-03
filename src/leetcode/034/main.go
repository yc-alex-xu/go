package main

func searchRange(nums []int, target int) []int {
	//如果数组中不存在目标值，返回 [-1, -1]
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l, h, n := 0, len(nums), len(nums)
	for l < h {
		mid := l + (h-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			h = mid
		} else { //相等
			for i := mid; i >= 0 && nums[i] == target; i-- {
				l = i
			}
			for i := mid; i < n && nums[i] == target; i++ {
				h = i
			}
			return []int{l, h}
		} //if-else

	} //for
	return []int{-1, -1}

}
