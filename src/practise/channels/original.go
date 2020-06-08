package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c) //g1
	go sum(a[len(a)/2:], c) //g2,g1和g2的调度次序不定
	x, y := <-c, <-c        // receive from c

	fmt.Println(x, y, x+y)
}
