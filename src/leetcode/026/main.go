package main

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	//想象成两个数组，不过起始地址相同
	src, dst := 1, 0
	for src < n {
		//只有differtnt value 才需要copy
		if nums[src] != nums[dst] {
			dst++
			nums[dst] = nums[src]
		}
		src++
	}
	//因为dst是index,长度需要+1
	return dst + 1
}
