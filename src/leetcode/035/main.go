package main

func searchInsert(nums []int, target int) int {
	n := len(nums)
	l, h := 0, n-1
	for l <= h {
		mid := l + (h-l)/2
		v := nums[mid]
		if target == v {
			return mid
		} else if target < v {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	/*退出循环的时候，l 肯定比 h 大
	if l < n && target > nums[l] {
		return l + 1
	}
	*/
	return l
}
