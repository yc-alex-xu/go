package main

/* producer-consumer problem in Go */

import (
	"fmt"
	"runtime"
)

/*
Go中channel可以是只读、只写、同时可读写的。
//定义只读的channel
read_only := make (<-chan int)
//定义只写的channel
write_only := make (chan<- int)
//可同时读写
read_write := make (chan int)
*/

func produce(ch chan<- int, done chan<- bool) {
	for i := 0; i < 10; i++ {
		ch <- i
		runtime.Gosched()
	}
	done <- true
}

func consume(ch <-chan int, done <-chan bool, quit chan<- bool) {
	for {
		select {
		case i := <-ch:
			fmt.Println(i)
			runtime.Gosched()
		case <-done:
			quit <- true
			break
		}
	}
}

func main() {
	ch := make(chan int)
	done := make(chan bool)
	go produce(ch, done)
	quit := make(chan bool)
	go consume(ch, done, quit)
	<-quit
}
