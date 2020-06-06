/*
defer 对closure影响好像因go 版本而异
*/
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("normal func=>", i)
		defer func(n int) {
			fmt.Println("lamda func=>", n*2)
		}(i) //
	}
}
