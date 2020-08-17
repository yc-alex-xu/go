package main

//双指针法
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	max := 0
	for l < r {
		volume := 0
		//在l,r两根柱子围成的空间倒水
		if height[l] < height[r] {
			//有个有趣的前提假设：柱子本身不占体积
			volume = height[l] * (r - l)
			l++
		} else {
			volume = height[r] * (r - l)
			r--
		}
		if volume > max {
			max = volume
		}
	}
	return max
}
