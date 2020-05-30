package main

import (
	"fmt"
	"practise/math"
)

const (
	i      = 100
	pi     = 3.1415
	prefix = "Go_"
)

var (
	num_g    = 2
	isActive bool // 全局变量声明

)

func main() {
	num := float64(num_g)
	//` 括起的字符串为Raw字符串，即字符串在代码中的形式就是打印时的形式，它没有字符转义，换行也将原样输出。
	s := `你
		好
		world
		`
	fmt.Printf("%v Sqrt(%v) = %v\n", s, num_g, math.Sqrt(num))
}
