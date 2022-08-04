package main

import (
	"fmt"
)

type node struct {
	data  int
	left  *node
	right *node
}

type binaryTree struct {
	root *node
}

var (
	stack []*node
)

func newnode(val int) *node {
	return &node{
		data: val,
	}
}

func (b *binaryTree) hasNext() bool {
	return len(stack) > 0
}

func (b *binaryTree) next() int {

	temp := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	traverseLeft(temp.right)
	return temp.data

}

func traverseLeft(n *node) {
	for n != nil {
		stack = append(stack, n)
		n = n.left
	}

}

func bstIterator(root *node) {
	traverseLeft(root)
}

func main() {
	fmt.Println("Binary Tree")

	b := binaryTree{
		newnode(10),
	}

	b.root.left = newnode(4)
	b.root.left.left = newnode(1)

	bstIterator(b.root)
	fmt.Println(b.hasNext())
	fmt.Println(b.next())
	fmt.Println(b.next())
	fmt.Println(b.hasNext())
	fmt.Println(b.next())
	fmt.Println(b.hasNext())

}
