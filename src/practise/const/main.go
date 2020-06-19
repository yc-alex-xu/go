package main

import "fmt"

/*
the previous expression and its type can be used again in a group
*/
func seqConst() {
	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d) // "1 1 2 2"
}

func shiftConst() {
	type Flags byte
	const (
		FlagUp Flags = 1 << iota
		FlagBroadcast
		FlagLoopback
		FlagPointToPoint
		FlagMulticast
	)
	//"1 10 100 1000 10000"
	fmt.Printf("%b %b %b %b %b\n", FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)
}

func main() {
	seqConst()
	shiftConst()
}
