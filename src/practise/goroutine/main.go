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
	go say("world") //开一个新的Goroutines执行
	say("hello")    //当前Goroutines执行
}
