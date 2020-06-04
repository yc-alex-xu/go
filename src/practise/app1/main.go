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
	numG     = 2
	isActive bool // 全局变量声明

)

func main() {
	num := float64(numG)
	//` 括起的字符串为Raw字符串，即字符串在代码中的形式就是打印时的形式，它没有字符转义，换行也将原样输出。
	s := `Hi
		你
		好
		world
		`
	fmt.Printf("Sqrt(%v) = %v\n", numG, math.Sqrt(num))
	a := []byte(s)
	b := a
	a[0] = 'h'
	fmt.Println(string(b))
	m := make(map[int]string)
	m[0] = "John"
	m[1] = "Mary"
	fmt.Println(m)
	m2 := map[string]string{
		"+86": "China",
		"+1":  "North America",
	}
	fmt.Println(m2["+86"])
	for k, v := range m2 {
		fmt.Println(k, v)
	}
}
