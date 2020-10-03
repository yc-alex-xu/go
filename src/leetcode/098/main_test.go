package main

import "testing"
import . "leetcode/comm"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	l := &TreeNode{Val: 1}
	r := &TreeNode{Val: 3}
	tree1 := &TreeNode{Val: 2, Left: l, Right: r}

	l = &TreeNode{Val: 3}
	r = &TreeNode{Val: 6}
	r = &TreeNode{Val: 4, Left: l, Right: r}
	l = &TreeNode{Val: 1}
	tree2 := &TreeNode{Val: 5, Left: l, Right: r}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{root: tree1},
			want: true,
		},
		{
			name: "tc2",
			args: args{root: tree2},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
