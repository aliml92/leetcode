package datastructure

// Queque implemented with singly-linked list
//
//   head              tail
//    []-[]-[]-[]-[]-[]-[]
//
//

type Node[T any] struct {
	value *T
	next  *Node[T]
}

type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
	len  int
}

func NewQueue[T any]() *Queue[T] {
	return new(Queue[T])
}

func (q *Queue[T]) Enqueue(item *T) {
	n := &Node[T]{
		value: item,
		next:  nil,
	}
	if q.len == 0 {
		q.head = n
	} else {
		q.tail.next = n
	}
	q.tail = n
	q.len++
}

// Popleft removes and returns the first element of the queue
func (q *Queue[T]) Dequeue() *T {
	val := q.head.value
	nextPtr := q.head.next
	q.head = nil
	q.head = nextPtr
	q.len--
	return val
}

func (q Queue[T]) Len() int {
	return q.len
}
