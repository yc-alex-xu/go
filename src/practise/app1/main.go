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
	num_g    float64
	isActive bool // 全局变量声明

)

func main() {
	num := 2.0
	num_g = num
	fmt.Printf("你好，Sqrt(%v) = %v\n", num_g, math.Sqrt(num_g))
}
