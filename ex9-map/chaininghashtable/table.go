package chaininghashtable

import (
	"fmt"
	"strings"
)

const (
	DefaultM = 7
)

type ChainingHashTable struct {
	n       int
	m       int
	buckets [][]int
}

func New() *ChainingHashTable {
	return &ChainingHashTable{
		m:       DefaultM,
		buckets: make([][]int, DefaultM),
	}
}

func (t *ChainingHashTable) String() string {
	var builder strings.Builder
	builder.WriteString("----------------------------------------\n")
	for i, bucket := range t.buckets {
		if bucket == nil {
			continue
		}
		builder.WriteString(fmt.Sprintf("%d: %v\n", i, bucket))
	}

	return builder.String()
}

func (t *ChainingHashTable) hashFn(num int) int {
	return num % t.m
}

func (t *ChainingHashTable) Search(num int) bool {
	hash := t.hashFn(num)
	for _, key := range t.buckets[hash] {
		if key == num {
			return true
		}
	}
	return false
}

func (t *ChainingHashTable) Add(num int) bool {
	if t.Search(num) {
		return false
	}

	hash := t.hashFn(num)
	t.buckets[hash] = append(t.buckets[hash], num)
	t.n++
	return true
}

func (t *ChainingHashTable) Remove(num int) bool {
	if !t.Search(num) {
		return false
	}

	hash := t.hashFn(num)
	bucket := t.buckets[hash]

	i := 0
	for ; i < len(bucket); i++ {
		if bucket[i] == num {
			break
		}
	}

	bucket[i] = bucket[len(bucket)-1]
	t.buckets[hash] = bucket[:len(bucket)-1]
	t.n--

	return true
}
