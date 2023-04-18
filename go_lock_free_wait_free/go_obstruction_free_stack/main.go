package main

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type Stack struct {
	head *Node
}

func New(data interface{}) *Node {
	return &Node{data: data}
}

func (s *Stack) Push(data interface{}) {
}
func (s *Stack) Pop() {
}

func main() {
	fmt.Println("vim-go")
}
