package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
		defer func(n int) { fmt.Println("lamda=>", n) }(i)
	}

}
