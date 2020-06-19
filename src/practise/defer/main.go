/*
defer 对closure影响好像因go 版本而异
*/
package main

import (
	"fmt"
)

func deferSimple() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("normal func=>", i)
		defer func() {
			fmt.Println("lamda func loop variable =>", i)
		}()
		defer func(n int) {
			fmt.Println("lamda func param=>", n)
		}(i) //
	}

}

func deferComplicated() {
	a := 1
	b := 2
	//calc("10", a, b) 并不会被defer
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	deferSimple()
	deferComplicated()
}
