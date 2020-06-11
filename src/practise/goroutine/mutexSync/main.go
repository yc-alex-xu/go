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
var mutexHello sync.Mutex
var mutexWorld sync.Mutex

func sayHello(name string) {
	for i := 0; i < 5; i++ {
		mutexHello.Lock()
		fmt.Println("goroutine:", name, "loop:", i)
		mutexWorld.Unlock()
	}
}

func sayWorld(name string) {
	for i := 0; i < 5; i++ {
		mutexWorld.Lock()
		fmt.Println("goroutine:", name, "loop:", i)
		mutexHello.Unlock()
	}
}

func main() {
	fmt.Println("cpu number: ", runtime.NumCPU())
	mutexWorld.Lock()
	go sayWorld("world") //g1:开一个新的Goroutines执行
	sayHello("Hello")
	mutexHello.Lock()
	fmt.Println("goodbye")
}
