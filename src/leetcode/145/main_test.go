package leetcode

import (
	. "leetcode/comm"
	"reflect"
	"testing"
)

/*
   1
    \
     2
    /
   3
*/
func tree1() *TreeNode {
	l := &TreeNode{Val: 3}
	r := &TreeNode{Val: 2, Left: l}
	return &TreeNode{Val: 1, Right: r}
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

func Test_postorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "tc1",
			args: args{root: tree1()},
			want: []int{3, 2, 1},
		},
		{
			name: "tc2",
			args: args{root: tree2()},
			want: []int{9, 15, 7, 20, -10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
