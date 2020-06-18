package main

import (
	"fmt"
)

type node struct {
	l, r *node
	val  int
}

func buildTree(pre []int, in []int) (r *node) {
	if len(pre) == 0 {
		return nil
	}
	r = new(node)
	r.val = pre[0]
	var idx int
	for ; idx < len(in); idx++ {
		if in[idx] == pre[0] {
			break
		}
	}
	r.l = buildTree(pre[1:idx+1], in[:idx])
	r.r = buildTree(pre[idx+1:], in[idx+1:])
	return

}

func traverse(r *node) {
	if r != nil {
		fmt.Print("val:", r.val)
	} else {
		return
	}

	if r.l != nil {
		fmt.Print("\tleft:", r.l.val)
	}
	if r.r != nil {
		fmt.Println("\tright:", r.r.val)
	}
	fmt.Println()
	traverse(r.l)
	traverse(r.r)

}

func main() {
	pre := []int{43, 12, 8, 32, 16, 35, 78, 56, 88, 83, 121, 97}
	in := []int{8, 12, 16, 32, 35, 43, 56, 78, 83, 88, 97, 121}

	r := buildTree(pre, in)
	traverse(r)
}
