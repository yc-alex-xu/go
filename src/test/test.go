package main

//二维的二分法
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	low, high := matrix[0][0], matrix[n-1][n-1]
	for low < high {
		mid := low + (high-low)/2
		//因为count()中用的标准是<=mid,
		//count(,mid)==k时，没法确定mid到底是第k个还是k+1个
		if count(matrix, mid, n) >= k {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

//统计矩阵左上角（因为行、列有序）小于mid的元素数目
func count(matrix [][]int, mid, n int) int {
	//从最左下角开始==>最右上角
	row, col := n-1, 0
	num := 0
	for row >= 0 && col < n {
		if matrix[row][col] <= mid {
			num += row + 1 //[0][col]`[row][col]自然都<=mid`
			col++
		} else {
			row--
		}
	}
	return num
}
