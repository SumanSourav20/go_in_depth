package main

import (
	"cmp"
	"fmt"
)

type orderableFunc[T any] func(t1, t2 T) int

type Tree[T any] struct {
	f    orderableFunc[T]
	root *Node[T]
}

type Node[T any] struct {
	val         T
	left, right *Node[T]
}

func NewTree[T any](f orderableFunc[T]) Tree[T] {
	return Tree[T]{
		f: f,
	}
}

func (tree *Tree[T]) add(val T) {
	tree.root = tree.root.add(val, tree.f)
}

func (node *Node[T]) add(val T, f orderableFunc[T]) *Node[T] {
	if node == nil {
		return &Node[T]{val: val}
	}
	switch r := f(val, node.val); {
	case r <= -1:
		node.left = node.left.add(val, f)
	case r >= 1:
		node.right = node.right.add(val, f)
	}

	return node
}

func (tree Tree[T]) contains(val T) bool {
	return tree.root.contains(val, tree.f)
}

func (node *Node[T]) contains(val T, f orderableFunc[T]) bool {
	if node == nil {
		return false
	}
	switch r := f(val, node.val); {
	case r <= -1:
		return node.left.contains(val, f)
	case r >= 1:
		return node.right.contains(val, f)
	}
	return true
}

func main() {

	intTree := NewTree(cmp.Compare[int])

	intTree.add(1)
	intTree.add(3)
	intTree.add(5)
	intTree.add(9)

	fmt.Println(intTree.contains(5))
}
