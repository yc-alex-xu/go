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

func main() {
	fmt.Println("cpu number: ", runtime.NumCPU())
	go say("world") //g1:开一个新的Goroutines执行
	say("hello")    //g2:当前Goroutines执行
	/*即使是g1,g2相互谦让，如果没有下一句的话，也是不能保证g1都能执行完
	$ go run main.go
	cpu number:  4
	goroutine: hello loop: 0
	goroutine: world loop: 0
	goroutine: hello loop: 1
	goroutine: world loop: 1
	goroutine: hello loop: 2
	goroutine: hello loop: 3
	goroutine: hello loop: 4
	*/
	fmt.Println("goodbye")
}
