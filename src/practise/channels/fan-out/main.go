package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, tasksCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task, ok := <-tasksCh
		if !ok {
			return
		}
		d := time.Duration(task) * time.Millisecond
		time.Sleep(d)
		fmt.Println("worker:", id, "\ttask:", task)
	}
}

func pool(wg *sync.WaitGroup, workers, tasks int) {
	tasksCh := make(chan int)

	for i := 0; i < workers; i++ {
		go worker(i, tasksCh, wg)
	}

	for i := 0; i < tasks; i++ {
		tasksCh <- i
	}

	close(tasksCh)
}

/*
can test how many goroutines can be created because goroutine stack exhuasted
*/
func main() {
	const (
		numGoroutine     = 36000
		numTask      int = 1.5 * numGoroutine
	)
	var wg sync.WaitGroup
	wg.Add(numGoroutine)
	go pool(&wg, numGoroutine, numTask)
	wg.Wait()
}
