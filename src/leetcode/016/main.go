package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]
	diffMin := abs(result - target)

	for i := 0; i < len(nums)-2; i++ {
		l, h := i+1, len(nums)-1
		for l < h {
			sum := nums[i] + nums[l] + nums[h]
			switch {
			case sum == target:
				return sum
			case sum > target:
				h--
			case sum < target:
				l++
			}
			if t := abs(sum - target); t < diffMin {
				result = sum
				diffMin = t
			}
		}
	}
	return result
}

func abs(v int) int {
	if v < 0 {
		v = -v
	}
	return v
}
