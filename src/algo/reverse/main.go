package main

import (
	"fmt"
)

type node struct {
	val  int
	next *node
}

func recurse(h *node) (nh *node) {
	next := h.next
	if next != nil {
		nh = recurse(next)
		next.next = h
	} else {
		nh = h
	}
	return
}

func reverse(h *node) (nh *node) {
	var prev *node
	prev = nil

	for p := h; p != nil; {
		next := p.next
		p.next = prev
		prev = p
		p = next
	}
	nh = prev
	return
}

func display(h *node) {
	for h != nil {
		fmt.Print(h.val, "\t")
		h = h.next
	}
	fmt.Println()
}

func main() {
	var next *node
	next = nil
	var p *node
	for i := 10; i > 0; i-- {
		p = new(node)
		p.next = next
		p.val = i
		next = p
	}
	display(p)
	nh := reverse(p)
	display(nh)

	p = recurse(nh)
	nh.next = nil
	display(p)
}
