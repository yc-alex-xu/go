package leetcode

import (
	. "leetcode/comm"
	"testing"
)

/*
     1
    / \
   2   3
*/
func tree1() *TreeNode {
	l := &TreeNode{Val: 2}
	r := &TreeNode{Val: 3}
	return &TreeNode{Val: 1, Left: l, Right: r}
}

/*
   -10
   / \
  9  20
    /  \
   15   7
*/
func tree2() *TreeNode {
	l := &TreeNode{Val: 15}
	r := &TreeNode{Val: 7}
	r = &TreeNode{Val: 20, Left: l, Right: r}
	l = &TreeNode{Val: 9}
	return &TreeNode{Val: -10, Left: l, Right: r}
}

func Test_maxPathSum(t *testing.T) {
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
			want: 6,
		},
		{
			name: "tc2",
			args: args{root: tree2()},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
