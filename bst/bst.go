package main

import "fmt"

type intTree struct {
	val         int
	left, right *intTree
}

func (it *intTree) insert(val int) *intTree {
	if it == nil {
		return &intTree{val: val}
	}
	if it.val > val {
		it.left = it.left.insert(val)
	} else {
		it.right = it.right.insert(val)
	}
	return it
}

func (it *intTree) contains(val int) bool {
	switch {
	case it == nil:
		return false
	case it.val > val:
		return it.left.contains(val)
	case it.val < val:
		return it.right.contains(val)
	default:
		return true
	}
}

func main() {
	it := intTree{val: 1}
	it.insert(5)
	it.insert(0)
	it.insert(3)
	it.insert(7)
	fmt.Println(it.contains(3))
	fmt.Println(it.contains(4))
}
