package main

func longestPalindrome(s string) string {
	return expand(s)
}

func expand(s string) string {
	if s == "" {
		return ""
	}
	l, r := 0, 0
	//c center of Palindrome
	for c := 0; c < len(s); c++ {
		lOdd, rOdd := expandAroundCenter(s, c, c)
		if t := rOdd - lOdd; t > r-l {
			l, r = lOdd, rOdd
		}
		lEven, rEven := expandAroundCenter(s, c, c+1)
		if t := rEven - lEven; t > r-l {
			l, r = lEven, rEven
		}
		//既然右边已经到顶了，c再右移的话Palindrome长度只可能减少
		if r == len(s)-1 {
			break
		}
	}
	return s[l : r+1]
}

func expandAroundCenter(s string, l, r int) (int, int) {

	for l >= 0 && r < len(s) && s[l] == s[r] {
		l, r = l-1, r+1
	}
	return l + 1, r - 1
}
