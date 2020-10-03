package main

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			//因为idx = v-1
			//而v= n+1 => idx = n 是无效的index
			//nums[]当然可能有>=n+1的值，但是这时，
			//firstMissingPositive 肯定是<n的，因为只有n个element
			nums[i] = n + 1
		}
	}
	//nums[i] 为负，表示i+1的存在
	//所以nums[]只能表示1~n 是否存在
	for i := 0; i < n; i++ {
		if v := abs(nums[i]); v <= n {
			//只要nums[]中的值各不相同，只会取负一次
			nums[v-1] = -abs(nums[v-1])
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	//1~n 都在,那nums[]就刚好包含1~n
	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
