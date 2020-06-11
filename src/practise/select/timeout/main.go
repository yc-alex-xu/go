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
			case <-c:
				Println("unexpected signal come!")
			case <-time.After(1 * time.Second):
				Println("timeout")
				o <- true
				break
			}
		}
	}()
	v := <-o //如丢弃收到的，写成 <-o即可
	Println("received ", v)
}
