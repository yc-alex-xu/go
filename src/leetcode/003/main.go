package main

func lengthOfLongestSubstring(s string) int {
	return slideWin(s)
}

func slideWin(s string) int {
	/*
	   用滑动窗口窗口法写了后，发现速度慢，看了下别人的题解都是用string方法，问题时候出在这里呢？试用投机的方式，
	   将map改为ascii字符的数组就取得了运行时间从16ms->0ms的惊人进步，看来go的map还是不够快．
	   m := map[byte]int{}
	*/
	m := [128]int{}
	n := len(s)

	r, max := 0, 0
	for l := 0; l < n; l++ {
		//右指针已经完成一次探索，需要尝试移动左指针了
		if l != 0 {
			// 左指针向右移动一格
			m[s[l-1]]--
		}
		//右指针开始探索,直到到一个无效值
		for r < n {
			if v := s[r]; m[v] == 0 {
				m[v]++
				r++
			} else {
				break
			}
		}
		if l := r - l; l > max {
			max = l
		}
	}
	return max
}
