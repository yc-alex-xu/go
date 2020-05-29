package main

import (
	"fmt"
	"time"
)

func sum(a []int, name string, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
	fmt.Println(name, "end")

}

func main() {
	a := []int{7, 2, 8, -9, 4, 0} //slice
	fmt.Println("the slice is:", a)
	chan2sum := make(chan int)                 //定义一个channel
	go sum(a[:len(a)/2], "1st", chan2sum) //go_1
	x := <-chan2sum
	go sum(a[len(a)/2:], "2nd", chan2sum) //go_2
	//如果go_2和go_1 放一起，到时谁先被blocked中被唤醒就不确定了
	fmt.Println("sum of 1/2 slice =", x)
	y := <-chan2sum // receive from channel
	fmt.Println("sum of another 1/2 slice=", y)
	time.Sleep(time.Duration(1) * time.Second)
}
