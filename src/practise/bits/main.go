package main

import "fmt"

func main() {
	//^ 取反
	var a, b byte = 0, 0xff
	a, b = ^a, ^b
	fmt.Printf("%b %b\n", a, b)
	var c, d int8 = 0, -1
	fmt.Printf("%X %X %X\n", c, d, byte(d))
	c, d = ^c, ^d
	fmt.Printf("%x %x %x\n", c, byte(c), d)

	//&^ 将指定位置上的值设置为 0
	//3 的二进制为 11，对应这两位被清空； 4的二进制为100，对应的一位被清空
	fmt.Printf("%b %b\n", a&^3, a&^4)
}
