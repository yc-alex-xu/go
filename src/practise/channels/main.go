package main

import "fmt"

func sum(a []int, name string, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	fmt.Println(name, "send back sum")
	c <- total // send total to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0} //slice
	fmt.Println("the slice is:", a)
	chan2sum := make(chan int)                     //定义一个channel时，也需要定义发送到channel的值的类型
	go sum(a[:len(a)/2], "the 1st half", chan2sum) //goroutine被调度的次序不太确定
	go sum(a[len(a)/2:], "the 2nd half", chan2sum)
	x, y := <-chan2sum, <-chan2sum // receive from channel

	fmt.Println("one  half=", x, "the another half=", y)
}
