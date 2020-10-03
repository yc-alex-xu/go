package main

import (
	. "leetcode/comm"
	"testing"
)

func tree1() *TreeNode {
	l := &TreeNode{Val: 15}
	r := &TreeNode{Val: 7}
	r = &TreeNode{Val: 20, Left: l, Right: r}
	l = &TreeNode{Val: 9}
	root := &TreeNode{Val: 3, Left: l, Right: r}
	return root
}

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{root: tree1()},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
