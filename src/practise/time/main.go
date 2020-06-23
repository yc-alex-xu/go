package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Now()
	fmt.Println(t0, "\n", t0.Day(), "\n", t0.Local())
	t1 := time.NewTicker(1 * time.Second)

	select {
	case <-t1.C:
		fmt.Println(time.Now())
		return
	}

}
