package main

import (
	"fmt"
	"strings"
)

type Queue []string

func (q *Queue) Push(v string) {
	*q = append(*q, v)
}

func (q *Queue) Pop() string {
	if len(*q) <= 0 {
		return ""
	}

	r := (*q)[0]
	*q = (*q)[1:]
	return r
}

func (q Queue) String() string {
	buf := strings.Builder{}
	buf.WriteByte('[')
	for i, v := range q {
		if i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(v)
	}
	buf.WriteByte(']')

	return buf.String()
}

func main() {
	var q Queue
	q.Pop()
	q.Push("a")
	q.Push("b")
	q.Push("c")
	q.Pop()
	q.Pop()
	fmt.Println(q)
}
