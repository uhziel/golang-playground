package linearprobinghashtable

import (
	"fmt"
	"strings"
)

const (
	DefaultM = 7
)

type LinearProbingHashTable struct {
	n    int
	m    int
	keys []*int
}

type optFn func(t *LinearProbingHashTable)

func New(opts ...optFn) *LinearProbingHashTable {
	t := &LinearProbingHashTable{
		m: DefaultM,
	}

	for _, opt := range opts {
		opt(t)
	}

	t.keys = make([]*int, t.m)

	return t
}

func WithM(m int) optFn {
	return func(t *LinearProbingHashTable) {
		t.m = m
	}
}

func (t *LinearProbingHashTable) String() string {
	var builder strings.Builder
	builder.WriteString("----------------------------------------\n")
	for i, key := range t.keys {
		if key == nil {
			continue
		}
		builder.WriteString(fmt.Sprintf("%d: %v\n", i, *key))
	}

	return builder.String()
}

func (t *LinearProbingHashTable) hashFn(num int) int {
	return num % t.m
}

func (t *LinearProbingHashTable) resize(size int) {
	tmp := New(WithM(size))
	for _, key := range t.keys {
		if key == nil {
			continue
		}
		tmp.Add(*key)
	}
	t.keys = tmp.keys
	t.m = tmp.m
}

func (t *LinearProbingHashTable) Search(num int) bool {
	hash := t.hashFn(num)

	for i := hash; t.keys[i] != nil; i = (i + 1) % t.m {
		if *(t.keys[i]) == num {
			return true
		}
	}

	return false
}

func (t *LinearProbingHashTable) Add(num int) bool {
	if t.Search(num) {
		return false
	}

	if t.n > t.m/2 {
		t.resize(2 * t.m)
	}

	hash := t.hashFn(num)
	i := hash
	for ; t.keys[i] != nil; i = (i + 1) % t.m {
	}
	t.keys[i] = &num
	t.n++
	return true
}

func (t *LinearProbingHashTable) Remove(num int) bool {
	if !t.Search(num) {
		return false
	}

	hash := t.hashFn(num)
	i := hash
	for ; t.keys[i] != nil; i = (i + 1) % t.m {
		if *(t.keys[i]) == num {
			t.keys[i] = nil
			break
		}
	}
	for i := (i + 1) % t.m; t.keys[i] != nil; i = (i + 1) % t.m {
		tmp := *(t.keys[i])
		t.n--
		t.keys[i] = nil
		t.Add(tmp)
	}
	t.n--

	if t.n < t.m/8 {
		t.resize(t.m / 2)
	}

	return true
}
