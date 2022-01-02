package queue

type Node struct{}

type Queue []*Node

func NewQueue(cap int) *Queue {
	var q Queue = make([]*Node, 0, cap)
	return &q
}
func (q *Queue) Enqueue(node *Node) {
	*q = append(*q, node)
}

func (q *Queue) Dequeue() *Node {
	out := (*q)[0]
	*q = (*q)[1:]
	return out
}

func (q *Queue) Len() int {
	return len(*q)
}

func (q *Queue) Peek() *Node {
	return (*q)[0]
}

func (q *Queue) IsFull() {}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
