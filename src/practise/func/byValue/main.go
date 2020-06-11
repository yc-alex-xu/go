package main

import "fmt"

func incr1(a int) int {
	a++
	return a
}

func main() {
	x := 3
	fmt.Println("x=", x)
	x1 := incr1(x)
	fmt.Println("x=", x) // 应该输出"x = 3"
}
