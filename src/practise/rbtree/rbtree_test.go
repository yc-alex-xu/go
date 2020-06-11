package rbtree

import (
	"reflect"
	"testing"
)

type Thing struct {
	Key Int
	Val string
}

func (t Thing) Less(item Item) bool {
	return t.Key < item.(Thing).Key
}

func (t Thing) More(item Item) bool {
	return t.Key > item.(Thing).Key
}

type Int int

func (i Int) Less(item Item) bool {
	return i < item.(Thing).Key
}

func (i Int) More(item Item) bool {
	return i > item.(Thing).Key
}

func TestInsertAndRetreivalRedBlack(t *testing.T) {
	tree := New()

	tree.Put(Thing{Int(6), "foo"})
	if tree.Size() != 1 {
		t.Fatal("Value not inserted correctly")
	}

	tree.Put(Thing{Int(7), "bar"})

	if tree.Size() != 2 {
		t.Fatal("Value not inserted correctly", tree.Size())
	}

	tree.Put(Thing{Int(7), "baz"})

	if tree.Size() != 2 {
		t.Fatal("Value not inserted correctly")
	}

	item, ok := tree.Find(Int(6))
	if !ok {
		t.Fatal("Value not found")
	}

	if item.(Thing).Val != "foo" {
		t.Fatal("Value not retreived correctly")
	}

	item, ok = tree.Find(Int(7))
	if !ok {
		t.Fatal("Value not found")
	}

	if item.(Thing).Val != "baz" {
		t.Fatal("Value not retreived correctly")
	}
}

func BenchmarkRedBlack(b *testing.B) {
	var (
		keys = make([]int, 0, b.N)
		tree = New()
	)

	for i := 0; i < b.N*3; i++ {
		keys = append(keys, i)
		tree.Put(Thing{Int(i), "foo"})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = tree.Find(Int(keys[i*3]))
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *RedBlackTree
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_put(t *testing.T) {
	type args struct {
		n    *Node
		item Item
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := put(tt.args.n, tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("put() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotateRight(t *testing.T) {
	type args struct {
		n *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
