package main

/* producer-consumer problem in Go */

import (
	"fmt"
)

/*
//定义只读的channel
read_only := make (<-chan int)
//定义只写的channel
write_only := make (chan<- int)
//可同时读写
read_write := make (chan int)
*/

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("produce ", i)
	}
	close(ch)
}

func consume(ch <-chan int, quit chan<- bool) {
	for i := range ch {
		fmt.Println("consume ", i)
	}
	quit <- true

}

func main() {
	ch := make(chan int, 2)
	go produce(ch)
	quit := make(chan bool)
	go consume(ch, quit)
	<-quit
}
