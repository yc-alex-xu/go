package main

import "fmt"

type testInt func(int) bool // 声明了一个函数类型

func isOdd(integer int) bool {
	return (integer%2 != 0)
}

func isEven(integer int) bool {
	return integer%2 == 0
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("slice:", slice)
	odd := filter(slice, isOdd) // 函数当做值来传递了
	/*alex: c/c++ 中是忌讳called func 回传在它那儿分配的空间的,
	也许go因为对每一变量的使用都进行跟踪，所以这块空间不会被复用
	*/
	even := filter(slice, isEven)
	fmt.Println("Odd elements:", odd)
	fmt.Println("Even elements:: ", even)
}
