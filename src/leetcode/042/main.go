package main

func trap(height []int) int {
	l, h := 0, len(height)-1
	maxL, maxH := 0, 0
	ans := 0
	for l < h {
		if height[l] < height[h] {
			if t := height[l]; t > maxL {
				maxL = t
			} else {
				//积水
				ans += maxL - t
			}
			l++
			continue
		} //if

		if t := height[h]; t > maxH {
			maxH = t
		} else {
			ans += maxH - t
		}
		h--
	} //for
	return ans
}
