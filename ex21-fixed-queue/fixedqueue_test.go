package fixedqueue

import (
	"fmt"
	"testing"
)

func TestFixed(t *testing.T) {
	q := New(3)
	for i := 0; i < q.Cap(); i++ {
		q.Push(i)
	}
	if q.Len() != q.Cap() {
		t.Errorf("q.Len() != q.Cap() len:%d cap:%d", q.Len(), q.Cap())
	}

	for e := q.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	q.Push(100)
	if q.Len() != q.Cap() {
		t.Errorf("after push overflow q.Len() != q.Cap() len:%d cap:%d", q.Len(), q.Cap())
	}
	for e := q.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
