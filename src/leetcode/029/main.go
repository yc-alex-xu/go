package main

func divide(dividend int, divisor int) int {
	negative := false
	if divisor < 0 {
		negative = !negative
		divisor = -divisor
	}
	if dividend < 0 {
		negative = !negative
		dividend = -dividend
	}
	ans := cal(dividend, divisor)
	if negative {
		return -ans
	}
	const max = (2 << 31) - 1
	if ans > max {
		return max
	}
	return ans
}

//全是正数
func cal(dividend int, divisor int) int {
	sum := 0
	for dividend >= divisor {
		shift := 0
		for dividend >= divisor<<shift {
			shift++
		}
		shift--
		sum += pow(shift)
		dividend -= divisor << shift
	}
	return sum
}

func pow(i int) int {
	ans := 1
	for ; i > 0; i-- {
		ans = ans << 1
	}
	return ans
}
