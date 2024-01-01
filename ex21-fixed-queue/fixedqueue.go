package fixedqueue

import (
	"container/list"
	"fmt"
)

type FixedQueue struct {
	capacity int
	l        *list.List
}

type Element = list.Element

func New(capacity int) *FixedQueue {
	if capacity <= 0 {
		panic(fmt.Sprintf("capacity <= 0 capacity=%d", capacity))
	}

	return &FixedQueue{
		capacity: capacity,
		l:        list.New(),
	}
}

func (q *FixedQueue) Push(v any) {
	if q.l.Len() >= q.capacity {
		q.l.Remove(q.l.Front())
	}
	q.l.PushBack(v)
}

func (q *FixedQueue) Len() int {
	return q.l.Len()
}

func (q *FixedQueue) Cap() int {
	return q.capacity
}

func (q *FixedQueue) Front() *Element {
	return q.l.Front()
}
