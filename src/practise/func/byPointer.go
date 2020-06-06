package main

import "fmt"

func add1(a *int) int {
	*a++
	return *a
}

func main() {
	x := 3
	fmt.Println("x = ", x)
	x1 := add1(&x) // 调用 add1(&x) 传x的地址
	fmt.Println("x+1 = ", x1)
	fmt.Println("x = ", x) // 应该输出 "x = 4"
}
