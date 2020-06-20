package main

import (
	"fmt"
	"runtime"
)

func say(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println("goroutine:", name, "loop:", i)
		runtime.Gosched() //让CPU把时间片让给别人

	}
}

/*
$ go run main.go
cpu number:  4
goroutine: hello loop: 0
goroutine: world loop: 0
goroutine: hello loop: 1
goroutine: hello loop: 2
goroutine: hello loop: 3
goroutine: hello loop: 4
goodbye
*/
func main() {
	fmt.Println("cpu number: ", runtime.NumCPU())
	go say("world")
	say("hello")
	fmt.Println("goodbye")
}
