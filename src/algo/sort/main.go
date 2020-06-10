package main

import "fmt"

func bubble(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func insert(arr []int) {
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		var j int
		for j = i; j > 0; j-- {
			if arr[j-1] > tmp {
				arr[j] = arr[j-1]
			} else {
				break
			}
		}
		arr[j] = tmp
	}
}

func display(arr []int) {
	for _, v := range arr {
		fmt.Print(v, "\t")
	}
	fmt.Println()
}

func main() {
	arr := []int{2, 3, 1, 7, 9, 8, 6, 5, 4, 0}
	display(arr)
	//bubble(arr)
	insert(arr)
	display(arr)
}
