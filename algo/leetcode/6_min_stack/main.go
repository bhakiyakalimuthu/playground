package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	m := NewMinstack()
	m.push(-1)
	m.push(5)
	m.push(3)
	fmt.Printf("stack : %v, top : %d, min : %d\n", m, m.top(), m.min())
	m.pop()
	fmt.Printf("stack : %v, top : %d, min : %d\n", m, m.top(), m.min())
	m.push(-10)

	fmt.Printf("stack : %v, top : %d, min : %d\n", m, m.top(), m.min())
}

type MinStack struct {
	stack    []int
	minStack []int
}

func NewMinstack() *MinStack {
	return &MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}

}

// push
func (m *MinStack) push(item int) {
	if m.stack == nil {
		m.stack = make([]int, 0)
	}
	m.stack = append(m.stack, item)
	if len(m.minStack) == 0 {
		m.minStack = append(m.minStack, item)
	} else {
		if item < m.minStack[len(m.minStack)-1] {
			m.minStack = append(m.minStack, item)
		}
	}

}

// pop
func (m *MinStack) pop() int {
	top := m.stack[len(m.stack)-1]
	m.stack = m.stack[:len(m.stack)-1]
	if top == m.minStack[len(m.minStack)-1] {
		m.minStack = m.minStack[:len(m.minStack)]

	}
	return top
}

// top
func (m *MinStack) top() int {
	return m.stack[len(m.stack)-1]
}

// min
func (m *MinStack) min() int {
	return m.minStack[len(m.minStack)-1]
}
