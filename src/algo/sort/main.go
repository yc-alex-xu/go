package main

import (
	"fmt"
	"math/rand"
)

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

func merge(a []int, b []int) []int {

	var r = make([]int, len(a)+len(b))
	var i = 0
	var j = 0

	for i < len(a) && j < len(b) {

		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}

	}

	for i < len(a) {
		r[i+j] = a[i]
		i++
	}
	for j < len(b) {
		r[i+j] = b[j]
		j++
	}

	return r

}

//a divide and conquer algorithm
func mergesort(items []int) []int {

	if len(items) < 2 {
		return items

	}
	var middle = len(items) / 2
	var a = mergesort(items[:middle])
	var b = mergesort(items[middle:])
	return merge(a, b)

}

/*selectionSort 应该是边界条件最清晰的排序算法了
generally performs worse than the similar insertion sort.
Selection sort is noted for its simplicity, and it has performance advantages
over more complicated algorithms in certain situations, particularly where auxiliary memory is limited.
*/
func selectionSort(arr []int) {

	for i := 0; i < len(arr); i++ {
		min := i
		//选出第i小的,优点交换次数比较少；缺点不能提前退出
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if i != min {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
}

/*
Shell's method, is an in-place comparison sort.
It can be seen as either a generalization of sorting by exchange (bubble sort)
 or sorting by insertion (insertion sort).
*/
func shellSort(arr []int) {
	//步长=1 就大功告成
	for d := int(len(arr) / 2); d > 0; d /= 2 {
		for i := d; i < len(arr); i++ {
			for j := i; j >= d && arr[j-d] > arr[j]; j -= d {
				arr[j], arr[j-d] = arr[j-d], arr[j]
			}
		}
	}
	return
}

/*
Quicksort (sometimes called partition-exchange sort) is an efficient sorting algorithm,
serving as a systematic method for placing the elements of an array in order.
Developed by Tony Hoare in 1959 and published in 1961, it is still a commonly used algorithm for sorting.
*/

func quickSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	median := arr[rand.Intn(len(arr))]

	lowPart := make([]int, 0, len(arr))
	highPart := make([]int, 0, len(arr))
	middlePart := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < median:
			lowPart = append(lowPart, item)
		case item == median:
			middlePart = append(middlePart, item)
		case item > median:
			highPart = append(highPart, item)
		}
	}

	lowPart = quickSort(lowPart)
	highPart = quickSort(highPart)
	return append(append(lowPart, middlePart...), highPart...)
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
	nArr := mergesort(arr)
	display(nArr)

	//bubble(arr)
	//insert(arr)
	//selectionSort(arr)
	//shellSort(arr)
	display(arr)
	nArrQuick := quickSort(arr)
	display(nArrQuick)

}
