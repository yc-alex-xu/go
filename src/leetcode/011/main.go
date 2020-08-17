package main

func reverse(x int) int {
	ans := 0
	for x != 0 {
		ans = ans*10 + x%10
		x = x / 10
	}
	if ans < -(1<<31) || ans > 1<<31-1 {
		ans = 0
	}
	return ans
}
