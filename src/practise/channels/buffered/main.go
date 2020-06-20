package main

import (
	"fmt"
	"runtime"
)

var numG int

func sum(a []int, name string, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	fmt.Println(name, "->")
	c <- total // send total to c
	// only main goroutine + self left
	if runtime.NumGoroutine() == 2 {
		close(c)
	}
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0} //slice
	fmt.Println("the slice is:", a)
	chan2sum := make(chan int, 2)          //指定channel buffer size
	go sum(a[:len(a)/2], "low", chan2sum)  //goroutine被调度的次序不太确定
	go sum(a[len(a)/2:], "high", chan2sum) //sender must close channel, otherwise deadlock
	for sum := range chan2sum {
		fmt.Println("received\t", sum)
	}
	fmt.Println("channel closed")
}
