package main

import "fmt"

type BinaryTree struct {
	root *Node
}

type Node struct {
	right, left *Node
	data        int64
}

func (b *BinaryTree) insert(data int64) {
	if b.root == nil {
		b.root = &Node{data: data}
	} else {
		b.root.insert(data)
	}
	return
}

func (n *Node) insert(data int64) {
	if n == nil {
		return
	}
	if data <= n.data {
		if n.left == nil {
			n.left = &Node{data: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {

			n.right = &Node{data: data}
		} else {
			n.right.insert(data)
		}
	}

}

func main() {
	fmt.Println("vim-go")
}
