package main

func maxProfit(prices []int) int {
	//只能买卖一次，就只能抓最大的一次波段(用max表示)
	if len(prices) == 0 {
		return 0
	}
	min, max := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if t := prices[i] - min; t > max {
			max = t
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return max
}
