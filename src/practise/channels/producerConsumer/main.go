package main

/* producer-consumer problem in Go */

import (
	"fmt"
	"runtime"
)

var done = make(chan bool)
var quit = make(chan bool)
var ch = make(chan int)

func produce() {
	for i := 0; i < 10; i++ {
		ch <- i
		runtime.Gosched()
	}
	done <- true
}

func consume() {
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
	go produce()
	go consume()
	<-quit
}
