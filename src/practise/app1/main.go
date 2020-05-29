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
	fmt.Printf("你好，Sqrt(%v) = %v\n", num_g, math.Sqrt(num))
}
