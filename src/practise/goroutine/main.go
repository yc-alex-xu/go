package main

import (
	"fmt"
	"runtime"
)

func say(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println(name, "loop:", i)
		runtime.Gosched() //让CPU把时间片让给别人
	}
}

/*
go$ go run src/practise/goroutine/main.go
cpu number:  4
hello loop: 0
hello loop: 1
hello loop: 2
hello loop: 3
hello loop: 4
goodbye
go$ go run src/practise/goroutine/main.go
cpu number:  4
hello loop: 0
world loop: 0
hello loop: 1
hello loop: 2
hello loop: 3
hello loop: 4
goodbye

*/
func main() {
	fmt.Println("cpu number: ", runtime.NumCPU())
	go say("world")
	say("hello")
	fmt.Println("goodbye")
}
