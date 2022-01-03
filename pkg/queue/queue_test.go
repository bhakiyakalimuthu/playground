package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue(10)
	q.Enqueue(&Node{element: 1})
	q.Enqueue(&Node{element: 2})
	q.Enqueue(&Node{element: 3})
	q.Enqueue(&Node{element: 4})
	// length
	assert.Equal(t, 4, q.Len())

	// deque
	n := q.Dequeue()
	t.Logf("deque 1 : %d", n.element.(int))
	assert.Equal(t, 3, q.Len())
	// deque
	n = q.Dequeue()
	t.Logf("deque 2 : %d", n.element.(int))
	assert.Equal(t, 2, q.Len())

	assert.Equal(t, 3, q.Peek().element.(int))

}
