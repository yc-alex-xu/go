package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	//前l个字符时候相同
	isCommonPrefix := func(l int) bool {
		str0, count := strs[0][:l], len(strs)
		for i := 1; i < count; i++ {
			if strs[i][:l] != str0 {
				return false
			}
		}
		return true
	}
	//找出最短字符串
	lenShortest := len(strs[0])
	for _, s := range strs {
		if len(s) < lenShortest {
			lenShortest = len(s)
		}
	}

	//二分试探可能的长度
	low, high := 0, lenShortest
	for low < high {
		mid := (high-low+1)/2 + low
		if isCommonPrefix(mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return strs[0][:low]
}
