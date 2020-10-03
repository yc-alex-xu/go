package main

func strStr(haystack string, needle string) int {

	l := len(needle)
	for i := 0; i <= len(haystack)-l; i++ {
		if haystack[i:i+l] == needle {
			return i
		}
	}
	return -1
}
