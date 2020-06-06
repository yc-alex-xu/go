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
	/*go中并没有waitpid()这种函数，可以用全局变量来判断goroutine有没有结束
	主协程的代码并不需要消耗 forcePreemptNS （默认为 10 ms）时长的资源，而主线程也没有主动把资源让出来，因此子协程没有运行。
	$ go run goodbye.go
	cpu number:  4
	goodbye
	*/
	fmt.Println("goodbye")
}
