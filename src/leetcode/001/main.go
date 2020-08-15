package main

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i, v := range nums {
		dict[v] = i
	}
	for i, v := range nums {
		if j, ok := dict[target-v]; ok && j != i { //found
			return []int{i, j}
		}
	}
	return nil
}
