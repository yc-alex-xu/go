package main

import (
	"fmt"
	"time"
)

func main() {
	var Ball int
	table := make(chan int)
	//can modify the number of players, ther is some intresting findings
	for i := 0; i < 8; i++ {
		go player(i, table)
	}
	table <- Ball
	time.Sleep(1 * time.Second)
	<-table
	fmt.Println("main goroutine wakeup,stop the game")
}

func player(id int, table chan int) {
	for {
		ball := <-table
		fmt.Println("plyer", id, "\treceived ball\t", ball)
		ball++
		time.Sleep(100 * time.Millisecond)
		table <- ball
		fmt.Println("plyer", id, "\tsend ball\t", ball)
	}
}
