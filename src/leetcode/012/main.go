package main

//体会一下没有　进位制　的罗马数学是多么折磨人
func intToRoman(num int) string {
	//设计映射表，实际就是设计优先匹配顺序
	//由于go 中map 的顺序是不确定的,所以改用两个数组
	ints := [...]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := [...]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	//贪心法
	ans := ""
	for idx, v := range ints {
		n := num / v
		for i := 0; i < n; i++ {
			ans += romans[idx]
		}
		num -= n * v
	}
	return ans
}
