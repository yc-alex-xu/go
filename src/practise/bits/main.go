package main

import "fmt"

func main() {
	//^   取补数
	var a, b byte
	a, b = 1, ^a
	var c, d uint8
	c, d = 1, ^c
	fmt.Printf("%b %b\n", b, d)

	//&^ 将指定位置上的值设置为 0
	//3 的二进制为 11，对应这两位被清空； 4的二进制为100，对应的一位被清空
	fmt.Printf("%b %b\n", b&^3, d&^4)
}
