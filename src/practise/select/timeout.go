package main

import (
	. "fmt"
	"time"
)

func main() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				Println(v)
			case <-time.After(5 * time.Second):
				Println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o
}
