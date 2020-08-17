package main

func isPalindrome(x int) bool {
	//个位数
	if x >= 0 && x < 10 {
		return true
	}
	//负数，或者以0结尾
	if x < 0 || x%10 == 0 {
		return false
	}

	//循环右移到shift中
	shift := 0
	//如果是回文，当shift的位数(长度) >= x 的位数时退出
	for x > shift {
		shift = shift*10 + x%10
		x /= 10
	}

	// 当数字长度为奇数时，我们可以通过 shift/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，循环的末尾我们可以得到 x = 12，shift = 123，
	return x == shift || x == shift/10
}
