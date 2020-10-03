package main

func maxProfit(prices []int) int {
	sum := 0
	//当天可以卖出后再买入，则可以抓住所有波段
	for i := 1; i < len(prices); i++ {
		if t := prices[i] - prices[i-1]; t > 0 {
			sum += t
		}
	}
	return sum
}

//学习dp 周火入魔了，忽略了这么简单的规律，每一波行情都抓住即可
func dpWay2(prices []int) int {
	l := len(prices)
	dp := make([]int, l+1)

	dp[0] = 0
	dp[1] = 0
	for i := 1; i < l; i++ {
		max := 0
		for j := 0; j < i; j++ {
			//如这次卖出股票
			if t := dp[j] + prices[i] - prices[j]; t > max {
				max = t
			}
		}
		if dp[i] > max {
			dp[i+1] = dp[i] //继续不持有
		} else {
			dp[i+1] = max //卖出
		}
	}
	max := dp[l-1]
	if max < dp[l] {
		max = dp[l]
	}
	return max
}

func dpWay(prices []int) int {
	//dp[i][j] 表示的是已经走第i步，持有第j步买入的股票,其值为盈利
	//dp[i][i+1]表示不持有股票
	l := len(prices)
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		//走到第i步时,可以不持有股票或者持有第0～i只股票
		dp[i] = make([]int, i+2)
	}

	dp[0][0] = 0
	dp[0][1] = 0
	for i := 1; i < l; i++ {
		max := 0
		for j := 0; j < i; j++ {
			dp[i][j] = dp[i-1][j] //继续持有原股票
			//如这次卖出股票
			if t := dp[i-1][j] + prices[i] - prices[j]; t > max {
				max = t
			}
		}
		dp[i][i] = dp[i-1][i] //加持
		if dp[i-1][i] > max {
			dp[i][i+1] = dp[i-1][i] //继续不持有
		} else {
			dp[i][i+1] = max //卖出
		}
	}
	max := dp[l-1][l-1]
	if max < dp[l-1][l] {
		max = dp[l-1][l]
	}
	return max
}
