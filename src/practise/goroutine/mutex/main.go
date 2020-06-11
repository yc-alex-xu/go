package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
$ go doc sync Mutex
package sync // import "sync"

type Mutex struct {
	// Has unexported fields.
}
    A Mutex is a mutual exclusion lock. The zero value for a Mutex is an
    unlocked mutex.

    A Mutex must not be copied after first use.

func (m *Mutex) Lock()
func (m *Mutex) Unlock()
*/
var mutex sync.Mutex

func say(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println("goroutine:", name, "loop:", i)
	}
	mutex.Unlock()

}

func main() {
	fmt.Println("cpu number: ", runtime.NumCPU())
	mutex.Lock()
	go say("world") //g1:开一个新的Goroutines执行
	mutex.Lock()
	fmt.Println("goodbye")
}
