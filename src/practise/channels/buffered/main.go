package main

import (
	"fmt"
	"runtime"
	"time"
)

func sum(a []int, name string, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	t := time.Now()
	c <- total
	fmt.Println(name, "return from <- cost", time.Since(t))
	// only main goroutine + self left
	if runtime.NumGoroutine() == 2 {
		close(c)
	}
}

func main() {
	fmt.Println("cpu number: ", runtime.NumCPU(), runtime.GOMAXPROCS(runtime.NumCPU()))
	a := []int{7, 2, 8, -9, 4, 0} //slice
	fmt.Println("the slice is:", a)
	const (
		bufSz = 3
	)
	chan2sum := make(chan int, bufSz)
	go sum(a[:len(a)/2], "low", chan2sum)
	go sum(a[len(a)/2:], "high", chan2sum) //sender must close channel, otherwise deadlock
	for sum := range chan2sum {
		fmt.Println("received\t", sum)
	}
	fmt.Println("channel closed")
}
