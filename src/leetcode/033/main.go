package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, h := 0, len(nums)
	for l < h {
		mid := l + (h-l)/2
		if nums[mid] == target {
			return mid
		}
		//左半边有序
		if nums[mid] > nums[l] {
			if target >= nums[l] && target < nums[mid] {
				h = mid
			} else {
				l = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[h-1] {
				l = mid + 1
			} else {
				h = mid
			}
		}
	} //for
	return -1
}
