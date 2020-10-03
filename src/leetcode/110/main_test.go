package leetcode

import (
	. "leetcode/comm"
	"testing"
)

func tree1() *TreeNode {
	l := &TreeNode{Val: 15}
	r := &TreeNode{Val: 7}
	r = &TreeNode{Val: 20, Left: l, Right: r}
	l = &TreeNode{Val: 9}
	return &TreeNode{Val: 3, Left: l, Right: r}
}

func tree2() *TreeNode {
	l := &TreeNode{Val: 4}
	r := &TreeNode{Val: 4}
	l = &TreeNode{Val: 3, Left: l, Right: r}
	r = &TreeNode{Val: 3}
	l = &TreeNode{Val: 2, Left: l, Right: r}
	r = &TreeNode{Val: 2}
	return &TreeNode{Val: 1, Left: l, Right: r}
}

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{root: tree1()},
			want: true,
		},
		{
			name: "tc2",
			args: args{root: tree2()},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
