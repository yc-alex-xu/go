package main

func maxSubArray(nums []int) int {
	//return dpMethod(nums)
	return get(nums, 0, len(nums)-1).mSum
}
func findMaxpostive(nums []int) int {
	bPositiveState := true
	sum, max := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			if bPositiveState {
				if sum > max {
					max = sum
				}
				bPositiveState = false
			}
			continue
		}
		if !bPositiveState {
			bPositiveState = true
			sum = 0
		}
		sum += nums[i]
	}
	if sum > max {
		return sum
	}
	return max

}

func dpMethod(nums []int) int {
	// 构造子问题，dp[n]代表以n结尾的最大子序和
	// 状态初始化
	dp, max := nums, nums[0]
	for i := 1; i < len(dp); i++ {
		//状态转移方程
		if dp[i-1] > 0 {
			dp[i] = dp[i] + dp[i-1]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum+r.lSum)
	rSum := max(r.rSum, r.iSum+l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum+r.lSum)
	return Status{lSum, rSum, mSum, iSum}
}

//线段树求解
//我们定义一个操作 get(a, l, r) 表示查询 aa 序列 [l, r][l,r] 区间内的最大子段和
func get(nums []int, l, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	lSub := get(nums, l, m)
	rSub := get(nums, m+1, r)
	return pushUp(lSub, rSub)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
Status lSum 表示 [l, r][l,r] 内以 l 为左端点的最大子段和
rSum 表示 [l, r][l,r] 内以 r 为右端点的最大子段和
mSum 表示 [l, r][l,r] 内的最大子段和
iSum 表示 [l, r][l,r] 的区间和
*/
type Status struct {
	lSum, rSum, mSum, iSum int
}
