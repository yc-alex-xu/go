package main

import (
	"fmt"
	"practise/math"
	"runtime"
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

func init() { // initialization of package
	//` 括起的字符串为Raw字符串，即字符串在代码中的形式就是打印时的形式，它没有字符转义，换行也将原样输出。
	s := `Hi
你
好`
	a := []byte(s)
	b := a
	a[0] = 'h'
	fmt.Println(string(b))
	fmt.Println(runtime.GOOS)
	go mapTest()

}

func mapTest() {
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

func main() {
	num := float64(numG)
	fmt.Printf("Sqrt(%v) = %v\n", numG, math.Sqrt(num))
}
