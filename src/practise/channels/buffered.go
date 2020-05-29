package main

import "fmt"

var num_goroutine int = 2

func sum(a []int, name string, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	fmt.Println(name, "send back sum")
	c <- total // send total to c
	if num_goroutine -= 1; num_goroutine == 0 {
		close(c)
	}
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0} //slice
	fmt.Println("the slice is:", a)
	chan2sum := make(chan int, 2)                  //指定channel buffer size
	go sum(a[:len(a)/2], "the 1st half", chan2sum) //goroutine被调度的次序不太确定
	go sum(a[len(a)/2:], "the 2nd half", chan2sum) //sender must close channel, otherwise deadlock
	for sum := range chan2sum {
		fmt.Println("sum of the half", sum)
	}

}
